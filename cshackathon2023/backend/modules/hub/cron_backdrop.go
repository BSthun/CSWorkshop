package ihub

import (
	"backend/functions/backdrop"
	"backend/modules"
	"backend/types/extern"
)

func CronBackdrop() {
	state := backdrop.GetBackdrop()

	modules.Hub.BackdropClient.Emit(&extern.OutboundMessage{
		Event:   extern.BackdropUpdateEvent,
		Payload: state,
	})
}
