package main

import (
	"fmt"
	"net/http"

	"github.com/vili-ping/go-metrics/internal/server/handlers"
)

func main() {
	fmt.Println("Server is running!!!")

	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/update/", handlers.UpdateMetrics)

	return http.ListenAndServe(":8080", mux)
}
