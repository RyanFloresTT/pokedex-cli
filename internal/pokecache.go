package internal

import (
	"sync"
	"time"
)

type Cache struct {
	cacheMap map[string]cacheEntry
	mu       sync.Mutex
	killTime time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(killTime time.Duration) *Cache {
	c := &Cache{
		cacheMap: make(map[string]cacheEntry),
		killTime: killTime,
	}

	go c.reapLoop()
	return c
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	newEntry := cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
	c.cacheMap[key] = newEntry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, exists := c.cacheMap[key]
	if !exists {
		return nil, false
	}

	return entry.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.killTime)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			c.mu.Lock()
			now := time.Now()
			for key, entry := range c.cacheMap {
				if now.Sub(entry.createdAt) > c.killTime {
					delete(c.cacheMap, key)
				}
			}
			c.mu.Unlock()
		}
	}
}
