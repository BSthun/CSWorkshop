package hub

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"backend/modules/config"
	"backend/modules/db/branchModel"
	"backend/modules/db/model"
)

func Create(user *model.User) (string, *gorm.DB) {
	sqlDB, err := sql.Open("mysql", config.C.CoreMySqlDsn)
	if err != nil {
		logrus.WithField("e", err).Error("UNABLE TO LOAD MYSQL DATABASE")
		sentry.CaptureException(err)
		return "", nil
	}

	// Create database for user if not exists
	dbName := strings.Replace(config.C.BranchMySqlDbName, "{{ID}}", fmt.Sprintf("%03d", *user.Id), 1)
	dbDsn := strings.Replace(config.C.BranchMySqlDsn, "{{DB_NAME}}", dbName, 1)
	if _, err := sqlDB.Exec(fmt.Sprintf(rol"CREATE DATABASE IF NOT EXISTS %s", dbName)); err != nil {
		logrus.WithField("e", err).Error("UNABLE TO CREATE DATABASE FOR USER")
		sentry.CaptureException(err)
		return "", nil
	}

	if err := sqlDB.Close(); err != nil {
		logrus.WithField("e", err).Error("UNABLE TO CLOSE MYSQL DATABASE")
		sentry.CaptureException(err)
	}

	// Open SQL connection
	gormLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             100 * time.Millisecond,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	dialector := mysql.New(
		mysql.Config{
			DSN:               dbDsn,
			DefaultStringSize: 255,
		},
	)

	if db, err := gorm.Open(dialector, &gorm.Config{
		Logger: gormLogger,
	}); err != nil {
		logrus.WithField("e", err).Error("UNABLE TO LOAD GORM MYSQL DATABASE")
		sentry.CaptureException(err)
		return "", nil
	} else {
		if err := db.AutoMigrate(
			new(branchModel.Lv1Track),
			new(branchModel.Lv2Track),
			new(branchModel.Lv2Album),
			new(branchModel.Lv2Artist),
			new(branchModel.Lv3Track),
			new(branchModel.Lv3Album),
			new(branchModel.Lv3Artist),
			new(branchModel.Lv3AlbumArtist),
			new(branchModel.Lv3Activity),
		); err != nil {
			logrus.WithField("e", err).Error("UNABLE TO MIGRATE DATABASE")
		}

		return dbDsn, db
	}
}
