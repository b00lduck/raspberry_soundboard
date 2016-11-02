package persistence

import (
	log "github.com/Sirupsen/logrus"
	"sync"
	"encoding/json"
	"io/ioutil"
)

type SoundList struct {
	Sounds[] Sound
}

type Sound struct {
	SoundFile   string
	ImageFile   string
	HasImage    bool
	Count       int
	Temperature int
}

var mutex = &sync.Mutex{}

var state = SoundList{
	Sounds: make([]Sound,0)}

func Init() {
	Load()
}

func LoadSounds(directory string) {
	mutex.Lock()
	loadSoundsNolock(directory)
	mutex.Unlock()
}

func State() SoundList {
	return state
}

func JsonState() []byte {
	soundList := State()
	myJson, err := json.Marshal(soundList)
	if err != nil {
		log.Error(err)
	}
	return myJson
}

func IncCounter(filename string) {
	mutex.Lock()
	defer mutex.Unlock()

	k, found := getSoundIndex(filename)
	if found {
		log.WithField("key", k).Info("Increasing counter")
		state.Sounds[k].Count++
		state.Sounds[k].Temperature += 5
		log.WithField("count", state.Sounds[k].Count).Error("Increased sound count")
	} else {
		log.WithField("id", k).Error("Sound not found during counter increase")
	}
	persistNolock()
}

func Persist() error {
	mutex.Lock()
	defer mutex.Unlock()
	return persistNolock()
}

func Load() {
	mutex.Lock()
	defer mutex.Unlock()
	bytes, err := ioutil.ReadFile("database.json")
	if err != nil {
		log.Error(err)
	}
	err = json.Unmarshal(bytes, &state)
	if err != nil {
		log.Error(err)
	}
	log.WithField("numSounds", len(state.Sounds)).Info("Database loaded from disk")
	loadSoundsNolock("sounds")
}

func persistNolock() error {
	bytes, err := json.Marshal(state)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("database.json", bytes, 0644)
	if err != nil {
		return err
	}
	log.WithField("numSounds", len(state.Sounds)).Info("Database saved to disk")
	return nil
}

func loadSoundsNolock(directory string) {
	sounds := GetSounds(directory)
	for _, v := range sounds.Sounds {
		if _, found := getSoundIndex(v.SoundFile); !found {
			state.Sounds = append(state.Sounds, v)
		}
	}
	log.WithField("numSounds", len(state.Sounds)).Info("Sounds updated from sound folder")
	persistNolock()
}

func getSoundIndex(name string) (int, bool) {
	for k, v := range state.Sounds {
		if v.SoundFile == name {
			return k, true
		}
	}
	return 0, false
}
