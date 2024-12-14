package main

import (
	"churchflowx/config"
	"churchflowx/internal/routes"
	"log"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		AppName:       "churchflowx-be v1.0.1",
		JSONEncoder:   json.Marshal,   // specifying the json encoder better performance
		JSONDecoder:   json.Unmarshal, // specifying the json decoder for better performance
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Add your frontend domains here
		AllowHeaders: "Content-Type",
		AllowMethods: "GET, POST, PUT, DELETE",
	}))

	config.LoadEnv()
	config.SetupAppConfig(app)

	routes.SetupPeopleRoutes(app)
	routes.SetupFinanceRoutes(app)
	routes.SetupCollectionsRoutes(app)

	port := config.GetEnv("PORT", "8080")

	log.Printf("Starting server on port %s...\n", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
