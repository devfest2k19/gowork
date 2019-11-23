package main

import (
	"fmt"
	"github.com/devfest2k19/gowork/session2/5.crud-post/pkg/handlers"
	"github.com/devfest2k19/gowork/session2/5.crud-post/pkg/models"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func main() {

	handlers.PersonMap = make(map[int64]models.Person, 0)

	router := mux.NewRouter()

	server := http.Server{
		Addr:         ":8001",
		Handler:      router,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	router.Handle("/person", handlers.HandlerPost{}).
		Methods(http.MethodPost).
		Name("add-post").
		Headers("content-type", "application/json")

	fmt.Println("server is starting...")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
