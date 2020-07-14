package main

import (
	"log"

	"github.com/SamtheSaint/jamgo/cmd"
	"github.com/joho/godotenv"
)

func init() {
	// local environment variables stored in ".env"
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}
}

func main() {
	cmd.Execute()
}
