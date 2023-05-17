package main

import (
	"mock/modules/config"
	idb "mock/modules/db"
)

func main() {
	iconfig.Init()
	idb.Init()
}
