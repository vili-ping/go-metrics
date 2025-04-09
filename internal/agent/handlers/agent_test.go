package handlers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMetrics_CollectMetrics(t *testing.T) {
	type fields struct {
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

	tests := []struct {
		name   string
		fields fields
	}{
		{
			name:   "CollectMetrics should update metrics",
			fields: fields{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Metrics{
				Alloc:         tt.fields.Alloc,
				BuckHashSys:   tt.fields.BuckHashSys,
				Frees:         tt.fields.Frees,
				GCCPUFraction: tt.fields.GCCPUFraction,
				GCSys:         tt.fields.GCSys,
				HeapAlloc:     tt.fields.HeapAlloc,
				HeapIdle:      tt.fields.HeapIdle,
				HeapInuse:     tt.fields.HeapInuse,
				HeapObjects:   tt.fields.HeapObjects,
				HeapReleased:  tt.fields.HeapReleased,
				HeapSys:       tt.fields.HeapSys,
				LastGC:        tt.fields.LastGC,
				Lookups:       tt.fields.Lookups,
				MCacheInuse:   tt.fields.MCacheInuse,
				MCacheSys:     tt.fields.MCacheSys,
				MSpanInuse:    tt.fields.MSpanInuse,
				MSpanSys:      tt.fields.MSpanSys,
				Mallocs:       tt.fields.Mallocs,
				NextGC:        tt.fields.NextGC,
				NumForcedGC:   tt.fields.NumForcedGC,
				NumGC:         tt.fields.NumGC,
				OtherSys:      tt.fields.OtherSys,
				PauseTotalNs:  tt.fields.PauseTotalNs,
				StackInuse:    tt.fields.StackInuse,
				StackSys:      tt.fields.StackSys,
				Sys:           tt.fields.Sys,
				TotalAlloc:    tt.fields.TotalAlloc,
				PollCount:     tt.fields.PollCount,
				RandomValue:   tt.fields.RandomValue,
			}

			before := *m

			m.CollectMetrics()

			assert.NotEqual(t, before, *m, "metrics shout change after CollectMetrics")
		})
	}
}
