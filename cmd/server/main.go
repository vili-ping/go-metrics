package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Server is running!!!")

	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/update/", update)

	return http.ListenAndServe(":8080", mux)
}

func update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	if r.Header.Get("Content-Type") != "text/plain" || r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	pParts := strings.Split(r.URL.Path, "/")[1:]
	if len(pParts) != 4 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	mType, mValue := pParts[1], pParts[3]

	switch mType {
	case "gauge":
		_, err := strconv.ParseFloat(mValue, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	case "counter":
		_, err := strconv.ParseInt(mValue, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Record value"))
}
