package physics

import (
	"time"
	"github.com/b00lduck/raspberry_soundboard/persistence"
)

func Process(persistence *persistence.Persistence) {

	for {
		time.Sleep(1 * time.Second)

		persistence.Lock()
		changed := false
		state := persistence.State()
		for k,v := range state.Sounds {
			if (v.Temperature < 20) {
				state.Sounds[k].Temperature = 20
				state.Sounds[k].Overheated = false
				changed = true
			} else if (v.Temperature != 20) {
				oldTemp := float32(v.Temperature)
				diff := (oldTemp - 20.0) * 0.003
				newTemp := oldTemp - diff
				state.Sounds[k].Temperature = newTemp
				changed = true
				if newTemp < 50 {
					state.Sounds[k].Overheated = false
				}
			}
		}
		if changed {
			persistence.PersistNoLock()
		}
		persistence.Unlock()

	}

}