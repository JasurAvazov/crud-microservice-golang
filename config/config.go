package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	LogLevel         string
	HTTPPort         string
	HTTPHost         string
	PostgresHost     string
	PostgresPort     int
	PostgresDatabase string
	PostgresUser     string
	PostgresPassword string
	PostgresSSLMode  string
}

func (c *Config) PostgresURL() string {
	if c.PostgresUser == "" {
		return fmt.Sprintf("host=%s port=%d  dbname=%s sslmode=disable",
			c.PostgresHost,
			c.PostgresPort,
			c.PostgresDatabase)
	}
	if c.PostgresPassword == "" {
		return fmt.Sprintf("host=%s port=%d user=%s  dbname=%s sslmode=disable",
			c.PostgresHost,
			c.PostgresPort,
			c.PostgresUser,
			c.PostgresDatabase)
	}
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.PostgresHost,
		c.PostgresPort,
		c.PostgresUser,
		c.PostgresPassword,
		c.PostgresDatabase)
}

func Load() Config {
	if err := godotenv.Load(); //load .env file
		err != nil {
		log.Print("No .env file found")
	}
	config := Config{}

	config.LogLevel = cast.ToString(getOrReturnDefaultValue("LOG_LEVEL", "debug"))
	config.HTTPPort = cast.ToString(getOrReturnDefaultValue("HTTP_PORT", ":7077"))
	config.HTTPHost = cast.ToString(getOrReturnDefaultValue("SERVER_IP", "localhost"))

	config.PostgresHost = cast.ToString(getOrReturnDefaultValue("POSTGRES_HOST", "localhost"))
	config.PostgresPort = cast.ToInt(getOrReturnDefaultValue("POSTGRES_PORT", 5432))
	config.PostgresDatabase = cast.ToString(getOrReturnDefaultValue("POSTGRES_DATABASE", "postgres"))
	config.PostgresUser = cast.ToString(getOrReturnDefaultValue("POSTGRES_USER", "postgres"))
	config.PostgresPassword = cast.ToString(getOrReturnDefaultValue("POSTGRES_PASSWORD", "root"))
	config.PostgresSSLMode = cast.ToString(getOrReturnDefaultValue("POSTGRES_SSLMODE", "disable"))

	return config
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)

	if exists {
		return os.Getenv(key)
	}
	return defaultValue
}
