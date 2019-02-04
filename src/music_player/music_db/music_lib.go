package music_db

import (
	//"fmt"
)

type MusicLib struct {
	lib map[string] *Music
}

func NewMusicLib () *MusicLib {
	return &MusicLib{}
}
