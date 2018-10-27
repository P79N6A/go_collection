package music_db

import (
	"fmt"
)

type Music struct {
	singer string
	album string
	time string
	lyric string	
}

func NewMusic() *Music  {
	music := Music{}

	processor := NewProcessor("/Users/wangzijie/code/golang/src/music_player/data/")
	info := processor.ParseContent("hello.st1")
	music.singer = info[0]
	music.album = info[1]
	music.time = info[2]
	for i := 3 ; i < len(info) ; i++ {
		music.lyric += info[i]
	}
	return &music
}

func (music *Music) GetInfo() {
	fmt.Println(music)
} 
