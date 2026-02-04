-- Seed data for gaming.type_tercio table
-- This file inserts the initial type tercio records

INSERT INTO gaming.type_tercio (id, name, slug) VALUES
(1, 'Casa', 'casa'),
(2, 'Banca', 'banca'),
(3, 'Usuario', 'usuario'),
(4, 'user_group', 'user_group')
ON CONFLICT (id) DO NOTHING;