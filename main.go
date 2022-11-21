package main

import (
	"log"

	"go-admin/src/database"
	"go-admin/src/routes"
	"go-admin/src/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()
	database.AutoMigrate()
    app := fiber.New()
	services.Setup()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)
	

    log.Fatal(app.Listen(":8000"))
}