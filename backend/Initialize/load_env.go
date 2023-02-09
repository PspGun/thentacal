package Initialize

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Load_env() {
	if ok := os.Getenv("GO_ENV"); ok != "production" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
}
