package main

import (
	"log"

	"github.com/biskitsx/aws-s3-tutorial/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()
	uploadController := handler.NewUploadController()
	app.Post("/", uploadController.UploadImage)

	app.Listen(":7070")

}
