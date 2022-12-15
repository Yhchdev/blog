package simple_cache

import (
	"sync"
)

type Cache struct {
	items map[string][]byte
	lock  *sync.RWMutex
}

func New() *Cache {
	return &Cache{
		items: make(map[string][]byte, 2000),
		lock:  new(sync.RWMutex),
	}
}

func (c *Cache) Get(key string) []byte {
	// 取数据只要加读锁
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.items[key]
}

func (c *Cache) Set(key string, data []byte) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.items[key] = data
}
