package hub

type Client interface {
	Id() ConnectId
	Send([]byte)
	Close()
}

type ConnectId string
