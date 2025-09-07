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

	var user = model.User{
		Id:       1,
		Email:    "admin@gmail.com",
		Password: "Password",
	}

	var todo = model.Todo{
		Id:          1,
		Title:       "First list",
		Description: "Created by Absor use Go programming languange",
		UserId:      user.Id,
	}

	if err := db.Create(&user).Error; err != nil {
		log.Fatal(err.Error())
	}

	if err := db.Create(&todo).Error; err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Database seed successfully")
}
