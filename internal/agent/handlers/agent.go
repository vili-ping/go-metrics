package handlers

import (
	"fmt"
	"math/rand/v2"
	"runtime"

	"github.com/go-resty/resty/v2"
)

var client = resty.New()

func sendMetric(mType, name string, val string) error {
	url := fmt.Sprintf("http://localhost:8080/update/%s/%s/%s", mType, name, val)

	_, err := client.R().SetHeader("Content-Type", "text/plain").Post(url)
	if err != nil {
		return fmt.Errorf("metric send error %s: %v", name, err)
	}

	return nil
}

type Metrics struct {
	Alloc         string
	BuckHashSys   string
	Frees         string
	GCCPUFraction string
	GCSys         string
	HeapAlloc     string
	HeapIdle      string
	HeapInuse     string
	HeapObjects   string
	HeapReleased  string
	HeapSys       string
	LastGC        string
	Lookups       string
	MCacheInuse   string
	MCacheSys     string
	MSpanInuse    string
	MSpanSys      string
	Mallocs       string
	NextGC        string
	NumForcedGC   string
	NumGC         string
	OtherSys      string
	PauseTotalNs  string
	StackInuse    string
	StackSys      string
	Sys           string
	TotalAlloc    string
	PollCount     uint
	RandomValue   string
}

func (m *Metrics) CollectMetrics() {
	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)

	m.Alloc = fmt.Sprint(stats.Alloc)
	m.BuckHashSys = fmt.Sprint(stats.BuckHashSys)
	m.GCCPUFraction = fmt.Sprint(int(stats.GCCPUFraction))
	m.GCSys = fmt.Sprint(stats.GCSys)
	m.HeapAlloc = fmt.Sprint(stats.HeapAlloc)
	m.HeapIdle = fmt.Sprint(stats.HeapIdle)
	m.HeapInuse = fmt.Sprint(stats.HeapInuse)
	m.HeapObjects = fmt.Sprint(stats.HeapObjects)
	m.HeapReleased = fmt.Sprint(stats.HeapReleased)
	m.HeapSys = fmt.Sprint(stats.HeapSys)
	m.LastGC = fmt.Sprint(stats.LastGC)
	m.Lookups = fmt.Sprint(stats.Lookups)
	m.MCacheInuse = fmt.Sprint(stats.MCacheInuse)
	m.MCacheSys = fmt.Sprint(stats.MCacheSys)
	m.MSpanInuse = fmt.Sprint(stats.MSpanInuse)
	m.MSpanSys = fmt.Sprint(stats.MSpanSys)
	m.Mallocs = fmt.Sprint(stats.Mallocs)
	m.NextGC = fmt.Sprint(stats.NextGC)
	m.NumForcedGC = fmt.Sprint(stats.NumForcedGC)
	m.NumGC = fmt.Sprint(stats.NumGC)
	m.OtherSys = fmt.Sprint(stats.OtherSys)
	m.PauseTotalNs = fmt.Sprint(stats.PauseTotalNs)
	m.StackInuse = fmt.Sprint(stats.StackInuse)
	m.StackSys = fmt.Sprint(stats.StackSys)
	m.Sys = fmt.Sprint(stats.Sys)
	m.TotalAlloc = fmt.Sprint(stats.TotalAlloc)

	m.PollCount += 1
	m.RandomValue = fmt.Sprint(rand.IntN(10_000))
}

func (m Metrics) SendMetrics() {
	sendMetric("gauge", "Alloc", m.Alloc)
	sendMetric("gauge", "BuckHashSys", m.BuckHashSys)
	sendMetric("gauge", "Frees", m.Frees)
	sendMetric("gauge", "GCCPUFraction", m.GCCPUFraction)
	sendMetric("gauge", "GCSys", m.GCSys)
	sendMetric("gauge", "HeapAlloc", m.HeapAlloc)
	sendMetric("gauge", "HeapIdle", m.HeapIdle)
	sendMetric("gauge", "HeapInuse", m.HeapInuse)
	sendMetric("gauge", "HeapObjects", m.HeapObjects)
	sendMetric("gauge", "HeapReleased", m.HeapReleased)
	sendMetric("gauge", "HeapSys", m.HeapSys)
	sendMetric("gauge", "LastGC", m.LastGC)
	sendMetric("gauge", "Lookups", m.Lookups)
	sendMetric("gauge", "MCacheInuse", m.MCacheInuse)
	sendMetric("gauge", "MCacheSys", m.MCacheSys)
	sendMetric("gauge", "MSpanInuse", m.MSpanInuse)
	sendMetric("gauge", "MSpanSys", m.MSpanSys)
	sendMetric("gauge", "Mallocs", m.Mallocs)
	sendMetric("gauge", "NextGC", m.NextGC)
	sendMetric("gauge", "NumForcedGC", m.NumForcedGC)
	sendMetric("gauge", "NumGC", m.NumGC)
	sendMetric("gauge", "OtherSys", m.OtherSys)
	sendMetric("gauge", "PauseTotalNs", m.PauseTotalNs)
	sendMetric("gauge", "StackInuse", m.StackInuse)
	sendMetric("gauge", "StackSys", m.StackSys)
	sendMetric("gauge", "Sys", m.Sys)
	sendMetric("gauge", "TotalAlloc", m.TotalAlloc)
	sendMetric("counter", "PollCount", fmt.Sprint(m.PollCount))
	sendMetric("gauge", "RandomValue", m.RandomValue)
}
