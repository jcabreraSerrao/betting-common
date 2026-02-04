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
		&sql.Currency{},
		&sql.Country{},
		&sql.Hippodromes{},
		&sql.User{},
		&sql.Group{},
		&sql.Race{},
		&sql.Tercios{},
		&sql.Bet{},
		&sql.ParticipantsRace{},
		&sql.BetParticipants{},
		&sql.Transactions{},
		&sql.Roles{},
		&sql.Permissions{},
		&sql.RolesPermissions{},
		&sql.TypeTransaction{},
		&sql.TypeBet{},
		&sql.TypeBetGroup{},
		&sql.TypeTercio{},
		&sql.UserTercio{},
		&sql.UserGroup{},
		&sql.UserSubGroup{},
		&sql.SubGroup{},
		&sql.PaymentPlatform{},
		&sql.PaymentPlatformsGroup{},
		&sql.PaymentPlatformCurrency{},
		&sql.TypePaymentPlatform{},
		&sql.GroupCurrency{},
		&sql.GroupExchangeRate{},
		&sql.ExchangeRate{},
		&sql.Config{},
		&sql.ConfigRemate{},
		&sql.RemateEjemplares{},
		&sql.Board{},
		&sql.BoardRaceParada{},
		&sql.BoardOfficialGroup{},
		&sql.RetiredHorse{},
		&sql.RetiredHorseGroup{},
		&sql.RetiredOfficial{},
		&sql.RaceDividend{},
		&sql.RaceDividendGroup{},
		&sql.RaceDividendCycleConfig{},
		&sql.Refills{},
		&sql.Withdrawal{},
		&sql.ComboRemate{},
		&sql.ParticipantCombo{},
		&sql.TerciosCombo{},
		&sql.TerciosRemate{},
		&sql.NoValid{},
		&sql.ProcessedMessage{},
		&sql.ProcessedMessageArchive{},
		&sql.WhatsAppProxy{},
		&sql.WhatsAppSession{},
		&sql.ChatWatermark{},
		&sql.LogsGroup{},
		&sql.RequestLog{},
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
