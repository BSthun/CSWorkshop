package iconfig

import (
	"flag"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"

	"backend/utils/text"
)

func Init() *Config {
	// Initialize blank configuration struct
	conf := new(Config)

	// Check testing config
	var file = "config.yaml"
	if flag.Lookup("test.v") != nil {
		file = text.RelativePath("config.test.yaml")
	}

	// Load configurations to struct
	yml, err := os.ReadFile(file)
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
