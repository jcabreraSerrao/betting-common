-- Optimización sugerida para la vista saldo_tercio
-- Se agregan índices recomendados y una versión alternativa de la vista para mejorar performance en consultas frecuentes.

-- 1. Índices sugeridos (ejecutar una sola vez en la base de datos):
-- CREATE INDEX IF NOT EXISTS idx_transactions_id_tercios ON transactions.transactions(id_tercios);
-- CREATE INDEX IF NOT EXISTS idx_tercios_id ON gaming.tercios(id);
-- CREATE INDEX IF NOT EXISTS idx_tercios_id_user_tercio ON gaming.tercios(id_user_tercio);
-- CREATE INDEX IF NOT EXISTS idx_transactions_date_transaction ON transactions.transactions(date_transaction);
-- CREATE INDEX IF NOT EXISTS idx_transactions_rever ON transactions.transactions(rever);

-- 2. Vista optimizada (usa CTE para reducir joins y agrupa solo lo necesario):
CREATE OR REPLACE VIEW reports.saldo_tercio_optimization AS
WITH diarios AS (
  SELECT t.id AS id_tercios,
         COALESCE(SUM(CASE WHEN tt.type_calc = 'rest' THEN (tr.amount::decimal(18,8)) * -1 ELSE (tr.amount::decimal(18,8)) END), 0)::decimal(18,8) AS total
    FROM gaming.tercios t
    LEFT JOIN transactions.transactions tr ON t.id = tr.id_tercios AND tr.rever = false
    LEFT JOIN transactions.type_transaction tt ON tt.id = tr.id_type_transaction
   WHERE t.deleted_at IS NULL
   GROUP BY t.id
)
SELECT d.total,
       t2.id AS id_tercios,
       t2.name,
       t2.id_group,
       t2.status,
       ut.id AS id_user
  FROM diarios d
  JOIN gaming.tercios t2 ON t2.id = d.id_tercios
  LEFT JOIN security.user_tercio ut ON t2.id_user_tercio = ut.id;
