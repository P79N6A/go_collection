package main

import(
	"fmt"
	//"./music_db"
	//"./player"
)

func main() {
	cmd := GetInput()
	fmt.Println("get cmd:" + cmd)

	m := NewManager("/Users/wangzijie/code/golang/src/music_player/data/")
	/*for e := m.musicList.Front(); e != nil; e = e.Next() {
		music := e.Value.(*music_db.Music)
		fmt.Println(music.FileType)
		p := player.Factory(music.FileType)
		player.PlayMusic(p,music)
	}*/
	m.PlayMusic()
}
