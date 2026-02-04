-- Crear tabla para bet_participants
CREATE TABLE gaming.bet_participants (
    id SERIAL PRIMARY KEY,
    id_bet BIGINT NOT NULL REFERENCES gaming.bet(id) ON DELETE CASCADE,
    id_participant INTEGER NOT NULL REFERENCES gaming.participants_race(id) ON DELETE CASCADE,
    is_main BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

-- Índices para optimizar consultas
CREATE INDEX idx_bet_participants_id_bet ON gaming.bet_participants(id_bet);
CREATE INDEX idx_bet_participants_id_participant ON gaming.bet_participants(id_participant);