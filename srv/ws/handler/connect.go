package handler

import "gateway/utils/hub"

type Connect struct {
	Fd    hub.ConnectId `json:"fd"`
	Event int           `json:"event"`
}

func NewConnect(id hub.ConnectId) *Connect {
	return &Connect{
		Fd:    id,
		Event: 1,
	}
}
