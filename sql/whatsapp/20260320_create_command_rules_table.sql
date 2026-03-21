-- Creación de la tabla de reglas dinámicas para comandos de WhatsApp
-- Permite configurar aliases y similitud mínima de forma global o por grupo.

CREATE TABLE IF NOT EXISTS whatsapp.command_rules (
    id SERIAL PRIMARY KEY,
    command_type VARCHAR(50) NOT NULL, -- Ej: 'se_fue'
    aliases JSONB NOT NULL,           -- Ej: ["se fue", "se fue todo", "sefue"]
    min_similarity FLOAT DEFAULT 0.8, -- Umbral de similitud (0.0 a 1.0)
    group_id INTEGER REFERENCES security.group(id), -- NULL = Global, non-NULL = Override por grupo
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- Comentario para documentación
COMMENT ON COLUMN whatsapp.command_rules.command_type IS 'Tipo semántico del comando (ej: se_fue, juego, consigo)';
COMMENT ON COLUMN whatsapp.command_rules.aliases IS 'Lista de frases en formato JSON ["alias1", "alias2"]';
COMMENT ON COLUMN whatsapp.command_rules.group_id IS 'ID del grupo. Si es NULL, la regla aplica globalmente.';

-- Índices para búsqueda rápida y restricciones de unicidad
CREATE UNIQUE INDEX IF NOT EXISTS idx_unique_global_command ON whatsapp.command_rules (command_type) WHERE group_id IS NULL;
CREATE UNIQUE INDEX IF NOT EXISTS idx_unique_group_command ON whatsapp.command_rules (command_type, group_id) WHERE group_id IS NOT NULL;
CREATE INDEX IF NOT EXISTS idx_command_rules_type_group ON whatsapp.command_rules(command_type, group_id) WHERE deleted_at IS NULL;
