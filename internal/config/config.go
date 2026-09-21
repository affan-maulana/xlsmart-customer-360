package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Env      string
	Port     string
	LogLevel string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string

	JWTSecret string

	VQMBaseURL  string
	NBSSBaseURL string
}

func Load() *Config {
	// Load .env file (ignore error if not found)
	_ = godotenv.Load()

	return &Config{
		Env:      getEnv("ENV", "development"),
		Port:     getEnv("PORT", "8080"),
		LogLevel: getEnv("LOG_LEVEL", "info"),

		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "postgres"),
		DBName:     getEnv("DB_NAME", "ums_bff"),
		DBSSLMode:  getEnv("DB_SSLMODE", "disable"),

		JWTSecret: getEnv("JWT_SECRET", "your-secret-key"),

		VQMBaseURL:  getEnv("VQM_BASE_URL", "http://prod-view-query-manager-svc.comet.svc.cluster.local:7003"),
		NBSSBaseURL: getEnv("NBSS_BASE_URL", "http://localhost:7005"),
	}
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}

func (c *Config) DBConnStr() string {
	return "host=" + c.DBHost + " port=" + c.DBPort + " user=" + c.DBUser +
		" password=" + c.DBPassword + " dbname=" + c.DBName + " sslmode=" + c.DBSSLMode
}
