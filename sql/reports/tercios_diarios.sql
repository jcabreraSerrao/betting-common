-- Vista: reports.tercios_diarios
CREATE OR REPLACE VIEW reports.tercios_diarios AS
SELECT date_trunc('day', (date(tr.date_transaction))::timestamp with time zone) AS dia,
    tr.id_tercios,
    t.name,
    t.id AS id_tercio,
    (sum(
        CASE
            WHEN (tt.type_calc = 'rest') THEN ((tr.amount)::decimal(18,8) * -1)
            ELSE (tr.amount)::decimal(18,8)
        END))::decimal(18,8) AS total,
    t.id_group,
    t.id_user_tercio
FROM ((transactions.transactions tr
     LEFT JOIN transactions.type_transaction tt ON (tt.id = tr.id_type_transaction))
     LEFT JOIN gaming.tercios t ON (t.id = tr.id_tercios))
WHERE (tr.rever = false)
GROUP BY date_trunc('day', (date(tr.date_transaction))::timestamp with time zone), tr.id_tercios, t.name, t.id, t.id_group;
