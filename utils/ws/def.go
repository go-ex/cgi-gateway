package ws

import "gateway/utils/hub"

type Client interface {
	App() int
	Id() hub.ConnectId
	Send([]byte)
	WritePump()
	ReadPump()
}
