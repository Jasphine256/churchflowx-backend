package controllers

import (
	"churchflowx/internal/objects"
	"churchflowx/internal/services"

	"github.com/gofiber/fiber/v2"
)

// ############################# FUNDS HANDLERS #############################

func CreateFund(ctx *fiber.Ctx) error {
	user_id := ctx.Params("id")
	var fund objects.Fund
	err := ctx.BodyParser(&fund)
	if err != nil {
		return ResponseHandler(ctx, 500, fiber.Map{"message": "invalid body fields", "data": map[string]string{}})
	}
	success := services.AddFundToDb(user_id, fund)
	if !success {
		return ResponseHandler(ctx, 500, fiber.Map{"message": "failed", "data": map[string]string{}})
	}
	return ResponseHandler(ctx, 201, fiber.Map{"message": "created", "data": map[string]objects.Fund{"object": fund}})
}

func GetFund(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{"message": "value"})
}

func GetFunds(ctx *fiber.Ctx) error {
	type CurrentUserId struct {
		GID string
	}
	var current_user_id CurrentUserId
	err := ctx.QueryParser(&current_user_id)
	if err != nil {
		return ResponseHandler(ctx, 500, fiber.Map{"message": "invalid body fields passed", "data": map[string]string{}})
	}
	status, funds := services.GetFundsFromDb(current_user_id.GID)
	if !status {
		ResponseHandler(ctx, 500, fiber.Map{"message": "failed", "data": map[string]string{}})
	}
	return ctx.Status(200).JSON(fiber.Map{"message": "success", "data": fiber.Map{"objects": funds}})
}

func UpdateFund(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{"message": "value"})
}

func DeleteFund(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{"message": "value"})
}

// ############################# PAYMENTS HANDLERS #############################

func CreatePayment(ctx *fiber.Ctx) error {
	user_id := ctx.Params("id")
	var payment objects.Payment
	err := ctx.BodyParser(&payment)
	if err != nil {
		return ResponseHandler(ctx, 500, fiber.Map{"message": "invalid body fields supplied", "data": map[string]string{}})
	}
	success := services.AddPaymentToDb(user_id, payment)
	if !success {
		return ResponseHandler(ctx, 500, fiber.Map{"message": "failed", "data": map[string]string{}})
	}
	return ResponseHandler(ctx, 201, fiber.Map{"message": "created", "data": map[string]objects.Payment{"object": payment}})
}

func GetPayment(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{"message": "value"})
}

func GetPayments(ctx *fiber.Ctx) error {
	type CurrentUserId struct {
		GID string
	}
	var current_user_id CurrentUserId
	err := ctx.QueryParser(&current_user_id)
	if err != nil {
		return ResponseHandler(ctx, 500, fiber.Map{"message": "invalid body fields passed", "data": map[string]string{}})
	}
	status, payments := services.GetPaymentsFromDb(current_user_id.GID)
	if !status {
		ResponseHandler(ctx, 500, fiber.Map{"message": "failed", "data": map[string]string{}})
	}
	return ctx.Status(200).JSON(fiber.Map{"message": "success", "data": fiber.Map{"objects": payments}})
}

func UpdatePayment(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{"message": "value"})
}

func DeletePayment(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{"message": "value"})
}
