package persistence

import (
	"io/ioutil"
	log "github.com/Sirupsen/logrus"
	"strings"
	"os"
)

func GetSounds(directoy string) SoundList {

	categories := GetCategoryDirs(directoy)
	allSounds := make([]Sound, 0)

	for _, v := range categories {
		allSounds = append(allSounds, GetSoundsOfCategory(directoy, v)...)
	}

	return SoundList{allSounds}
}

func GetCategoryDirs(directory string) []string {

	categoryDirs := make([]string, 0)

	dir, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Error(err)
	}

	for _, v := range dir {
		if v.IsDir() {
			categoryDirs = append(categoryDirs, v.Name())
		}
	}

	return categoryDirs
}

func GetSoundsOfCategory(directory string, category string) []Sound {
	sounds := make([]Sound, 0)

	dir, err := ioutil.ReadDir(directory + "/" + category)
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
					Deleted: false,
					Category: category}
				filenameWithoutExt := filename[:len(filename)-4]
				pngFilename := filenameWithoutExt + ".png"
				if _, err := os.Stat(directory + "/" + category + "/" + pngFilename); os.IsNotExist(err) {
					jpgFilename := filenameWithoutExt + ".jpg"
					if _, err := os.Stat(directory + "/" + category + "/" + jpgFilename); os.IsNotExist(err) {
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

	return sounds
}