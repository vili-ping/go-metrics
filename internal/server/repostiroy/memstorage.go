package repostiroy

import (
	"fmt"
	"strings"
	"sync"
)

type memStorage struct {
	sync.RWMutex
	Vals map[string]string
}

func NewMemStorage() *memStorage {
	return &memStorage{Vals: make(map[string]string)}
}

func (m *memStorage) SetMetric(key, val string) {
	m.Lock()
	defer m.Unlock()

	m.Vals[key] = val
}

func (m *memStorage) DeleteMetric(key string) error {
	m.Lock()
	defer m.Unlock()

	_, exists := m.Vals[key]

	if !exists {
		return ErrStorageKeyIsNotExist
	}

	delete(m.Vals, key)

	return nil
}

func (m *memStorage) GetMetric(key string) (string, error) {
	m.RLock()
	defer m.RUnlock()

	val, exists := m.Vals[key]

	if !exists {
		return "", ErrStorageKeyIsNotExist
	}

	return val, nil
}

func (m *memStorage) GetAllMetrics() string {
	m.RLock()
	defer m.RUnlock()

	var sb strings.Builder
	for k, v := range m.Vals {
		fmt.Println(k, v)
		sb.WriteString(fmt.Sprintf("%s=%s\n", k, v))
	}
	return sb.String()
}
