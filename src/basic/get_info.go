package main

import (
	"fmt"
	"os"
	"io"
	"bufio"
	"encoding/json"
)

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func list(fileName string, target string) {
	fmt.Println("==========" + target + "==========")
	file, err := os.Open(fileName)
    checkErr(err)
    reader := bufio.NewReader(file)

    var f interface{}

    for {
        line, err := reader.ReadString('\n')
        if err != nil || io.EOF == err {
            break
        }

		b := []byte(line)
		err = json.Unmarshal(b, &f)
		checkErr(err)

		m := f.(map[string]interface{})
		fmt.Println(m[target])
    }

}

func main() {
	fileName := "file-20181205T1216"
	list(fileName, "md5")
	list(fileName, "sha1")
	list(fileName, "sha256")
	list(fileName, "link")
}
