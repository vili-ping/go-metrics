package storageservice

import (
	"errors"
	"strconv"

	"fmt"

	"github.com/vili-ping/go-metrics/internal/server/repostiroy"
)

type storage interface {
	SetMetric(key, val string)
	GetMetric(key string) (val string, err error)
	GetAllMetrics() string
	DeleteMetric(key string) error
}

type service struct {
	storage storage
}

func newService(s storage) *service {
	return &service{storage: s}
}

var memStorage = repostiroy.NewMemStorage()
var serviceStorage = newService(memStorage)

func SetCounter(key, mType string, val int) {
	counterMetric, err := serviceStorage.storage.GetMetric(key)
	if errors.Is(err, repostiroy.ErrStorageKeyIsNotExist) {
		counterMetric = "0"
	}

	counter, _ := strconv.Atoi(counterMetric)
	counter += val

	serviceStorage.storage.SetMetric(key, fmt.Sprint(counter))
}

func SetGauge(key, mType string, val float64) {
	serviceStorage.storage.SetMetric(key, fmt.Sprint(val))
}

func GetMetric(key string) (string, error) {
	return serviceStorage.storage.GetMetric(key)
}

func GetAllMetrics() string {
	return serviceStorage.storage.GetAllMetrics()
}
