package pokecache

import (
	"reflect"
	"testing"
	"time"
)

func TestGetCache(t *testing.T) {
	tests := []struct {
		name string
		want *PokeCache
	}{

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCache(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCache() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPokeCache_Write(t *testing.T) {

	tests := []struct {
		name     string
		key      string
		value    []byte
		expected []byte
	}{
		{
			key:      "Pass",
			value:    []byte("value1"),
			expected: []byte("value1"),
		},
		{
			key:      "nil Value",
			value:    nil,
			expected: nil,
		},
		{
			key:      "Empty Value",
			value:    []byte(""),
			expected: []byte(""),
		},
	}

	for _, tt := range tests {
		t.Run(t.Name()+" "+tt.key, func(t *testing.T) {
			// t.Parallel()
			GetCache().Write(tt.key, tt.value)

			// Check if the key exists in the cache
			value, ok := GetCache().Read(tt.key)
			if !ok {
				t.Errorf("Write() failed to write key %s", tt.key)
			}
			if !reflect.DeepEqual(value, tt.expected) {
				t.Errorf("Write() failed to write or it is wrong value %s expected %c", value, tt.expected)
			}
		})
	}
}

func TestPokeCache_Read(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name  string
		pc    *PokeCache
		args  args
		want  []byte
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.pc.Read(tt.args.key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PokeCache.Read() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("PokeCache.Read() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPokeCache_Delete(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		pc   *PokeCache
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.pc.Delete(tt.args.key)
		})
	}
}

func TestPokeCache_Clear(t *testing.T) {
	tests := []struct {
		name string
		pc   *PokeCache
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.pc.Clear()
		})
	}
}

func TestPokeCache_isEntryExpired(t *testing.T) {
	type args struct {
		entryTime time.Time
	}
	tests := []struct {
		name string
		pc   *PokeCache
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pc.isEntryExpired(tt.args.entryTime); got != tt.want {
				t.Errorf("PokeCache.isEntryExpired() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPokeCache_cacheMaintenance(t *testing.T) {
	tests := []struct {
		name string
		pc   *PokeCache
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.pc.cacheMaintenance()
		})
	}
}
