package main

import(
	//"fmt"
	"./music_db"
)

func main() {
	processor := music_db.NewProcessor("/Users/wangzijie/code/golang/src/music_player/data/")
	processor.LoadAllMusic()
	music := processor.GenerateMusic("hello.st1")
	music.ShowInfo()
}
