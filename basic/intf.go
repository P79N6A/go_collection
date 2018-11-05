package main

import (
	"fmt"
)

type Car interface {
	run()
}

type Benz struct {
	brand string
}

type Ford struct {
	brand string
}

func (b Benz) run() {
	fmt.Println("benz run")
}

func (f Ford) run() {
	fmt.Println("ford run")
}

func main() {
	var car Car
	b := Benz{"benz"}
	car = b
	car.run()
}
