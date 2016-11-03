package websocket

import (
	log "github.com/Sirupsen/logrus"
	"github.com/b00lduck/raspberry_soundboard/persistence"
)

type Hub struct {
	clients map[*Client]bool
	broadcast chan bool
	register chan *Client
	unregister chan *Client
	persistence *persistence.Persistence
}

func NewHub(persistence *persistence.Persistence) *Hub {
	return &Hub{
		broadcast:  make(chan bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
		persistence: persistence,
	}
}

func (h *Hub) Broadcast() {
	h.broadcast <- true
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			log.Info("register client")
			h.clients[client] = true
			client.send <- h.persistence.JsonState()
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				log.Info("unregister client")
				delete(h.clients, client)
				close(client.send)
			}
		case <-h.broadcast:
			log.Info("broadcast to all clients")
			for client := range h.clients {
				select {
				case client.send <- h.persistence.JsonState():
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
