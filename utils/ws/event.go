package ws

type Event interface {
	OnConnect(client Client)
	OnMessage(client Client, message []byte)
	OnClose(client Client)
	OnError(client Client, err error)
}
