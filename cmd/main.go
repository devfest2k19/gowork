package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/devfest2k19/gowork/pkg/models"

	_ "net/http/pprof"

	"github.com/devfest2k19/gowork/pkg/handlers"
	gorillamux "github.com/gorilla/mux"
	"github.com/pickme-go/log"
)

func main() {

	router := gorillamux.NewRouter()

	handlers.PersonMap = make(map[int64]models.Person)

	//mux := http.NewServeMux()

	server := http.Server{
		Addr:         ":8001",
		Handler:      router,
		WriteTimeout: 2 * time.Second,
		ReadTimeout:  2 * time.Second,
	}

	router.Handle("/person", handlers.HandlePost{}).
		Methods(http.MethodPost).
		Name("add-person")

	go func() {
		log.Info(http.ListenAndServe(`:6060`, nil))
	}()
	fmt.Println("server is starting")

	server.ListenAndServe()

}
