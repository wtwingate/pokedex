package pokecache

import (
	"fmt"
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

func NewCache(interval time.Duration) *Cache {
	entryMap := make(map[string]cacheEntry)
	cache := Cache{
		Entry: entryMap,
	}
	cache.reapLoop(interval)
	return &cache
}

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
		fmt.Println("*** data from cache ***")
		return entry.data, true
	}
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	go func() {
		for range ticker.C {
			c.Mu.Lock()
			for k, v := range c.Entry {
				if time.Since(v.created) > interval {
					delete(c.Entry, k)
				}
			}
			c.Mu.Unlock()
		}
	}()
}
