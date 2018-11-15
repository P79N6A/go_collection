package main

import (
    "fmt"
    "os"
    "io"
    "io/ioutil"
    "bufio"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func ReadByIoutil(name string) {
    contents, err := ioutil.ReadFile(name)
    check(err)
    // the type of contents is []byte
    // need to transfer it to string
    fmt.Println(string(contents))
}

func ReadByOs(name string) {
    file, err := os.Open(name)
    check(err)

    buf := make([]byte,1024)
    contents, err := file.Read(buf)
    check(err)
    fmt.Printf("len is %d\n",contents)
    fmt.Println(string(buf))
}

func ReadByLine(name string) {
    file, err := os.Open(name)
    check(err)
    reader := bufio.NewReader(file)
    for {
        line, err := reader.ReadString('\n')
        if err != nil || io.EOF == err {
            break
        }
        fmt.Print(line)
    }
}

func main() {
    name := "./data/test.txt"
    //1
    fmt.Println("=====func ReadByIoutil=====")
    ReadByIoutil(name)
    //2
    fmt.Println("=====func ReadByOs=====")
    ReadByOs(name)
    //3
    fmt.Println("=====func ReadByLine=====")
    ReadByLine(name)
}
