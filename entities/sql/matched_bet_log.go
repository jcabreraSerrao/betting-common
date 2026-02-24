package sql

import (
	"bytes"
	"encoding/json"
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

// MatchedBetLog almacena el registro principal de una apuesta procesada (Original)
type MatchedBetLog struct {
	OriginalMessageID string          `gorm:"primaryKey;column:original_message_id;type:varchar(255)" json:"original_message_id"` // Equivale también a MongoID en la tabla Bet
	Tercio1ID         uint            `gorm:"column:tercio_1_id;index" json:"tercio_1_id"`
	Tercio1Name       string          `gorm:"column:tercio_1_name;type:varchar(255)" json:"tercio_1_name"`
	RaceID            uint64          `gorm:"column:race_id;index:idx_matched_race_group_status;index:idx_matched_race" json:"race_id"`
	Race              Race            `gorm:"foreignKey:RaceID" json:"race"`
	GroupID           uint            `gorm:"column:group_id;index:idx_matched_race_group_status;index:idx_matched_group" json:"group_id"`
	Group             Group           `gorm:"foreignKey:GroupID" json:"group"`
	TotalAmount       decimal.Decimal `gorm:"column:total_amount;type:decimal(18,8)" json:"total_amount"`
	RemainingAmount   decimal.Decimal `gorm:"column:remaining_amount;type:decimal(18,8)" json:"remaining_amount"`
	BetType           string          `gorm:"column:bet_type;type:varchar(255)" json:"bet_type"`
	Status            string          `gorm:"column:status;type:varchar(50);index:idx_matched_race_group_status;index:idx_matched_status" json:"status"`
	OriginalParsed    JSONBMap        `gorm:"column:original_parsed;type:jsonb" json:"original_parsed"`
	Matches           MatchInfos      `gorm:"column:matches;type:jsonb;index:idx_matched_matches,type:gin" json:"matches"`
	CreatedAt         time.Time       `gorm:"column:created_at;index" json:"created_at"`
	UpdatedAt         time.Time       `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt         gorm.DeletedAt  `gorm:"index" json:"-"`
}

func (MatchedBetLog) TableName() string {
	return "whatsapp.matched_bet_logs"
}

// TestMatchResult almacena pruebas de matched bets
type TestMatchResult struct {
	ID                uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	OriginalMessageID string    `gorm:"column:original_message_id;type:varchar(255);index" json:"original_message_id"`
	MatcherMessageID  string    `gorm:"column:matcher_message_id;type:varchar(255);index" json:"matcher_message_id"`
	OriginalBody      string    `gorm:"column:original_body;type:text" json:"original_body"`
	MatcherBody       string    `gorm:"column:matcher_body;type:text" json:"matcher_body"`
	GroupID           uint      `gorm:"column:group_id;index" json:"group_id"`
	ParsedBet         JSONBMap  `gorm:"column:parsed_bet;type:jsonb" json:"parsed_bet"`
	Status            string    `gorm:"column:status;type:varchar(50)" json:"status"`
	Error             string    `gorm:"column:error;type:text" json:"error"`
	CreatedAt         time.Time `gorm:"column:created_at" json:"created_at"`
}

func (TestMatchResult) TableName() string {
	return "whatsapp.test_match_results"
}

// Tipos auxiliares para JSONB
type JSONBMap map[string]interface{}

func (j *JSONBMap) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}
	return json.Unmarshal(bytes, j)
}

func (j JSONBMap) Value() (interface{}, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}

// MatchInfo encapsula los detalles de quien casa la apuesta
type MatchInfo struct {
	MatcherMessageID string          `json:"matcher_message_id"`
	MatcherBody      string          `json:"matcher_body"`
	ParticipantID    uint            `json:"participant_id"`
	ParticipantName  string          `json:"participant_name"`
	AmountMatched    decimal.Decimal `json:"amount_matched"`
	MatcherParsed    JSONBMap        `json:"matcher_parsed"` // Resultado del parseo de la respuesta
	Timestamp        time.Time       `json:"timestamp"`
}

type MatchInfos []MatchInfo

func (j *MatchInfos) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}
	b, ok := value.([]byte)
	if !ok {
		return nil
	}
	// Tolerancia a ambos formatos: arreglo [] u objeto {} (datos legacy)
	trimmed := bytes.TrimSpace(b)
	if len(trimmed) > 0 && trimmed[0] == '{' {
		var single MatchInfo
		if err := json.Unmarshal(b, &single); err != nil {
			return err
		}
		*j = MatchInfos{single}
		return nil
	}
	return json.Unmarshal(b, j)
}

func (j MatchInfos) Value() (interface{}, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}
