package response

import (
	"context"
	"errors"

	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
)

type ErrorInstance struct {
	Message string
	Code    string
	Err     error
}

func (v *ErrorInstance) Error() string {
	return v.Message
}

func Error(c *fiber.Ctx, critical bool, message string, args2 ...any) *ErrorInstance {
	if c != nil {
		// * Prepare sentry tracing
		sentryCtx := c.Locals("sentry").(context.Context)
		sentryHub := sentry.GetHubFromContext(sentryCtx)

		// * Capture critical error to sentry
		if critical {
			if len(args2) == 1 {
				if err, ok := args2[0].(error); ok {
					sentryHub.CaptureException(err)
				} else {
					sentryHub.CaptureException(errors.New(message))
				}
			} else {
				sentryHub.CaptureException(errors.New(message))
			}
		}
	}

	// * Return error instance
	if len(args2) == 1 {
		if code, ok := args2[0].(string); ok {
			return &ErrorInstance{
				Message: message,
				Code:    code,
				Err:     nil,
			}
		}
		if err, ok := args2[0].(error); ok {
			return &ErrorInstance{
				Message: message,
				Code:    "",
				Err:     err,
			}
		}
	}

	if len(args2) == 2 {
		if code, ok := args2[0].(string); ok {
			if err, ok := args2[1].(error); ok {
				return &ErrorInstance{
					Message: message,
					Code:    code,
					Err:     err,
				}
			}
		}
	}

	return &ErrorInstance{
		Message: message,
		Code:    "",
		Err:     nil,
	}
}
