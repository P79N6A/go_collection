package main

import(
	"fmt"
)

func main() {
	c := new(Controller)

	m := NewManager("/Users/wangzijie/code/golang/src/music_player/data/")
	//m.AddMusic("hello.st1")
	//m.PlayMusic()
	for {
		cmd := c.GetInput()
		switch cmd {
			case "a" :
				name := c.GetInput()
				m.AddMusic(name)
			case "p" :
				m.PlayMusic()
			default :
				fmt.Println("cmd err")
		}
	}
}
