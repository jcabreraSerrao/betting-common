-- =============================================================================
-- ASIGNACIÓN DE PERMISOS A ROLES ADMIN
-- =============================================================================
-- Fecha: 2024-10-10
-- Descripción: Asigna todos los permisos existentes a los roles con nombre 'admin'
--              en la tabla security.roles_permissions, evitando duplicados.
-- =============================================================================

-- Insertar permisos para roles 'admin' que no existan ya
INSERT INTO security.roles_permissions (id_roles, id_permission, status, created_at, updated_at)
SELECT r.id, p.id, true, now(), now()
FROM security.roles r
CROSS JOIN security.permissions p
WHERE r.name = 'admin'
AND NOT EXISTS (
    SELECT 1 FROM security.roles_permissions rp
    WHERE rp.id_roles = r.id AND rp.id_permission = p.id
);

-- =============================================================================
-- VERIFICACIÓN
-- =============================================================================

-- Verificar asignaciones realizadas
SELECT
    r.name as rol,
    COUNT(rp.id_permission) as permisos_asignados,
    STRING_AGG(p.alias, ', ') as lista_permisos
FROM security.roles r
JOIN security.roles_permissions rp ON r.id = rp.id_roles
JOIN security.permissions p ON rp.id_permission = p.id
WHERE r.name = 'admin'
GROUP BY r.id, r.name
ORDER BY r.name;

-- Mensaje de confirmación
SELECT
    COUNT(*) as total_asignaciones_realizadas,
    '✅ Permisos asignados correctamente a roles admin' as resultado
FROM security.roles_permissions rp
JOIN security.roles r ON rp.id_roles = r.id
WHERE r.name = 'admin';