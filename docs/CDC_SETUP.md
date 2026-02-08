# Configuración de CDC para Producción y QA

Este documento detalla los pasos necesarios para configurar correctamente el sistema de Change Data Capture (CDC) en los entornos de Producción y QA. El objetivo es utilizar una **única conexión de replicación** para evitar el agotamiento de `max_wal_senders` en PostgreSQL.

## 1. Limpieza de Slots "Zombies" (Importante)

Antes de aplicar la nueva configuración, es crucial eliminar los slots de replicación antiguos que puedan estar ocupando recursos innecesariamente. Ejecute los siguientes comandos SQL en la base de datos:

```sql
-- Eliminar slots antiguos individuales si existen
SELECT pg_drop_replication_slot('cdc_slot_bets') WHERE EXISTS (SELECT 1 FROM pg_replication_slots WHERE slot_name = 'cdc_slot_bets');
SELECT pg_drop_replication_slot('cdc_slot_race') WHERE EXISTS (SELECT 1 FROM pg_replication_slots WHERE slot_name = 'cdc_slot_race');
SELECT pg_drop_replication_slot('cdc_slot_board') WHERE EXISTS (SELECT 1 FROM pg_replication_slots WHERE slot_name = 'cdc_slot_board');
SELECT pg_drop_replication_slot('cdc_slot_process') WHERE EXISTS (SELECT 1 FROM pg_replication_slots WHERE slot_name = 'cdc_slot_process');
SELECT pg_drop_replication_slot('cdc_slot_retired') WHERE EXISTS (SELECT 1 FROM pg_replication_slots WHERE slot_name = 'cdc_slot_retired');
SELECT pg_drop_replication_slot('cdc_slot_activations') WHERE EXISTS (SELECT 1 FROM pg_replication_slots WHERE slot_name = 'cdc_slot_activations');
SELECT pg_drop_replication_slot('cdc_slot') WHERE EXISTS (SELECT 1 FROM pg_replication_slots WHERE slot_name = 'cdc_slot');
SELECT pg_drop_replication_slot('cdc_slot_main') WHERE EXISTS (SELECT 1 FROM pg_replication_slots WHERE slot_name = 'cdc_slot_main');
```

## 2. Configuración Unificada

El nuevo sistema utiliza una **única publicación** (`betting_cdc_pub`) que agrupa todas las tablas necesarias y un **único slot de replicación** (`betting_main_slot`).

### 2.1 Crear Publicación

Ejecute el siguiente SQL para crear la publicación que incluye todas las tablas del dominio de apuestas:

```sql
-- Eliminar publicación anterior si existe para asegurar estado limpio
DROP PUBLICATION IF EXISTS corridor_cdc_pub;
DROP PUBLICATION IF EXISTS betting_cdc_pub;

-- Crear publicación unificada
CREATE PUBLICATION betting_cdc_pub FOR TABLE 
    gaming.bet, 
    gaming.races_process_groups, 
    gaming.board_race_group, 
    gaming.retired_horse_group, 
    gaming.race, 
    gaming.group_race_activations;
```

**Nota:** Es fundamental que todas las tablas tengan `REPLICA IDENTITY FULL` configurado si se requiere recibir la fila completa en actualizaciones (UPDATE). Para `group_race_activations` esto es mandatorio:

```sql
ALTER TABLE gaming.group_race_activations REPLICA IDENTITY FULL;
ALTER TABLE gaming.bet REPLICA IDENTITY FULL;
```

### 2.2 Crear Slot de Replicación

Finalmente, cree el slot de replicación lógica que utilizará el servicio `cdc-horse`. Este slot debe crearse una única vez.

```sql
SELECT pg_create_logical_replication_slot('betting_main_slot', 'pgoutput');
```

## 3. Verificación

Para verificar que la configuración se aplicó correctamente:

**Verificar Publicación:**
```sql
SELECT * FROM pg_publication_tables WHERE pubname = 'betting_cdc_pub';
```
Debe listar las 6 tablas mencionadas anteriormente.

**Verificar Slot:**
```sql
SELECT slot_name, plugin, active FROM pg_replication_slots WHERE slot_name = 'betting_main_slot';
```

## 4. Configuración del Servicio (Go)

El servicio `cdc-horse` está configurado por defecto para conectarse a:
- **Publicación:** `betting_cdc_pub`
- **Slot:** `betting_main_slot`

No se requiere configuración adicional en el código si la base de datos está preparada como se indica arriba.
