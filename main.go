package main

import (
	"net/http"
	log "github.com/Sirupsen/logrus"
	"github.com/b00lduck/raspberry_soundboard/endpoints"
	"github.com/b00lduck/raspberry_soundboard/persistence"
)



func main() {

	syncChan := make(chan int)

	persistence.Init()

	endpoints.InitImage()
	endpoints.InitPlay(syncChan)
	endpoints.InitWebsocket(syncChan)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}

}