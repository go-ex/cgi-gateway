package main

import (
	"flag"
	"gateway/constants"
	"gateway/srv/api"
	"gateway/srv/ws"
	"io/ioutil"
	"log"
	"os"
)

var addr = flag.String("addr", ":8080", "服务端口配置")
var conf = flag.String("conf", "config.json", "配置文件路径")

func init() {
	flag.Parse()

	_, err := os.Stat(*conf)
	if err == nil {
		if os.IsNotExist(err) == false {
			f, _ := os.Open(*conf)
			str, _ := ioutil.ReadAll(f)
			constants.NewConfig(str)

			log.Printf("load config file %s\n", *conf)
		}
	} else {
		log.Printf("无法加载配置文件 %s\n", *conf)
	}
}

func main() {
	servers := []constants.Server{
		ws.NewWebsocketServer(),
		api.NewHttpServer(),
	}

	for _, server := range servers {
		server.Start()
	}

	for _, server := range servers {
		go server.Run()
	}

	err := constants.Router.Run(*addr)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
