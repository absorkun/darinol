package main

import (
	"fmt"
	"log"

	"github.com/absorkun/darinol/database"
	"github.com/absorkun/darinol/model"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	var db = database.New()

	if err := db.AutoMigrate(&model.User{}); err != nil {
		log.Fatal(err.Error())
	}
	if err := db.AutoMigrate(&model.Todo{}); err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Database migrate successfully")
}
