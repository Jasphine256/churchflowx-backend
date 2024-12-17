package controllers

import (
	"churchflowx/internal/objects"
	"churchflowx/internal/services"
	"log"

	"github.com/gofiber/fiber/v2"
)

func ResponseHandler(ctx *fiber.Ctx, status int, mapped_data fiber.Map) error {
	return ctx.Status(status).JSON(mapped_data)
}

// ############################# TASKS HANDLERS #############################

func CreateTask(ctx *fiber.Ctx) error {

	user_id := ctx.Params("id")
	var new_task objects.Task
	err := ctx.BodyParser(&new_task)
	if err != nil {
		log.Fatal(err)
		return ResponseHandler(ctx, 500, fiber.Map{"message": "failed to parse body", "data": map[string]string{}})
	}
	success := services.AddTaskToDb(user_id, new_task)
	if !success {
		return ResponseHandler(ctx, 500, fiber.Map{"message": "failed to add task", "data": map[string]string{}})
	}
	return ResponseHandler(ctx, 201, fiber.Map{"message": "created", "data": map[string]objects.Task{"object": new_task}})
}

func GetTask(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{"message": "value"})
}

func GetTasks(ctx *fiber.Ctx) error {
	type CurrentUserId struct {
		GID string
	}
	var current_user_id CurrentUserId
	err := ctx.QueryParser(&current_user_id)
	if err != nil {
		return ResponseHandler(ctx, 500, fiber.Map{"message": "invalid body fields passed", "data": map[string]string{}})
	}
	status, tasks := services.GetTasksFromDb(current_user_id.GID)
	if !status {
		ResponseHandler(ctx, 500, fiber.Map{"message": "failed", "data": map[string]string{}})
	}
	return ctx.Status(200).JSON(fiber.Map{"message": "success", "data": fiber.Map{"objects": tasks}})
}

func UpdateTask(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{"message": "value"})
}

func DeleteTask(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{"message": "value"})
}

// ############################# PLANS HANDLERS #############################

func CreatePlan(ctx *fiber.Ctx) error {
	user_id := ctx.Params("id")
	var plan objects.Plan
	err := ctx.BodyParser(&plan)
	if err != nil {
		log.Fatal(err)
		return ResponseHandler(ctx, 500, fiber.Map{"message": "invalid body fields passed", "data": map[string]string{}})
	}
	success := services.AddPlanToDb(user_id, plan)
	if !success {
		return ResponseHandler(ctx, 500, fiber.Map{"message": "failed", "data": map[string]string{}})
	}
	return ResponseHandler(ctx, 201, fiber.Map{"message": "created", "data": map[string]objects.Plan{"object": plan}})
}

func GetPlan(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{"message": "value"})
}

func GetPlans(ctx *fiber.Ctx) error {
	type CurrentUserId struct {
		GID string
	}
	var current_user_id CurrentUserId
	err := ctx.QueryParser(&current_user_id)
	if err != nil {
		return ResponseHandler(ctx, 500, fiber.Map{"message": "invalid body fields passed", "data": map[string]string{}})
	}
	status, plans := services.GetPlansFromDb(current_user_id.GID)
	if !status {
		ResponseHandler(ctx, 500, fiber.Map{"message": "failed", "data": map[string]string{}})
	}
	return ctx.Status(200).JSON(fiber.Map{"message": "success", "data": fiber.Map{"objects": plans}})
}

func UpdatePlan(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{"message": "value"})
}

func DeletePlan(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{"message": "value"})
}

// ############################ PROJECTS HANDLERS ###########################

func CreateProject(ctx *fiber.Ctx) error {
	user_id := ctx.Params("id")
	var project objects.Project
	err := ctx.BodyParser(&project)
	if err != nil {
		return ResponseHandler(ctx, 500, fiber.Map{"message": "invalid body fields supplied ", "data": map[string]string{}})
	}
	success := services.AddProjectToDb(user_id, &project)
	if !success {
		return ResponseHandler(ctx, 500, fiber.Map{"message": "failed", "data": map[string]string{}})
	}
	return ResponseHandler(ctx, 201, fiber.Map{"message": "created", "data": map[string]objects.Project{"object": project}})
}

func GetProject(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{"message": "value"})
}

func GetProjects(ctx *fiber.Ctx) error {
	type CurrentUserId struct {
		GID string
	}
	var current_user_id CurrentUserId
	err := ctx.QueryParser(&current_user_id)
	if err != nil {
		return ResponseHandler(ctx, 500, fiber.Map{"message": "invalid body fields passed", "data": map[string]string{}})
	}
	status, projects := services.GetProjectsFromDb(current_user_id.GID)
	if !status {
		ResponseHandler(ctx, 500, fiber.Map{"message": "failed", "data": map[string]string{}})
	}
	return ctx.Status(200).JSON(fiber.Map{"message": "success", "data": fiber.Map{"objects": projects}})
}

func UpdateProject(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{"message": "value"})
}

func DeleteProject(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{"message": "value"})
}
