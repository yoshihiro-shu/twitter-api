package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	twitter_service "github.com/yoshihiro-shu/twitter/twitter/service"
)

func main() {
	loadEnvFile()

	t := twitter_service.New()

	data, err := t.GetUserTweetTimeline()
	if err != nil {
		log.Fatalln(err)
	}

	// cmd := command.NewApplicationCommand()

	// log.Fatalln(cmd.Execute())

	fmt.Println(data)

}

func loadEnvFile() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
