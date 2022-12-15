package simple_cache

import (
	"reflect"
	"testing"
)

func TestCache_Get(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		c    Cache
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Get(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCache_Set(t *testing.T) {
	type args struct {
		key  string
		data []byte
	}
	tests := []struct {
		name string
		c    Cache
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestCache_getShard(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		c    Cache
		args args
		want *ShardCache
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.getShard(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getShard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewCache(t *testing.T) {
	tests := []struct {
		name string
		want *Cache
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCache(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCache() = %v, want %v", got, tt.want)
			}
		})
	}
}
