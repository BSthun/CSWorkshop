package ihub

import "backend/modules/config"

type Hub struct {
	Conf *iconfig.Config
}

func Init(conf *iconfig.Config) *Hub {
	hub := &Hub{
		Conf: conf,
	}

	return hub
}
