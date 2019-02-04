package main

import(
    "log"
    "time"
)

func print(i int, c chan int) {
    log.Printf("%d print start, send to channel", i )
    time.Sleep(3*time.Second)
    log.Printf("%d print end", i)
    c <- i
}

func main() {
    c := make(chan int)
    go print(1, c)
    go print(2, c)
    go print(3, c)

    for i := 0; i < 3; i++ {
    log.Printf("get chan data:%d", <-c)
    }
}
