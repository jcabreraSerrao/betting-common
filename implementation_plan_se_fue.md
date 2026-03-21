# Feature: Apuestas Propias con Cancelación "Se Fue"

## Descripción del Problema

Actualmente el sistema **solo procesa mensajes de juego si son una respuesta (reply)** a otro mensaje. Si alguien pone un mensaje de "juego 3-5 15.000" *sin responder a nadie*, se descarta silenciosamente.

Queremos implementar dos comportamientos nuevos interrelacionados:

### 1. Apuestas Propias (sin reply → estado `pending`)
Cuando `tercio A` pone `juego 3-5 15.000` (sin reply), el sistema debe:
- Detectar que es un mensaje de juego sin reply
- Guardarlo en `matched_bet_logs` con estado **`pending`** y `remaining_amount = -1`
- **NO** llamar a TurfLex ni a `auto_bet` (el NLU se omite, solo se registra el mensaje crudo)
- Cuando alguien responda ese mensaje, el flujo normal procesa por primera vez el NLU

### 2. Cancelación con "Se Fue" / "Se Fue Todo"
El dueño original puede cancelar con `se fue` o `se fue todo`. El comportamiento depende de si hay reply:

| Variante | Reply | Resultado |
|----------|-------|-----------|
| `se fue` o `se fue todo` **CON reply** | Sí (a una jugada específica) | Cierra **solo** esa jugada (status → `closed`) |
| `se fue` o `se fue todo` **SIN reply** | No | Cierra **todas** las jugadas `pending`/`open` del tercio en el grupo |

- No responde nada al grupo en ningún caso
- Cuando otro tercio intente tomar una jugada cerrada: verifica `closed` **ANTES** de `remaining_amount` → responde `NADA ⛔`

---

## Análisis del Flujo Actual

```
Mensaje llega por Kafka → dispatchCommand() (consumer.go)
  → isGameCommand(firstWord) → si es juego/consigo
  → Si quoted == nil → se DESCARTA (línea 379-388)  ← PUNTO DE CAMBIO
  → Si quoted != nil → handleGame → GameUseCase.HandleGameCommand
    → Si tiene reply → MatchBetUseCase.ProcessBetMatch
      → resolveRootAndData → busca en matched_bet_logs
      → validateMatchRequest → verifica status y remaining
      → NLU parser → auto_bet
```

---

## Nuevo Ciclo de Estados

| Estado | Significado |
|--------|-------------|
| `pending` | Apuesta propia registrada, esperando que alguien la tome |
| `open` | En proceso de ser tomada (al menos un match parcial) |
| `processing` | Bloqueada temporalmente durante un match concurrente |
| `completed` | Completamente tomada (remaining = 0) |
| `closed` | Cancelada por el dueño con "se fue" |

---

## Cambios Propuestos

---

### Componente 1: `betting-common` — Entidad

#### [MODIFY] [matched_bet_log.go](file:///home/yesus/personal/caballos/betting-common/entities/sql/matched_bet_log.go)

Agregar constantes de estado para evitar strings mágicos:

```go
const (
    MatchedBetStatusPending    = "pending"
    MatchedBetStatusOpen       = "open"
    MatchedBetStatusProcessing = "processing"
    MatchedBetStatusCompleted  = "completed"
    MatchedBetStatusClosed     = "closed"
)
```

---

### Componente 1.5: `betting-common` — Reglas de Comandos en DB

#### [NEW] Migración SQL (`sql/whatsapp/20260320_create_command_rules_table.sql`)

```sql
CREATE TABLE whatsapp.command_rules (
    id SERIAL PRIMARY KEY,
    command_type VARCHAR(50) NOT NULL, -- Ej: 'se_fue'
    aliases JSONB NOT NULL,           -- Ej: ["se fue", "se fue todo", "sefue"]
    min_similarity FLOAT DEFAULT 0.8, -- Umbral de similitud (0.0 a 1.0)
    group_id INTEGER REFERENCES security.group(id), -- NULL = Global
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- Índices para búsqueda rápida en dos fases
CREATE INDEX idx_command_rules_type_group ON whatsapp.command_rules(command_type, group_id) WHERE deleted_at IS NULL;
```

#### [NEW] [command_rule.go](file:///home/yesus/personal/caballos/betting-common/entities/sql/command_rule.go)

```go
package sql

import (
    "gorm.io/gorm"
)

type CommandRule struct {
    gorm.Model
    CommandType   string   `gorm:"column:command_type;not null" json:"command_type"`
    Aliases       JSONBMap `gorm:"column:aliases;type:jsonb;not null" json:"aliases"`
    MinSimilarity float64  `gorm:"column:min_similarity;default:0.8" json:"min_similarity"`
    GroupID       *uint    `gorm:"column:group_id" json:"group_id"` // Nullable: Global rule
    IsActive      bool     `gorm:"column:is_active;default:true" json:"is_active"`
}

func (CommandRule) TableName() string {
    return "whatsapp.command_rules"
}
```

---

### Componente 2: `consumer-evolution` — Interfaces de Dominio

#### [MODIFY] interfaz `MatchedBetRepository` en `src/core/domain/interface`

Agregar métodos para las operaciones de cierre:

```go
// Para "se fue" CON reply: valida dueño y cierra solo ese log
GetByOriginalIDAndTercio1JID(ctx context.Context, msgID string, tercio1JID string) (*MatchedBetLog, error)

// Para "se fue" SIN reply: cierra todos los logs pending/open del tercio en el grupo
CloseAllByTercio(ctx context.Context, groupID uint, tercio1JID string) (int64, error)
```

---

### Componente 3: `consumer-evolution` — GameUseCase

#### [MODIFY] [game.use-case.go](file:///home/yesus/personal/caballos/consumer-evolution/src/core/use-cases/game/game.use-case.go)

Agregar a la interfaz `GameUseCase`:

```go
type GameUseCase interface {
    HandleGameCommand(ctx context.Context, input GameInputDTO) error
    HandleOwnBetRegistration(ctx context.Context, input OwnBetInputDTO) error
    HandleSeFue(ctx context.Context, input SeFueInputDTO) error
}
```

**`HandleOwnBetRegistration`**:
- Obtiene grupo y carrera activa del cache
- Crea `MatchedBetLog` con `status = "pending"`, `remaining_amount = -1`, `Tercio1ID` del participante
- NO llama NLU ni auto_bet
- Persiste via `matchedBetRepo.Upsert()`

**`HandleSeFue`** (con reply — cierra una jugada):
- Recibe `QuotedMessageID` (el mensaje al que respondió)
- Busca el log con `GetByOriginalIDAndTercio1JID` para validar que quien lo cierra es el dueño
- Si coincide → `matchedBetRepo.SetStatus(msgID, "closed")`
- No publica ningún mensaje al grupo

**`HandleSeFueTodo`** (sin reply — cierra todas las jugadas del tercio):
- Recibe solo el `GroupID` y el `Participant` (quien envió el mensaje)
- Llama a `matchedBetRepo.CloseAllByTercio(ctx, groupID, tercio1JID)`
- Cierra todos los logs con status `pending` u `open` del tercio en ese grupo
- No publica ningún mensaje al grupo

---

### Componente 4: `consumer-evolution` — MatchBetUseCase

#### [MODIFY] [match_bet.use-case.go](file:///home/yesus/personal/caballos/consumer-evolution/src/core/use-cases/game/matcher/match_bet.use-case.go)

En `validateMatchRequest`, agregar verificación de `closed` **antes** de la verificación de `completed`:

```go
func (u *matchBetUseCase) validateMatchRequest(...) error {
    if aggLog != nil {
        // NUEVO: apuesta cancelada por el dueño
        if aggLog.Status == "closed" {
            return fmt.Errorf("NADA ⛔")
        }
        // Existente: completada
        if aggLog.Status == "completed" || (...) {
            return fmt.Errorf("NADA ⛔")
        }
        ...
    }
    return nil
}
```

> [!IMPORTANT]
> El estado `pending` se convierte automáticamente en `open`/`completed` cuando alguien responde, porque `initOrUpdateLog` establece `"processing"` y `finalizeMatch` actualiza al estado final. No se requiere lógica adicional de transición.

---

### Componente 5: `consumer-evolution` — Kafka Consumer

#### [MODIFY] [consumer.go](file:///home/yesus/personal/caballos/consumer-evolution/src/infrastructure/entry-points/kafka/inbound/consumer.go)

**Cambio 1**: En lugar de descartar silenciosamente cuando `quoted == nil && isDirectGame`, llamar a `handleOwnBet`:

```go
// Antes:
if quoted == nil && isDirectGame {
    return nil  // descarte silencioso
}

// Después:
if quoted == nil && isDirectGame {
    return c.handleOwnBet(ctx, remoteJid, participant, text, messageId, msg)
}
```

**Cambio 2**: Agregar detección de "se fue" antes del dispatch de handlers. La presencia o ausencia de `quoted` determina el alcance:

```go
if isSeFueCommand(text) {
    if quoted != nil {
        // Cierra solo la jugada citada
        return c.handleSeFue(ctx, remoteJid, participant, text, messageId, quoted, msg)
    }
    // Sin reply: cierra TODAS las jugadas del tercio en el grupo
    return c.handleSeFueTodo(ctx, remoteJid, participant, text, messageId, msg)
}
```

**Nuevas funciones helper**:

```go
func (c *KafkaConsumer) isSeFueCommand(ctx context.Context, text string, groupID uint) bool {
    t := strings.ToLower(strings.TrimSpace(text))

    // FASE 1: Búsqueda de regla específica por grupo
    rule, err := c.commandRuleRepo.GetActiveRule(ctx, "se_fue", &groupID)
    if err != nil || rule == nil {
        // FASE 2: Fallback a regla global (group_id IS NULL)
        rule, _ = c.commandRuleRepo.GetActiveRule(ctx, "se_fue", nil)
    }

    if rule == nil {
        return false
    }

    // Convertir Aliases de JSONBMap a []string (asumiendo que es un array JSON)
    // ... lógica de desempaquetado de rule.Aliases ...

    for _, target := range aliases {
        distance := fuzzy.LevenshteinDistance(t, target)
        maxLen := len(t)
        if len(target) > maxLen { maxLen = len(target) }
        if maxLen == 0 { continue }

        similarity := 1.0 - (float64(distance) / float64(maxLen))
        if similarity >= rule.MinSimilarity {
            return true
        }
    }
    return false
}

#### Estrategia de Caché (Redis)

Para evitar latencia de base de datos en cada mensaje:
1. **Calentamiento**: Al arrancar el consumer, se cargan todas las reglas activas en Redis.
2. **Estructura de llaves**:
   - `whatsapp:cmd_rule:global:{type}`
   - `whatsapp:cmd_rule:group:{groupID}:{type}`
3. **Invalidación**: Implementar un pequeño TTL (ej: 5-10 min) o un webhook de invalidación si la configuración cambia desde un panel administrativo.
func (c *KafkaConsumer) handleOwnBet(ctx context.Context, remoteJid, participant, text, messageId string, msg WhatsAppMessage) error {
    return c.gameUseCase.HandleOwnBetRegistration(ctx, game.OwnBetInputDTO{
        RemoteJid:   remoteJid,
        Participant: participant,
        RawText:     text,
        MessageID:   messageId,
        GroupID:     msg.GroupID,
    })
}

func (c *KafkaConsumer) handleSeFue(ctx context.Context, remoteJid, participant, text, messageId string, quoted *WhatsAppQuotedMsg, msg WhatsAppMessage) error {
    return c.gameUseCase.HandleSeFue(ctx, game.SeFueInputDTO{
        RemoteJid:       remoteJid,
        Participant:     participant,
        QuotedMessageID: quoted.MessageID,
        MessageID:       messageId,
        GroupID:         msg.GroupID,
    })
}
```

---

## Flujo Visual Completo

```
CASO 1 — Apuesta Propia (sin reply)
────────────────────────────────────
Tercio A: "juego 3-5 15.000"
  → isDirectGame && quoted == nil
  → handleOwnBet() → HandleOwnBetRegistration()
  → MatchedBetLog { status: "pending", remaining: -1, tercio1: A }
  [Sin NLU, sin auto_bet]

CASO 2 — Alguien responde la apuesta
──────────────────────────────────────
Tercio B: "j 10.000" (reply a Tercio A)
  → ProcessBetMatch() → encuentra log "pending"
  → validateMatchRequest: NOT closed, NOT completed → OK
  → NLU parsea QuotedBody por primera vez
  → auto_bet ejecutado
  → pending → processing → open/completed

CASO 3a — Dueño cancela con "se fue" + reply (solo esa jugada)
────────────────────────────────────────────────────────────────
Tercio A: "se fue" (reply a su propio juego)
  → isSeFueCommand && quoted != nil
  → HandleSeFue() → valida Tercio A == Tercio1 del log citado
  → SetStatus(msgID, "closed") → sin respuesta al grupo

CASO 3b — Dueño cancela con "se fue" sin reply (todas sus jugadas)
────────────────────────────────────────────────────────────────────
Tercio A: "se fue" o "se fue todo" (sin reply)
  → isSeFueCommand && quoted == nil
  → HandleSeFueTodo() → busca TODOS los logs del tercio en el grupo
    con status IN ("pending", "open")
  → SetStatus para cada uno → "closed" → sin respuesta al grupo

CASO 4 — Intento de tomar apuesta cerrada
───────────────────────────────────────────
Tercio C: "j 5.000" (reply post "se fue")
  → validateMatchRequest: status == "closed" → "NADA ⛔"
```

---

## Consideraciones de Diseño

> [!NOTE]
> **¿Por qué NO parsear NLU al registrar la apuesta propia?** El NLU necesita contexto de ambos mensajes (original + respuesta). Al registrar sin reply, no hay taker aún. El primer taker activa el NLU de forma natural con el QuotedBody del mensaje original.

> [!WARNING]
> **"Se fue" sin reply** actúa como un cierre masivo: afecta **todos** los logs `pending`/`open` del tercio en el grupo. El repo necesita un método nuevo `CloseAllByTercio(ctx, groupID, tercio1JID)` que haga un `UPDATE ... WHERE group_id = ? AND tercio1_jid = ? AND status IN ('pending', 'open')`.

> [!IMPORTANT]
> **Anti-self-close**: En `HandleSeFue` se debe verificar que quien cierra sea el `Tercio1` del log. Si otra persona responde con "se fue" (no el dueño), se ignora silenciosamente.

> [!TIP]
> **Robustez de Comandos (Similitud %)**: Se utiliza una lista de comandos permitidos y un cálculo de **porcentaje de similitud** basado en la distancia de Levenshtein.
> - **Umbral recomendado**: 80% de similitud.
> - **Fórmula**: `1.0 - (distancia / max(longitud_input, longitud_target))`.
> - **Ventaja**: Permite escalar fácilmente añadiendo nuevos alias o comandos a la lista sin modificar la lógica principal.

---

## Plan de Verificación

### Tests Existentes

```bash
cd /home/yesus/personal/caballos/consumer-evolution && go test ./...
```

- `parser/verify_parser_test.go` — cubre parseo NLU (no afectado)
- `parser/verify_pollution_test.go` — cubre contaminación de parseo (no afectado)
- `inbound/retry_consumer_test.go` — cubre retry topic (puede verse afectado si hay cambio en dispatch)

### Nuevos Tests a Escribir

**`game/game_own_bet_test.go`**: `TestHandleOwnBetRegistration` y `TestHandleSeFue`
- Verifica que `HandleOwnBetRegistration` crea un log con status `pending`
- Verifica que `HandleSeFue` cambia status a `closed` solo si el Participant es Tercio1
- Verifica que `HandleSeFue` ignora si el Participant NO es Tercio1

**`matcher/match_validate_test.go`**: `TestValidateMatchRequest_ClosedStatus`
- Verifica que `validateMatchRequest` retorna `NADA ⛔` cuando status es `closed`

### Verificación Manual

1. Enviar "juego 3-5 15.000" sin reply → DB: `matched_bet_logs` con `status = pending`
2. Responder ese mensaje con "j 10.000" → verificar que `auto_bet` se ejecuta y status pasa a `open`/`completed`
3. Enviar "juego 3-5 15.000" sin reply → responderlo el mismo usuario con "se fue" → DB: `status = closed`
4. Intentar responder mensajes con `status = closed` → verificar respuesta `NADA ⛔`
