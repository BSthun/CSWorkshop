package response

import (
	"context"

	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
)

type SuccessResponse struct {
	Success bool   `json:"success"`
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

type GenericSuccessResponse[T any] struct {
	Success bool   `json:"success"`
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Data    T      `json:"data,omitempty"`
}

func Success(c *fiber.Ctx, args1 any, args2 ...any) *SuccessResponse {
	sentryCtx := c.Locals("sentry").(context.Context)
	if sentryCtx.Value("span") != nil {
		sentryCtx.Value("span").(*sentry.Span).Status = sentry.SpanStatusOK
		sentryCtx.Value("span").(*sentry.Span).Finish()
	}

	if message, ok := args1.(string); ok {
		if len(args2) == 0 {
			return &SuccessResponse{
				Success: true,
				Message: message,
			}
		}
		if message2, ok := args2[0].(string); ok {
			return &SuccessResponse{
				Success: true,
				Code:    message,
				Message: message2,
			}
		} else {
			return &SuccessResponse{
				Success: true,
				Code:    message,
				Data:    args2[0],
			}
		}
	}

	return &SuccessResponse{
		Success: true,
		Data:    args1,
	}
}
