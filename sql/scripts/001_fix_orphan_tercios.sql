-- 001_fix_orphan_tercios.sql
-- Este script crea un UserTercio para cada Tercio que no tiene uno asignado,
-- permitiendo que la migración de contactos sea consistente.

-- DO $$
-- DECLARE
--     r RECORD;
--     new_user_id INT;
-- BEGIN
--     FOR r IN 
--         SELECT id, name, phone_number, id_telegram 
--         FROM gaming.tercios 
--         WHERE id_user_tercio IS NULL 
--           AND (phone_number <> '' OR id_telegram <> '')
--     LOOP
--         -- Insertar un UserTercio base
--         INSERT INTO security.user_tercio (name, status, slug, created_at, updated_at)
--         VALUES (r.name, true, LOWER(REPLACE(r.name, ' ', '-')), NOW(), NOW())
--         RETURNING id INTO new_user_id;

--         -- Vincular el tercio al nuevo usuario
--         UPDATE gaming.tercios 
--         SET id_user_tercio = new_user_id 
--         WHERE id = r.id;
        
--         RAISE NOTICE 'Vinculado tercio % (%) a nuevo UserTercio %', r.id, r.name, new_user_id;
--     END LOOP;
-- END $$;
