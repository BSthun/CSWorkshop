package ihub

import (
	"github.com/sirupsen/logrus"

	"backend/types/extern"
	"backend/types/model"
	"backend/types/payload"
	"backend/utils/value"
)

func InitialState(session *Session) {
	// * Query enrollment
	var enrollment *model.Enrollment
	if result := b.DB.First(&enrollment, session.Id); result.Error != nil {
		logrus.Warn(result.Error)
		return
	}

	// * Handle invalid db state
	if *enrollment.DbValid == false {
		// * Send initial message
		_ = session.Conn.WriteJSON(&extern.OutboundMessage{
			Event: extern.LabStateEvent,
			Payload: &payload.LabState{
				DbValid:         value.Ptr(false),
				TaskTitle:       nil,
				TaskDescription: nil,
				TaskTags:        nil,
				Query:           nil,
				QueryPassed:     nil,
				QueryError:      nil,
				Result:          nil,
			},
		})
		return
	}

	// * Otherwise, send initial message
	_ = session.Conn.WriteJSON(&extern.OutboundMessage{
		Event: extern.LabStateEvent,
		Payload: &payload.LabState{
			DbValid:         value.Ptr(true),
			TaskTitle:       nil,
			TaskDescription: nil,
			TaskTags:        nil,
			Query:           nil,
			QueryPassed:     nil,
			QueryError:      nil,
			Result:          nil,
		},
	})
}
