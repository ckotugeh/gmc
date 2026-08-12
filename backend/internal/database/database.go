package database

import (
	"fmt"
	"log"
	"os"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDB connects to Postgres using either a full DATABASE_URL connection
// string or individual DB_* environment variables.
func ConnectDB() {
	dsn := os.Getenv("DATABASE_URL")

	if dsn == "" {
		// Fall back to individual DB_* vars
		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		user := os.Getenv("DB_USER")
		password := os.Getenv("DB_PASSWORD")
		dbname := os.Getenv("DB_NAME")
		sslmode := os.Getenv("DB_SSLMODE")

		if host == "" || user == "" || dbname == "" {
			log.Fatal("Database not configured: set DATABASE_URL or DB_HOST/DB_USER/DB_PASSWORD/DB_NAME/DB_PORT env vars.")
		}

		if port == "" {
			port = "5432"
		}
		if sslmode == "" {
			sslmode = "require"
		}

		dsn = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			host, port, user, password, dbname, sslmode,
		)
	}

	pgConfig := postgres.Config{
		DSN: dsn,
	}

	// PgBouncer (port 6543, transaction mode) doesn't support prepared
	// statements — switch to simple query protocol when detected.
	if os.Getenv("DB_USE_POOLER") == "true" || strings.Contains(dsn, ":6543") {
		pgConfig.PreferSimpleProtocol = true
	}

	database, err := gorm.Open(postgres.New(pgConfig), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	DB = database
	log.Println("✅ Connected to database successfully")
}
