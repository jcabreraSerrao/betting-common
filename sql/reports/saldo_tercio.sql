-- Vista: reports.saldo_tercio
DROP VIEW IF EXISTS reports.saldo_tercio CASCADE;

CREATE VIEW reports.saldo_tercio AS
SELECT COALESCE(sum(
    CASE
        WHEN tt.type_calc = 'rest' THEN (tr.amount::decimal(18,8) * -1)
        ELSE tr.amount::decimal(18,8)
    END), 0::decimal(18,8))::decimal(18,8) AS total,
    t2.id AS id_tercios,
    t2.name,
    t2.id_group,
    t2.status,
    t2.token,
    t2.phone_number,
    t2.id_telegram,
    COALESCE(ut.id, 0) AS id_user
FROM gaming.tercios t2
LEFT JOIN transactions.transactions tr ON t2.id = tr.id_tercios AND tr.rever = false
LEFT JOIN transactions.type_transaction tt ON tt.id = tr.id_type_transaction
LEFT JOIN security.user_tercio ut ON t2.id_user_tercio = ut.id
WHERE t2.deleted_at IS NULL
GROUP BY t2.id, t2.name, t2.id_group, t2.status, t2.token, t2.phone_number, t2.id_telegram, ut.id;


