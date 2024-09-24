package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadENVVar() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error env not found")
	}
}
