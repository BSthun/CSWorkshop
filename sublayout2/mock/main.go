package main

import (
	"mock/commands"
	"mock/modules"
	"mock/modules/config"
	idb "mock/modules/db"
)

func main() {
	modules.Conf = iconfig.Init()
	modules.DB = idb.Init()
	commands.Run()
}
