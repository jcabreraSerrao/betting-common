-- Script para insertar 100,000 registros simulados en transactions.transactions
-- Basado en el JSON de ejemplo, con amount variable (10-1000) e id_type_transaction alternando entre 3 y 4

INSERT INTO transactions.transactions (
    id_tercios,
    amount,
    commission,
    rever,
    rol,
    id_transaction,
    id_bet,
    id_type_transaction,
    id_race,
    date_transaction,
    id_remate,
    currency_id,
    exchange_rate,
    amount_usd,
    created_at,
    updated_at,
    deleted_at
)
SELECT
    CASE WHEN gs.i % 4 = 0 THEN 195 WHEN gs.i % 4 = 1 THEN 196 WHEN gs.i % 4 = 2 THEN 197 ELSE 198 END AS id_tercios,  -- Distribuye entre 195,196,197,198
    (random() * 990 + 10)::numeric(18,8) AS amount,  -- Variable: 10-1000
    0::numeric(18,8) AS commission,  -- Fijo
    false AS rever,  -- Fijo
    CASE WHEN random() > 0.3 THEN 'receptor' ELSE 'emisor' END AS rol,  -- Más 'receptor' (70%) que 'emisor' (30%)
    NULL AS id_transaction,  -- Fijo
    NULL AS id_bet,  -- Fijo
    CASE WHEN random() > 0.3 THEN 3 ELSE 4 END AS id_type_transaction,  -- Más 3 (70%) que 4 (30%)
    NULL AS id_race,  -- Fijo
    NOW() AS date_transaction,  -- Timestamp actual
    NULL AS id_remate,  -- Fijo
    2 AS currency_id,  -- Fijo
    NULL AS exchange_rate,  -- Fijo
    NULL AS amount_usd,  -- Fijo
    NOW() AS created_at,  -- Timestamp actual
    NOW() AS updated_at,  -- Timestamp actual
    NULL AS deleted_at  -- Fijo
FROM generate_series(1, 100000) AS gs(i);  -- 100,000 registros, IDs auto-generados

-- Nota: Ajusta el rango de IDs si ya existen registros. Ejecuta con precaución en producción.