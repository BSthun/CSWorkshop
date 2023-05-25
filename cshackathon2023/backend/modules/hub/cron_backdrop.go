package ihub

import (
	"backend/functions/backdrop"
	"backend/types/extern"
)

func CronBackdrop() {
	state := backdrop.GetBackdrop()

	for _, backdropClientConnection := range Hub.BackdropClientConnections {
		// * Send message to backdrop client
		_ = backdropClientConnection.WriteJSON(&extern.OutboundMessage{
			Event:   extern.BackdropUpdateEvent,
			Payload: state,
		})
	}
}
