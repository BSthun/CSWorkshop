package middlewares

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"

	"backend/modules"
	"backend/types/common"

	"backend/utils/network"
)

func SentryRecover(b *modules.Base, hub *sentry.Hub, c *fiber.Ctx, e *error) {
	if err := recover(); err != nil {
		eventID := hub.RecoverWithContext(
			context.WithValue(context.Background(), sentry.RequestContextKey, c),
			err,
		)
		if eventID != nil {
			hub.Flush(2 * time.Second)
		}
		if b.Conf.Environment == 1 {
			panic(err)
		}
		*e = fmt.Errorf("%v", err)
	}
}

func Sentry(b *modules.Base) fiber.Handler {
	return func(c *fiber.Ctx) (e error) {
		hub := sentry.CurrentHub().Clone()
		scope := hub.Scope()
		scope.SetRequest(network.ConvertRequest(c.Context()))
		scope.SetRequestBody(c.Body())

		// Configure scope
		hub.ConfigureScope(func(scope *sentry.Scope) {
			if c.Locals("l") != nil {
				claims := c.Locals("l").(*jwt.Token).Claims.(*common.UserClaims)
				scope.SetUser(sentry.User{
					ID:        strconv.FormatUint(uint64(*claims.UserId), 10),
					Email:     "",
					IPAddress: c.Get("X-Forwarded-For", c.IP()),
					Username:  "",
					Name:      "",
					Segment:   "",
					Data:      nil,
				})
			}
		})

		sentryCtx := context.TODO()
		sentryCtx = sentry.SetHubOnContext(sentryCtx, hub)

		// Start a transaction
		span := sentry.StartSpan(sentryCtx, "http.server", sentry.TransactionName(fmt.Sprintf("%s (%s)", c.Path(), c.Method())))
		sentryCtx = context.WithValue(sentryCtx, "span", span)
		c.Locals("sentry", sentryCtx)

		defer SentryRecover(b, hub, c, &e)
		return c.Next()
	}
}
