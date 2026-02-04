package constants

// TypeTransactionIdConst maps transaction types to their numeric IDs
var TypeTransactionIdConst = map[string]int{
	"ganador":                 1,
	"perdedor":                2,
	"recarga":                 3,
	"retiro":                  4,
	"bloqueo":                 5,
	"tie":                     1,
	"win":                     2,
	"lose":                    5,
	"lose_middle":             6,
	"win_middle":              7,
	"jugado":                  9,
	"sistema":                 10,
	"bet-winner-create":       11,
	"bet-winner-pago":         12,
	"bet-winner-reembolso":    13,
	"initial-sum":             14,
	"initial-rest":            15,
	"transaction-delete-rest": 16,
	"transaction-delete-sum":  17,
	"transfer-banca-sum":      18,
	"transfer-banca-rest":     19,
}

const (
	// BancaTercioID is the fixed ID for the banca tercio
	BancaTercioID = 4

	// TypeTercio IDs
	TypeTercioCasaID      = 1
	TypeTercioBancaID     = 2
	TypeTercioUsuarioID   = 3
	TypeTercioUserGroupID = 4
)
