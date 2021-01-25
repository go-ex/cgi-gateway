package main

import (
	"flag"
	"gateway/constants"
	"gateway/srv/api"
	"gateway/srv/ws"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8080", "api service address")

func main() {
	flag.Parse()

	servers := []constants.Server{
		ws.NewWebsocketServer(),
		api.NewHttpServer(),
	}

	for _, server := range servers {
		server.Start()
	}

	for _, server := range servers {
		server.Run()
	}

	log.Println("ListenAndServe: ", *addr)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
