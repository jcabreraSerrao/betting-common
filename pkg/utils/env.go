package utils

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	Database struct {
		URL string `env:"DATABASE_URL"`
	}
	Redis struct {
		URL string `env:"REDIS_URL"`
	}
	MongoDB struct {
		URI      string `env:"MONGODB_URI"`
		DATABASE string `env:"MONGODB_DATABASE"`
	}
	RabbitMQ struct {
		URL string `env:"RABBITMQ_URL"`
	}
	JWT struct {
		SECRET string `env:"JWT_SECRET"`
		EXPIRE string `env:"JWT_EXPIRE"`
	}
	Server struct {
		PORT string `env:"PORT"`
		ENV  string `env:"ENV"`
	}
}

var (
	configInstance *Config
	once           sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		_ = godotenv.Load()
		configInstance = &Config{}
		configInstance.Database.URL = os.Getenv("DATABASE_URL")
		configInstance.Redis.URL = os.Getenv("REDIS_URL")
		configInstance.MongoDB.URI = getEnvDefault("MONGODB_URI", "mongodb://localhost:27017")
		configInstance.MongoDB.DATABASE = getEnvDefault("MONGODB_DATABASE", "betting_common")
		configInstance.RabbitMQ.URL = getEnvDefault("RABBITMQ_URL", "amqp://guest:guest@localhost:5672/")
		configInstance.JWT.SECRET = os.Getenv("JWT_SECRET")
		configInstance.JWT.EXPIRE = getEnvDefault("JWT_EXPIRE", "24h")
		configInstance.Server.PORT = getEnvDefault("PORT", "8080")
		configInstance.Server.ENV = getEnvDefault("ENV", "development")

		if configInstance.Database.URL == "" {
			log.Println("WARNING: DATABASE_URL is not set")
		}
	})
	return configInstance
}

func getEnvDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func (c *Config) Validate() error {
	if c.Database.URL == "" {
		return fmt.Errorf("DATABASE_URL is required")
	}
	return nil
}
