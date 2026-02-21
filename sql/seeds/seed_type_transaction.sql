-- Seed data for transactions.type_transaction table
-- This file inserts the initial type transaction records

INSERT INTO transactions.type_transaction (id, name, type_calc, classification, alias, slug, created_at, updated_at, deleted_at) VALUES
(1, 'tie', '', '', '', 'tie', NULL, NULL, NULL),
(2, 'win', 'sum', '', '', 'win', NULL, NULL, NULL),
(5, 'lose', 'rest', '', '', 'lose', NULL, NULL, NULL),
(6, 'lose_middle', 'rest', '', '', 'lose_middle', NULL, NULL, NULL),
(3, 'recarga', 'sum', '', '', 'recarga', NULL, NULL, NULL),
(4, 'retiro', 'rest', '', '', 'retiro', NULL, NULL, NULL),
(8, 'bloqueo', 'rest', '', '', 'bloqueo', NULL, NULL, NULL),
(7, 'win_middle', 'sum', '', '', 'win_middle', NULL, NULL, NULL),
(9, 'jugado', 'rest', NULL, NULL, 'jugado', NULL, NULL, NULL),
(10, 'sistema', 'rest', NULL, NULL, 'sistema', NULL, NULL, NULL),
(11, 'bet-winner-create', 'sum', NULL, NULL, 'bet-winner-create', NULL, NULL, NULL),
(12, 'bet-winner-pago', 'rest', NULL, NULL, 'bet-winner-pago', NULL, NULL, NULL),
(13, 'bet-winner-reembolso', 'rest', NULL, NULL, 'bet-winner-reembolso', NULL, NULL, NULL),
(14, 'initial', 'sum', NULL, NULL, 'initial-sum', NULL, NULL, NULL),
(15, 'initial', 'rest', NULL, NULL, 'initial-rest', NULL, NULL, NULL),
(16, 'transaction-delete-rest', 'rest', NULL, NULL, 'transaction-delete-rest', NULL, NULL, NULL),
(17, 'transaction-delete-sum', 'sum', NULL, NULL, 'transaction-delete-sum', NULL, NULL, NULL),
(18, 'transfer-banca-sum', 'sum', NULL, NULL, 'transfer-banca-sum', NULL, NULL, NULL),
(19, 'transfer-banca-rest', 'rest', NULL, NULL, 'transfer-banca-rest', NULL, NULL, NULL),
(20, 'reverso-self-sum', 'sum', NULL, NULL, 'reverso-self-sum', NULL, NULL, NULL),
(21, 'reverso-self-rest', 'rest', NULL, NULL, 'reverso-self-rest', NULL, NULL, NULL),
(22, 'reverso-interno-sum', 'sum', NULL, NULL, 'reverso-interno-sum', NULL, NULL, NULL),
(23, 'reverso-interno-rest', 'rest', NULL, NULL, 'reverso-interno-rest', NULL, NULL, NULL),
(24, 'reverso-externo-rest', 'rest', NULL, NULL, 'reverso-externo-rest', NULL, NULL, NULL),
(25, 'reverso-log', '', NULL, NULL, 'reverso-log', NULL, NULL, NULL)
ON CONFLICT (id) DO NOTHING;



