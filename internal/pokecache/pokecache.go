package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	Entry map[string]cacheEntry
	Mu    sync.Mutex
}

type cacheEntry struct {
	created time.Time
	data    []byte
}

func NewCache(interval time.Duration) Cache {}

func (c *Cache) Add(key string, data []byte) {
	c.Mu.Lock()
	defer c.Mu.Unlock()

	c.Entry[key] = cacheEntry{
		created: time.Now(),
		data:    data,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.Mu.Lock()
	defer c.Mu.Unlock()

	entry, ok := c.Entry[key]
	if !ok {
		return []byte{}, false
	} else {
		return entry.data, true
	}
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
}
