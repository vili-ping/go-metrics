package storage

import (
	"errors"
	"fmt"
	"strconv"
)

type Gauge float64
type Counter int64

type MemStorage struct {
	Vals map[string]string
}

type Storage interface {
	SetMetric(key string, mType string, val string) (err error)
	GetMetric(key string) (val string, err error)
	DeleteMetric(key string) error
}

func (m *MemStorage) SetMetric(key string, mType string, val string) error {
	if mType == "gauge" {
		m.Vals[key] = val
		return nil
	}

	if mType == "counter" {
		oldValue, _ := strconv.Atoi(m.Vals[key])
		newValue, _ := strconv.Atoi(val)
		oldValue += newValue

		m.Vals[key] = fmt.Sprint(oldValue)
		return nil
	}

	return errors.New("type metrics is not support")
}

func (m *MemStorage) DeleteMetric(key string) error {
	_, exists := m.Vals[key]

	if !exists {
		return errors.New("storage key is not exists")
	}

	delete(m.Vals, key)

	return nil
}

func (m MemStorage) GetMetric(key string) (string, error) {
	val, exists := m.Vals[key]

	if !exists {
		return "", errors.New("storage key is not exists")
	}

	return val, nil
}
