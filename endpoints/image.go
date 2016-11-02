package endpoints

import (
	"os"
	"net/http"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
)

func InitImage() {
	http.HandleFunc("/api/image/", imageHandler)
}

func imageHandler(w http.ResponseWriter, r *http.Request) {
	filename := "sounds/" + r.RequestURI[11:]
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		log.Error(filename)
		w.WriteHeader(404)
		return
	}
	image, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Error(err)
		w.WriteHeader(500)
	}
	w.Write(image)
}
