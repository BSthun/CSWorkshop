package ihub

import (
	"github.com/sirupsen/logrus"

	"backend/types/model"
)

func InitialState(session *Session) {
	// * Query enrollment
	var enrollment *model.Enrollment
	if err := b.DB.Find(&enrollment, session.Id); err != nil {
		logrus.Warn(err)
		return
	}
}
