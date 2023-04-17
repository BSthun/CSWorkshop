package enroll

import (
	"backend/modules"
	ihub "backend/modules/hub"
	"backend/types/model"
	"backend/utils/text"
)

func ActLoadEnrollmentSession(enrollment *model.Enrollment) (*ihub.Session, error) {
	// * Check current enrollment session
	session, ok := modules.Hub.Sessions[*enrollment.Id]
	if !ok {
		// * Create new session
		session = &ihub.Session{
			Id:         enrollment.Id,
			Credential: enrollment.User.Credential,
			DbName:     enrollment.DbName,
			Token:      text.Random(text.RandomSet.UpperAlphaNum, 16),
			Conn:       nil,
			ConnMutex:  nil,
		}
		modules.Hub.Sessions[*enrollment.Id] = session
		modules.Hub.SessionDbNameMap[*enrollment.DbName] = session
	}

	return session, nil
}
