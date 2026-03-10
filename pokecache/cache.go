package pokecache

import (
	"sync"
	"time"
)
type Cache struct {
	entries	map[string]cacheEntry
	mu 		sync.Mutex
}

type cacheEntry struct {
	createdAt	time.Time
	val			[]byte
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry := cacheEntry{
		createdAt: time.Now(),
		val: val,
	}

	c.entries[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if entry, ok := c.entries[key]; ok {
		return entry.val, true
	}

	return nil, false
}

func (c *Cache) reapLoop(interval time.Duration) {
	for {
		time.Sleep(interval)

		c.mu.Lock()

		for key, entry := range c.entries {
			if(time.Since(entry.createdAt) > interval) {
				delete(c.entries, key)
			}
		}
		c.mu.Unlock()
	}
}

func NewCache(interval time.Duration) *Cache {
	var cache Cache

	cache.entries = make(map[string]cacheEntry)

	go cache.reapLoop(interval)

	return &cache
}