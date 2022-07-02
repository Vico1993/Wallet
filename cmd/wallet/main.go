package main

import (
	"Vico1993/Wallet/domain/cmd"
	"Vico1993/Wallet/domain/config"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	config.InitConfig()

	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	cmd.Execute()
}