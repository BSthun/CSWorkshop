package main

import (
	"github.com/sirupsen/logrus"

	"backend/modules"
)

func main() {
	// * Initialize modules
	Init()

	// * Run the server
	err := modules.Fiber.Listen(modules.Conf.Address)
	if err != nil {
		logrus.Fatal(err.Error())
	}
}
