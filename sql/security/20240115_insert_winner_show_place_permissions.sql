-- Inserción de permisos para el módulo Winner/Show/Place
-- Fecha: 2024-01-15
-- alias: Inserta las claves de permisos necesarias para la gestión de configuraciones de apuestas mínimas y dividendos

-- Insertar permisos para configuración de apuestas mínimas
INSERT INTO security.permissions ("key", alias, "group", created_at)
SELECT v.key, v.alias, v.group_name, v.created_at
FROM (VALUES
  ('minbet.config.create', 'Crear configuraciones de min-bet por country+currency para grupos', 'config-winner', now()),
  ('minbet.config.create.group', 'Crear configuraciones de min-bet con alcance explícito por group (scoped)', 'config-winner', now()),
  ('minbet.config.update', 'Actualizar configuraciones de min-bet por country+currency para grupos', 'config-winner', now()),
  ('minbet.config.view', 'Consultar configuraciones de min-bet por country+currency para grupos', 'config-winner', now()),
  ('minbet.config.delete', 'Eliminar configuraciones de min-bet por country+currency para grupos', 'config-winner', now())
) AS v(key, alias, group_name, created_at)
WHERE NOT EXISTS (
  SELECT 1 FROM security.permissions p WHERE p."key" = v.key
);

-- Insertar permisos para gestión de dividendos
INSERT INTO security.permissions ("key", alias, "group", created_at)
SELECT v.key, v.alias, v.group_name, v.created_at
FROM (VALUES
  ('dividend.config.create', 'Crear configuraciones de dividendos por hipódromo y ciclo', 'config-winner', now()),
  ('dividend.config.update', 'Actualizar configuraciones de dividendos por hipódromo y ciclo', 'config-winner', now()),
  ('dividend.config.view', 'Consultar configuraciones de dividendos por hipódromo y ciclo', 'config-winner', now()),
  ('dividend.config.delete', 'Eliminar configuraciones de dividendos por hipódromo y ciclo', 'config-winner', now()),
  ('dividend.official.load', 'Cargar dividendos oficiales por carrera', 'config-winner', now()),
  ('dividend.official.update', 'Actualizar dividendos oficiales por carrera', 'config-winner', now()),
  ('dividend.official.view', 'Consultar dividendos oficiales por carrera', 'config-winner', now()),
  ('dividend.config.bulk', 'Operaciones masivas en configuraciones de dividendos', 'config-winner', now()),
  ('dividend.report.view', 'Ver reportes de dividendos', 'config-winner', now()),
  ('dividend.audit.view', 'Auditoría de cambios en dividendos', 'config-winner', now())
) AS v(key, alias, group_name, created_at)
WHERE NOT EXISTS (
  SELECT 1 FROM security.permissions p WHERE p."key" = v.key
);

-- Insertar permisos para apuestas Winner/Show/Place
INSERT INTO security.permissions ("key", alias, "group", created_at)
SELECT v.key, v.alias, v.group_name, v.created_at
FROM (VALUES
  ('bet.winner.create', 'Crear apuestas tipo Winner', 'betting-winner', now()),
  ('bet.show.create', 'Crear apuestas tipo Show', 'betting-winner', now()),
  ('bet.place.create', 'Crear apuestas tipo Place', 'betting-winner', now()),
  ('bet.winner.view', 'Consultar apuestas tipo Winner', 'betting-winner', now()),
  ('bet.show.view', 'Consultar apuestas tipo Show', 'betting-winner', now()),
  ('bet.place.view', 'Consultar apuestas tipo Place', 'betting-winner', now()),
  ('bet.winner.collect', 'Cobrar tickets de apuestas Winner', 'betting-winner', now()),
  ('bet.show.collect', 'Cobrar tickets de apuestas Show', 'betting-winner', now()),
  ('bet.place.collect', 'Cobrar tickets de apuestas Place', 'betting-winner', now()),
  ('bet.winner.cancel', 'Cancelar apuestas tipo Winner', 'betting-winner', now()),
  ('bet.show.cancel', 'Cancelar apuestas tipo Show', 'betting-winner', now()),
  ('bet.place.cancel', 'Cancelar apuestas tipo Place', 'betting-winner', now()),
  ('bet.bulk.validate', 'Validar múltiples tickets', 'betting-winner', now()),
  ('bet.bulk.collect', 'Cobrar múltiples tickets', 'betting-winner', now())
) AS v(key, alias, group_name, created_at)
WHERE NOT EXISTS (
  SELECT 1 FROM security.permissions p WHERE p."key" = v.key
);

-- Insertar permisos para reportes y estadísticas
INSERT INTO security.permissions ("key", alias, "group", created_at)
SELECT v.key, v.alias, v.group_name, v.created_at
FROM (VALUES
  ('bet.report.daily', 'Ver reportes diarios de apuestas', 'reports-winner', now()),
  ('bet.report.monthly', 'Ver reportes mensuales de apuestas', 'reports-winner', now()),
  ('bet.stats.view', 'Ver estadísticas generales de apuestas', 'reports-winner', now()),
  ('bet.stats.by_race', 'Ver estadísticas por carrera', 'reports-winner', now()),
  ('bet.stats.by_hippodrome', 'Ver estadísticas por hipódromo', 'reports-winner', now()),
  ('bet.stats.export', 'Exportar estadísticas de apuestas', 'reports-winner', now())
) AS v(key, alias, group_name, created_at)
WHERE NOT EXISTS (
  SELECT 1 FROM security.permissions p WHERE p."key" = v.key
);

-- Comentarios para documentación
COMMENT ON TABLE security.permissions IS 'Tabla de permisos del sistema, incluye nuevos permisos para Winner/Show/Place betting';

-- Confirmar inserción
SELECT 'Permisos insertados correctamente para el módulo Winner/Show/Place' as resultado;