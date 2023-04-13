package idbInit

import (
	"database/sql"
	"time"

	"github.com/sirupsen/logrus"

	"backend/modules"
)

func Connect() *sql.DB {
	// Open MySQL connection
	conn, err := sql.Open("mysql", modules.Conf.CoreMySqlDsn)
	if err != nil {
		logrus.WithField("e", err).Fatal("UNABLE TO OPEN MYSQL DATABASE")
	}

	// Configure connection pool
	conn.SetMaxIdleConns(10)
	conn.SetMaxOpenConns(100)
	conn.SetConnMaxLifetime(time.Hour)

	return conn
}
