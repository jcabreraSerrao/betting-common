-- Migración local: añadir columna percentage con default 100 y constraint
-- Ejecutar solo en entornos locales

BEGIN;

ALTER TABLE config.group_country_min_bet_config
ADD COLUMN IF NOT EXISTS percentage integer NOT NULL DEFAULT 100 CHECK (percentage BETWEEN 1 AND 100);

COMMIT;

-- Rollback:
-- ALTER TABLE config.group_country_min_bet_config DROP COLUMN percentage;
