package main

import (
	"log"
    _"sync"
    "time"
)


func add(x, y int, c chan int) {
	time.Sleep(3*time.Second)
    c <- (x + y)
}

func main() {
    sum := 0
	c := make(chan int)
	go add(10, 20, c)
	go add(30, 40, c)
	go add(50, 60, c)

    for i:=0; i<3; i++ {
        sum += <-c
    }
    log.Printf("sum is %d", sum)

}
