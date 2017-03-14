package endpoints

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/b00lduck/raspberry_soundboard/persistence"
)

func InitPlay(persistence *persistence.Persistence) {
	http.HandleFunc("/api/play/", func(w http.ResponseWriter, r *http.Request) {

		log.WithField("URI", r.RequestURI).Info("Incoming play request")

		filename := r.RequestURI[10:]

		err := playSound(filename, persistence)

		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}

	})
}

func playSound(filename string, persistence *persistence.Persistence) error {
	if strings.HasSuffix(filename, ".mp3") {

		if persistence.IsPlayable(filename) {

			cat := persistence.GetCategory(filename)

			filenameWithPath := "sounds/" + cat + "/" + filename
			log.WithField("filename", filename).Info("playing sound")
			if _, err := os.Stat(filenameWithPath); os.IsNotExist(err) {
				log.Error(filenameWithPath)
				return fmt.Errorf("Not found")
			}

			go func() {
				cmd := exec.Command("omxplayer", "-o", "both", filenameWithPath)
				err := cmd.Run()
				if err != nil {
					log.Error(err)
				}
			}()

			persistence.IncCounter(filename)

		}

	} else if strings.HasSuffix(filename, ".youtube") {
		cat := persistence.GetCategory(filename)

		filenameWithPath := "sounds/" + cat + "/" + filename
		log.WithField("filename", filename).Info("playing video")
		if _, err := os.Stat(filenameWithPath); os.IsNotExist(err) {
			log.Error(filenameWithPath)
			return fmt.Errorf("Not found")
		}

		b, _ := ioutil.ReadFile(filenameWithPath)
		url := string(b)
		url := strings.TrimSpace(url)

		go func() {
			cmd := exec.Command("omxyoutube", url)
			err := cmd.Run()
			if err != nil {
				log.Error(err)
			}
		}()

		persistence.IncCounter(filename)

	} else {
		return fmt.Errorf("no .mp3 suffix")
	}
	return nil
}
