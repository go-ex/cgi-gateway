package handler

import (
	"gateway/utils/json"
	"gateway/utils/ws"
	"log"
)

type notServer struct {
	Event int    `json:"event"`
	Msg   string `json:"msg"`
}

func NotServer(client ws.Client) {
	log.Println("api 服务离线")

	handler := &notServer{
		Event: 0,
		Msg:   "api 服务离线",
	}
	client.Send(json.StructToString(handler))
}
