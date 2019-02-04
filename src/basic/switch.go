package main

import (
    "fmt"
)

func init() {
    fmt.Println("init switch.go")
}

func switchString(str string) string {
    var flag string
    switch str {
    case "hello" :
        fmt.Println("this is hello")
        flag = "hello ok"
    case "hi" :
        fmt.Println("this is hi")
        flag = "hi ok"
    default :
        fmt.Println("other")
        flag = "other ok"
    }
    return flag
}

func switchInt(i int) int {
    var flag int
    switch {
    case i > 0 && i <= 5 :
        fmt.Println("i is 0 to 5")
        flag = 0
    case i > 5 && i <= 10 :
        fmt.Println("i is 5 to 10")
        flag = 5
    default :
        fmt.Println("other")
        flag = -1
    }
    return flag
}

func main() {
    switchString("hello")
    switchInt(100)
}
