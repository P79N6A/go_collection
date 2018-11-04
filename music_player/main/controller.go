package main

import(
	"fmt"
)

type Controller struct {
}

func (c *Controller) GetInput() string {
	var MusicName string
	fmt.Scanln(&MusicName)
	return MusicName
}
