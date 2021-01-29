package hub

type Client interface {
	Id() ConnectId
	Send([]byte)
}

type ConnectId string
