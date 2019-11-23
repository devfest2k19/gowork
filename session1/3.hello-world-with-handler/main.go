package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {

	mux := http.NewServeMux()

	server := http.Server{
		Addr:         ":8001",
		Handler:      mux,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	mux.HandleFunc("/helloFunc", func(writer http.ResponseWriter, request *http.Request) {
		msg, _ := ioutil.ReadAll(request.Body)
		fmt.Fprintf(writer, "endpoint: %v method: %v , message received : %v \n",request.URL.String(), request.Method, string(msg))
	})

	mux.Handle("/helloHandler", handler{})

	fmt.Println("server is starting...")

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

type handler struct {
}

func (handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	msg, _ := ioutil.ReadAll(request.Body)
	fmt.Fprintf(writer, "endpoint: %v method: %v , message received : %v \n",request.URL.String(), request.Method, string(msg))
}
