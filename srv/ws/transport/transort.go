package transport

type Transport interface {
	NewRequest(url string) Request
}

type Request interface {
	Call(message []byte) error
}

func NewTransport() Transport {
	return NewHttpTransport()
}
