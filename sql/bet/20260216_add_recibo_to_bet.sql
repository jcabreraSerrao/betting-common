-- Migración: añadir columna recibo a gaming.bet
-- Fecha: 2026-02-16

BEGIN;

ALTER TABLE gaming.bet
ADD COLUMN IF NOT EXISTS recibo boolean DEFAULT NULL;

COMMENT ON COLUMN gaming.bet.recibo IS 'Invierte la lógica de ganancia/pérdida: si true, gana todo cuando gana y pierde solo el porcentaje cuando pierde';

COMMIT;

-- Rollback:
-- ALTER TABLE gaming.bet DROP COLUMN IF EXISTS recibo;
