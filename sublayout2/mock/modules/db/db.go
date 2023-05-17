package idb

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"mock/modules"
	"mock/types/model"
	"os"
	"time"
)

func Init() *gorm.DB {
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
			LogLevel:                  gormLogLevel[modules.Conf.LogLevel],
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	// Open SQL connection
	conn := Connect()
	dialector := mysql.New(
		mysql.Config{
			Conn: conn,
		},
	)
	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		logrus.WithField("e", err).Fatal("UNABLE TO OPEN GORM MYSQL DATABASE")
	}

	// Initialize model migrations
	if modules.Conf.MysqlMigrate {
		if err := db.AutoMigrate(
			new(model.User),
			new(model.Country),
			new(model.Album),
			new(model.Artist),
			new(model.AlbumArtist),
			new(model.Track),
			new(model.PlaylistTrack),
			new(model.Concert),
			new(model.ConcertTicket),
		); err != nil {
			logrus.WithField("e", err).Fatal("UNABLE TO MIGRATE GORM MODEL")
		}
	}

	logrus.Info("INITIALIZED MYSQL CONNECTION")

	return db
}
