-- Vista: reports.v_saldo_tercio_live
-- Reemplaza reports.saldo_tercio como vista principal de saldo.
-- Diferencias clave vs saldo_tercio:
--   1. El balance se calcula desde el ÚLTIMO CIERRE de jornada, no desde el inicio.
--      saldo_actual = snapshot.balance + delta(snapshot_at → NOW())
--   2. Incluye token, phone_number, id_telegram, id_user de gaming.tercios/user_tercio
--      para mantener compatibilidad con los repositorios existentes.
--   3. Si no hay jornada activa el grupo no aparece en la vista.
--   4. Si no hay snapshot previo (primera jornada): saldo_base = 0, delta desde -infinity.
--
-- En Go: db.Where("id_group = ?", groupID).Find(&[]entities.ViewSaldoTercioLive{})

DROP VIEW IF EXISTS reports.v_saldo_tercio_live;

CREATE OR REPLACE VIEW reports.v_saldo_tercio_live AS
WITH jornada_activa AS (
    SELECT id, id_group, open_date
    FROM gaming.working_days
    WHERE open = true
),
ultimo_snapshot AS (
    -- Snapshot más reciente por tercio, de la ÚLTIMA jornada CERRADA
    SELECT DISTINCT ON (wds.id_tercio, wds.id_group)
        wds.id_tercio,
        wds.id_group,
        wds.balance                    AS saldo_base,
        wds.ganancia_comission_parada  AS comission_base,
        wds.snapshot_at
    FROM gaming.working_day_snapshots wds
    INNER JOIN gaming.working_days wd ON wd.id = wds.id_working_day
    WHERE wd.open = false
    ORDER BY wds.id_tercio, wds.id_group, wds.snapshot_at DESC
),
delta_desde_cierre AS (
    -- Movimientos desde el snapshot_at del ÚLTIMO CIERRE hasta ahora (TIMESTAMPTZ exacto)
    SELECT
        t.id                                              AS id_tercio,
        t.id_group,
        COALESCE(
            SUM(
                CASE
                    WHEN tt.type_calc = 'rest' THEN tr.amount * -1
                    ELSE tr.amount
                END
            ), 0
        )::DECIMAL(18,8)                                  AS delta_balance,
        COALESCE(SUM(tr.commission), 0)::DECIMAL(18,8)   AS delta_comission
    FROM gaming.tercios t
    LEFT JOIN ultimo_snapshot us
           ON us.id_tercio = t.id
          AND us.id_group  = t.id_group
    LEFT JOIN transactions.transactions tr
           ON tr.id_tercios  = t.id
          AND tr.rever        = false
          -- Punto de inicio: snapshot_at del último cierre (NO open_date de jornada activa)
          AND tr.created_at  >= COALESCE(us.snapshot_at, '-infinity'::TIMESTAMPTZ)
          AND tr.created_at  <  NOW()
    LEFT JOIN transactions.type_transaction tt
           ON tt.id = tr.id_type_transaction
    WHERE t.deleted_at IS NULL
    GROUP BY t.id, t.id_group
)
SELECT
    -- Campos de saldo (nuevos)
    COALESCE(us.saldo_base,      0::DECIMAL(18,8)) +
    COALESCE(dc.delta_balance,   0::DECIMAL(18,8))                              AS saldo_actual,
    COALESCE(us.comission_base,  0::DECIMAL(18,8)) +
    COALESCE(dc.delta_comission, 0::DECIMAL(18,8))                              AS ganancia_comission_total,
    ja.open_date                                                                AS jornada_open_date,
    us.snapshot_at                                                              AS ultimo_snapshot_at,

    -- Campos de tercios (compatibilidad con ViewSaldoTercio antigua)
    t.id                                                                        AS id_tercio,
    t.name,
    t.id_group,
    t.status,
    COALESCE(t.token, '')                                                       AS token,
    COALESCE(t.phone_number, '')                                                AS phone_number,
    COALESCE(t.id_telegram, '')                                                 AS id_telegram,
    COALESCE(ut.id, 0)                                                          AS id_user
FROM gaming.tercios t
LEFT JOIN jornada_activa ja   ON ja.id_group   = t.id_group
LEFT JOIN  ultimo_snapshot us  ON us.id_tercio  = t.id AND us.id_group = t.id_group
LEFT JOIN  delta_desde_cierre dc ON dc.id_tercio = t.id AND dc.id_group = t.id_group
LEFT JOIN  security.user_tercio ut ON t.id_user_tercio = ut.id
WHERE t.deleted_at IS NULL;
