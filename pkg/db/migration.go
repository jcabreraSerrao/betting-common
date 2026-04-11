package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jcabreraSerrao/betting-common/entities/sql"
	"gorm.io/gorm"
)

// RunMigrations executes AutoMigrate for all registered entities and creates necessary schemas
func RunMigrations(db *gorm.DB) error {
	// Ensure PostgreSQL is configured for logical replication (CDC)
	ensureLogicalReplication(db)

	// Create common schemas
	schemas := []string{"gaming", "security", "config", "transactions", "payments", "reports", "evolution", "whatsapp"}
	for _, schema := range schemas {
		if err := db.Exec("CREATE SCHEMA IF NOT EXISTS " + schema).Error; err != nil {
			log.Printf("Error creating schema %s: %v", schema, err)
			return err
		}
	}

	// Enable extension for UUIDs
	if err := db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`).Error; err != nil {
		log.Printf("Error creating uuid-ossp extension: %v", err)
		return err
	}

	// Pre-create tables involved in circular dependencies to avoid "relation does not exist" errors
	preCreateTables := []struct {
		table  string
		schema string
	}{
		{"transactions", "transactions"},
		{"config_remate", "gaming"},
		{"bet", "gaming"},
	}

	for _, tc := range preCreateTables {
		query := "CREATE TABLE IF NOT EXISTS " + tc.schema + "." + tc.table + " (id bigserial PRIMARY KEY)"
		if err := db.Exec(query).Error; err != nil {
			log.Printf("Error pre-creating table %s.%s: %v", tc.schema, tc.table, err)
			return err
		}
	}

	// Order is important for FKs
	err := db.AutoMigrate(
		&sql.Config{},
		&sql.Currency{},
		&sql.Country{},
		&sql.ExchangeRate{},
		&sql.Hippodromes{},
		&sql.LogsGroup{},
		&sql.RequestLog{},
		&sql.User{},
		&sql.Group{},
		&sql.GroupExcludedParticipant{},
		&sql.BancaGroupLink{},
		&sql.SubGroup{},
		&sql.Roles{},
		&sql.Permissions{},
		&sql.RolesPermissions{},
		&sql.TypeTercio{},
		&sql.TypeBetGroup{},
		&sql.TypeBet{},
		&sql.TypeTransaction{},
		&sql.TypePaymentPlatform{},
		&sql.PaymentPlatformsGroup{},
		&sql.FlagNotAproved{},
		&sql.ChatWatermark{},
		&sql.ProcessedMessage{},
		&sql.ProcessedMessageArchive{},
		&sql.WhatsAppProxy{},
		&sql.WhatsAppSession{},
		&sql.GroupCountryMinBetConfig{},
		&sql.RaceDividendCycleConfig{},
		&sql.RaceDividendConfig{},
		&sql.RaceDividendConfigRange{},
		&sql.RaceDividendRange{},
		&sql.PaymentPlatform{},
		&sql.PaymentPlatformCurrency{},
		&sql.GroupCurrency{},
		&sql.GroupExchangeRate{},
		&sql.Race{},
		&sql.ParticipantsRace{},
		&sql.SettlementRaceEstimate{},
		&sql.SettlementRaceEstimateDetail{},
		&sql.RacesProcessGroup{},
		&sql.GroupRaceActivation{},
		&sql.Tercios{},
		&sql.TercioContact{},
		&sql.TercioReverso{},
		&sql.UserTercio{},
		&sql.WorkingDay{},
		&sql.WorkingDaySnapshot{},
		&sql.GroupWhatsAppConfig{},
		&sql.ConfigRemate{},
		&sql.UserGroup{},
		&sql.UserSubGroup{},
		&sql.RemateEjemplares{},
		&sql.Board{},
		&sql.BoardRaceParada{},
		&sql.BoardConfiRemate{},
		&sql.RetiredHorseGroup{},
		&sql.RetiredHorse{},
		&sql.Bet{},
		&sql.BetParticipants{},
		&sql.NoValid{},
		&sql.RaceDividend{},
		&sql.BoardOfficialGroup{},
		&sql.RetiredOfficial{},
		&sql.RaceDividendGroup{},
		&sql.MatchedBetLog{},
		&sql.TestMatchResult{},
		&sql.WhatsappMessageLog{},
		&sql.WhatsappMatchAttempt{},
		&sql.Polla{},
		&sql.PollaRace{},
		&sql.PollaInvalidHorse{},
		&sql.CommandRule{},
		&sql.Transactions{},
		&sql.TerciosRemate{},
		&sql.Refills{},
		&sql.Withdrawal{},
		&sql.ComboRemate{},
		&sql.ParticipantCombo{},
		&sql.TerciosCombo{},
		&sql.RaceGroupCommission{},
		&sql.PollaParticipant{},
		&sql.PollaSelection{},
	)
	if err != nil {
		log.Printf("Error migrating first batch: %v", err)
		return err
	}

	if err != nil {
		log.Printf("Error running migrations: %v", err)
		return err
	}

	// Ensure unique constraint for match_attempts upsert
	if err := db.Exec("CREATE UNIQUE INDEX IF NOT EXISTS idx_wa_attempt_quoted_id ON whatsapp.match_attempts (quoted_message_id)").Error; err != nil {
		log.Printf("Error creating unique index on match_attempts: %v", err)
	}

	log.Println("Migrations executed successfully")

	// Automatically run seeds from organized directories
	seedDirs := []string{"seeds", "security", "bet", "race", "tercios", "transaction", "functions", "reports"}
	for _, dir := range seedDirs {
		path := "sql/" + dir
		files, err := os.ReadDir(path)
		if err != nil {
			log.Printf("Warning: could not read seed directory %s: %v", path, err)
			continue
		}

		for _, file := range files {
			if !file.IsDir() && (len(file.Name()) > 4 && file.Name()[len(file.Name())-4:] == ".sql") {
				seedFile := path + "/" + file.Name()
				if err := RunSeed(db, seedFile); err != nil {
					log.Printf("Error running seed %s: %v", seedFile, err)
					return err
				}
			}
		}
	}

	return nil
}

// RunSeed executes a SQL file from the specified path
func RunSeed(db *gorm.DB, path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	if err := db.Exec(string(data)).Error; err != nil {
		log.Printf("Error executing seed %s: %v", path, err)
		return err
	}

	log.Printf("Seed executed: %s", path)
	return nil
}

// ensureLogicalReplication attempts to configure PostgreSQL for logical decoding (CDC).
// This requires superuser privileges and a server restart to take effect if changes are made.
func ensureLogicalReplication(db *gorm.DB) {
	commands := []struct {
		key   string
		value string
	}{
		{"wal_level", "logical"},
		{"max_replication_slots", "10"},
		{"max_wal_senders", "10"},
	}

	for _, cmd := range commands {
		// Check current value
		var currentValue string
		err := db.Raw(fmt.Sprintf("SHOW %s", cmd.key)).Scan(&currentValue).Error
		if err == nil && currentValue == cmd.value {
			continue // Already configured correctly
		}

		// Attempt to update
		query := fmt.Sprintf("ALTER SYSTEM SET %s = '%s'", cmd.key, cmd.value)
		if err := db.Exec(query).Error; err != nil {
			log.Printf("Warning: Could not set %s to %s automatically: %v. This may require manual configuration for CDC.", cmd.key, cmd.value, err)
		} else {
			log.Printf("PostgreSQL configuration updated: %s = %s. A RESTART of the database is required to apply changes.", cmd.key, cmd.value)
		}
	}
}
