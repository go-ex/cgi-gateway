package handler

import "gateway/utils/hub"

type message struct {
	Fd      hub.ConnectId `json:"fd"`
	Event   int           `json:"event"`
	Payload string        `json:"payload"`
}

func NewMessage(id hub.ConnectId, msg []byte) *message {
	return &message{
		Fd:      id,
		Event:   2,
		Payload: string(msg),
	}
}
