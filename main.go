package main

import (
	"net/http"
	log "github.com/Sirupsen/logrus"
	"github.com/b00lduck/raspberry_soundboard/endpoints"
	"github.com/b00lduck/raspberry_soundboard/persistence"
)



func main() {

	persistence.Init()

	endpoints.InitImage()
	hub := endpoints.InitWebsocket()
	endpoints.InitPlay(hub)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}

}