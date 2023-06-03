package iconfig

import (
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"

	"backend/utils/text"
)

func Init() *Config {
	// Initialize blank configuration struct
	conf := new(Config)

	// Load configurations to struct
	yml, err := os.ReadFile("config.yaml")
	if err != nil {
		logrus.Fatal("UNABLE TO READ YAML CONFIGURATION FILE")
	}
	if err := yaml.Unmarshal(yml, conf); err != nil {
		logrus.Fatal("UNABLE TO PARSE YAML CONFIGURATION FILE")
	}

	// Validate configurations
	if err := text.Validator.Struct(conf); err != nil {
		logrus.Fatal("INVALID CONFIGURATION: " + err.Error())
	}

	// Apply log level configuration
	logrus.SetLevel(logrus.Level(conf.LogLevel))
	spew.Config = spew.ConfigState{Indent: "  "}

	return conf
}
