package main

import (
	"net/http"
	log "github.com/Sirupsen/logrus"
	"github.com/b00lduck/raspberry_soundboard/endpoints"
	"github.com/b00lduck/raspberry_soundboard/persistence"
	"github.com/b00lduck/raspberry_soundboard/websocket"
	"github.com/b00lduck/raspberry_soundboard/physics"
)

func main() {

	persistence := persistence.NewPersistence()
	hub := websocket.NewHub(persistence)
	persistence.PersistCallback = hub.Broadcast
	go hub.Run()
	http.HandleFunc("/api/websocket", func(w http.ResponseWriter, r *http.Request) {
		websocket.ServeWs(hub, w, r)
	})

	endpoints.InitImage()
	endpoints.InitPlay(persistence)
	http.Handle("/", http.FileServer(http.Dir("frontend/build")))

	go physics.Process(persistence)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}

}