package config

import (
	"github.com/gofiber/fiber/v2"
)

func SetupAppConfig(app *fiber.App) {

	app.Static("/static", "./public")
	// => http://localhost:3000/static/hello.html looks in public folder

}
