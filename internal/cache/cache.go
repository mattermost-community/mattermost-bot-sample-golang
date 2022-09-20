package cache

import (
	"sync"
)

type Cache struct {
	mu   sync.RWMutex
	data map[string]interface{}
}

func NewCache() *Cache {
	return &Cache{}
}

func (c *Cache) Set(key string, data interface{}) {
	c.mu.Lock()
	c.data[key] = data
	c.mu.Unlock()
}

func (c *Cache) Get(key string) (data interface{}, ok bool) {
	c.mu.RLock()
	data, ok = c.data[key]
	c.mu.RUnlock()
	return
}

func (c *Cache) Remove(key string) {
	c.mu.Lock()
	delete(c.data, key)
	c.mu.Unlock()
}

func (c *Cache) ClearCache() {
	c.mu.Lock()
	c.data = make(map[string]interface{})
	c.mu.Unlock()
}
