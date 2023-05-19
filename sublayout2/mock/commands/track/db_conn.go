package track

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func DbConn(dsn string) {
	gormLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             100 * time.Millisecond,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	// Open SQL connection
	dialector := mysql.Open(dsn)
	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: gormLogger,
	})

	if err != nil {
		log.Fatal(err)
	}

	Fetch(db)
}
