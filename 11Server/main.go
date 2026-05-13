package main

import (
	"log"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func main() {
	http.HandleFunc("/say_hello", sayHello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Print("ListenAndServe: ", err)
		return
	}
}
