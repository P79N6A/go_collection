package player

import (
	"fmt"
	"time"
	"strings"
	"../music_db"
)

type MusicPlayer interface {
	GetSuffix() string
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

func (st1Player St1Player) GetSuffix() string {
	return st1Player.suffix
}

func NewSt2Player() *St2Player {
	return &St2Player{"@@@@@"}
}

func (st2Player St2Player) GetSuffix() string {
	return st2Player.suffix
}

func NewSt3Player() *St3Player {
	return &St3Player{"#####"}
}

func (st3Player St3Player) GetSuffix() string {
	return st3Player.suffix
}

//实现多态
func Factory(t string) MusicPlayer {
	switch t {
		case "type:st1":
			return NewSt1Player()
		case "type:st2":
			return NewSt2Player()
		case "type:st3":
			return NewSt3Player()
		default:
			panic("no such player")
	}
}

func PlayMusic(musicPlayer MusicPlayer, music *music_db.Music) {
	fmt.Println("start to play ",music.Name)

	line := strings.Split(music.Lyric,"\n")
	for i := 1 ; i < len(line) ; i++ {
		fmt.Println(line[i]+" "+musicPlayer.GetSuffix())
		time.Sleep(time.Second*1)
	}
}

