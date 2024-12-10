package shared

import (
	"sync"

	"golang.org/x/net/websocket"
)

var (
	websocketClients = make(map[*websocket.Conn]bool)
	mu               = &sync.RWMutex{}
)

func AddWebsocketClient(client *websocket.Conn) {
	mu.Lock()
	defer mu.Unlock()
	if _, exists := websocketClients[client]; !exists {
		websocketClients[client] = true
	}
}

func DeleteWebsocketClient(client *websocket.Conn) {
	mu.Lock()
	defer mu.Unlock()
	delete(websocketClients, client)
}

func GetWebsocketClients() map[*websocket.Conn]bool {
	mu.Lock()
	defer mu.Unlock()
	return websocketClients
}
