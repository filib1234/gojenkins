package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write(sayHello())
	})

	http.HandleFunc("/bye", func(writer http.ResponseWriter, request *http.Request) {
		bye := sayByeBye()
		writer.Write(bye)
		log.Fatal(bye)
	})

	log.Fatal(http.ListenAndServe(":3333", nil))
}

func sayByeBye() []byte {
	return []byte("bye bye")
}

func sayHello() []byte {
	return []byte("hello jenkins")
}
