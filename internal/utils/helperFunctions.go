package utils

import (
	"fmt"
	"github.com/jinzhu/configor"
	"github.com/joho/godotenv"
	"log"
	"math/rand"
	"os"
	"streams/internal/model"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func LoadConfig() model.Config {

	var config model.Config
	err := configor.Load(&config, "./config/config.yml")

	if err != nil {
		fmt.Println("File not found", err)
		os.Exit(-1)
	}
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("Go Environment:", os.Getenv("CONFIGOR_ENV"))

	return config
}
