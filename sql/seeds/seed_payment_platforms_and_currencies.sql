-- Seed script: currencies, payment platform types, banks and exchanges
-- Idempotent script: checks existence before inserting

DO $$
DECLARE
    v_bank_type_id INTEGER;
    v_exchange_type_id INTEGER;
BEGIN
    -- 1) Insert core currencies if they don't exist (verifica cada código antes de insertar)

    IF NOT EXISTS (SELECT 1 FROM config.currency WHERE code = 'USD') THEN
        INSERT INTO config.currency (code, name, symbol, is_crypto, status, created_at, updated_at)
        VALUES ('USD', 'US Dollar', '$', false, true, now(), now());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM config.currency WHERE code = 'BS') THEN
        INSERT INTO config.currency (code, name, symbol, is_crypto, status, created_at, updated_at)
        VALUES ('BS', 'Venezuelan Bolívar', 'Bs.', false, true, now(), now());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM config.currency WHERE code = 'EUR') THEN
        INSERT INTO config.currency (code, name, symbol, is_crypto, status, created_at, updated_at)
        VALUES ('EUR', 'Euro', '€', false, true, now(), now());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM config.currency WHERE code = 'USDT') THEN
        INSERT INTO config.currency (code, name, symbol, is_crypto, status, created_at, updated_at)
        VALUES ('USDT', 'Tether USD', 'USDT', true, true, now(), now());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM config.currency WHERE code = 'BTC') THEN
        INSERT INTO config.currency (code, name, symbol, is_crypto, status, created_at, updated_at)
        VALUES ('BTC', 'Bitcoin', '₿', true, true, now(), now());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM config.currency WHERE code = 'ETH') THEN
        INSERT INTO config.currency (code, name, symbol, is_crypto, status, created_at, updated_at)
        VALUES ('ETH', 'Ethereum', 'Ξ', true, true, now(), now());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM config.currency WHERE code = 'USDC') THEN
        INSERT INTO config.currency (code, name, symbol, is_crypto, status, created_at, updated_at)
        VALUES ('USDC', 'USD Coin', 'USDC', true, true, now(), now());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM config.currency WHERE code = 'BNB') THEN
        INSERT INTO config.currency (code, name, symbol, is_crypto, status, created_at, updated_at)
        VALUES ('BNB', 'BNB', 'BNB', true, true, now(), now());
    END IF;

    -- Agregadas para cobertura de países
    IF NOT EXISTS (SELECT 1 FROM config.currency WHERE code = 'COP') THEN
        INSERT INTO config.currency (code, name, symbol, is_crypto, status, created_at, updated_at)
        VALUES ('COP', 'Colombian Peso', 'COP$', false, true, now(), now());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM config.currency WHERE code = 'BRL') THEN
        INSERT INTO config.currency (code, name, symbol, is_crypto, status, created_at, updated_at)
        VALUES ('BRL', 'Brazilian Real', 'R$', false, true, now(), now());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM config.currency WHERE code = 'PEN') THEN
        INSERT INTO config.currency (code, name, symbol, is_crypto, status, created_at, updated_at)
        VALUES ('PEN', 'Peruvian Sol', 'S/', false, true, now(), now());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM config.currency WHERE code = 'CLP') THEN
        INSERT INTO config.currency (code, name, symbol, is_crypto, status, created_at, updated_at)
        VALUES ('CLP', 'Chilean Peso', 'CLP$', false, true, now(), now());
    END IF;

    -- 2) Ensure TypePaymentPlatform 'banco' exists (get id or create)
    SELECT id INTO v_bank_type_id FROM payments.type_payment_platform WHERE name = 'banco' LIMIT 1;
    IF v_bank_type_id IS NULL THEN
        INSERT INTO payments.type_payment_platform (name, created_at, updated_at)
        VALUES ('banco', now(), now())
        RETURNING id INTO v_bank_type_id;
    END IF;

    -- 3) Ensure TypePaymentPlatform 'exchange' exists
    SELECT id INTO v_exchange_type_id FROM payments.type_payment_platform WHERE name = 'exchange' LIMIT 1;
    IF v_exchange_type_id IS NULL THEN
        INSERT INTO payments.type_payment_platform (name, created_at, updated_at)
        VALUES ('exchange', now(), now())
        RETURNING id INTO v_exchange_type_id;
    END IF;

    -- 4) Insert banks (use code_ibp as unique key check) into payments."PaymentPlatform"
    -- If your DB used a different table name/casing adjust the identifiers accordingly.

    IF NOT EXISTS (SELECT 1 FROM payments."PaymentPlatform" WHERE code_ibp = '0156') THEN
        INSERT INTO payments."PaymentPlatform" (code_ibp, name, id_type_payment_platform, created_at, updated_at)
        VALUES ('0156', '100%BANCO', v_bank_type_id, now(), now());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM payments."PaymentPlatform" WHERE code_ibp = '0196') THEN
        INSERT INTO payments."PaymentPlatform" (code_ibp, name, id_type_payment_platform, created_at, updated_at)
        VALUES ('0196', 'ABN AMRO BANK', v_bank_type_id, now(), now());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM payments."PaymentPlatform" WHERE code_ibp = '0172') THEN
        INSERT INTO payments."PaymentPlatform" (code_ibp, name, id_type_payment_platform, created_at, updated_at)
        VALUES ('0172', 'BANCAMIGA BANCO MICROFINANCIERO, C.A.', v_bank_type_id, now(), now());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM payments."PaymentPlatform" WHERE code_ibp = '0171') THEN
        INSERT INTO payments."PaymentPlatform" (code_ibp, name, id_type_payment_platform, created_at, updated_at)
        VALUES ('0171', 'BANCO ACTIVO BANCO COMERCIAL, C.A.', v_bank_type_id, now(), now());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM payments."PaymentPlatform" WHERE code_ibp = '0166') THEN
        INSERT INTO payments."PaymentPlatform" (code_ibp, name, id_type_payment_platform, created_at, updated_at)
        VALUES ('0166', 'BANCO AGRICOLA', v_bank_type_id, now(), now());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM payments."PaymentPlatform" WHERE code_ibp = '0175') THEN
        INSERT INTO payments."PaymentPlatform" (code_ibp, name, id_type_payment_platform, created_at, updated_at)
        VALUES ('0175', 'BANCO BICENTENARIO', v_bank_type_id, now(), now());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM payments."PaymentPlatform" WHERE code_ibp = '0128') THEN
        INSERT INTO payments."PaymentPlatform" (code_ibp, name, id_type_payment_platform, created_at, updated_at)
        VALUES ('0128', 'BANCO CARONI, C.A. BANCO UNIVERSAL', v_bank_type_id, now(), now());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM payments."PaymentPlatform" WHERE code_ibp = '0164') THEN
        INSERT INTO payments."PaymentPlatform" (code_ibp, name, id_type_payment_platform, created_at, updated_at)
        VALUES ('0164', 'BANCO DE DESARROLLO DEL MICROEMPRESARIO', v_bank_type_id, now(), now());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM payments."PaymentPlatform" WHERE code_ibp = '0102') THEN
        INSERT INTO payments."PaymentPlatform" (code_ibp, name, id_type_payment_platform, created_at, updated_at)
        VALUES ('0102', 'BANCO DE VENEZUELA S.A.I.C.A.', v_bank_type_id, now(), now());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM payments."PaymentPlatform" WHERE code_ibp = '0114') THEN
        INSERT INTO payments."PaymentPlatform" (code_ibp, name, id_type_payment_platform, created_at, updated_at)
        VALUES ('0114', 'BANCO DEL CARIBE C.A.', v_bank_type_id, now(), now());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM payments."PaymentPlatform" WHERE code_ibp = '0149') THEN
        INSERT INTO payments."PaymentPlatform" (code_ibp, name, id_type_payment_platform, created_at, updated_at)
        VALUES ('0149', 'BANCO DEL PUEBLO SOBERANO C.A.', v_bank_type_id, now(), now());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM payments."PaymentPlatform" WHERE code_ibp = '0163') THEN
        INSERT INTO payments."PaymentPlatform" (code_ibp, name, id_type_payment_platform, created_at, updated_at)
        VALUES ('0163', 'BANCO DEL TESORO', v_bank_type_id, now(), now());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM payments."PaymentPlatform" WHERE code_ibp = '0176') THEN
        INSERT INTO payments."PaymentPlatform" (code_ibp, name, id_type_payment_platform, created_at, updated_at)
        VALUES ('0176', 'BANCO ESPIRITO SANTO, S.A.', v_bank_type_id, now(), now());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM payments."PaymentPlatform" WHERE code_ibp = '0115') THEN
        INSERT INTO payments."PaymentPlatform" (code_ibp, name, id_type_payment_platform, created_at, updated_at)
        VALUES ('0115', 'BANCO EXTERIOR C.A.', v_bank_type_id, now(), now());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM payments."PaymentPlatform" WHERE code_ibp = '0003') THEN
        INSERT INTO payments."PaymentPlatform" (code_ibp, name, id_type_payment_platform, created_at, updated_at)
        VALUES ('0003', 'BANCO INDUSTRIAL DE VENEZUELA.', v_bank_type_id, now(), now());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM payments."PaymentPlatform" WHERE code_ibp = '0173') THEN
        INSERT INTO payments."PaymentPlatform" (code_ibp, name, id_type_payment_platform, created_at, updated_at)
        VALUES ('0173', 'BANCO INTERNACIONAL DE DESARROLLO, C.A.', v_bank_type_id, now(), now());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM payments."PaymentPlatform" WHERE code_ibp = '0105') THEN
        INSERT INTO payments."PaymentPlatform" (code_ibp, name, id_type_payment_platform, created_at, updated_at)
        VALUES ('0105', 'BANCO MERCANTIL C.A.', v_bank_type_id, now(), now());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM payments."PaymentPlatform" WHERE code_ibp = '0191') THEN
        INSERT INTO payments."PaymentPlatform" (code_ibp, name, id_type_payment_platform, created_at, updated_at)
        VALUES ('0191', 'BANCO NACIONAL DE CREDITO', v_bank_type_id, now(), now());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM payments."PaymentPlatform" WHERE code_ibp = '0116') THEN
        INSERT INTO payments."PaymentPlatform" (code_ibp, name, id_type_payment_platform, created_at, updated_at)
        VALUES ('0116', 'BANCO OCCIDENTAL DE DESCUENTO.', v_bank_type_id, now(), now());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM payments."PaymentPlatform" WHERE code_ibp = '0138') THEN
        INSERT INTO payments."PaymentPlatform" (code_ibp, name, id_type_payment_platform, created_at, updated_at)
        VALUES ('0138', 'BANCO PLAZA', v_bank_type_id, now(), now());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM payments."PaymentPlatform" WHERE code_ibp = '0108') THEN
        INSERT INTO payments."PaymentPlatform" (code_ibp, name, id_type_payment_platform, created_at, updated_at)
        VALUES ('0108', 'BANCO PROVINCIAL BBVA', v_bank_type_id, now(), now());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM payments."PaymentPlatform" WHERE code_ibp = '0104') THEN
        INSERT INTO payments."PaymentPlatform" (code_ibp, name, id_type_payment_platform, created_at, updated_at)
        VALUES ('0104', 'BANCO VENEZOLANO DE CREDITO S.A.', v_bank_type_id, now(), now());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM payments."PaymentPlatform" WHERE code_ibp = '0168') THEN
        INSERT INTO payments."PaymentPlatform" (code_ibp, name, id_type_payment_platform, created_at, updated_at)
        VALUES ('0168', 'BANCRECER S.A. BANCO DE DESARROLLO', v_bank_type_id, now(), now());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM payments."PaymentPlatform" WHERE code_ibp = '0134') THEN
        INSERT INTO payments."PaymentPlatform" (code_ibp, name, id_type_payment_platform, created_at, updated_at)
        VALUES ('0134', 'BANESCO BANCO UNIVERSAL', v_bank_type_id, now(), now());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM payments."PaymentPlatform" WHERE code_ibp = '0177') THEN
        INSERT INTO payments."PaymentPlatform" (code_ibp, name, id_type_payment_platform, created_at, updated_at)
        VALUES ('0177', 'BANFANB', v_bank_type_id, now(), now());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM payments."PaymentPlatform" WHERE code_ibp = '0146') THEN
        INSERT INTO payments."PaymentPlatform" (code_ibp, name, id_type_payment_platform, created_at, updated_at)
        VALUES ('0146', 'BANGENTE', v_bank_type_id, now(), now());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM payments."PaymentPlatform" WHERE code_ibp = '0174') THEN
        INSERT INTO payments."PaymentPlatform" (code_ibp, name, id_type_payment_platform, created_at, updated_at)
        VALUES ('0174', 'BANPLUS BANCO COMERCIAL C.A', v_bank_type_id, now(), now());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM payments."PaymentPlatform" WHERE code_ibp = '0190') THEN
        INSERT INTO payments."PaymentPlatform" (code_ibp, name, id_type_payment_platform, created_at, updated_at)
        VALUES ('0190', 'CITIBANK.', v_bank_type_id, now(), now());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM payments."PaymentPlatform" WHERE code_ibp = '0121') THEN
        INSERT INTO payments."PaymentPlatform" (code_ibp, name, id_type_payment_platform, created_at, updated_at)
        VALUES ('0121', 'CORP BANCA.', v_bank_type_id, now(), now());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM payments."PaymentPlatform" WHERE code_ibp = '0157') THEN
        INSERT INTO payments."PaymentPlatform" (code_ibp, name, id_type_payment_platform, created_at, updated_at)
        VALUES ('0157', 'DELSUR BANCO UNIVERSAL', v_bank_type_id, now(), now());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM payments."PaymentPlatform" WHERE code_ibp = '0151') THEN
        INSERT INTO payments."PaymentPlatform" (code_ibp, name, id_type_payment_platform, created_at, updated_at)
        VALUES ('0151', 'FONDO COMUN', v_bank_type_id, now(), now());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM payments."PaymentPlatform" WHERE code_ibp = '0601') THEN
        INSERT INTO payments."PaymentPlatform" (code_ibp, name, id_type_payment_platform, created_at, updated_at)
        VALUES ('0601', 'INSTITUTO MUNICIPAL DE CRÉDITO POPULAR', v_bank_type_id, now(), now());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM payments."PaymentPlatform" WHERE code_ibp = '0169') THEN
        INSERT INTO payments."PaymentPlatform" (code_ibp, name, id_type_payment_platform, created_at, updated_at)
        VALUES ('0169', 'MIBANCO BANCO DE DESARROLLO, C.A.', v_bank_type_id, now(), now());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM payments."PaymentPlatform" WHERE code_ibp = '0137') THEN
        INSERT INTO payments."PaymentPlatform" (code_ibp, name, id_type_payment_platform, created_at, updated_at)
        VALUES ('0137', 'SOFITASA', v_bank_type_id, now(), now());
    END IF;

    -- Add Zinly bank (no code provided) — check existence by name and insert with NULL code_ibp if missing
    IF NOT EXISTS (SELECT 1 FROM payments."PaymentPlatform" WHERE name = 'ZINLY') THEN
        INSERT INTO payments."PaymentPlatform" (code_ibp, name, id_type_payment_platform, created_at, updated_at)
        VALUES (NULL, 'ZINLY', v_bank_type_id, now(), now());
    END IF;

    -- 5) Insert exchanges (Binance, Bybit, LocalBitcoins) as payment platforms of type 'exchange'
    IF NOT EXISTS (SELECT 1 FROM payments."PaymentPlatform" WHERE code_ibp = 'EX01') THEN
        INSERT INTO payments."PaymentPlatform" (code_ibp, name, id_type_payment_platform, created_at, updated_at)
        VALUES ('EX01', 'BINANCE', v_exchange_type_id, now(), now());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM payments."PaymentPlatform" WHERE code_ibp = 'EX02') THEN
        INSERT INTO payments."PaymentPlatform" (code_ibp, name, id_type_payment_platform, created_at, updated_at)
        VALUES ('EX02', 'BYBIT', v_exchange_type_id, now(), now());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM payments."PaymentPlatform" WHERE code_ibp = 'EX03') THEN
        INSERT INTO payments."PaymentPlatform" (code_ibp, name, id_type_payment_platform, created_at, updated_at)
        VALUES ('EX03', 'AIRTM', v_exchange_type_id, now(), now());
    END IF;

END
$$ LANGUAGE plpgsql;

-- End of seed script
