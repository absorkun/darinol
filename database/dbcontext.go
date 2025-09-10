package database

import (
	"fmt"
	"log"
	"os"

	config "github.com/absorkun/darinol/config/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New() *gorm.DB {
	var (
		user     = os.Getenv("POSTGRES_USER")
		password = os.Getenv("POSTGRES_PASSWORD")
		dbname   = os.Getenv("POSTGRES_DBNAME")
		host     = os.Getenv("POSTGRES_HOST")
		port     = os.Getenv("POSTGRES_PORT")
		sslmode  = os.Getenv("POSTGRES_SSLMODE")
		timezone = os.Getenv("POSTGRES_TIMEZONE")
	)
	var dsn = fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=%s TimeZone=%s statement_cache_mode=describe",
		user, password, dbname, host, port, sslmode, timezone,
	)
	var db, err = gorm.Open(postgres.Open(dsn), config.DbConfig())
	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}
