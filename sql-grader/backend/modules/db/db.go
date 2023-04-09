package idb

import (
	"log"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"backend/modules"
	"backend/types/model"
)

func Init(b *modules.Base) {
	// Initialize GORM instance using previously opened SQL connection
	gormLogLevel := []logger.LogLevel{
		logger.Silent,
		logger.Error,
		logger.Error,
		logger.Warn,
		logger.Info,
		logger.Info,
		logger.Info,
	}

	gormLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             100 * time.Millisecond,
			LogLevel:                  gormLogLevel[b.Conf.LogLevel],
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	// Open SQL connection
	dialector := mysql.New(
		mysql.Config{
			DSN:               b.Conf.CoreMySqlDsn,
			DefaultStringSize: 255,
		},
	)

	if db, err := gorm.Open(dialector, &gorm.Config{
		Logger: gormLogger,
	}); err != nil {
		logrus.WithField("e", err).Fatal("UNABLE TO LOAD GORM MYSQL DATABASE")
	} else {
		b.DB = db
	}

	// Initialize model migrations
	if err := b.DB.AutoMigrate(
		new(model.User),
	); err != nil {
		logrus.WithField("e", err).Fatal("UNABLE TO MIGRATE GORM MODEL")
	}

	logrus.Info("INITIALIZED MYSQL CONNECTION")
}
