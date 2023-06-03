package websocket

import (
	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/websocket/v2"
)

func Register(router fiber.Router) {
	router.Use("/", func(c *fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	router.Get("/music", websocket.New(func(conn *websocket.Conn) {
		// c.Locals are added to the *websocket.Conn
		// log.Println(conn.Locals("allowed"))  // true
		// log.Println(conn.Params("id"))       // 123
		// log.Println(conn.Query("v"))         // 1.0
		// log.Println(conn.Cookies("session")) // ""

		// websocket.Conn bindings https://pkg.go.dev/github.com/fasthttp/websocket?tab=doc#pkg-index
		ServeMusicState(conn)
	}))

	router.Get("/backdrop", websocket.New(func(conn *websocket.Conn) {
		// c.Locals are added to the *websocket.Conn
		// log.Println(conn.Locals("allowed"))  // true
		// log.Println(conn.Params("id"))       // 123
		// log.Println(conn.Query("v"))         // 1.0
		// log.Println(conn.Cookies("session")) // ""

		// websocket.Conn bindings https://pkg.go.dev/github.com/fasthttp/websocket?tab=doc#pkg-index
		ServeBackdropState(conn)
	}))
}
