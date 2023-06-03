package ifiber

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"backend/types/response"
)

func NotFoundHandler(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
		Success: false,
		Message: fmt.Sprintf("%s %s not found", ctx.Method(), ctx.Path()),
		Error:   "404_NOT_FOUND",
	})
}
