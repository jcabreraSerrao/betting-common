-- Función para actualizar el saldo de todos los Tercios basado en la vista reports.saldo_tercio
-- Esta función itera sobre la vista SaldoUnificado y actualiza el amount de cada Tercios correspondiente

CREATE OR REPLACE FUNCTION transactions.update_all_tercios_balances()
RETURNS VOID AS $$
DECLARE
    rec RECORD;
BEGIN
    -- Iterar sobre cada registro de la vista saldo_tercio
    FOR rec IN SELECT id_tercios, total FROM reports.saldo_tercio LOOP
        -- Actualizar el saldo del Tercios correspondiente
        UPDATE gaming.tercios
        SET amount = rec.total,
            updated_at = NOW()
        WHERE id = rec.id_tercios;

        -- Opcional: Log o verificación si no se actualizó nada
        IF NOT FOUND THEN
            RAISE NOTICE 'No se encontró Tercios con ID % para actualizar.', rec.id_tercios;
        END IF;
    END LOOP;

    RAISE NOTICE 'Actualización de saldos completada para todos los Tercios.';
END;
$$ LANGUAGE plpgsql;

-- Ejemplo de uso:
-- SELECT update_all_tercios_balances();  -- Actualiza los saldos de todos los Tercios basados en la vista