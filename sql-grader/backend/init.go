package main

import (
	"backend/modules"
	"backend/modules/config"
	"backend/modules/db"
	"backend/modules/fiber"
	"backend/modules/firebase"
	"backend/modules/hub"
)

func Init() {
	modules.Conf = iconfig.Init()
	modules.Hub = ihub.Init(modules.Conf)
	modules.FirebaseApp, modules.FirebaseAuth = ifirebase.Init()
	modules.DB = idb.Init()
	modules.Fiber = ifiber.Init()
}
