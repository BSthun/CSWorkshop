package enroll

import (
	"database/sql"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"

	"backend/modules"
	ihub "backend/modules/hub"
	"backend/types/model"
	"backend/types/payload"
	"backend/utils/text"
)

func ActLoadEnrollmentSession(enrollment *model.Enrollment) (*ihub.Session, error) {
	// * Check current enrollment session
	session, ok := modules.Hub.Sessions[*enrollment.Id]
	if !ok {
		// * Create new database connection
		dsn := strings.Replace(modules.Conf.MysqlDsn, "{{DB_NAME}}", *enrollment.DbName, 1)
		// [UNUSED] Switching to user
		// dsn = fmt.Sprintf("%s:%s@%s", *enrollment.User.Credential.Username, *enrollment.User.Credential.Password, strings.Split(dsn, "@")[1])
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			logrus.WithField("e", err).Fatal("UNABLE TO OPEN MYSQL DATABASE")
			return nil, err
		}

		// * Create new session
		session = &ihub.Session{
			Id:          enrollment.Id,
			LabId:       enrollment.Lab.Id,
			UserId:      enrollment.User.Id,
			Db:          db,
			DbName:      enrollment.DbName,
			DbValid:     enrollment.DbValid,
			Token:       text.Random(text.RandomSet.UpperAlphaNum, 16),
			CurrentTask: nil,
			TaskResults: make(map[uint64]*payload.LabStateResult),
			Conn:        nil,
			ConnMutex:   new(sync.Mutex),
		}
		modules.Hub.Sessions[*enrollment.Id] = session
		modules.Hub.SessionDbNameMap[*enrollment.DbName] = session
	}

	return session, nil
}
