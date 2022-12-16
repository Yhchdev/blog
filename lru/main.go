package main

import "container/list"

type Cacheable interface {
	Key() string
	Size() int
}

type LRUCache struct {
	capacity int
	items    map[string]*LRUCacheItem // hash表  key => 双向链表指针
	list     *list.List
}

type LRUCacheItem struct {
	cacheable   Cacheable
	listElement *list.Element
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		items:    make(map[string]*LRUCacheItem, 9999),
		list:     list.New(),
	}
}

func (c *LRUCache) Get(key string) Cacheable {
	item, ok := c.items[key]
	if !ok {
		return nil
	}
	c.list.MoveToFront(item.listElement)
	return item.cacheable
}

func (c *LRUCache) Set(cacheable Cacheable) {
	// 清理空间
	for {
		if cacheable.Size() < c.capacity {
			break
		}
		c.prune()
	}

	item, ok := c.items[cacheable.Key()]
	if ok {
		item.cacheable = cacheable
		c.capacity -= cacheable.Size() - item.cacheable.Size()
		c.list.MoveToFront(item.listElement)
	} else {
		item := &LRUCacheItem{
			cacheable: cacheable,
		}
		item.listElement = c.list.PushBack(item)
		c.items[cacheable.Key()] = item

		c.capacity -= cacheable.Size()
	}
}

// 淘汰末尾的10个缓存
func (c *LRUCache) prune() {
	for i := 0; i < 10; i++ {
		tail := c.list.Back()
		if tail == nil {
			continue
		}
		item := c.list.Remove(tail).(*LRUCacheItem)
		delete(c.items, item.cacheable.Key())
		c.capacity += item.cacheable.Size()
	}
}

func main() {

}
