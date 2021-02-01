package handler

import "gateway/utils/hub"

type Close struct {
	Fd    hub.ConnectId `json:"fd"`
	Event int           `json:"event"`
}

func NewClose(id hub.ConnectId) *Connect {
	return &Connect{
		Fd:    id,
		Event: 3,
	}
}
