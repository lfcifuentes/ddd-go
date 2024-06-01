package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/lfcifuentes/ddd-go/cmd"
	"github.com/spf13/viper"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	// Bind environment variables
	viper.AutomaticEnv()

	// Execute root command
	cmd.Execute()
}
