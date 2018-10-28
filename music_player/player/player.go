package player

import (
	"fmt"
	"time"
	"strings"
	"../music_db"
)

type MusicPlayer interface {
	PlaySingleMusic(music *music_db.Music)
	//PlayAllMusic()
}

type St1Player struct {
	suffix string
}

type St2Player struct {
	suffix string
}

type St3Player struct {
	suffix string
}

func NewSt1Player() *St1Player {
	return &St1Player{"!!!!!"}
}

func NewSt2Player() *St2Player {
	return &St2Player{"@@@@@"}
}

func NewSt3Player() *St3Player {
	return &St3Player{"#####"}
}

func (st1Player *St1Player) PlaySingleMusic(music *music_db.Music) {
	fmt.Println("start to play ",music.Name)
	line := strings.Split(music.Lyric,"\n")
	for i := 0 ; i < len(line) ; i++ {
		fmt.Println(line[i])
		time.Sleep(time.Second*1)
	}
}

