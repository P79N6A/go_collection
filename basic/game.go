package main

import (
    "fmt"
)

type character interface {
    display()
    fight()
}

type behaviour interface {
    performBehave()
}

type fighter struct {
    behaviour
    name string
    level int
}

type knight struct {
    character
    behaviour
    name string
    level int
}

type swordBehaviour struct {
    attr string
}

type oxeBehaviour struct {
    attr string
}

func (s *swordBehaviour) performBehave() {
    fmt.Println("use sword to fight")
}

func (o *oxeBehaviour) performBehave() {
    fmt.Println("use oxe to fight")
}

func (f *fighter) display() {
    fmt.Println("my name is "+f.name)
}

func (f *fighter) fight() {
    f.performBehave()
}

func main() {
    var c character
    c = &fighter{&oxeBehaviour{attr : "short distance weapon"},"fighter_1",10}
    c.display()
    c.fight()
    //sb := swordBehaviour{attr : "short distance weapon"}
    //c.behaviour = sb
    //c.fight()
    // how to set embeded var?
}
