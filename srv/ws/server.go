package ws

import (
	"gateway/constants"
	"gateway/utils/hub"
	"gateway/utils/ws"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
)

type Server struct {
	hub *hub.Hub
}

func (s *Server) Start() {
	s.hub = hub.GetHub()

	go s.hub.Run()

	constants.Router.GET("/", func(c *gin.Context) {
		// 确定app
		app, _ := strconv.Atoi(c.DefaultQuery("app", "1"))

		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Println(err)
			return
		}

		client := ws.NewClient(app, conn, GetEvent())
		s.hub.Register(client)

		// Allow collection of memory referenced by the caller by doing all work in
		// new goroutines.
		go client.WritePump()
		go client.ReadPump()
	})
}

func (s *Server) Run() {

}

func NewWebsocketServer() constants.Server {
	return &Server{}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
