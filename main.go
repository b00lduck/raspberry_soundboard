package main

import (
	"fmt"
	"net/http"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"strings"
	"os"
	"os/exec"
)

func handler(w http.ResponseWriter, r *http.Request) {

	log.WithField("requestURi", r.RequestURI).Info("HTTP")

	if strings.HasPrefix(r.RequestURI, "/play") {
		handlePlay(w, r)
	} else if strings.HasPrefix(r.RequestURI, "/images") {
		handleImage(w, r)
	} else {
		handleIndex(w, r)
	}

}

func handleImage(w http.ResponseWriter, r *http.Request) {
	filename := "sounds/" + r.RequestURI[8:]
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

func handlePlay(w http.ResponseWriter, r *http.Request) {
	filename := "sounds/" + r.RequestURI[6:]
	log.WithField("filename", filename).Info("playing sound")
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		log.Error(filename)
		w.WriteHeader(404)
		return
	}
	cmd := exec.Command("omxplayer", "-o", "hdmi", filename)
	err := cmd.Run()
	if err != nil {
		log.Error(err)
	}
	http.Redirect(w, r, "/", 307)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	dir, err := ioutil.ReadDir("sounds")
	if err != nil {
		log.Error(err)
	}

	for _, v := range dir {
		if !v.IsDir() {
			filename := v.Name()
			if (strings.HasSuffix(filename, ".mp3")) {

				fmt.Fprintf(w, "<a href=\"/play/" + filename + "\">")

				filenameWithoutExt := filename[:len(filename)-4]

				pngFilename := filenameWithoutExt + ".png"
				if _, err := os.Stat("sounds/" + pngFilename); os.IsNotExist(err) {
					jpgFilename := filenameWithoutExt + ".jpg"
					if _, err := os.Stat("sounds/" + jpgFilename); os.IsNotExist(err) {
						fmt.Fprintf(w, filename)
					} else {
						fmt.Fprintf(w, "<img src=\"/images/" + jpgFilename + "\">")
					}
				} else {
					fmt.Fprintf(w, "<img src=\"/images/" + pngFilename + "\">")
				}

				fmt.Fprintf(w, "</a>")
			}
		}
	}
}

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}