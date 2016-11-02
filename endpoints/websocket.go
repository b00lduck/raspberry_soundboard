package endpoints

import (
	"github.com/gorilla/websocket"
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/b00lduck/raspberry_soundboard/persistence"
	log "github.com/Sirupsen/logrus"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 2048,
	WriteBufferSize: 2048,
	EnableCompression: true,
	CheckOrigin: func(r *http.Request) bool { return true },
}

var numClients = 0

func InitWebsocket(syncChan chan int) {
	http.HandleFunc("/api/websocket", func(w http.ResponseWriter, r *http.Request) {
		wsConn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer wsConn.Close()

		numClients++
		log.WithField("numClients", numClients).Info("Client subscribed")

		for {
			soundList := persistence.State()

			myJson, err := json.Marshal(soundList)
			if err != nil {
				log.Error(err)
				break
			}
			log.Info("Delivering state to client")
			err = wsConn.WriteMessage(websocket.TextMessage, myJson)
			if err != nil {
				log.Error(err)
				break
			}
			<-syncChan
		}

		numClients--
		log.WithField("numClients", numClients).Info("Client unsubscribed")

	})
}