package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {

	router := mux.NewRouter()

	server := http.Server{
		Addr:         ":8001",
		Handler:      router,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	router.Handle("/hello", handler{}).
		Methods(http.MethodGet).
		Name("get-hello")

	router.Handle("/hello", handler{}).
		Methods(http.MethodPost).
		Name("post-hello").
		Headers("content-type", "application/json")

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
	fmt.Fprintf(writer, "endpoint: %v method: %v , message received : %v \n", request.URL.String(), request.Method, string(msg))
}
