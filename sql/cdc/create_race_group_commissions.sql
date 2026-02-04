-- Índices para mejorar rendimiento
CREATE INDEX IF NOT EXISTS idx_race_group_bet_info_race ON gaming.race_group_bet_info(id_race);
CREATE INDEX IF NOT EXISTS idx_race_group_bet_info_group ON gaming.race_group_bet_info(id_group);
CREATE INDEX IF NOT EXISTS idx_race_group_bet_info_date ON gaming.race_group_bet_info(race_date);
ALTER TABLE gaming.race_group_bet_info ADD CONSTRAINT unique_race_group UNIQUE (id_race, id_group);