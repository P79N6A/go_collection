package main

import(
    "log"
    "time"
    "sync"
)

func print(i int, wg *sync.WaitGroup) {
    log.Printf("%d print start", i)
    time.Sleep(3*time.Second)
    log.Printf("%d print end", i)
    wg.Done()
}

func main() {
    var wg sync.WaitGroup
    wg.Add(3)

    go print(1, &wg)
    go print(2, &wg)
    go print(3, &wg)

    wg.Wait()
}
