-- Seed independiente para plataforma de pago "SIN BANCO"
-- Idempotente: crea el tipo 'banco' si falta y la plataforma si no existe

DO $$
DECLARE
    v_bank_type_id INTEGER;
BEGIN
    -- Asegurar que el tipo de plataforma 'banco' existe
    SELECT id INTO v_bank_type_id FROM payments.type_payment_platform WHERE name = 'banco' LIMIT 1;
    IF v_bank_type_id IS NULL THEN
        INSERT INTO payments.type_payment_platform (name, created_at, updated_at)
        VALUES ('banco', now(), now())
        RETURNING id INTO v_bank_type_id;
    END IF;

    -- Insertar plataforma "SIN BANCO" si no existe (usa code_ibp como clave única)
    IF NOT EXISTS (SELECT 1 FROM payments."PaymentPlatform" WHERE name = 'SIN BANCO') THEN
        INSERT INTO payments."PaymentPlatform" (code_ibp, name, id_type_payment_platform, created_at, updated_at)
        VALUES ('9999', 'SIN BANCO', v_bank_type_id, now(), now());
    END IF;
END $$;