-- =====================================================
-- Script: Configuraciones por defecto para Módulo Polla
-- Descripción: Valores preconfigurados para auto-rellenar configuraciones de polla por grupo
-- Tabla: config.config (clave-valor por grupo)
-- Moneda: Moneda principal del sistema
-- =====================================================

-- NOTA: Este script debe ejecutarse DESPUÉS de crear los grupos
-- Las configuraciones se insertan SOLO si no existen previamente

BEGIN;

-- 1. Monto de entrada a la polla (entry_fee_polla)
-- Valor por defecto: $1.00 en moneda principal
INSERT INTO config.config (id_group, key, value, created_at, updated_at)
SELECT DISTINCT g.id, 'entry_fee_polla', '1.00', NOW(), NOW()
FROM security.groups g
WHERE NOT EXISTS (
    SELECT 1 FROM config.config c 
    WHERE c.id_group = g.id AND c.key = 'entry_fee_polla'
);

-- 2. Comisión de la casa (commission_polla)
-- Valor por defecto: 15%
INSERT INTO config.config (id_group, key, value, created_at, updated_at)
SELECT DISTINCT g.id, 'commission_polla', '15', NOW(), NOW()
FROM security.groups g
WHERE NOT EXISTS (
    SELECT 1 FROM config.config c 
    WHERE c.id_group = g.id AND c.key = 'commission_polla'
);

-- 3. Porcentaje del premio para 1er lugar (prize_first_percent_polla)
-- Valor por defecto: 47% del pool distribuible
INSERT INTO config.config (id_group, key, value, created_at, updated_at)
SELECT DISTINCT g.id, 'prize_first_percent_polla', '47', NOW(), NOW()
FROM security.groups g
WHERE NOT EXISTS (
    SELECT 1 FROM config.config c 
    WHERE c.id_group = g.id AND c.key = 'prize_first_percent_polla'
);

-- 4. Porcentaje del premio para 2do lugar (prize_second_percent_polla)
-- Valor por defecto: 16% del pool distribuible
INSERT INTO config.config (id_group, key, value, created_at, updated_at)
SELECT DISTINCT g.id, 'prize_second_percent_polla', '16', NOW(), NOW()
FROM security.groups g
WHERE NOT EXISTS (
    SELECT 1 FROM config.config c 
    WHERE c.id_group = g.id AND c.key = 'prize_second_percent_polla'
);

-- 5. Porcentaje del premio para 3er lugar (prize_third_percent_polla)
-- Valor por defecto: 7% del pool distribuible
-- NOTA: 47% + 16% + 7% = 70% distribuido, 30% queda como reserva/fondo
INSERT INTO config.config (id_group, key, value, created_at, updated_at)
SELECT DISTINCT g.id, 'prize_third_percent_polla', '7', NOW(), NOW()
FROM security.groups g
WHERE NOT EXISTS (
    SELECT 1 FROM config.config c 
    WHERE c.id_group = g.id AND c.key = 'prize_third_percent_polla'
);

-- 6. Jackpot para 30 puntos perfectos (jackpot_polla)
-- Valor por defecto: $10,000.00 en moneda principal
-- IMPORTANTE: Este se paga en la MONEDA PRINCIPAL del sistema
INSERT INTO config.config (id_group, key, value, created_at, updated_at)
SELECT DISTINCT g.id, 'jackpot_polla', '10000.00', NOW(), NOW()
FROM security.groups g
WHERE NOT EXISTS (
    SELECT 1 FROM config.config c 
    WHERE c.id_group = g.id AND c.key = 'jackpot_polla'
);

COMMIT;

-- =====================================================
-- Verificación de configuraciones insertadas
-- =====================================================
SELECT 
    g.name AS grupo,
    c.key AS configuracion,
    c.value AS valor
FROM config.config c
INNER JOIN security.groups g ON g.id = c.id_group
WHERE c.key IN (
    'entry_fee_polla',
    'commission_polla',
    'prize_first_percent_polla',
    'prize_second_percent_polla',
    'prize_third_percent_polla',
    'jackpot_polla'
)
ORDER BY g.name, c.key;
