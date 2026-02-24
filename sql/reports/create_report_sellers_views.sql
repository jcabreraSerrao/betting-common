-- Vista unificada para reportes de vendedores en el esquema 'reports'
-- Esta vista agrupa premios, reembolsos y pendientes para optimizar consultas
-- Se incluyen joins con tercios y user_group para obtener el nombre del usuario asociado
-- La currency se obtiene desde race -> hippodrome -> country -> currency
-- El amount/sales se maneja en una vista separada (sellers_sales)

DROP VIEW IF EXISTS reports.sellers_general_report CASCADE;
CREATE OR REPLACE VIEW reports.sellers_general_report AS
SELECT
    b.id_group,
    DATE(r.open_date AT TIME ZONE 'America/Caracas') AS date,
    curr.name AS currency_name,
    ug.name AS user_name,
    b.id_tercio,
    SUM(COALESCE(sred.payout, 0)) AS monto_prizes,
    SUM(COALESCE(sred.refund, 0)) AS monto_refunds,
    SUM(CASE WHEN b.status = 'pending' THEN COALESCE(sred.payout, 0) ELSE 0 END) AS monto_pending,
    SUM(CASE WHEN b.status = 'pending' THEN COALESCE(sred.refund, 0) ELSE 0 END) AS monto_pending_refunds
FROM gaming.bet b
INNER JOIN gaming.type_bet_group tbg ON tbg.id = b.id_type_bet_group
INNER JOIN gaming.type_bet tb ON tb.id = tbg.id_type_bet
INNER JOIN gaming.race r ON r.id = b.id_race
INNER JOIN config.hipodromos h ON h.id = r.hippodrome_group
INNER JOIN config.country co ON co.id = h.country_id
INNER JOIN config.currency curr ON curr.id = co.currency_id
LEFT JOIN gaming.settlement_race_estimate_details sred ON sred.bet_id = b.id
LEFT JOIN gaming.tercios ter ON ter.id = b.id_tercio
LEFT JOIN security.user_group ug ON ug.id_tercio = ter.id
WHERE tb.classification = 'winner_show_place' 
  AND b.cancel = false
GROUP BY b.id_group, DATE(r.open_date AT TIME ZONE 'America/Caracas'), curr.id, curr.name, ug.name, b.id_tercio;

-- Vista separada para ventas (sales/amount)
-- Filtra por tipo de transacción 'bet-winner-create', no revertidas
DROP VIEW IF EXISTS reports.sellers_sales CASCADE;
CREATE OR REPLACE VIEW reports.sellers_sales AS
SELECT
    b.id_group,
    DATE(r.open_date AT TIME ZONE 'America/Caracas') AS date,
    curr.name AS currency_name,
    ug.name AS user_name,
    b.id_tercio,
    SUM(t.amount) AS monto_total_apues
FROM gaming.bet b
INNER JOIN gaming.type_bet_group tbg ON tbg.id = b.id_type_bet_group
INNER JOIN gaming.type_bet tb ON tb.id = tbg.id_type_bet
INNER JOIN gaming.race r ON r.id = b.id_race
INNER JOIN config.hipodromos h ON h.id = r.hippodrome_group
INNER JOIN config.country co ON co.id = h.country_id
INNER JOIN config.currency curr ON curr.id = co.currency_id
INNER JOIN transactions.transactions t ON t.id_bet = b.id
INNER JOIN transactions.type_transaction tt ON tt.id = t.id_type_transaction
LEFT JOIN gaming.tercios ter ON ter.id = b.id_tercio
LEFT JOIN security.user_group ug ON ug.id_tercio = ter.id
WHERE tb.classification = 'winner_show_place' 
  AND b.cancel = false
  AND t.rever = false
  AND tt.slug = 'bet-winner-create'
GROUP BY b.id_group, DATE(r.open_date AT TIME ZONE 'America/Caracas'), curr.id, curr.name, ug.name, b.id_tercio;