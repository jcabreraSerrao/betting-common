-- Semilla inicial para reglas de comandos de WhatsApp
-- Regla global para el comando de cancelación "se fue"

INSERT INTO whatsapp.command_rules (command_type, aliases, min_similarity, group_id, is_active, created_at, updated_at)
VALUES (
    'se_fue', 
    '{"se fue": true, "se fue todo": true, "sf": true, "sef": true, "sefu": true, "sefue": true}'::jsonb, 
    0.85, 
    NULL, 
    true, 
    NOW(), 
    NOW()
)
ON CONFLICT (command_type) WHERE (group_id IS NULL) 
DO UPDATE SET 
    aliases = EXCLUDED.aliases,
    min_similarity = EXCLUDED.min_similarity,
    updated_at = NOW();
