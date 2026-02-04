-- =============================================================================
-- INSERCIÓN DE PERMISOS GENERALES DEL SISTEMA
-- =============================================================================
-- Fecha: 2024-01-15
-- Descripción: Inserta permisos básicos para funcionalidades principales
--              del sistema de apuestas de caballos
-- =============================================================================

-- =============================================================================
-- 1. PERMISOS PARA GESTIÓN DE CARRERAS
-- =============================================================================
INSERT INTO security.permissions ("key", alias, "group", created_at, updated_at)
SELECT v.key, v.alias, v.group_name, v.created_at, v.updated_at
FROM (VALUES
  ('race.list', 'Listar Carreras', 'race', now(), now()),
  ('race.create', 'Crear Carreras', 'race', now(), now()),
  ('race.report', 'Generar Reportes de Carreras', 'race', now(), now()),
  ('race.revert', 'Revertir Carreras', 'race', now(), now())
) AS v(key, alias, group_name, created_at, updated_at)
WHERE NOT EXISTS (
  SELECT 1 FROM security.permissions p WHERE p."key" = v.key
);

-- =============================================================================
-- 2. PERMISOS PARA GESTIÓN DE APUESTAS
-- =============================================================================
INSERT INTO security.permissions ("key", alias, "group", created_at, updated_at)
SELECT v.key, v.alias, v.group_name, v.created_at, v.updated_at
FROM (VALUES
  ('bet.list', 'Listar Apuestas', 'bet', now(), now()),
  ('bet.create', 'Crear Apuestas', 'bet', now(), now()),
  ('bet.delete', 'Cancelar Apuestas', 'bet', now(), now()),
  ('bet.close', 'Procesar Apuestas', 'bet', now(), now()),
  ('bet.revert', 'Revertir Apuestas', 'bet', now(), now())

) AS v(key, alias, group_name, created_at, updated_at)
WHERE NOT EXISTS (
  SELECT 1 FROM security.permissions p WHERE p."key" = v.key
);

-- =============================================================================
-- 3. PERMISOS PARA GESTIÓN DE PIZARRAS
-- =============================================================================
INSERT INTO security.permissions ("key", alias, "group", created_at, updated_at)
SELECT v.key, v.alias, v.group_name, v.created_at, v.updated_at
FROM (VALUES
  ('board.create', 'Crear Pizarra', 'board', now(), now()),
  ('board.revert', 'Revertir Pizarra', 'board', now(), now())

) AS v(key, alias, group_name, created_at, updated_at)
WHERE NOT EXISTS (
  SELECT 1 FROM security.permissions p WHERE p."key" = v.key
);

-- =============================================================================
-- 4. PERMISOS PARA GESTIÓN DE JORNADAS DE TRABAJO
-- =============================================================================
INSERT INTO security.permissions ("key", alias, "group", created_at, updated_at)
SELECT v.key, v.alias, v.group_name, v.created_at, v.updated_at
FROM (VALUES
  ('working_day.list', 'Listar Jornadas de Trabajo', 'working_day', now(), now())
) AS v(key, alias, group_name, created_at, updated_at)
WHERE NOT EXISTS (
  SELECT 1 FROM security.permissions p WHERE p."key" = v.key
);

-- =============================================================================
-- 5. PERMISOS PARA GESTIÓN DE ROLES
-- =============================================================================
INSERT INTO security.permissions ("key", alias, "group", created_at, updated_at)
SELECT v.key, v.alias, v.group_name, v.created_at, v.updated_at
FROM (VALUES
  ('role.list', 'Listar Roles', 'role', now(), now()),
  ('role.create', 'Crear Roles', 'role', now(), now()),
  ('role.update', 'Actualizar Roles', 'role', now(), now())
) AS v(key, alias, group_name, created_at, updated_at)
WHERE NOT EXISTS (
  SELECT 1 FROM security.permissions p WHERE p."key" = v.key
);

-- =============================================================================
-- 6. PERMISOS PARA GESTIÓN DE GRUPOS E HIPÓDROMOS
-- =============================================================================
INSERT INTO security.permissions ("key", alias, "group", created_at, updated_at)
SELECT v.key, v.alias, v.group_name, v.created_at, v.updated_at
FROM (VALUES
  ('hipo_group.list', 'Listar Grupos de Hipódromos', 'hipo_group', now(), now()),
  ('hipo_group.update', 'Actualizar Grupos de Hipódromos', 'hipo_group', now(), now()),
  ('hipo_group.disable', 'Deshabilitar/Habilitar Grupos de Hipódromos', 'hipo_group', now(), now()),
  ('subgroup.list', 'Listar Subgrupos', 'subGroup', now(), now()),
  ('subgroup.create', 'Crear Subgrupos', 'subGroup', now(), now()),
  ('subgroup.update', 'Actualizar Subgrupos', 'subGroup', now(), now()),
  ('subgroup.disable', 'Deshabilitar/Habilitar Subgrupos', 'subGroup', now(), now()),
  ('hipo_subgroup.update', 'Actualizar Subgrupos de Hipódromos', 'hipo_subGroup', now(), now())
) AS v(key, alias, group_name, created_at, updated_at)
WHERE NOT EXISTS (
  SELECT 1 FROM security.permissions p WHERE p."key" = v.key
);

-- =============================================================================
-- 7. PERMISOS PARA GESTIÓN DE TRANSACCIONES Y TERCIOS
-- =============================================================================
INSERT INTO security.permissions ("key", alias, "group", created_at, updated_at)
SELECT v.key, v.alias, v.group_name, v.created_at, v.updated_at
FROM (VALUES
  ('transaction_tercio.list', 'Listar Transacciones de Tercios', 'transaction_tercio', now(), now())
) AS v(key, alias, group_name, created_at, updated_at)
WHERE NOT EXISTS (
  SELECT 1 FROM security.permissions p WHERE p."key" = v.key
);

-- =============================================================================
-- 8. PERMISOS PARA GESTIÓN DE TERCIOS
-- =============================================================================
INSERT INTO security.permissions ("key", alias, "group", created_at, updated_at)
SELECT v.key, v.alias, v.group_name, v.created_at, v.updated_at
FROM (VALUES
  ('tercio.list', 'Listar Tercios', 'tercio', now(), now()),
  ('tercio.update', 'Actualizar Tercios', 'tercio', now(), now()),
  ('tercio.create', 'Crear Tercios', 'tercio', now(), now()),
  ('tercio.delete', 'Eliminar Tercios', 'tercio', now(), now()),
  ('tercio.recargar', 'Recargar Tercios', 'tercio', now(), now()),
  ('tercio.retirar', 'Retirar Tercios', 'tercio', now(), now()),
  ('tercio.activate', 'Activar/desactivar Tercios', 'tercio', now(), now())
) AS v(key, alias, group_name, created_at, updated_at)
WHERE NOT EXISTS (
  SELECT 1 FROM security.permissions p WHERE p."key" = v.key
);

-- =============================================================================
-- 9. PERMISOS PARA GESTIÓN DE MONEDAS
-- =============================================================================
INSERT INTO security.permissions ("key", alias, "group", created_at, updated_at)
SELECT v.key, v.alias, v.group_name, v.created_at, v.updated_at
FROM (VALUES
  ('currency.currency', 'Gestionar Monedas', 'currency', now(), now()),
  ('currency.config', 'Configurar Monedas', 'currency', now(), now())
) AS v(key, alias, group_name, created_at, updated_at)
WHERE NOT EXISTS (
  SELECT 1 FROM security.permissions p WHERE p."key" = v.key
);

-- =============================================================================
-- PERMISOS PARA GESTIÓN DE TRANSACCIONES BANCA
-- =============================================================================
INSERT INTO security.permissions ("key", alias, "group", created_at, updated_at)
SELECT v.key, v.alias, v.group_name, v.created_at, v.updated_at
FROM (VALUES
  ('banca.close_group', 'Cerrar Grupo (Transferir a Banca)', 'banca', now(), now()),
  ('banca.close_individual', 'Cerrar Tercio Individual', 'banca', now(), now())
) AS v(key, alias, group_name, created_at, updated_at)
WHERE NOT EXISTS (
  SELECT 1 FROM security.permissions p WHERE p."key" = v.key
);

-- =============================================================================
-- VERIFICACIÓN Y DOCUMENTACIÓN
-- =============================================================================

-- Actualizar comentario de la tabla
--COMMENT ON TABLE security.permissions IS 'Tabla de permisos del sistema - incluye permisos generales y específicos del módulo Winner/Show/Place';

-- Verificar permisos insertados por grupo
SELECT
  "group" as grupo_permisos,
  COUNT(*) as cantidad_permisos,
  STRING_AGG(alias, ', ') as permisos
FROM security.permissions
WHERE "group" IN ('race', 'bet', 'board', 'working_day', 'role', 'hipo_group', 'subGroup', 'hipo_subGroup', 'transaction_tercio', 'tercio', 'currency')
GROUP BY "group"
ORDER BY "group";

-- Mensaje de confirmación
SELECT
  COUNT(*) as total_permisos_generales_insertados,
  '✅ Permisos generales del sistema insertados correctamente' as resultado
FROM security.permissions
WHERE "group" IN ('race', 'bet', 'board', 'working_day', 'role', 'hipo_group', 'subGroup', 'hipo_subGroup', 'transaction_tercio', 'tercio', 'currency');


