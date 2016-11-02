// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package websocket

import (
	"github.com/b00lduck/raspberry_soundboard/persistence"
	log "github.com/Sirupsen/logrus"
)

// hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) BroadcastState() {
	h.broadcast <- persistence.JsonState()
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			log.Info("register client")
			h.clients[client] = true
			client.send <- persistence.JsonState()
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				log.Info("unregister client")
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			log.Info("broadcast to all clients")
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
