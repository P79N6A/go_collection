package main

import(
	//"fmt"
	"net/http"
)

func printHelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world go web."))
}

func main() {
	http.HandleFunc("/",printHelloWorld)
	http.ListenAndServe(":8888",nil)
}
