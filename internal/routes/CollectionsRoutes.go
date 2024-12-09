package routes

import (
	"churchflowx/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupCollectionsRoutes(app *fiber.App) {

	users := app.Group("/collections")

	// #####################  ROUTES FOR /users/* #####################

	users.Route("/tasks", func(api fiber.Router) {

		api.Get("/", func(c *fiber.Ctx) error {
			return controllers.GetTasks(c)
		})

		api.Get("/:id", func(c *fiber.Ctx) error {
			return controllers.GetTask(c)
		})

		api.Post("/new", func(c *fiber.Ctx) error {
			return controllers.CreateTask(c)
		})

		api.Patch("/:id", func(c *fiber.Ctx) error {
			return controllers.UpdateTask(c)
		})

		api.Delete("/:id", func(c *fiber.Ctx) error {
			return controllers.DeleteTask(c)
		})
	})

	users.Route("/plans", func(api fiber.Router) {

		api.Get("/", func(c *fiber.Ctx) error {
			return controllers.GetPlans(c)
		})

		api.Get("/:id", func(c *fiber.Ctx) error {
			return controllers.GetPlan(c)
		})

		api.Post("/new", func(c *fiber.Ctx) error {
			return controllers.CreatePlan(c)
		})

		api.Patch("/:id", func(c *fiber.Ctx) error {
			return controllers.UpdatePlan(c)
		})

		api.Delete("/:id", func(c *fiber.Ctx) error {
			return controllers.DeletePlan(c)
		})
	})

	users.Route("/projects", func(api fiber.Router) {

		api.Get("/", func(c *fiber.Ctx) error {
			return controllers.GetProjects(c)
		})

		api.Get("/:id", func(c *fiber.Ctx) error {
			return controllers.GetProject(c)
		})

		api.Post("/new", func(c *fiber.Ctx) error {
			return controllers.CreateProject(c)
		})

		api.Patch("/:id", func(c *fiber.Ctx) error {
			return controllers.UpdateProject(c)
		})

		api.Delete("/:id", func(c *fiber.Ctx) error {
			return controllers.DeleteProject(c)
		})
	})
}
