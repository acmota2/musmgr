package main

import (
	db "backend/db"
	"backend/model"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	loadDb()
}

func loadDb() {
	db.Connect()
	db.DB.AutoMigrate(&model.Category{}, &model.Event{}, &model.Song{})
}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal(err)
	}
}
