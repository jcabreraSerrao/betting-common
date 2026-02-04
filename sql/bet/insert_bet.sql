-- insert_bet.sql
-- Inserciones idempotentes para tipos de apuesta y su grupo
-- Actualizado para incluir Winner, Show, Place según documentación tipos_apuestas.md

-- type_bet: inserta solo si el id no existe
-- Usar esquema 'gaming' (las entidades definen gaming.type_bet)
INSERT INTO gaming.type_bet (name, type_calc, classification)
SELECT v.name, v.type_calc, v.classification
FROM (VALUES
  -- Winner/Show/Place betting types (según documentación tipos_apuestas.md)
  ( 'winner', NULL, 'winner_show_place'),
  ( 'show', NULL, 'winner_show_place'),
  ( 'place', NULL, 'winner_show_place'),
  -- Tipos de apuesta existentes (
  ( '1P', NULL,  'jugada'),
  ( '1PY2N', NULL,  'jugada'),
  ( '2N', NULL,  'jugada'),
  ( '2PY2N', NULL,  'jugada'),
  ( '2P', NULL,  'jugada'),
  ( '3N', NULL,  'jugada'),
  ( '3PY3N', NULL,  'jugada'),
  ( '3P', NULL,  'jugada'),
  ( '3PY4N', NULL,  'jugada'),
  ( '4N', NULL,  'jugada'),
  ( '4PY4N', NULL,  'jugada'),
  ( '4P', NULL,  'jugada'),
  ( '4PY5N', NULL,  'jugada'),
  ( '5N', NULL,  'jugada'),
  ( '5PY5N', NULL,  'jugada'),
  ( '5P', NULL,  'jugada'),
  ( '2PY3N', NULL, 'jugada'),
  ( 'PP', 'pp', 'jugada'),
  ( 'Directa', 'N/A', 'ganancia'),
  ( '@1.5', '1.5', 'ganancia'),
  ( '@2', '2', 'ganancia'),
  ( '@2.5', '2.5', 'ganancia'),
  ( '@3', '3', 'ganancia'),
  ( '@3.5', '3.5', 'ganancia'),
  ( '@4', '4', 'ganancia'),
  ( '@4.5', '4.5', 'ganancia'),
  ( '@5', '5', 'ganancia'),
  ( '@5.5', '5.5', 'ganancia'),
  ( '@6', '6', 'ganancia'),
  ( '@6.5', '6.5', 'ganancia'),
  ( '@7', '7', 'ganancia'),
  ( '@7.5', '7.5', 'ganancia'),
  ( '@8', '8', 'ganancia'),
  ( '@8.5', '8.5', 'ganancia'),
  ( '@9', '9', 'ganancia'),
  ( '@9.5', '9.5', 'ganancia')
) AS v( name, type_calc, classification)
WHERE NOT EXISTS (SELECT 1 FROM gaming.type_bet tb WHERE tb.name = v.name);

-- Usar esquema 'gaming' para la tabla de asociación
INSERT INTO gaming.type_bet_group (id_type_bet, alias, id_group)
SELECT v.id_type_bet, v.alias, v.id_group
FROM (VALUES
  -- Winner/Show/Place associations (según documentación tipos_apuestas.md)
  (1, 'Winner', 2),
  (2, 'Show', 2),
  (3, 'Place', 2),
  -- Asociaciones legacy existentes (IDs ajustados)
  (4, '1P',  2),
  (5, '1PY2N',  2),
  (6, '2N',  2),
  (7, '2PY2N',  2),
  (8, '2P',  2),
  (9, '3N',  2),
  (10, '3PY3N',  2),
  (11, '3P',  2),
  (12, '3PY4N',  2),
  (13, '4N',  2),
  (14, '4PY4N',  2),
  (15, '4P',  2),
  (16, '4PY5N',  2),
  (17, '5N',  2),
  (18, '5PY5N',  2),
  (19, '5P',  2),
  (20, 'Directa',  2),
  (21, '@2',  2),
  (22, '@3',  2),
  (23, '@4',  2),
  (24, '@5',  2),
  (25, '@6',  2),
  (26, '@7',  2),
  (27, '@8',  2),
  (28, '@9',  2),
  (29, 'PP',  2),
  (30, '2PY3N',  2)
) AS v(id_type_bet, alias, id_group)
WHERE NOT EXISTS (
  SELECT 1 FROM gaming.type_bet_group tbg WHERE tbg.id_type_bet = v.id_type_bet AND tbg.id_group = v.id_group  
) 
AND EXISTS (
  -- Solo insertar si el grupo referenciado ya existe en security.group
  SELECT 1 FROM security."group" g WHERE g.id = v.id_group
);