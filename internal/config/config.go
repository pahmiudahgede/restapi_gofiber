package config

import (
	"log"
	"os"
)

var (
	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string

	JWTSecret string

	ServerHost string
	ServerPort string
)

func InitConfig() {

	ServerHost = os.Getenv("SERVER_HOST")
	ServerPort = os.Getenv("SERVER_PORT")
	DBHost = os.Getenv("DB_HOST")
	DBPort = os.Getenv("DB_PORT")
	DBName = os.Getenv("DB_NAME")
	DBUser = os.Getenv("DB_USER")
	DBPassword = os.Getenv("DB_PASSWORD")
	JWTSecret = os.Getenv("JWT_KEY_SECRET")

	if ServerHost == "" || ServerPort == "" || JWTSecret == "" {
		log.Fatal("Required environment variables are not set!")
	}
}
