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
	"html/template"
	"github.com/b00lduck/raspberry_soundboard/templates"
	"math/rand"
	"sort"
)

var mutex = &sync.Mutex{}

type Sound struct {
	SoundFile string
	ImageFile string
	HasImage  bool
	Count     int
}

func handler(w http.ResponseWriter, r *http.Request) {

	log.WithField("requestURi", r.RequestURI).Info("HTTP")

	if strings.HasPrefix(r.RequestURI, "/play") {
		handlePlay(w, r)
	} else if strings.HasPrefix(r.RequestURI, "/random") {
		handleRandomPlay(w, r)
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

func handleRandomPlay(w http.ResponseWriter, r *http.Request) {

	sounds := getSounds()

	filename := "sounds/" + sounds[rand.Intn(len(sounds))].SoundFile

	err := play(filename)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	http.Redirect(w, r, "/", 307)
}

func handlePlay(w http.ResponseWriter, r *http.Request) {

	filename := "sounds/" + r.RequestURI[6:]

	err := play(filename)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	http.Redirect(w, r, "/", 307)

}

func play(filename string) error {
	if strings.HasSuffix(filename, ".mp3") {

		log.WithField("filename", filename).Info("playing sound")
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			log.Error(filename)
			return fmt.Errorf("Not found")

		}
		incCounter(filename)
		go func() {
			cmd := exec.Command("omxplayer", "-o", "hdmi", filename)
			err := cmd.Run()
			if err != nil {
				log.Error(err)
			}
		}()
	} else {
		return fmt.Errorf("no .mp3 suffix")
	}
	return nil
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

	sounds := getSounds()

	t := template.New("html")
	t, err := t.Parse(templates.MainTemplate)
	if err != nil {
		log.Error(err)
	}

	err = t.Execute(w, sounds)
	if err != nil {
		log.Error(err)
	}

}

type ByNumPlayed []Sound

func (s ByNumPlayed) Len() int {
	return len(s)
}
func (s ByNumPlayed) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByNumPlayed) Less(i, j int) bool {
	return s[i].Count > s[j].Count
}

func getSounds() []Sound {
	sounds := make([]Sound, 0)

	dir, err := ioutil.ReadDir("sounds")
	if err != nil {
		log.Error(err)
	}

	for _, v := range dir {
		if !v.IsDir() {
			filename := v.Name()
			if (strings.HasSuffix(filename, ".mp3")) {
				newSound := Sound{
					SoundFile: filename,
					HasImage: true,
					Count: getCounter("sounds/" + filename),
				}
				filenameWithoutExt := filename[:len(filename)-4]
				pngFilename := filenameWithoutExt + ".png"
				if _, err := os.Stat("sounds/" + pngFilename); os.IsNotExist(err) {
					jpgFilename := filenameWithoutExt + ".jpg"
					if _, err := os.Stat("sounds/" + jpgFilename); os.IsNotExist(err) {
						newSound.HasImage = false
					} else {
						newSound.ImageFile = jpgFilename
					}
				} else {
					newSound.ImageFile = pngFilename
				}
				sounds = append(sounds, newSound)
			}
		}
	}

	sort.Sort(ByNumPlayed(sounds))

	return sounds
}

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}