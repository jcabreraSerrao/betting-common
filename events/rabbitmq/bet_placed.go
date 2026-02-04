package rabbitmq

import "time"

// BetPlacedEvent is triggered when a new bet is created.
type BetPlacedEvent struct {
	BetID     uint64    `json:"bet_id"`
	RaceID    uint64    `json:"race_id"`
	UserID    uint64    `json:"user_id"`
	Amount    int64     `json:"amount"`
	Timestamp time.Time `json:"timestamp"`
}
