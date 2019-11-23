package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	mux := http.NewServeMux()

	server := http.Server{
		Addr:         ":8001",
		Handler:      mux,
		WriteTimeout: 2 * time.Second,
		ReadTimeout:  2 * time.Second,
	}

	mux.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello from server\n")
	})

	mux.Handle("/hello2", handler{})

	fmt.Println("server is starting")

	server.ListenAndServe()

}

type handler struct {

}

func(handler)ServeHTTP( w http.ResponseWriter,r  *http.Request){
	fmt.Fprintf(w, "Hello2 from server\n" )
}