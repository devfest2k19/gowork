package main

import (
	"fmt"
	gorillamux "github.com/gorilla/mux"
	"net/http"
	"time"
)

func main() {

	router := gorillamux.NewRouter()

	//mux := http.NewServeMux()

	server := http.Server{
		Addr:         ":8001",
		Handler:      router,
		WriteTimeout: 2 * time.Second,
		ReadTimeout:  2 * time.Second,
	}

	router.Handle("/hello", handler{}).
		Methods(http.MethodGet).
		Name("hello-get")

	router.Handle("/hello", handlerpost{}).
		Methods(http.MethodPost).
		Name("hello-post")

	fmt.Println("server is starting")

	server.ListenAndServe()

}

type handler struct {
}

func (handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello,handler from server method:%v \n", r.Method)
}

type handlerpost struct {
}

func (handlerpost) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello,handlerpost from server method:%v \n", r.Method)
}