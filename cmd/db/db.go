package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	var args = os.Args

	if len(args) > 1 {
		var dsn = os.Getenv("DATABASE_INIT_URI")
		if dsn == "" {
			dsn = "user:pass@tcp(host:port)/db"
		}

		var db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})

		if err != nil {
			log.Fatal(err.Error())
		}
		if args[1] == "up" {
			if err := db.Exec("CREATE DATABASE IF NOT EXISTS darinol").Error; err != nil {
				log.Fatal(err.Error())
			}

			fmt.Println("Database created successfully")
		}
		if args[1] == "down" {
			if err := db.Exec("DROP DATABASE darinol").Error; err != nil {
				log.Fatal(err.Error())
			}

			fmt.Println("Database deleted successfully")
		}
	} else {
		fmt.Println("Nothing happen, add 'up' or 'down' argument for create or delete database")
	}
}
