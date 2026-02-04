-- Función para actualizar el saldo de Tercios basada en el tipo de transacción
-- Parámetros:
--   tercios_id: ID del registro en gaming.tercios
--   p_amount: Monto a sumar o restar (tipo DECIMAL(18,8))
--   type_transaction_id: ID del tipo de transacción en transactions.type_transaction
-- Retorna: Toda la fila actualizada de gaming.tercios

CREATE OR REPLACE FUNCTION transactions.update_tercios_balance(
    tercios_id BIGINT,
    p_amount DECIMAL(18,8),
    type_transaction_id BIGINT
)
RETURNS TABLE(
    id BIGINT,
    name TEXT,
    slug TEXT,
    status BOOLEAN,
    amount DECIMAL(18,8),
    amount_lock DECIMAL(18,8),
    id_group BIGINT,
    id_type_tercio BIGINT,
    id_user_tercio BIGINT,
    id_telegram TEXT,
    token TEXT,
    casa BOOLEAN,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
) AS $$
DECLARE
    calc_type TEXT;
    current_amount DECIMAL(18,8);
    new_amount DECIMAL(18,8);
BEGIN
    -- Obtener el tipo de cálculo del TypeTransaction
    SELECT type_calc INTO calc_type
    FROM transactions.type_transaction
    WHERE id = type_transaction_id;

    -- Verificar que el TypeTransaction existe
    IF NOT FOUND THEN
        RAISE EXCEPTION 'El tipo de transacción con ID % no existe.', type_transaction_id;
    END IF;

    -- Bloquear el registro Tercios para evitar concurrencia y obtener el amount actual
    SELECT amount INTO current_amount
    FROM gaming.tercios
    WHERE gaming.tercios.id = tercios_id
    FOR UPDATE;

    -- Verificar que el registro Tercios existe
    IF NOT FOUND THEN
        RAISE EXCEPTION 'El registro Tercios con ID % no existe.', tercios_id;
    END IF;

    -- Calcular el nuevo saldo basado en el tipo de cálculo
    IF calc_type = 'sum' THEN
        new_amount := current_amount + p_amount;
    ELSIF calc_type = 'rest' THEN
        new_amount := current_amount - p_amount;
    ELSE
        RAISE EXCEPTION 'Tipo de cálculo desconocido: %. Use "sum" o "rest".', calc_type;
    END IF;

    -- Actualizar el saldo
    UPDATE gaming.tercios
    SET amount = new_amount,
        updated_at = NOW()
    WHERE gaming.tercios.id = tercios_id;

    -- Devolver toda la fila actualizada
    RETURN QUERY SELECT 
        t.id, t.name, t.slug, t.status, t.amount, t.amount_lock, 
        t.id_group, t.id_type_tercio, t.id_user_tercio, t.id_telegram, 
        t.token, t.casa, t.created_at, t.updated_at, t.deleted_at
    FROM gaming.tercios t WHERE t.id = tercios_id;
END;
$$ LANGUAGE plpgsql;

-- Ejemplo de uso:
-- SELECT * FROM public.update_tercios_balance(1, 100.50, 1);  -- Retorna toda la fila del Tercios actualizado