package ws

import (
	"gateway/utils/ws"
	"log"
)

type event struct {
}

func (e *event) OnConnect(client ws.Client) {
	log.Println("OnConnect")
}

func (e *event) OnMessage(client ws.Client, message []byte) {
	log.Println(message)
}

func (e *event) OnClose(client ws.Client) {
	log.Println("OnClose")
}

func (e *event) OnError(client ws.Client, err error) {
	log.Println("OnError")
}

var ev ws.Event

func GetEvent() ws.Event {
	if ev == nil {
		ev = &event{}
	}
	return ev
}
