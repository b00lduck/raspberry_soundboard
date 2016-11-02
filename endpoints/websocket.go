package endpoints

import (
	"net/http"
	"github.com/b00lduck/raspberry_soundboard/websocket"
)

func InitWebsocket() *websocket.Hub {
	hub := websocket.NewHub()
	go hub.Run()
	http.HandleFunc("/api/websocket", func(w http.ResponseWriter, r *http.Request) {
		websocket.ServeWs(hub, w, r)
	})
	return hub
}
