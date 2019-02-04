package main

import(
	"fmt"
)

func main() {
	c := new(Controller)

	m := NewManager("/Users/wangzijie/code/golang/src/music_player/data/")

	for {
		ShowTips()
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
			case "d" :
				fmt.Println("delete music: ")
				name := c.GetInput()
				m.DeleteMusic(name)
			case "r" :
				fmt.Println("rollback music: ")
				name := c.GetInput()
				m.RollbackMusic(name)
			case "g" :
				m.ShowDeletedList()
			default :
				fmt.Println("cmd err")
		}
	}
}

func ShowTips() {
	fmt.Println("please input cmd")
	fmt.Println("a	:	add music in play list")
	fmt.Println("p	:	play music")
	fmt.Println("s	:	show music in play list")
	fmt.Println("d	:	delete music in play list")
	fmt.Println("r	:	rollback music in deleted list")
	fmt.Println("g	:	show deleted list")
}
