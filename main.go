package main

import (
	"fmt"
	"net/http"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"strings"
	"os"
	"os/exec"
	"strconv"
	"sync"
)

var mutex = &sync.Mutex{}

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

	if strings.HasSuffix(filename, ".mp3") {

		log.WithField("filename", filename).Info("playing sound")
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			log.Error(filename)
			w.WriteHeader(404)
			return
		}
		http.Redirect(w, r, "/", 307)
		cmd := exec.Command("omxplayer", "-o", "hdmi", filename)
		err := cmd.Start()
		if err != nil {
			log.Error(err)
		} else {
			incCounter(filename)
		}

	} else {
		log.Error("no .mp3 suffix")
	}
}

func incCounter(filename string) {

	log.WithField("filename", filename).Info("Increasing counter")

	mutex.Lock()
	intCount := getCounter(filename)
	log.WithField("count", intCount).Info("Old count")
	intCount++

	ioutil.WriteFile(filename + ".count", []byte(fmt.Sprintf("%d", intCount)), 0644)
	mutex.Unlock()
}

func getCounter(filename string) int {

	countfile := filename + ".count"

	if _, err := os.Stat(countfile); os.IsNotExist(err) {
		return 0
	}

	count, err := ioutil.ReadFile(countfile)
	intCount, err := strconv.Atoi(string(count))
	if err != nil {
		log.Error(err)
	}
	return intCount
}


func handleIndex(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "<html><body style=\"font-family: arial, helvetica\">")

	dir, err := ioutil.ReadDir("sounds")
	if err != nil {
		log.Error(err)
	}

	for _, v := range dir {
		if !v.IsDir() {
			filename := v.Name()
			if (strings.HasSuffix(filename, ".mp3")) {
				fmt.Fprintf(w, "<div style=\"border: 1px solid black; margin: 3px; padding: 3px; float: left\"><div><a href=\"/play/" + filename + "\">")

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

				fmt.Fprintf(w, "</a></div><div style=\"padding-top: 3px;\">played %d times</div></div>", getCounter("sounds/" + filename))
			}
		}
	}

	fmt.Fprint(w, "</body></html>")
}

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}