package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]cacheEntry
	mu      *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.entries[key]
	if ok {
		return entry.val, ok
	}
	return nil, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		expirationLimit := time.Now().Add(-interval)
		c.reap(expirationLimit)
	}
}

func (c *Cache) reap(limit time.Time) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k, v := range c.entries {
		if v.createdAt.Before(limit) {
			delete(c.entries, k)
		}
	}
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		entries: map[string]cacheEntry{},
		mu:      &sync.Mutex{},
	}
	go cache.reapLoop(interval)
	return cache
}
