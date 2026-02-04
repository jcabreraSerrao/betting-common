-- Crear tabla para almacenar comisiones y conteo de apuestas por carrera y grupo
CREATE TABLE IF NOT EXISTS gaming.race_group_bet_info (
    id SERIAL PRIMARY KEY,
    id_race BIGINT NOT NULL REFERENCES gaming.race(id),
    id_group INT NOT NULL REFERENCES security.group(id),
    total_commission DECIMAL(18,8) DEFAULT 0,
    bet_count INT DEFAULT 0,
    total_amount DECIMAL(18,8) DEFAULT 0,
    race_date TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(id_race, id_group)
);

-- Índices para mejorar rendimiento
CREATE INDEX IF NOT EXISTS idx_race_group_bet_info_race ON gaming.race_group_bet_info(id_race);
CREATE INDEX IF NOT EXISTS idx_race_group_bet_info_group ON gaming.race_group_bet_info(id_group);
CREATE INDEX IF NOT EXISTS idx_race_group_bet_info_date ON gaming.race_group_bet_info(race_date);
ALTER TABLE gaming.race_group_bet_info ADD CONSTRAINT unique_race_group UNIQUE (id_race, id_group);