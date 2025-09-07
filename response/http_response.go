package response

import (
	"github.com/gofiber/fiber/v3"
)

func Ok(c fiber.Ctx, data any) error {
	return c.Status(200).JSON(SuccessStruct{Success: true, Data: data})
}

func Created(c fiber.Ctx, data any) error {
	return c.Status(201).JSON(SuccessStruct{Success: true, Data: data})
}

func NoContent(c fiber.Ctx) error {
	return c.Status(204).JSON(SuccessStruct{Success: true})
}

func BadRequest(c fiber.Ctx, message string) error {
	return c.Status(400).JSON(FailedStruct{Success: false, Message: message})
}

func NotFound(c fiber.Ctx, message string) error {
	return c.Status(404).JSON(FailedStruct{Success: false, Message: message})
}

func InternalServerError(c fiber.Ctx, message string) error {
	return c.Status(500).JSON(FailedStruct{Success: false, Message: message})
}
