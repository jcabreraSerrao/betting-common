-- Crear tabla type_tercio
CREATE TABLE IF NOT EXISTS gaming.type_tercio (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    slug VARCHAR(255) NOT NULL UNIQUE,
    status BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

-- Agregar columna id_type_tercio a la tabla tercios
ALTER TABLE gaming.tercios
ADD COLUMN IF NOT EXISTS id_type_tercio INTEGER REFERENCES gaming.type_tercio(id);

-- Crear índice para la nueva columna
CREATE INDEX IF NOT EXISTS idx_tercios_type_tercio ON gaming.tercios(id_type_tercio);
