-- seed_countries_hipodromos.sql
-- Inserta países y hipódromos iniciales si no existen.
-- Usa códigos de moneda existentes en config.currency (BS, USD, COP, BRL, PEN, CLP).

BEGIN;

-- Insertar países (solo si no existen por code)
INSERT INTO config.country (name, code, currency_id, status, created_at, updated_at)
SELECT * FROM (
  VALUES
    ('Venezuela','VE',(SELECT id FROM config.currency WHERE code='BS'), true, '2025-06-27 20:11:36.976-04', '2025-06-27 20:11:36.976-04'),
    ('United States','US',(SELECT id FROM config.currency WHERE code='USD'), true, now(), now()),
    ('Colombia','CO',(SELECT id FROM config.currency WHERE code='COP'), true, now(), now()),
    ('Brazil','BR',(SELECT id FROM config.currency WHERE code='BRL'), true, now(), now()),
    ('Peru','PE',(SELECT id FROM config.currency WHERE code='PEN'), true, now(), now()),
    ('Chile','CL',(SELECT id FROM config.currency WHERE code='CLP'), true, now(), now())
) AS v(name, code, currency_id, status, created_at, updated_at)
WHERE NOT EXISTS (SELECT 1 FROM config.country c WHERE c.code = v.code);

-- Hipódromo de Venezuela (La Rinconada) — inserta solo si no existe
INSERT INTO config.hipodromos (name, slug, country_id, status, created_at, updated_at)
SELECT 'La Rinconada','la-rinconada', (SELECT id FROM config.country WHERE code='VE'), true, '2025-06-27 20:11:36.976-04', '2025-06-27 20:11:36.976-04'
WHERE NOT EXISTS (SELECT 1 FROM config.hipodromos h WHERE lower(h.name) = lower('La Rinconada'));

-- Hipódromos indicados como de United States (idempotente)
INSERT INTO config.hipodromos (name, slug, country_id, status, created_at, updated_at)
SELECT v.name, v.slug, (SELECT id FROM config.country WHERE code='US'), true, now(), now()
FROM (VALUES
  ('Camarero','camarero'),
  ('Belterra Park','belterra-park'),
  ('Canterbury Park','canterbury-park'),
  ('Colonial Downs','colonial-downs'),
  ('Del Mar','del-mar'),
  ('Delaware Park','delaware-park'),
  ('Delta Downs','delta-downs'),
  ('Hawthorne','hawthorne'),
  ('Parx Racing','parx-racing'),
  ('Penn National','penn-national'),
  ('Retama Park','retama-park'),
  ('Charles Town','charles-town'),
  ('Evangeline','evangeline'),
  ('Saratoga','saratoga'),
  ('Thistledown','thistledown'),
  ('Woodbine','woodbine'),
  ('Emerald Downs','emerald-downs'),
  ('Horsemen''s Park','horsemens-park'),
  ('Gulfstream Park','gulfstream-park'),
  ('Monmouth Park','monmouth-park'),
  ('Prairie Meadows','prairie-meadows'),
  ('North Dakota','north-dakota'),
  ('Wyoming Downs','wyoming-downs'),
  ('Ellis Park','ellis-park'),
  ('Fairmount Park','fairmount-park'),
  ('Mountaineer','mountaineer'),
  ('Finger Lakes','finger-lakes'),
  ('Louisiana Downs','louisiana-downs'),
  ('Horseshoe Indianapolis','horseshoe-indianapolis'),
  ('Century Mile','century-mile'),
  ('Grande Prairie','grande-prairie'),
  ('Hastings','hastings'),
  ('Oneida Fair','oneida-fair'),
  ('Presque Isle Downs','presque-isle-downs')
) AS v(name, slug)
WHERE NOT EXISTS (SELECT 1 FROM config.hipodromos h WHERE lower(h.name) = lower(v.name))
  AND (SELECT id FROM config.country WHERE code = 'US') IS NOT NULL;

COMMIT;
