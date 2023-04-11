package idb

import (
	"database/sql"
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

	// Oprn MySQL connection
	conn, err := sql.Open("mysql", modules.Conf.CoreMySqlDsn)
	if err != nil {
		logrus.WithField("e", err).Fatal("UNABLE TO OPEN MYSQL DATABASE")
	}

	// Configure connection pool
	conn.SetMaxIdleConns(10)
	conn.SetMaxOpenConns(100)
	conn.SetConnMaxLifetime(time.Hour)

	// Open SQL connection
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
	if err := db.AutoMigrate(
		new(model.User),
	); err != nil {
		logrus.WithField("e", err).Fatal("UNABLE TO MIGRATE GORM MODEL")
	}

	logrus.Info("INITIALIZED MYSQL CONNECTION")

	return db
}
