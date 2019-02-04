package music_db

import (
	"fmt"
)

type Music struct {
	Name string
	FileType string
	Singer string
	Album string
	Time string
	Lyric string
}

func NewMusic() *Music  {
	music := Music{}
	return &music
}

func (music *Music) ShowInfo() {
	fmt.Println(music)
}
