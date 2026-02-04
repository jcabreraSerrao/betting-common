-- Seed para crear tercio 'vendedor' individual para cada user_group sin tercio asociado

-- Insertar tercio 'vendedor' para cada user_group que no tenga tercio asociado
INSERT INTO gaming.tercios (name, slug, status, amount, amount_lock, id_group, id_type_tercio)
SELECT 'vendedor', 'vendedor-' || ug.id, true, 0, 0, ug.id_group, 4
FROM security.user_group ug
WHERE ug.id_tercio IS NULL;

-- Actualizar los user_group para asociar el tercio correspondiente usando el slug único
UPDATE security.user_group
SET id_tercio = (
    SELECT t.id
    FROM gaming.tercios t
    WHERE t.slug = 'vendedor-' || security.user_group.id
    AND t.name = 'vendedor'
)
WHERE security.user_group.id_tercio IS NULL;