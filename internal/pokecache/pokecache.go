package pokecache

import (
	"sync"
	"time"
)

type Cache struct{
	cache map[string] cacheEntry
	mux *sync.Mutex
}

type cacheEntry struct {
	val []byte
	createdAt time.Time
}

func NewCache(interval time.Duration) Cache {

	c := Cache {
		cache: make(map[string] cacheEntry),
		mux: &sync.Mutex{},
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) reapLoop(interval time.Duration){
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(interval)
	}
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.cache[key] = cacheEntry {
		val: val,
		createdAt: time.Now().UTC(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool){
	c.mux.Lock()
	defer c.mux.Unlock()
	data, ok := c.cache[key]
	if !ok {
		return nil, false	
	}

	return data.val, true
}

func (c *Cache) reap(interval time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	fiveMinPast := time.Now().UTC().Add(-interval)
	for k, v := range c.cache {
		if v.createdAt.Before(fiveMinPast) {
			delete(c.cache, k)
		}
	}
}