package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/vili-ping/go-metrics/internal/config/serverconfig"
	"github.com/vili-ping/go-metrics/internal/server/handlers"
)

var config = serverconfig.ParseConfig()

func main() {
	fmt.Printf("Server is running on %s\n", config.Address)

	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	r := chi.NewRouter()

	r.Get("/", handlers.GetMetrics)
	r.Get("/value/{type}/{name}", handlers.GetMetric)
	r.Post("/update/{type}/{name}/{value}", handlers.UpdateMetrics)

	return http.ListenAndServe(config.Address, r)
}
