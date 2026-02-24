-- Vista: reports.saldo_tercios_user_group
DROP VIEW IF EXISTS reports.saldo_tercios_user_group CASCADE;
CREATE OR REPLACE VIEW reports.saldo_tercios_user_group AS
SELECT COALESCE(sum(
    CASE
        WHEN tt.type_calc = 'rest' THEN (tr.amount::decimal(18,8) * -1)
        ELSE tr.amount::decimal(18,8)
    END), 0::decimal(18,8))::decimal(18,8) AS total,
    t2.id AS id_tercios,
    t2.name,
    t2.id_group,
    t2.status,
    ut.id AS id_user,
    tr.currency_id AS id_currency
FROM gaming.tercios t2
LEFT JOIN transactions.transactions tr ON t2.id = tr.id_tercios AND tr.rever = false
LEFT JOIN transactions.type_transaction tt ON tt.id = tr.id_type_transaction
LEFT JOIN security.user_tercio ut ON t2.id_user_tercio = ut.id
WHERE t2.deleted_at IS NULL
GROUP BY t2.id, t2.name, t2.id_group, t2.status, ut.id, tr.currency_id;