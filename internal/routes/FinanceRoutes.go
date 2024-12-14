package routes

import (
	"churchflowx/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupFinanceRoutes(app *fiber.App) {

	users := app.Group("/finance")

	// #####################  ROUTES FOR /users/* #####################

	users.Route("/funds", func(api fiber.Router) {

		api.Get("/", func(c *fiber.Ctx) error {
			return controllers.GetFunds(c)
		})

		api.Get("/:id", func(c *fiber.Ctx) error {
			return controllers.GetFund(c)
		})

		api.Post("/new/:id", func(c *fiber.Ctx) error {
			return controllers.CreateFund(c)
		})

		api.Patch("/:id", func(c *fiber.Ctx) error {
			return controllers.UpdateFund(c)
		})

		api.Delete("/:id", func(c *fiber.Ctx) error {
			return controllers.DeleteFund(c)
		})
	})

	users.Route("/payments", func(api fiber.Router) {

		api.Get("/", func(c *fiber.Ctx) error {
			return controllers.GetPayments(c)
		})

		api.Get("/:id", func(c *fiber.Ctx) error {
			return controllers.GetPayment(c)
		})

		api.Post("/new/:id", func(c *fiber.Ctx) error {
			return controllers.CreatePayment(c)
		})

		api.Patch("/:id", func(c *fiber.Ctx) error {
			return controllers.UpdatePayment(c)
		})

		api.Delete("/:id", func(c *fiber.Ctx) error {
			return controllers.DeletePayment(c)
		})
	})
}
