package main

import (
	"net/http"
	"./controller"
)

func main() {
	controller.StartUp()
	http.ListenAndServe(":8888",nil)

}
