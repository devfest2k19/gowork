package cmd

import (
	"../pkg/handlers"
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


	router.Handle("/person", handlers.HandlePost{}).
		Methods(http.MethodPost).
		Name("add-person")

	fmt.Println("server is starting")

	server.ListenAndServe()

}
