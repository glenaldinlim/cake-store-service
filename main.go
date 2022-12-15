package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello World")
	}
	server := http.Server{
		Addr:    "0.0.0.0:8090",
		Handler: handler,
	}

	log.Println("server running on http://0.0.0.0:8090/")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
