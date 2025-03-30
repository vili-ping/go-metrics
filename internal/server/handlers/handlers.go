package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/vili-ping/go-metrics/internal/storage"
)

var memStorage = &storage.MemStorage{Vals: make(map[string]string)}
var storageInstance storage.Storage = memStorage

func UpdateMetrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	if r.Header.Get("Content-Type") != "text/plain" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	mType, mKey, mValue := chi.URLParam(r, "type"), chi.URLParam(r, "name"), chi.URLParam(r, "value")

	switch mType {
	case "gauge":
		_, err := strconv.ParseFloat(mValue, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = storageInstance.SetMetric(mKey, mType, mValue)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	case "counter":
		_, err := strconv.ParseInt(mValue, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = storageInstance.SetMetric(mKey, mType, mValue)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func GetMetric(w http.ResponseWriter, r *http.Request) {
	mName := chi.URLParam(r, "name")
	value, err := storageInstance.GetMetric(mName)

	fmt.Println(value, err)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Write([]byte(value))
}

func GetMetrics(w http.ResponseWriter, r *http.Request) {
	fmt.Println(storageInstance.GetAllMetrics())
	w.Write([]byte(storageInstance.GetAllMetrics()))
}
