package ifiber

import (
	"context"
	"strings"

	"github.com/getsentry/sentry-go"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"backend/types/response"
	"backend/utils/text"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	// Defer sentry store
	var sentryCtx context.Context
	var sentryHub *sentry.Hub
	var sentryTrace map[string]any
	if c.Locals("sentry") == nil {
		goto switcher
	}

	sentryCtx = c.Locals("sentry").(context.Context)
	sentryHub = sentry.GetHubFromContext(sentryCtx)
	defer func() {
		sentryHub.Scope().SetContext("error", sentryTrace)
		if sentryCtx.Value("span") != nil {
			sentryCtx.Value("span").(*sentry.Span).Status = sentry.SpanStatusAborted
			sentryCtx.Value("span").(*sentry.Span).Finish()
		}
	}()

switcher:

	// * Internal Error Instance
	if e, ok := err.(*response.ErrorInstance); ok {
		if e.Code == "" {
			e.Code = "GENERIC_ERROR"
		}

		if e.Err != nil {
			sentryTrace = map[string]any{
				"code":    e.Code,
				"message": e.Message,
				"error":   e.Err.Error(),
			}
			return c.Status(fiber.StatusBadRequest).JSON(&response.ErrorResponse{
				Success: false,
				Code:    e.Code,
				Message: e.Message,
				Error:   e.Err.Error(),
			})
		}

		sentryTrace = map[string]any{
			"code":    e.Code,
			"message": e.Message,
		}
		return c.Status(fiber.StatusBadRequest).JSON(&response.ErrorResponse{
			Success: false,
			Code:    e.Code,
			Message: e.Message,
		})
	}

	// * Fiber Error
	if e, ok := err.(*fiber.Error); ok {
		sentryTrace = map[string]any{
			"fiber": e.Message,
		}
		return c.Status(e.Code).JSON(&response.ErrorResponse{
			Success: false,
			Code:    strings.ReplaceAll(strings.ToUpper(e.Error()), " ", "_"),
			Message: e.Error(),
			Error:   e.Error(),
		})
	}

	// Validation Error
	if e, ok := err.(validator.ValidationErrors); ok {
		var lists []string
		for _, err := range e {
			lists = append(lists, text.DescribeValidator(err))
			break
		}

		message := strings.Join(lists[:], ", ")

		sentryTrace = map[string]any{
			"message": message,
		}
		return c.Status(fiber.StatusBadRequest).JSON(&response.ErrorResponse{
			Success: false,
			Code:    "VALIDATION_FAILED",
			Message: "Validation failed on field " + message,
			Error:   e.Error(),
		})
	}

	// * Unknown error
	sentryTrace = map[string]any{
		"error": err.Error(),
	}
	return c.Status(fiber.StatusInternalServerError).JSON(&response.ErrorResponse{
		Success: false,
		Code:    "UNKNOWN_SERVER_SIDE_ERROR",
		Message: "Unknown server side error",
		Error:   err.Error(),
	})
}
