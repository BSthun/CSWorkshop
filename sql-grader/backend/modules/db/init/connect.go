package idbInit

import (
	"database/sql"
	"strings"
	"time"

	"github.com/sirupsen/logrus"

	"backend/modules"
)

func Connect() *sql.DB {
	// Open MySQL connection
	dsn := strings.Replace(modules.Conf.MysqlDsn, "{{DB_NAME}}", modules.Conf.MysqlDb, 1)
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		logrus.WithField("e", err).Fatal("UNABLE TO OPEN MYSQL DATABASE")
	}

	// Configure connection pool
	conn.SetMaxIdleConns(10)
	conn.SetMaxOpenConns(100)
	conn.SetConnMaxLifetime(time.Hour)

	return conn
}
