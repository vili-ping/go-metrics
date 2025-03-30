package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/vili-ping/go-metrics/internal/storage"
)

var memStorage = &storage.MemStorage{Vals: make(map[string]string)}
var storageInstance storage.Storage = memStorage

func UpdateMetrics(w http.ResponseWriter, r *http.Request) {
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

	mType, mKey, mValue := pParts[1], pParts[2], pParts[3]

	fmt.Println(mType, mKey, mValue)

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
	w.Write([]byte("Record value"))
}
