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
	/*processor := NewProcessor("/Users/wangzijie/code/golang/src/music_player/data/")
	info := processor.ParseContent(name)
	music.name = info[0]
	music.fileType = info[1]
	music.singer = info[2]
	music.album = info[3]
	music.time = info[4]
	for i := 5 ; i < len(info) ; i++ {
		music.lyric += info[i]
	}*/
	return &music
}

func (music *Music) ShowInfo() {
	fmt.Println(music)
}
