-- Migration: Create settlement_race_estimate_details table
-- Date: 2025-11-11
-- Description: Table to store detailed bet information for settlement race estimates

CREATE TABLE gaming.settlement_race_estimate_details (
    id SERIAL PRIMARY KEY,
    settlement_race_estimate_id INTEGER NOT NULL REFERENCES gaming.settlement_race_estimates(id) ON DELETE CASCADE,
    bet_id INTEGER NOT NULL REFERENCES gaming.bet(id) ON DELETE CASCADE,
    payout DECIMAL(15,2) NOT NULL DEFAULT 0.00,
    refund DECIMAL(15,2) NOT NULL DEFAULT 0.00,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

-- Indexes for performance
CREATE INDEX idx_settlement_race_estimate_details_estimate_id ON gaming.settlement_race_estimate_details(settlement_race_estimate_id);
CREATE INDEX idx_settlement_race_estimate_details_bet_id ON gaming.settlement_race_estimate_details(bet_id);
CREATE INDEX idx_settlement_race_estimate_details_deleted_at ON gaming.settlement_race_estimate_details(deleted_at);