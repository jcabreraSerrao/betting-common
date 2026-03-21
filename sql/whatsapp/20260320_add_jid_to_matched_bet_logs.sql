-- Migración para añadir tercio_1_jid a matched_bet_logs
-- Esto permite cierres masivos por JID sin joins.

ALTER TABLE whatsapp.matched_bet_logs 
ADD COLUMN IF NOT EXISTS tercio_1_jid VARCHAR(255);

CREATE INDEX IF NOT EXISTS idx_matched_bet_logs_tercio_jid ON whatsapp.matched_bet_logs(tercio_1_jid);
