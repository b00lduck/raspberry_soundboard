package persistence

import (
	"io/ioutil"
	log "github.com/Sirupsen/logrus"
	"strings"
	"os"
)

func GetSounds(directory string) SoundList {
	sounds := make([]Sound, 0)

	dir, err := ioutil.ReadDir(directory)
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
					Count: 0,
					Temperature: 20,
					Overheated: false,
					Deleted: false}
				filenameWithoutExt := filename[:len(filename)-4]
				pngFilename := filenameWithoutExt + ".png"
				if _, err := os.Stat(directory + "/" + pngFilename); os.IsNotExist(err) {
					jpgFilename := filenameWithoutExt + ".jpg"
					if _, err := os.Stat(directory + "/" + jpgFilename); os.IsNotExist(err) {
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

	return SoundList{sounds}
}