package main

import(
	//"fmt"
	"./music_db"
	"./player"
)

func main() {
	processor := music_db.NewProcessor("/Users/wangzijie/code/golang/src/music_player/data/")
	processor.LoadAllMusic()
	music := processor.GenerateMusic("hello.st1")
	music.ShowInfo()
	st1Player := player.NewSt1Player()
	st1Player.PlaySingleMusic(music)
}
