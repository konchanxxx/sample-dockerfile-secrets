package main

import (
	"log"
	"os"
)

func main() {
	passwordFile := os.Getenv("DB_PASSWORD_FILE")
	password, err := os.ReadFile(passwordFile)
	if err != nil {
		log.Fatalf("Failed to read password: %s", err)
	}

	log.Printf("Successfully read password: %s", password)

	passwordOtherFile := os.Getenv("DB_PASSWORD_OTHER_FILE")
	passwordOther, err := os.ReadFile(passwordOtherFile)
	if err != nil {
		log.Fatalf("Failed to read other password: %s", err)
	}

	log.Printf("Successfully read other password: %s", passwordOther)
}
