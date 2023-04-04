package db

import (
	"encoding/base64"
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"gorm.io/gorm"

	"backend/modules/db/model"
	"backend/utils/value"
)

func preset() {
	var clientCount int64
	if result := DB.Model(new(model.Client)).Count(&clientCount); result.Error != nil {
		logrus.Fatal("UNABLE TO CHECK GAME COUNT FOR PRESET MIGRATION")
	}

	if clientCount == 0 {
		var data struct {
			Clients []*model.Client `yaml:"clients"`
		}
		if bytes, err := os.ReadFile("./clients.yaml"); err != nil {
			logrus.WithField("e", err).Fatal("UNABLE TO READ CLIENT FILE")
		} else {
			if err := yaml.Unmarshal(bytes, &data); err != nil {
				logrus.WithField("e", err).Fatal("UNABLE TO PARSE CLIENT FILE")
			}
		}

		if err := DB.Transaction(func(tx *gorm.DB) error {
			for _, client := range data.Clients {
				client.Authorization = value.Ptr(base64.StdEncoding.EncodeToString([]byte(*client.SpotifyClientId + ":" + *client.SpotifyClientSecret)))
				if result := tx.Create(client); result.Error != nil {
					return result.Error
				}
			}
			return nil
		}); err != nil {
			logrus.WithField("e", err).Fatal("UNABLE TO INSERT CLIENT PRESETS")
		}
	}
}
