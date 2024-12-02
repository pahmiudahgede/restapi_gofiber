package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"rijik.id/restapi_gofiber/internal/api"
	"rijik.id/restapi_gofiber/internal/config"
	"rijik.id/restapi_gofiber/internal/connection"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.InitConfig()

	connection.InitDB()
}

func main() {
	app := fiber.New()

	api.RegisterRoutes(app)

	log.Fatal(app.Listen(":" + os.Getenv("SERVER_PORT")))
}
