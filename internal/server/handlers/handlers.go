package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/vili-ping/go-metrics/internal/server/service/storage"
)

var memStorage = storage.NewMemStorage()
var service = storage.NewService(memStorage)

func UpdateMetrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	mType, mKey, mValue := chi.URLParam(r, "type"), chi.URLParam(r, "name"), chi.URLParam(r, "value")

	switch mType {
	case "gauge":
		_, err := strconv.ParseFloat(mValue, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = service.Storage.SetMetric(mKey, mType, mValue)
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

		err = service.Storage.SetMetric(mKey, mType, mValue)
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
	value, err := service.Storage.GetMetric(mName)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Write([]byte(value))
}

func GetMetrics(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(service.Storage.GetAllMetrics()))
}
