package persistence

import (
	log "github.com/Sirupsen/logrus"
	"github.com/HouzuoGuo/tiedot/db"
)

const SOUNDS_COLLECTION = "sounds"

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
	Deleted     bool
	Category    string
}

type Persistence struct {
	UpdateCallback func()
	sounds	       *db.Col
}

func NewPersistence() *Persistence {

	// create db
	db, err := db.OpenDB("database")
	if err != nil {
		log.Fatal(err)
	}

	// create collection
	if err := db.Create(SOUNDS_COLLECTION); err != nil {
		log.Fatal(err)
	}

	sounds := db.Use(SOUNDS_COLLECTION)

	obj := Persistence{
		sounds: sounds,
		UpdateCallback: nil}

	//go obj.SaveThread()
	return &obj
}

//func (p *Persistence) SaveThread() {
//	for {
		/*
		p.mutex.Lock()
		p.loadSoundsNolock("sounds")
		p.persistNoLock()
		p.mutex.Unlock()
		time.Sleep(15 * time.Second)
		*/
//	}
//}

/*
func (p *Persistence) State() *SoundList {
	sort.Sort(ByNumPlayed(p.state.Sounds))
	return p.state
}

func (p *Persistence) IsPlayable(filename string) bool {
	k, found := p.getSoundIndex(filename, p.state.Sounds)
	if found {
		return (!p.state.Sounds[k].Overheated) && (!p.state.Sounds[k].Deleted)
	}
	return false
}

func (p *Persistence) GetCategory(filename string) string {
	k, found := p.getSoundIndex(filename, p.state.Sounds)
	if found {
		return (p.state.Sounds[k].Category)
	}
	return ""
}

func (p *Persistence) IncCounter(filename string) {

	k, found := p.getSoundIndex(filename, p.state.Sounds)
	changed := false
	if found {
		if (!p.state.Sounds[k].Overheated) {
			p.state.Sounds[k].Count++
			p.state.Sounds[k].Temperature += 100.0
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

/*
func (p *Persistence) Load() {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	jsonBytes, err := ioutil.ReadFile("database.json")
	if err != nil {
		log.Error(err)
	}
	err = json.Unmarshal(jsonBytes, &p.state)
	if err != nil {
		log.Error(err)
	}
	log.WithField("numSounds", len(p.state.Sounds)).Info("Database loaded from disk")
}
*/

/*
func (p *Persistence) persistNoLock() error {
	jsonbytes, err := json.Marshal(p.state)
	if err != nil {
		return err
	}

	if !bytes.Equal(jsonbytes, p.oldJsonState) {
		err = ioutil.WriteFile("database.json", jsonbytes, 0644)
		if err != nil {
			return err
		}
		log.WithField("numSounds", len(p.state.Sounds)).Info("Database saved to disk")
		p.oldJsonState = jsonbytes
	}
	return nil
}
*/


func (p *Persistence) loadSounds(directory string) error {


	log.WithField("numSounds", p.sounds.ApproxDocCount()).Debug("Now updating sounds from sound folder...")

	// add new sounds
	sounds := GetSounds(directory)
	for _, v := range sounds.Sounds {

		p.sounds.ForEachDoc(func(id int, doc []byte) {



		});

		tmpSound := Sound{}
		p.scribble.Read(SOUNDS_COLLECTION, v.SoundFile, &tmpSound)

		if _, found := p.getSoundIndex(v.SoundFile, p.state.Sounds); !found {
			p.state.Sounds = append(p.state.Sounds, v)
			log.WithField("soundFile", v.SoundFile).Info("added new sound")
		}

		if index, found := p.getSoundIndex(v.SoundFile, p.state.Sounds); found && p.state.Sounds[index].Category != v.Category {
			p.state.Sounds[index].Category = v.Category
			log.WithField("soundFile", v.SoundFile).WithField("newCategory", v.Category).Info("changed category of sound")
		}

		if index, found := p.getSoundIndex(v.SoundFile, p.state.Sounds); found && p.state.Sounds[index].Deleted {
			oldCount := p.state.Sounds[index].Count
			oldTemp := p.state.Sounds[index].Temperature
			p.state.Sounds[index] = v
			p.state.Sounds[index].Count = oldCount
			p.state.Sounds[index].Temperature = oldTemp
			log.WithField("soundFile", v.SoundFile).Info("added new sound")
		}
	}

	// delete nonexsisting sounds
	for k, v := range p.state.Sounds {
		if _, found := p.getSoundIndex(v.SoundFile, sounds.Sounds); !v.Deleted && !found {
			p.state.Sounds[k].Deleted = true
			log.WithField("soundFile", p.state.Sounds[k].SoundFile).Info("removed sound")
		}
	}

	// delete double sounds
	for k, v := range p.state.Sounds {
		if _, found := p.getSoundIndex(v.SoundFile, sounds.Sounds); !v.Deleted && !found {
			p.state.Sounds[k].Deleted = true
			log.WithField("soundFile", p.state.Sounds[k].SoundFile).Info("removed sound")
		}
	}

	log.WithField("numSounds", len(p.state.Sounds)).Debug("Sounds updated from sound folder")
}
