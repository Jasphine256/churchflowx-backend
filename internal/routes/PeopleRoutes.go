package routes

import (
	"churchflowx/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupPeopleRoutes(app *fiber.App) {

	users := app.Group("/people")

	// #####################  ROUTES FOR /users/* #####################

	users.Route("/ministers", func(api fiber.Router) {

		api.Get("/", func(c *fiber.Ctx) error {
			return controllers.GetMinisters(c)
		})

		api.Get("/:id", func(c *fiber.Ctx) error {
			return controllers.GetMinister(c)
		})

		api.Post("/new", func(c *fiber.Ctx) error {
			return controllers.CreateMinister(c)
		})

		api.Patch("/:id", func(c *fiber.Ctx) error {
			return controllers.UpdateMinister(c)
		})

		api.Delete("/:id", func(c *fiber.Ctx) error {
			return controllers.DeleteMinister(c)
		})
	})

	users.Route("/members", func(api fiber.Router) {

		api.Get("/", func(c *fiber.Ctx) error {
			return controllers.GetMembers(c)
		})

		api.Get("/:id", func(c *fiber.Ctx) error {
			return controllers.GetMember(c)
		})

		api.Post("/new", func(c *fiber.Ctx) error {
			return controllers.CreateMember(c)
		})

		api.Patch("/:id", func(c *fiber.Ctx) error {
			return controllers.UpdateMember(c)
		})

		api.Delete("/:id", func(c *fiber.Ctx) error {
			return controllers.DeleteMember(c)
		})
	})

	users.Route("/visitors", func(api fiber.Router) {

		api.Get("/", func(c *fiber.Ctx) error {
			return controllers.GetVisitors(c)
		})

		api.Get("/:id", func(c *fiber.Ctx) error {
			return controllers.GetVisitor(c)
		})

		api.Post("/new", func(c *fiber.Ctx) error {
			return controllers.CreateVisitor(c)
		})

		api.Patch("/:id", func(c *fiber.Ctx) error {
			return controllers.UpdateVisitor(c)
		})

		api.Delete("/:id", func(c *fiber.Ctx) error {
			return controllers.DeleteVisitor(c)
		})
	})

	users.Route("/pastors", func(api fiber.Router) {

		api.Get("/", func(c *fiber.Ctx) error {
			return controllers.GetPastors(c)
		})

		api.Get("/:id", func(c *fiber.Ctx) error {
			return controllers.GetPastor(c)
		})

		api.Post("/new", func(c *fiber.Ctx) error {
			return controllers.CreatePastor(c)
		})

		api.Patch("/:id", func(c *fiber.Ctx) error {
			return controllers.UpdatePastor(c)
		})

		api.Delete("/:id", func(c *fiber.Ctx) error {
			return controllers.DeletePastor(c)
		})
	})

	users.Route("/admins", func(api fiber.Router) {

		api.Get("/", func(c *fiber.Ctx) error {
			return controllers.GetAdmin(c)
		})

		api.Post("/new", func(c *fiber.Ctx) error {
			return controllers.CreateAdmin(c)
		})

		api.Patch("/:id", func(c *fiber.Ctx) error {
			return controllers.UpdateAdmin(c)
		})

		api.Delete("/:id", func(c *fiber.Ctx) error {
			return controllers.DeleteAdmin(c)
		})
	})
}
