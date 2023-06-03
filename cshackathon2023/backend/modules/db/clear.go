package idb

import "github.com/sirupsen/logrus"

func Clear() {
	// * Get connection
	conn := Connect()

	// * Get database name
	var dbName string
	if err := conn.QueryRow("SELECT DATABASE()").Scan(&dbName); err != nil {
		logrus.WithField("e", err).Fatal("UNABLE TO GET DATABASE NAME")
	}

	// * Drop database
	if _, err := conn.Exec("DROP DATABASE " + dbName); err != nil {
		logrus.WithField("e", err).Fatal("UNABLE TO DROP DATABASE")
	}

	// * Recreate database
	if _, err := conn.Exec("CREATE DATABASE " + dbName); err != nil {
		logrus.WithField("e", err).Fatal("UNABLE TO CREATE DATABASE")
	}

	logrus.Info("CLEARED MYSQL DATABASE")
}
