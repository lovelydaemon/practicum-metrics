package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMemStorage_Get(t *testing.T) {
	storage := NewMemStorage()

	storage.storage["test"] = 123

	tests := []struct {
		name      string
		key       string
		wantValue any
		wantOK    bool
	}{
		{
			name:      "found",
			key:       "test",
			wantValue: 123,
			wantOK:    true,
		},
		{
			name:      "not found",
			key:       "test2",
			wantValue: nil,
			wantOK:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := storage.Get(tt.key)

			assert.Equal(t, tt.wantValue, got)
			assert.Equal(t, tt.wantOK, ok)
		})
	}
}

func TestMemStorage_Save(t *testing.T) {
	storage := NewMemStorage()

	storage.storage["test"] = 123

	type args struct {
		key   string
		value any
	}

	tests := []struct {
		name      string
		args      args
		wantValue any
		wantError bool
	}{
		{
			name: "new save",
			args: args{
				key:   "example",
				value: 123,
			},
			wantValue: 123,
			wantError: false,
		},
		{
			name: "rewrite existing key",
			args: args{
				key:   "test",
				value: 456,
			},
			wantValue: 456,
			wantError: false,
		},
		{
			name: "empty key",
			args: args{
				key:   "",
				value: 1,
			},
			wantValue: nil,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := storage.Save(tt.args.key, tt.args.value)
			if tt.wantError {
				assert.EqualError(t, err, ErrStorageEmptyKey.Error())
				return
			}

			require.NoError(t, err)

			gotValue, _ := storage.Get(tt.args.key)
			assert.Equal(t, tt.wantValue, gotValue)
		})
	}
}
