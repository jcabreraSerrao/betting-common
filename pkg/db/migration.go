package db

import (
	"log"
	"os"

	"github.com/jcabreraSerrao/betting-common/entities/sql"
	"gorm.io/gorm"
)

// RunMigrations executes AutoMigrate for all registered entities and creates necessary schemas
func RunMigrations(db *gorm.DB) error {
	// Create common schemas
	schemas := []string{"gaming", "security", "config", "transactions", "payments", "reports"}
	for _, schema := range schemas {
		if err := db.Exec("CREATE SCHEMA IF NOT EXISTS " + schema).Error; err != nil {
			log.Printf("Error creating schema %s: %v", schema, err)
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
		&sql.BancaGroupLink{},
		&sql.SubGroup{},
		&sql.Roles{},
		&sql.Permissions{},
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
		&sql.SettlementRaceEstimate{},
		&sql.SettlementRaceEstimateDetail{},
		&sql.RacesProcessGroup{},
		&sql.GroupRaceActivation{},
		&sql.Tercios{},
		&sql.GroupWhatsAppConfig{},
		&sql.ConfigRemate{},
		&sql.UserGroup{},
		&sql.UserSubGroup{},
		&sql.RolesPermissions{},
		&sql.RemateEjemplares{},
		&sql.Board{},
		&sql.BoardRaceParada{},
		&sql.RetiredHorseGroup{},
		&sql.RetiredHorse{},
		&sql.Bet{},
		&sql.BetParticipants{},
		&sql.NoValid{},
		&sql.RaceDividend{},
		&sql.BoardOfficialGroup{},
		&sql.RetiredOfficial{},
		&sql.RaceDividendGroup{},
		&sql.Transactions{},
		&sql.TerciosRemate{},
		&sql.Refills{},
		&sql.Withdrawal{},
		&sql.ComboRemate{},
		&sql.ParticipantCombo{},
		&sql.TerciosCombo{},
		&sql.RaceGroupCommission{},
	)

	if err != nil {
		log.Printf("Error running migrations: %v", err)
		return err
	}

	log.Println("Migrations executed successfully")

	// Automatically run seeds from organized directories
	seedDirs := []string{"seeds", "security", "bet", "race", "tercios", "transaction"}
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
