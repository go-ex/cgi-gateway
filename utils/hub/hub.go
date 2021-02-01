package hub

// 信息总线
type Hub struct {
	// 记录所有的连接
	clients map[ConnectId]Client

	// 全局定义的广播
	broadcast chan []byte

	// 注册线
	register chan Client

	// 取消线
	unregister chan Client
}

var hub *Hub

// 获取总线单例
func GetHub() *Hub {
	if hub == nil {
		hub = &Hub{
			broadcast:  make(chan []byte),
			register:   make(chan Client),
			unregister: make(chan Client),
			clients:    make(map[ConnectId]Client),
		}
	}

	return hub
}

func (h *Hub) Unregister(c Client) {
	h.unregister <- c
}

func (h *Hub) Register(c Client) {
	h.register <- c
}

func (h *Hub) SendAll(msg []byte) {
	h.broadcast <- msg
}

func (h *Hub) Send(id ConnectId, msg []byte) {
	client, ok := h.clients[id]
	if ok {
		client.Send(msg)
	}
}

func (h *Hub) Close(id ConnectId) {
	client, ok := h.clients[id]
	if ok {
		client.Close()

		delete(h.clients, client.Id())
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client.Id()] = client
		case client := <-h.unregister:
			if _, ok := h.clients[client.Id()]; ok {
				delete(h.clients, client.Id())
			}
		case message := <-h.broadcast:
			for _, client := range h.clients {
				client.Send(message)
			}
		}
	}
}
