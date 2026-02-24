-- Función: gaming.fn_saldo_tercio_range
-- Calcula el balance neto y las comisiones de los tercios de un grupo en un rango de timestamps exacto.
--
-- Uso al CERRAR jornada:
--   p_desde = snapshot_at de la jornada anterior cerrada (o -infinity para la primera jornada)
--   p_hasta = NOW() en el momento del cierre
--   → balance_delta + snapshot_previo.Balance = nuevo balance a guardar en WorkingDaySnapshot
--
-- Uso para SALDO EN TIEMPO REAL:
--   p_desde = snapshot_at del último cierre (NO open_date de la jornada activa)
--   p_hasta = NOW()
--   La vista v_saldo_tercio_live usa COALESCE(us.snapshot_at, '-infinity') automáticamente.
--
-- IMPORTANTE: siempre TIMESTAMPTZ (nunca DATE) para evitar problemas al cruzar medianoche.

CREATE OR REPLACE FUNCTION gaming.fn_saldo_tercio_range(
    p_group_id   INT,
    p_desde      TIMESTAMPTZ,   -- open_date de la jornada activa (o snapshot_at de la jornada anterior)
    p_hasta      TIMESTAMPTZ    -- close_date al cerrar, o NOW() para saldo en vivo
)
RETURNS TABLE (
    id_tercio          BIGINT,
    name               TEXT,
    balance_delta      DECIMAL(18,8),  -- movimiento neto en el rango (puede ser negativo)
    ganancia_comission DECIMAL(18,8)   -- suma de commission del rango
) AS $$
BEGIN
    RETURN QUERY
    SELECT
        t.id                                                        AS id_tercio,
        t.name::TEXT,
        COALESCE(
            SUM(
                CASE
                    WHEN tt.type_calc = 'rest' THEN tr.amount * -1
                    ELSE tr.amount
                END
            ), 0
        )::DECIMAL(18,8)                                            AS balance_delta,
        COALESCE(SUM(tr.commission), 0)::DECIMAL(18,8)             AS ganancia_comission
    FROM gaming.tercios t
    LEFT JOIN transactions.transactions tr
           ON tr.id_tercios  = t.id
          AND tr.rever        = false
          AND tr.created_at  >= p_desde   -- límite inferior exacto (TIMESTAMPTZ)
          AND tr.created_at  <  p_hasta   -- límite superior exacto (TIMESTAMPTZ)
    LEFT JOIN transactions.type_transaction tt
           ON tt.id = tr.id_type_transaction
    WHERE t.id_group    = p_group_id
      AND t.deleted_at IS NULL
    GROUP BY t.id, t.name;
END;
$$ LANGUAGE plpgsql STABLE;

-- Índices requeridos (verificar existencia antes de ejecutar)
CREATE INDEX IF NOT EXISTS idx_transactions_id_tercios ON transactions.transactions(id_tercios);
CREATE INDEX IF NOT EXISTS idx_transactions_created_at ON transactions.transactions(created_at);
CREATE INDEX IF NOT EXISTS idx_transactions_rever      ON transactions.transactions(rever);
