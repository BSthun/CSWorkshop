package extern

type OutboundMessage struct {
	Event   OutboundEvent `json:"event"`
	Payload any           `json:"payload"`
}

type OutboundEvent string

const (
	ConnectionSwitchEvent OutboundEvent = "general/switch"
	EchoEvent             OutboundEvent = "general/echo"
	PingEvent             OutboundEvent = "general/ping"
	ErrorEvent            OutboundEvent = "general/error"
	BackdropUpdateEvent   OutboundEvent = "backdrop/update"
)
