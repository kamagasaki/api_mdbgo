package main

import (
	"api/config"
	"api/package/middelware"
	"api/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"os"
	"strings"
)

func init() {
	f, err := os.OpenFile("requests.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(f)
}

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(middelware.RequestLog)
	router.SetupRoutes(app)

	app.Use(cors.New(cors.Config{
		AllowOrigins:     strings.Join(config.GetAllowedOrigins(), ","),
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	port := config.Config("PORT")

	err := app.Listen(":" + port)
	if err != nil {
		log.Fatalf("Error starting the server: %v", err)
	}
}
