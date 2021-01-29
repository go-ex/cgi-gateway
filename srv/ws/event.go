package ws

import (
	"bytes"
	"gateway/constants"
	"gateway/srv/ws/handler"
	"gateway/utils/json"
	"gateway/utils/ws"
	"log"
	"net/http"
)

type event struct {
}

func (e *event) OnConnect(client ws.Client) {
	url := constants.GetAppAgent(client.App())

	resp, err := http.Post(
		url,
		"application/json",
		bytes.NewReader(json.StructToString(&handler.Connect{Event: 1})),
	)

	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Println("api 服务离线")
		client.Send([]byte("{\"event\":0,\"msg\":\"服务离线\""))
	}
}

func (e *event) OnMessage(client ws.Client, message []byte) {
	url := constants.GetAppAgent(client.App())

	resp, err := http.Post(
		url,
		"application/json",
		bytes.NewReader(message),
	)

	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Println("api 服务离线")
		client.Send([]byte("{\"event\":0,\"msg\":\"服务离线\""))
	}
}

func (e *event) OnClose(client ws.Client) {
	url := constants.GetAppAgent(client.App())

	resp, err := http.Post(
		url,
		"application/json",
		bytes.NewReader(json.StructToString(&handler.Close{Event: 3})),
	)

	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Println("api 服务离线")
		client.Send([]byte("{\"event\":0,\"msg\":\"服务离线\""))
	}
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
