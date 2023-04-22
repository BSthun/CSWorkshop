package main

import (
	"github.com/sirupsen/logrus"

	"backend/modules"
	"backend/modules/config"
	"backend/modules/db/init"
	"backend/modules/fiber"
	"backend/modules/firebase"
	"backend/modules/hub"
)

func main() {
	// * Initialize modules
	modules.Conf = iconfig.Init()
	modules.FirebaseApp, modules.FirebaseAuth = ifirebase.Init()
	modules.SqlDB, modules.DB = idbInit.Init()
	modules.Hub = ihub.Init(modules.SqlDB, modules.DB)
	modules.Fiber = ifiber.Init()

	// * Run the server
	err := modules.Fiber.Listen(modules.Conf.Address)
	if err != nil {
		logrus.Fatal(err.Error())
	}
}
