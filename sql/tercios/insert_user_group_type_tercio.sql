-- Insertar tipo de tercio 'user_group'
INSERT INTO gaming.type_tercio (name, slug) VALUES
('user_group', 'user_group')
ON CONFLICT (name) DO NOTHING;