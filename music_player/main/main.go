package main

import(
	"fmt"
)

func main() {
	c := new(Controller)

	m := NewManager("/Users/wangzijie/code/golang/src/music_player/data/")

	for {
		fmt.Println("please input cmd")
		fmt.Println("a	:	add music in play list")
		fmt.Println("p	:	play music")
		fmt.Println("s	:	show music in play list")
		cmd := c.GetInput()
		switch cmd {
			case "a" :
				fmt.Println("add music: ")
				name := c.GetInput()
				m.AddMusic(name)
			case "p" :
				m.PlayMusic()
			case "s" :
				m.ShowPlayList()
			default :
				fmt.Println("cmd err")
		}
	}
}
