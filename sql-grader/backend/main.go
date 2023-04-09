package main

import (
	"backend/modules"
	"backend/modules/config"
	"backend/modules/db"
	"backend/modules/fiber"
	"backend/modules/firebase"
)

func main() {
	base := modules.NewBase()
	base.Conf = iconfig.Init()
	ifirebase.Init(base)
	idb.Init(base)
	fiber.Init(base)
}
