package ws

import "gateway/utils/hub"

type Client interface {
	Id() hub.ConnectId
	Send([]byte)
	WritePump()
	ReadPump()
}
