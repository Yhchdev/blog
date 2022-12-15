package simple_cache

import (
	"crypto/sha1"
	"fmt"
	"sync"
)

type Cache map[string]*ShardCache

type ShardCache struct {
	items map[string][]byte
	lock  *sync.RWMutex
}

func NewCache() *Cache {
	cache := make(Cache, 256)
	for i := 0; i < 256; i++ {
		cache[fmt.Sprintf("%02x", i)] = &ShardCache{
			items: make(map[string][]byte, 2000),
			lock:  new(sync.RWMutex),
		}
	}
	return &cache
}

func (c Cache) getShard(key string) *ShardCache {
	hasher := sha1.New()
	hasher.Write([]byte(key))
	shardKey := fmt.Sprintf("%x", hasher.Sum(nil))[0:2]
	return c[shardKey]
}

func (c Cache) Get(key string) []byte {
	// 取数据只要加读锁
	shard := c.getShard(key)
	shard.lock.RLock()
	defer shard.lock.RUnlock()
	return shard.items[key]
}

func (c Cache) Set(key string, data []byte) {
	shard := c.getShard(key)
	shard.lock.Lock()
	shard.lock.Unlock()
	shard.items[key] = data
}
