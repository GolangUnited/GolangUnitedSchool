package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/lozovoya/GolangUnitedSchool/internal/app"
)

func init() {
	// loads env vars from a .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	if err := app.Run(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
