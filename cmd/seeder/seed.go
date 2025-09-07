package main

import (
	"fmt"
	"log"

	"github.com/absorkun/darinol/database"
	"github.com/absorkun/darinol/model"
	_ "github.com/joho/godotenv/autoload"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	var db = database.New()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err.Error())
	}

	var user = model.User{
		Id:       1,
		Email:    "admin@gmail.com",
		Password: string(hashedPassword),
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
