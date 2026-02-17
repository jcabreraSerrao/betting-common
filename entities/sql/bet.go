package sql

import (
	"encoding/json"
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

// Bet representa una apuesta en el sistema de carreras de caballos.
// Contiene información sobre la cantidad apostada, la carrera, los tercios involucrados,
// el tipo de apuesta, y otros detalles relevantes.
type Bet struct {
	ID              uint64            `gorm:"primaryKey;autoIncrement" json:"id"`
	Ticket          string            `gorm:"column:ticket;type:varchar(255);index:idx_ticket_group" json:"ticket"` // Índice para buscar tickets por grupo
	Amount          decimal.Decimal   `gorm:"column:amount;type:decimal(18,8)" json:"amount"`
	RaceID          uint64            `gorm:"column:id_race;index:idx_group_race" json:"idRace"`
	Race            Race              `gorm:"foreignKey:RaceID;references:ID" json:"race"`
	TercioId        *uint             `gorm:"column:id_tercio;index"  json:"idTercio"`
	Tercio          *Tercios          `gorm:"foreignKey:TercioId;references:ID" json:"tercio"`
	TercioId2       *uint             `gorm:"column:id_tercio2;index" json:"idTercio2"`
	Tercio2         *Tercios          `gorm:"foreignKey:TercioId2;references:ID" json:"tercio2"`
	TypeBetGroupId  uint              `gorm:"column:id_type_bet_group;index" json:"idTypeBetGroup"`
	TypeBetGroup    TypeBetGroup      `gorm:"foreignKey:TypeBetGroupId;references:ID" json:"typeBetGroup"`
	TypeProfitId    *uint             `gorm:"column:types_profit" json:"idTypeProfit" `
	TypesProfit     TypeBetGroup      `gorm:"foreignKey:TypeProfitId;references:ID" json:"typesProfit"`
	Cancel          bool              `gorm:"column:cancel;default:false;not null" json:"cancel"`
	BetId           *uint             `gorm:"column:id_bet" json:"idBet" `
	Bet             *Bet              `gorm:"foreignKey:BetId;references:ID" json:"bet"`
	GroupId         uint              `gorm:"column:id_group;index;index:idx_group_race;index:idx_group_status;index:idx_ticket_group" json:"idGroup" `
	Group           Group             `gorm:"foreignKey:GroupId;references:ID" json:"group"`
	CancelAutomatic bool              `gorm:"column:cancel_automatic;default:false;not null" json:"cancel_automatic"`
	Transactions    []Transactions    `gorm:"foreignKey:BetsId;references:ID" json:"transactions"`
	BetParticipants []BetParticipants `gorm:"foreignKey:BetID" json:"betParticipants"`
	CurrencyID      *uint             `gorm:"column:currency_id;index" json:"currency_id"`
	Currency        *Currency         `gorm:"foreignKey:CurrencyID;references:ID" json:"currency"`
	CantidadJugadas *float64          `gorm:"column:cantidad_jugadas;type:decimal(15,6)" json:"cantidadJugadas"`
	Status          string            `gorm:"column:status;default:'pending';index:idx_group_status" json:"status"`
	NoValids        []NoValid         `gorm:"foreignKey:BetID" json:"noValids"`
	CreatedAt       time.Time         `gorm:"index:idx_created_at_desc,priority:1;index:idx_created_at_desc,sort:desc" json:"createdAt"`
	TextAutomatic   string            `gorm:"column:text_automatic;type:text" json:"text_automatic"`
	MongoID         string            `gorm:"column:mongo_id;type:varchar(255)" json:"mongo_id"`
	Recibo          *bool             `gorm:"column:recibo;default:null" json:"recibo"`
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName devuelve el nombre de la tabla en la base de datos para la entidad Bet.
// Este método es utilizado por GORM para mapear la estructura a la tabla correcta.
func (s *Bet) TableName() string {
	return "gaming.bet"
}

// MarshalJSON personaliza la serialización JSON de la entidad Bet.
// Si el ID es 0, devuelve "null". De lo contrario, serializa la estructura completa.
func (s *Bet) MarshalJSON() ([]byte, error) {
	if s.ID == 0 {
		return []byte("null"), nil
	}

	type Alias Bet
	return json.Marshal(
		&struct {
			*Alias
		}{
			Alias: (*Alias)(s),
		},
	)
}
