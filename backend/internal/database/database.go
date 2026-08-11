package database

import (
	"log"
	"os"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDB connects to your Supabase Postgres database. This app talks
// only to Supabase — there is no local/self-hosted Postgres fallback.
//
// Set DATABASE_URL to the connection string from your Supabase project:
// Settings -> Database -> Connection string (URI).
func ConnectDB() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is not set. Copy .env.example to .env and set it to your Supabase connection string (Settings -> Database -> Connection string).")
	}

	pgConfig := postgres.Config{
		DSN: dsn,
	}

	// Supabase's connection pooler (PgBouncer, port 6543, "Transaction" mode)
	// does not support prepared statements, which GORM's pgx driver uses by
	// default. If we detect the pooler (or the caller explicitly asks for it
	// via DB_USE_POOLER=true), switch to the simple query protocol so GORM
	// works correctly through it. The direct connection (port 5432) does not
	// need this.
	if os.Getenv("DB_USE_POOLER") == "true" || strings.Contains(dsn, ":6543") {
		pgConfig.PreferSimpleProtocol = true
	}

	database, err := gorm.Open(postgres.New(pgConfig), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to Supabase database: %v", err)
	}

	DB = database
	log.Println("Connected to Supabase database successfully")
}
