package main

import (
	"github.com/devfest2k19/gowork/session3/10.crud-with-metrics/metrics"
	"github.com/devfest2k19/gowork/session3/10.crud-with-metrics/pkg/handlers"
	"github.com/devfest2k19/gowork/session3/10.crud-with-metrics/pkg/models"
	"github.com/gorilla/mux"
	"github.com/pickme-go/log"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"time"
)

func main() {

	handlers.PersonMap = make(map[int64]models.Person, 0)

	metricsCounter := metrics.InitServiceLatencyCounter(`dev_fest`, `phone_book_crud`)

	router := mux.NewRouter()

	server := http.Server{
		Addr:         ":8001",
		Handler:      router,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	router.Handle("/person", handlers.HandlerPost{MetricsCounter:metricsCounter}).
		Methods(http.MethodPost).
		Name("create-person-info").
		Headers("content-type", "application/json")

	router.Handle("/person/{id}", handlers.HandlerGet{MetricsCounter:metricsCounter}).
		Methods(http.MethodGet).
		Name("get-person-info").
		Headers("content-type", "application/json")

	router.Handle("/person", handlers.HandlerGetAll{MetricsCounter:metricsCounter}).
		Methods(http.MethodGet).
		Name("get-person-info").
		Headers("content-type", "application/json")

	router.Handle("/person/{id}", handlers.HandlerPut{MetricsCounter:metricsCounter}).
		Methods(http.MethodPut).
		Name("update-person-info").
		Headers("content-type", "application/json")

	router.Handle("/person/{id}", handlers.HandlerDelete{MetricsCounter:metricsCounter}).
		Methods(http.MethodDelete).
		Name("delete-person-info").
		Headers("content-type", "application/json")

	router.Handle(`/metrics`, promhttp.Handler()).Methods(http.MethodGet)

	log.Info(`server is starting...`)
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
