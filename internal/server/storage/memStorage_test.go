package storage

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemStorage_SaveGauge(t *testing.T) {
	storage := NewMemStorage()

	type args struct {
		key   string
		value float64
	}
	tests := []struct {
		name      string
		args      args
		wantError error
	}{
		{
			name: "empty key",
			args: args{
				key:   "",
				value: 1.1,
			},
			wantError: ErrStorageEmptyKey,
		},
		{
			name: "successful saved",
			args: args{
				key:   "test",
				value: 1.1,
			},
			wantError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := storage.SaveGauge(tt.args.key, tt.args.value)
			if tt.wantError != nil {
				assert.ErrorIs(t, err, tt.wantError)
				return
			}
			assert.NoError(t, err)
		})
	}
}

func TestMemStorage_SaveCounter(t *testing.T) {
	storage := NewMemStorage()

	type args struct {
		key   string
		value int64
	}
	tests := []struct {
		name      string
		args      args
		wantError error
	}{
		{
			name: "empty key",
			args: args{
				key:   "",
				value: 1,
			},
			wantError: ErrStorageEmptyKey,
		},
		{
			name: "successful saved",
			args: args{
				key:   "test",
				value: 1,
			},
			wantError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := storage.SaveCounter(tt.args.key, tt.args.value)
			if tt.wantError != nil {
				assert.ErrorIs(t, err, tt.wantError)
				return
			}
			assert.NoError(t, err)
		})
	}
}

func TestMemStorage_GetGauge(t *testing.T) {
	storage := NewMemStorage()
	storage.gauge["test"] = 1.1

	tests := []struct {
		name string
		key  string
		want float64
		ok   bool
	}{
		{
			name: "not found",
			key:  "abcd",
			want: 0.0,
			ok:   false,
		},
		{
			name: "successful found",
			key:  "test",
			want: 1.1,
			ok:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := storage.GetGauge(tt.key)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.ok, ok)
		})
	}
}

func TestMemStorage_GetCounter(t *testing.T) {
	storage := NewMemStorage()
	storage.counter["test"] = 1

	tests := []struct {
		name string
		key  string
		want int64
		ok   bool
	}{
		{
			name: "not found",
			key:  "abcd",
			want: 0,
			ok:   false,
		},
		{
			name: "successful found",
			key:  "test",
			want: 1,
			ok:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := storage.GetCounter(tt.key)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.ok, ok)
		})
	}
}

func TestMemStorage_GetAll(t *testing.T) {
	storage := NewMemStorage()
	storage.counter["testCounter"] = 1
	storage.gauge["testGauge"] = 1.1

	want := map[string]any{
		"testCounter": int64(1),
		"testGauge":   float64(1.1),
	}

	t.Run("should return all data", func(t *testing.T) {
		result := storage.GetAll()
		if !reflect.DeepEqual(result, want) {
			t.Error("Maps are not equal")
		}
	})
}
