package main

import (
	"backend/modules/config"
	"backend/modules/db"
	"backend/modules/fiber"
	"backend/modules/sentry"
)

func main() {
	config.Init()
	sentry.Init()
	db.Init()
	//hub.Init()
	fiber.Init()
}
