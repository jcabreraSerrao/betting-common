package db

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/jcabreraSerrao/betting-common/pkg/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	dbInstance *gorm.DB
	dbOnce     sync.Once
)

// NewConnectionDB initializes and returns a GORM database connection
func NewConnectionDB() (*gorm.DB, error) {
	var err error
	config := utils.GetConfig()

	dbOnce.Do(func() {
		newLogger := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logger.Info,
				IgnoreRecordNotFoundError: true,
				ParameterizedQueries:      false,
				Colorful:                  true,
			},
		)

		dbInstance, err = gorm.Open(postgres.Open(config.Database.URL), &gorm.Config{
			Logger: newLogger,
		})

		if err != nil {
			log.Printf("Error connecting to database: %v", err)
			return
		}

		sqlDB, err := dbInstance.DB()
		if err != nil {
			log.Printf("Error getting SQL DB instance: %v", err)
			return
		}

		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxLifetime(time.Hour)
	})

	return dbInstance, err
}

// GetDB returns the singleton database instance
func GetDB() *gorm.DB {
	if dbInstance == nil {
		_, _ = NewConnectionDB()
	}
	return dbInstance
}
