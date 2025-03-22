package storage

type Gauge float64
type Counter int64

type MemStorage struct {
	Vals map[string]string
}

type StorageOperations interface {
	AddGauge(key string, gauge Gauge) (err error)
	AddCounter(key string, counter Counter) (err error)
}
