package pokecache

import (
	"time"
	"sync"
)

type Cache struct {
	mu			*sync.Mutex
	CacheEntry	map[string]*cacheEntry
}

type cacheEntry struct {
	CreatedAt	time.Time
	Val			[]byte
}

func NewCache(interval time.Duration) *Cache {
	cache := new(Cache)
	cache.mu = &sync.Mutex{}
	cache.CacheEntry = make(map[string]*cacheEntry)

	go cache.reapLoop(interval)

	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	entry := new(cacheEntry)
	entry.CreatedAt = time.Now()
	entry.Val = val

	c.CacheEntry[key] = entry

	return
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	value, ok := c.CacheEntry[key]
	if !ok {
		return []byte{}, false
	}
	return value.Val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for range ticker.C {
		for i, entry := range c.CacheEntry {
			timeDiff := time.Since(entry.CreatedAt)

			if timeDiff > interval {
				c.mu.Lock()
				defer c.mu.Unlock()

				delete(c.CacheEntry, i)

				c.mu.Unlock()
			}
		}
	}

	return
}