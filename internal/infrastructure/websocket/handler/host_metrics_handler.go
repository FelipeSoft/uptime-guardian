package handler

import (
	"encoding/json"
	"fmt"

	"github.com/FelipeSoft/uptime-guardian/internal/infrastructure/shared"
	"golang.org/x/net/websocket"
)

func HostMetricsWebsocketHandler(ws *websocket.Conn) {
	shared.AddWebsocketClient(ws)
	handshakeMsg, err := json.Marshal(map[string]interface{}{"message": "hello!"})
	if err != nil {
		fmt.Printf("fail on marshal: %s \n", err.Error())
	}
	err = websocket.Message.Send(ws, handshakeMsg)
	if err != nil {
		fmt.Printf("fail on websocket connection: %s \n", err.Error())
		shared.DeleteWebsocketClient(ws)
	}
}
