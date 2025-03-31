package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemStorage_SetMetric(t *testing.T) {
	type fields struct {
		Vals map[string]string
	}
	type args struct {
		key   string
		mType string
		val   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "Set Alloc metric",
			fields:  fields{Vals: make(map[string]string)},
			args:    args{key: "Alloc", mType: "gauge", val: "123"},
			wantErr: false,
		},
		{
			name:    "Set PollCount metric",
			fields:  fields{Vals: make(map[string]string)},
			args:    args{key: "PollCount", mType: "counter", val: "1"},
			wantErr: false,
		},
		{
			name:    "Set unsupported metric type",
			fields:  fields{Vals: make(map[string]string)},
			args:    args{key: "PollCount", mType: "pieman", val: "2345"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &memStorage{
				Vals: tt.fields.Vals,
			}

			err := m.SetMetric(tt.args.key, tt.args.mType, tt.args.val)
			if tt.wantErr {
				assert.EqualError(t, err, "type metrics is not support")
				return
			}

			assert.Equal(t, m.Vals[tt.args.key], tt.args.val, "metrics is not right exp: %v, get: %v", m.Vals[tt.args.key], tt.args.val)
		})
	}
}

func TestMemStorage_GetAllMetrics(t *testing.T) {
	type fields struct {
		Vals map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "Get all metrics in storage",
			fields: fields{Vals: map[string]string{"Alloc": "123"}},
			want:   "Alloc=123\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := memStorage{
				Vals: tt.fields.Vals,
			}

			result := m.GetAllMetrics()

			assert.Equal(t, tt.want, result, "Metrics convert string not right: exp: %s, get %s", tt.want, result)
		})
	}
}

func TestMemStorage_GetMetric(t *testing.T) {
	type fields struct {
		Vals map[string]string
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "Get metric by Alloc name",
			fields:  fields{Vals: map[string]string{"Alloc": "123"}},
			args:    args{key: "Alloc"},
			want:    "123",
			wantErr: false,
		},
		{
			name:    "Get eror when key is not exist",
			fields:  fields{Vals: map[string]string{"Alloc": "123"}},
			args:    args{key: "HeartKey"},
			want:    "123",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := memStorage{
				Vals: tt.fields.Vals,
			}
			got, err := m.GetMetric(tt.args.key)

			if tt.wantErr {
				assert.EqualError(t, err, "storage key is not exists", "eror is wrong")
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
