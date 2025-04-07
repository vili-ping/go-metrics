package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/vili-ping/go-metrics/internal/server/repostiroy"
	"github.com/vili-ping/go-metrics/internal/server/storageservice"
)

func UpdateMetrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	mType, mKey, mValue := chi.URLParam(r, "type"), chi.URLParam(r, "name"), chi.URLParam(r, "value")

	switch mType {
	case "gauge":
		gauge, err := strconv.ParseFloat(mValue, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		storageservice.SetGauge(mKey, mType, gauge)
	case "counter":
		counter, err := strconv.Atoi(mValue)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		storageservice.SetCounter(mKey, mType, counter)
	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func GetMetric(w http.ResponseWriter, r *http.Request) {
	mName := chi.URLParam(r, "name")
	value, err := storageservice.GetMetric(mName)

	if errors.Is(err, repostiroy.ErrStorageKeyIsNotExist) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte(value))
}

func GetMetrics(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(storageservice.GetAllMetrics()))
}
