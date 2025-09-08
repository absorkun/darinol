package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: db.exe [up|down|show|clear]")
		return
	}

	db, err := connect()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	switch os.Args[1] {
	case "up":
		createDatabase(db, "darinol")
	case "down":
		dropDatabase(db, "darinol")
	case "show":
		listDatabases(db)
	case "clear":
		clearStatements(db)
	default:
		fmt.Println("Unknown command. Use: up, down, show, or clear")
	}
}

func connect() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=%s TimeZone=%s statement_cache_mode=describe",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DBNAME"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_SSLMODE"),
		os.Getenv("POSTGRES_TIMEZONE"),
	)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
}

func createDatabase(db *gorm.DB, name string) {
	query := fmt.Sprintf("CREATE DATABASE %s", name)
	if err := db.Exec(query).Error; err != nil {
		log.Fatal("Error creating database:", err)
	}
	fmt.Println("✅ Database created:", name)
}

func dropDatabase(db *gorm.DB, name string) {
	query := fmt.Sprintf("DROP DATABASE IF EXISTS %s", name)
	if err := db.Exec(query).Error; err != nil {
		log.Fatal("Error dropping database:", err)
	}
	fmt.Println("🗑️ Database dropped:", name)
}

func listDatabases(db *gorm.DB) {
	type Database struct {
		Datname string
	}

	var databases []Database
	if err := db.Raw("SELECT datname FROM pg_database WHERE datistemplate = false").Scan(&databases).Error; err != nil {
		log.Fatal("Error listing databases:", err)
	}

	fmt.Println("📋 Available databases:")
	for _, d := range databases {
		fmt.Println("- " + d.Datname)
	}

	clearStatements(db)
}

func clearStatements(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Println("⚠️ Tidak bisa ambil sql.DB:", err)
		return
	}

	if _, err := sqlDB.Exec(`DEALLOCATE ALL`); err != nil {
		log.Println("⚠️ Gagal deallocate statements:", err)
	} else {
		fmt.Println("✅ Semua prepared statement di session ini sudah dibersihkan")
	}
}
