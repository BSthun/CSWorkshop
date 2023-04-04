package response

import (
	"context"

	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
)

type InfoResponse struct {
	Success bool   `json:"success"`
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

type GenericInfoResponse[T any] struct {
	Success bool   `json:"success"`
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Data    T      `json:"data,omitempty"`
}

func Info(c *fiber.Ctx, args1 any, args2 ...any) *InfoResponse {
	sentryCtx := c.Locals("sentry").(context.Context)
	if sentryCtx.Value("span") != nil {
		sentryCtx.Value("span").(*sentry.Span).Status = sentry.SpanStatusOK
		sentryCtx.Value("span").(*sentry.Span).Finish()
	}

	if message, ok := args1.(string); ok {
		if len(args2) == 0 {
			return &InfoResponse{
				Success: true,
				Message: message,
			}
		}
		if message2, ok := args2[0].(string); ok {
			return &InfoResponse{
				Success: true,
				Code:    message,
				Message: message2,
			}
		} else {
			return &InfoResponse{
				Success: true,
				Code:    message,
				Data:    args2[0],
			}
		}
	}

	return &InfoResponse{
		Success: true,
		Data:    args1,
	}
}
