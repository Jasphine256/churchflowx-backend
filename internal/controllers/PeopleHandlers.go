package controllers

import (
	"churchflowx/internal/objects"
	"churchflowx/internal/services"
	"log"

	"github.com/gofiber/fiber/v2"
)

// ############################# MINISTERS HANDLERS #############################

func CreateMinister(ctx *fiber.Ctx) error {
	var minister objects.Minister
	err := ctx.BodyParser(&minister)
	if err != nil {
		return ResponseHandler(ctx, 500, fiber.Map{"message": "invalid body fields supplied", "data": map[string]string{}})
	}
	success := services.AddMinisterToDb(&minister)
	if !success {
		return ResponseHandler(ctx, 500, fiber.Map{"message": "failed", "data": map[string]string{}})
	}
	return ResponseHandler(ctx, 201, fiber.Map{"message": "created", "data": map[string]objects.Minister{"object": minister}})
}

func GetMinister(ctx *fiber.Ctx) error {
	type CurrentUserId struct {
		GID string
	}
	var current_user_id CurrentUserId
	err := ctx.QueryParser(&current_user_id)
	if err != nil {
		return ResponseHandler(ctx, 500, fiber.Map{"message": "invalid query fields passed", "data": map[string]string{}})
	}
	minister_id := ctx.Params("id")
	status, minister := services.GetMinisterFromDb(current_user_id.GID, minister_id)
	if !status {
		ResponseHandler(ctx, 500, fiber.Map{"message": "failed", "data": map[string]string{}})
	}
	return ctx.Status(200).JSON(fiber.Map{"message": "success", "data": fiber.Map{"objects": minister}})
}

func GetMinisters(ctx *fiber.Ctx) error {
	type CurrentUserId struct {
		GID string
	}
	var current_user_id CurrentUserId
	err := ctx.QueryParser(&current_user_id)
	if err != nil {
		return ResponseHandler(ctx, 500, fiber.Map{"message": "invalid body fields passed", "data": map[string]string{}})
	}
	status, ministers := services.GetMinistersFromDb(current_user_id.GID)
	if !status {
		ResponseHandler(ctx, 500, fiber.Map{"message": "failed", "data": map[string]string{}})
	}
	return ctx.Status(200).JSON(fiber.Map{"message": "success", "data": fiber.Map{"objects": ministers}})
}

func UpdateMinister(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{"message": "value"})
}

func DeleteMinister(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{"message": "value"})
}

// ############################# MEMBER HANDLERS #############################

func CreateMember(ctx *fiber.Ctx) error {
	var member objects.Member
	err := ctx.BodyParser(&member)
	if err != nil {
		return ResponseHandler(ctx, 500, fiber.Map{"message": "invalid body fields supplied", "data": map[string]string{}})
	}
	success := services.AddMemberToDb(&member)
	if !success {
		return ResponseHandler(ctx, 500, fiber.Map{"message": "failed", "data": map[string]string{}})
	}
	return ResponseHandler(ctx, 201, fiber.Map{"message": "created", "data": map[string]objects.Member{"object": member}})
}

func GetMember(ctx *fiber.Ctx) error {
	type CurrentUserId struct {
		GID string
	}
	var current_user_id CurrentUserId
	err := ctx.QueryParser(&current_user_id)
	if err != nil {
		return ResponseHandler(ctx, 500, fiber.Map{"message": "invalid query fields passed", "data": map[string]string{}})
	}
	member_id := ctx.Params("id")
	status, member := services.GetMemberFromDb(current_user_id.GID, member_id)
	if !status {
		ResponseHandler(ctx, 500, fiber.Map{"message": "failed", "data": map[string]string{}})
	}
	return ctx.Status(200).JSON(fiber.Map{"message": "success", "data": fiber.Map{"objects": member}})
}

func GetMembers(ctx *fiber.Ctx) error {
	type CurrentUserId struct {
		GID string
	}
	var current_user_id CurrentUserId
	err := ctx.QueryParser(&current_user_id)
	if err != nil {
		return ResponseHandler(ctx, 500, fiber.Map{"message": "invalid body fields passed", "data": map[string]string{}})
	}
	status, members := services.GetMembersFromDb(current_user_id.GID)
	if !status {
		ResponseHandler(ctx, 500, fiber.Map{"message": "failed", "data": map[string]string{}})
	}
	return ctx.Status(200).JSON(fiber.Map{"message": "success", "data": fiber.Map{"objects": members}})
}

func UpdateMember(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{"message": "value"})
}

func DeleteMember(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{"message": "value"})
}

// ############################ VISITORS HANDLERS ###########################

func CreateVisitor(ctx *fiber.Ctx) error {
	var visitor objects.Visitor
	err := ctx.BodyParser(&visitor)
	if err != nil {
		return ResponseHandler(ctx, 500, fiber.Map{"message": "invalid body fields supplied", "data": map[string]string{}})
	}
	success := services.AddVisitorToDb(&visitor)
	if !success {
		return ResponseHandler(ctx, 500, fiber.Map{"message": "failed", "data": map[string]string{}})
	}
	return ResponseHandler(ctx, 201, fiber.Map{"message": "created", "data": map[string]objects.Visitor{"object": visitor}})
}

func GetVisitor(ctx *fiber.Ctx) error {
	type CurrentUserId struct {
		GID string
	}
	var current_user_id CurrentUserId
	err := ctx.QueryParser(&current_user_id)
	if err != nil {
		return ResponseHandler(ctx, 500, fiber.Map{"message": "invalid query fields passed", "data": map[string]string{}})
	}
	visitor_id := ctx.Params("id")
	status, visitor := services.GetVisitorFromDb(current_user_id.GID, visitor_id)
	if !status {
		ResponseHandler(ctx, 500, fiber.Map{"message": "failed", "data": map[string]string{}})
	}
	return ctx.Status(200).JSON(fiber.Map{"message": "success", "data": fiber.Map{"objects": visitor}})
}

func GetVisitors(ctx *fiber.Ctx) error {
	type CurrentUserId struct {
		GID string
	}
	var current_user_id CurrentUserId
	err := ctx.QueryParser(&current_user_id)
	if err != nil {
		return ResponseHandler(ctx, 500, fiber.Map{"message": "invalid body fields passed", "data": map[string]string{}})
	}
	status, visitors := services.GetVisitorsFromDb(current_user_id.GID)
	if !status {
		ResponseHandler(ctx, 500, fiber.Map{"message": "failed", "data": map[string]string{}})
	}
	return ctx.Status(200).JSON(fiber.Map{"message": "success", "data": fiber.Map{"objects": visitors}})
}

func UpdateVisitor(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{"message": "value"})
}

func DeleteVisitor(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{"message": "value"})
}

// ############################ PASTORS HANDLERS ###########################

func CreatePastor(ctx *fiber.Ctx) error {
	var pastor objects.Pastor
	err := ctx.BodyParser(&pastor)
	if err != nil {
		return ResponseHandler(ctx, 500, fiber.Map{"message": "invalid body fields supplied", "data": map[string]string{}})
	}
	success := services.AddPastorToDb(&pastor)
	if !success {
		return ResponseHandler(ctx, 500, fiber.Map{"message": "failed", "data": map[string]string{}})
	}
	return ResponseHandler(ctx, 201, fiber.Map{"message": "created", "data": map[string]objects.Pastor{"object": pastor}})
}

func GetPastor(ctx *fiber.Ctx) error {
	type CurrentUserId struct {
		GID string
	}
	var current_user_id CurrentUserId
	err := ctx.QueryParser(&current_user_id)
	if err != nil {
		return ResponseHandler(ctx, 500, fiber.Map{"message": "invalid query fields passed", "data": map[string]string{}})
	}
	pastor_id := ctx.Params("id")
	status, pastor := services.GetPastorFromDb(current_user_id.GID, pastor_id)
	if !status {
		ResponseHandler(ctx, 500, fiber.Map{"message": "failed", "data": map[string]string{}})
	}
	return ctx.Status(200).JSON(fiber.Map{"message": "success", "data": fiber.Map{"objects": pastor}})
}

func GetPastors(ctx *fiber.Ctx) error {
	type CurrentUserId struct {
		GID string
	}
	var current_user_id CurrentUserId
	err := ctx.QueryParser(&current_user_id)
	if err != nil {
		return ResponseHandler(ctx, 500, fiber.Map{"message": "invalid body fields passed", "data": map[string]string{}})
	}
	status, pastors := services.GetPastorsFromDb(current_user_id.GID)
	if !status {
		ResponseHandler(ctx, 500, fiber.Map{"message": "failed", "data": map[string]string{}})
	}
	return ctx.Status(200).JSON(fiber.Map{"message": "success", "data": fiber.Map{"objects": pastors}})
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
	type CurrentUserGid struct {
		GID string
	}
	var current_user_gid CurrentUserGid
	err := ctx.QueryParser(&current_user_gid)
	if err != nil {
		return ResponseHandler(ctx, 500, fiber.Map{"message": "invalid body fields passed", "data": map[string]string{}})
	}
	status, admins := services.GetAdminFromDb(current_user_gid.GID)
	if !status {
		ResponseHandler(ctx, 500, fiber.Map{"message": "failed", "data": map[string]string{}})
	}
	return ctx.Status(200).JSON(fiber.Map{"message": "success", "data": fiber.Map{"objects": admins}})
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
