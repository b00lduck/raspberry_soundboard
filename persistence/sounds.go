package persistence

import (
	"io/ioutil"
	log "github.com/Sirupsen/logrus"
	"strings"
	"os"
	"sort"
	"strconv"
)

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
					Count: getCountLegacy(directory + "/" + filename),
				}
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

	sort.Sort(ByNumPlayed(sounds))

	return SoundList{sounds}
}

func getCountLegacy(filename string) int {
	countfile := filename + ".count"

	_, err := os.Stat(countfile)
	if err != nil {
		log.Error(err)
		return 0
	}

	count, err := ioutil.ReadFile(countfile)
	if err != nil {
		log.Error(err)
		return 0
	}
	intCount, err := strconv.Atoi(string(count))
	if err != nil {
		log.Error(err)
		return 0
	}
	return intCount
}