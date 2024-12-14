package main

import (
	"churchflowx/config"
	"churchflowx/internal/routes"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		AppName:       "churchflowx-be v1.0.1",
		JSONEncoder:   json.Marshal,   // specifying the json encoder better performance
		JSONDecoder:   json.Unmarshal, // specifying the json decoder for better performance
	})

	config.LoadEnv()
	config.SetupAppConfig(app)

	routes.SetupPeopleRoutes(app)
	routes.SetupFinanceRoutes(app)
	routes.SetupCollectionsRoutes(app)

	app.Listen(":3210")
}
