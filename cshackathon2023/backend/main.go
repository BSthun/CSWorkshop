package main

import (
	"backend/modules"
	"backend/modules/config"
	"backend/modules/db"
	"backend/modules/fiber"
	"backend/modules/firebase"
	"backend/modules/hub"
	"github.com/sirupsen/logrus"
)

func main() {
	// * Initialize modules
	modules.Conf = iconfig.Init()
	modules.FirebaseApp, modules.FirebaseAuth = ifirebase.Init()
	modules.SqlDB, modules.DB = idb.Init()
	modules.Hub = ihub.Init()
	modules.Fiber = ifiber.Init()

	// * Run the server
	err := modules.Fiber.Listen(modules.Conf.Address)
	if err != nil {
		logrus.Fatal(err.Error())
	}
}
