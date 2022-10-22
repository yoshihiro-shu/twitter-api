package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/yoshihiro-shu/twitter/command"
	twitter_service "github.com/yoshihiro-shu/twitter/twitter/service"
)

func main() {
	loadEnvFile()

	t := twitter_service.New()

	cmd := command.NewApplicationCommand(t)

	cmd.Execute()

}

func loadEnvFile() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
