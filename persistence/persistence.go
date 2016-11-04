package persistence

import (
	log "github.com/Sirupsen/logrus"
	"sync"
	"encoding/json"
	"io/ioutil"
	"time"
)

type SoundList struct {
	Sounds[] Sound
}

type Sound struct {
	SoundFile   string
	ImageFile   string
	HasImage    bool
	Count       int
	Temperature float32
	Overheated  bool
}

type Persistence struct {
	UpdateCallback func()
	mutex          *sync.Mutex
	state          *SoundList
}

func NewPersistence() *Persistence {
	obj := Persistence{
		mutex: &sync.Mutex{},
		state: &SoundList{Sounds: make([]Sound,0)},
		UpdateCallback: nil}
	obj.Load()
	go obj.SaveThread()
	return &obj
}

func (p *Persistence) SaveThread() {
	for {
		p.Persist()
		time.Sleep(15 * time.Second)
	}
}

func (p *Persistence) LoadSounds(directory string) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.loadSoundsNolock(directory)
}

func (p *Persistence) State() *SoundList {
	return p.state
}

func (p *Persistence) Lock() {
	p.mutex.Lock()
}

func (p *Persistence) Unlock() {
	p.mutex.Unlock()
}

func (p *Persistence) JsonState() []byte {
	soundList := p.State()
	myJson, err := json.Marshal(soundList)
	if err != nil {
		log.Error(err)
	}
	return myJson
}

func (p *Persistence) IncCounter(filename string) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	k, found := p.getSoundIndex(filename)
	changed := false
	if found {
		if (!p.state.Sounds[k].Overheated) {
			p.state.Sounds[k].Count++
			p.state.Sounds[k].Temperature += 30.0
			log.WithField("count", p.state.Sounds[k].Count).
			    WithField("temp", p.state.Sounds[k].Temperature).
			    Info("Increased sound count and temperature")
			if p.state.Sounds[k].Temperature > 100.0 {
				log.WithField("temp", p.state.Sounds[k].Temperature).Warn("Sound now overheated! Bäm!")
				p.state.Sounds[k].Overheated = true
			}
			changed = true
		} else {
			log.WithField("temp", p.state.Sounds[k].Temperature).Warn("Sound too hot, cool down first")
		}
	} else {
		log.WithField("id", k).Error("Sound not found during counter increase")
	}
	if p.UpdateCallback != nil && changed {
		p.UpdateCallback()
	}
}

func (p *Persistence) Persist() error {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.persistNoLock()
}

func (p *Persistence) Load() {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	bytes, err := ioutil.ReadFile("database.json")
	if err != nil {
		log.Error(err)
	}
	err = json.Unmarshal(bytes, &p.state)
	if err != nil {
		log.Error(err)
	}
	log.WithField("numSounds", len(p.state.Sounds)).Info("Database loaded from disk")
	p.loadSoundsNolock("sounds")
}

func (p *Persistence) persistNoLock() error {
	bytes, err := json.Marshal(p.state)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("database.json", bytes, 0644)
	if err != nil {
		return err
	}
	log.WithField("numSounds", len(p.state.Sounds)).Info("Database saved to disk")
	return nil
}

func (p *Persistence) loadSoundsNolock(directory string) {
	sounds := GetSounds(directory)
	for _, v := range sounds.Sounds {
		if _, found := p.getSoundIndex(v.SoundFile); !found {
			p.state.Sounds = append(p.state.Sounds, v)
		}
	}
	log.WithField("numSounds", len(p.state.Sounds)).Info("Sounds updated from sound folder")
}

func (p *Persistence) getSoundIndex(name string) (int, bool) {
	for k, v := range p.state.Sounds {
		if v.SoundFile == name {
			return k, true
		}
	}
	return 0, false
}
