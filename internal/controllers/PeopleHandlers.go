package controllers

import (
	"churchflowx/internal/objects"
	"churchflowx/internal/services"
	"log"

	"github.com/gofiber/fiber/v2"
)

// ############################# MINISTERS HANDLERS #############################

func CreateMinister(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{"message": "value"})
}

func GetMinister(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{"message": "value"})
}

func GetMinisters(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{"message": "value"})
}

func UpdateMinister(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{"message": "value"})
}

func DeleteMinister(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{"message": "value"})
}

// ############################# MEMBER HANDLERS #############################

func CreateMember(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{"message": "value"})
}

func GetMember(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{"message": "value"})
}

func GetMembers(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{"message": "value"})
}

func UpdateMember(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{"message": "value"})
}

func DeleteMember(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{"message": "value"})
}

// ############################ VISITORS HANDLERS ###########################

func CreateVisitor(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{"message": "value"})
}

func GetVisitor(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{"message": "value"})
}

func GetVisitors(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{"message": "value"})
}

func UpdateVisitor(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{"message": "value"})
}

func DeleteVisitor(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{"message": "value"})
}

// ############################ PASTORS HANDLERS ###########################

func CreatePastor(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{"message": "value"})
}

func GetPastor(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{"message": "value"})
}

func GetPastors(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{"message": "value"})
}

func UpdatePastor(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{"message": "value"})
}

func DeletePastor(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{"message": "value"})
}

// ############################ ADMINS HANDLERS ###########################

func CreateAdmin(ctx *fiber.Ctx) error {
	var Admin objects.Admin
	err := ctx.BodyParser(&Admin)
	if err != nil {
		log.Fatal(err)
	}

	success := services.AddAdminToDb(&Admin)
	if !success {
		return ResponseHandler(ctx, 500, fiber.Map{"message": "error", "data": map[string]string{}})
	}
	return ResponseHandler(ctx, 201, fiber.Map{"message": "created", "data": map[string]objects.Admin{"object": Admin}})
}

func GetAdmin(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{"message": "value"})
}

func GetPAdmin(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{"message": "value"})
}

func UpdateAdmin(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{"message": "value"})
}

func DeleteAdmin(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{"message": "value"})
}
