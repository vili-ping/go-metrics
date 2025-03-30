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
			m := &MemStorage{
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
