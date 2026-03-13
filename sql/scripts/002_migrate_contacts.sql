-- 002_migrate_contacts.sql
-- Migra los datos de phone_number e id_telegram desde gaming.tercios 
-- hacia la nueva tabla gaming.tercio_contacts.

-- Asegurar que la tabla existe (por si no se ha corrido AutoMigrate)
CREATE TABLE IF NOT EXISTS gaming.tercio_contacts (
    id BIGSERIAL PRIMARY KEY,
    id_tercio BIGINT NOT NULL REFERENCES gaming.tercios(id),
    contact_type TEXT NOT NULL,
    contact_value TEXT NOT NULL,
    is_primary BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS idx_tercio_contacts_id_tercio ON gaming.tercio_contacts(id_tercio);
CREATE INDEX IF NOT EXISTS idx_tercio_contacts_value ON gaming.tercio_contacts(contact_value);

-- Migrar Contactos (Phones y Telegrams) en una sola operación
INSERT INTO gaming.tercio_contacts (id_tercio, contact_type, contact_value, is_primary)
SELECT id, 'PHONE', phone_number, TRUE
FROM gaming.tercios
WHERE phone_number IS NOT NULL AND phone_number <> ''
UNION ALL
SELECT id, 'TELEGRAM', id_telegram, CASE WHEN (phone_number IS NULL OR phone_number = '') THEN TRUE ELSE FALSE END
FROM gaming.tercios
WHERE id_telegram IS NOT NULL AND id_telegram <> ''
ON CONFLICT DO NOTHING;

