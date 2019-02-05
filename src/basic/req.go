package main

import (
    "net/http"
    "log"
    _"fmt"
    "io/ioutil"
)

const (
    stored_url = "http://sync1.dlc.zzyc.360es.cn:80/stored_query.php"
    cab_url = "http://dl.360safe.com/skylar6/download/v3/engine_360/360ave_ex_def_11_9_4_1.cab"
)

func get() {
    resp, err := http.Get(cab_url)
    if err != nil {
        log.Println("get err")
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Println("get body err")
    }

    fileName := "go_cab"
    err = ioutil.WriteFile(fileName, body, 0644) //body需要是二进制类型，若不是，用[]byte(xxx)强转
    //读：4，写：2，执行：1
    //ABCD，A通常为0，表示十进制；B表示用户权限，C表示组用户权限，D表示其他用户权限
    //一般目录赋予0755权限（即用户具有读写执行权限，组和其他用户具有读写权限），
    //文件赋予0644权限（即用户具有读写权限，组用户和其他用户具有只读权限）

    if err != nil {
        log.Println("write file err")
    }

}

func post() {


}

func main() {
    //log.Println(url)
    get()
}
