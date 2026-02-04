-- Script para autogenerar el campo code en la tabla security.group
-- 1. Crear función para generar códigos aleatorios de 6 caracteres
CREATE OR REPLACE FUNCTION gen_group_code()
RETURNS text AS $$
DECLARE
    chars TEXT := 'ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789';
    result TEXT;
    existe INT;
BEGIN
    LOOP
        result := '';
        FOR i IN 1..6 LOOP
            result := result || substr(chars, floor(random()*length(chars)+1)::int, 1);
        END LOOP;
        SELECT COUNT(1) INTO existe FROM security."group" WHERE code = result;
        IF existe = 0 THEN
            RETURN result;
        END IF;
        -- Si existe, repite el loop
    END LOOP;
END;
$$ LANGUAGE plpgsql;

-- 2. Agregar columna code si no existe
ALTER TABLE security."group"
ADD COLUMN IF NOT EXISTS code varchar(20) UNIQUE;

-- 3. Asignar default usando la función
ALTER TABLE security."group"
ALTER COLUMN code SET DEFAULT gen_group_code();

-- 4. Actualizar los registros existentes que no tengan code
UPDATE security."group"
SET code = gen_group_code()
WHERE code IS NULL OR code = '';

-- 5. Asegurar restricción UNIQUE solo si no existe
DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM pg_constraint WHERE conname = 'group_code_unique'
    ) THEN
        ALTER TABLE security."group"
        ADD CONSTRAINT group_code_unique UNIQUE (code);
    END IF;
END;
$$;
