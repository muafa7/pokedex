package pokecache

import (
	"testing"
	"time"
)

func TestCacheAddAndGet(t *testing.T) {
	cache := NewCache(1 * time.Minute)

	key := "test-key"
	value := []byte("test-value")

	cache.Add(key, value)

	got, found := cache.Get(key)
	if !found {
		t.Fatalf("expected to find key in cache")
	}

	if string(got) != string(value) {
		t.Fatalf("expected %s, got %s", value, got)
	}
}