package ws

import (
	"gateway/constants"
	"gateway/srv/ws/handler"
	"gateway/srv/ws/transport"
	"gateway/utils/json"
	"gateway/utils/ws"
	"log"
)

type event struct {
	transport transport.Transport
}

func (e *event) OnConnect(client ws.Client) {
	url := constants.GetAppAgent(client.App())
	request := e.transport.NewRequest(url)
	err := request.Call(json.StructToString(handler.NewConnect(client.Id())))

	if err != nil {
		log.Println(err)
		handler.NotServer(client)
	}
}

func (e *event) OnMessage(client ws.Client, message []byte) {
	url := constants.GetAppAgent(client.App())
	request := e.transport.NewRequest(url)
	err := request.Call(message)

	if err != nil {
		log.Println(err)
		handler.NotServer(client)
	}
}

func (e *event) OnClose(client ws.Client) {
	url := constants.GetAppAgent(client.App())
	request := e.transport.NewRequest(url)
	err := request.Call(json.StructToString(handler.NewClose(client.Id())))

	if err != nil {
		log.Println(err)
		handler.NotServer(client)
	}
}

func (e *event) OnError(client ws.Client, err error) {
	log.Println("OnError")
}

var ev ws.Event

func GetEvent() ws.Event {
	if ev == nil {
		ev = &event{
			transport: transport.NewTransport(),
		}
	}
	return ev
}
