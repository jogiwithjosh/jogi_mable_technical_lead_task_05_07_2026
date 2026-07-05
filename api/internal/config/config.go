package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName string
	AppEnv  string

	Port string

	LogLevel string

	AllowedOrigin string

	JWTSecret string

	ClickHouse ClickHouseConfig
	JWT        JWTConfig
	Queue      QueueConfig
}

type JWTConfig struct {
	Secret string
	Expiry time.Duration
}

type ClickHouseConfig struct {
	Host     string
	Port     int
	Database string
	Username string
	Password string
	Secure   bool
}

type QueueConfig struct {
	Size int `mapstructure:"size"`
}

func Load() *Config {
	_ = godotenv.Load()

	expiry, err := time.ParseDuration(
		getString("JWT_EXPIRY", "24h"),
	)
	if err != nil {
		expiry = 24 * time.Hour
	}

	cfg := &Config{
		AppName:       getString("APP_NAME", "api"),
		AppEnv:        getString("APP_ENV", "development"),
		Port:          getString("PORT", "8080"),
		LogLevel:      getString("LOG_LEVEL", "info"),
		AllowedOrigin: getString("ALLOWED_ORIGIN", "http://localhost:5173"),
		JWTSecret:     getString("JWT_SECRET", "secret"),
		ClickHouse: ClickHouseConfig{
			Host:     getString("CLICKHOUSE_HOST", "clickhouse"),
			Port:     getInt("CLICKHOUSE_PORT", 9000),
			Database: getString("CLICKHOUSE_DATABASE", "analytics"),
			Username: getString("CLICKHOUSE_USERNAME", "default"),
			Password: getString("CLICKHOUSE_PASSWORD", ""),
			Secure:   getBool("CLICKHOUSE_SECURE", ""),
		},
		JWT: JWTConfig{
			Secret: getString(
				"JWT_SECRET",
				"change-me",
			),
			Expiry: expiry,
		},
	}

	fmt.Printf("%+v", cfg)

	return cfg
}

func getString(key string, defaultValue string) string {
	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	return value
}

func getInt(key string, defaultValue int) int {
	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	number, err := strconv.Atoi(value)
	if err != nil {
		log.Fatalf("invalid integer for %s", key)
	}

	return number
}

func getBool(key string, defaultValue string) bool {
	return os.Getenv(key) == "true"
}
