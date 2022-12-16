package main

import "container/list"

type CacheAble interface {
	Key() string
	Size() int
}

type LRUCache struct {
	capacity int
	items    map[string]*LRUCacheItem // hash表  key => 双向链表指针
	list     *list.List
}

type LRUCacheItem struct {
	CacheAble   CacheAble
	listElement *list.Element
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		items:    make(map[string]*LRUCacheItem, 9999),
		list:     list.New(),
	}
}

func (c *LRUCache) Get(key string) CacheAble {
	item, ok := c.items[key]
	if !ok {
		return nil
	}
	c.list.MoveToFront(item.listElement)
	return item.CacheAble
}

func (c *LRUCache) Set(CacheAble CacheAble) {
	// 清理空间
	for {
		if CacheAble.Size() < c.capacity {
			break
		}
		c.prune()
	}

	item, ok := c.items[CacheAble.Key()]
	if ok {
		item.CacheAble = CacheAble
		c.capacity -= CacheAble.Size() - item.CacheAble.Size()
		c.list.MoveToFront(item.listElement)
	} else {
		item := &LRUCacheItem{
			CacheAble: CacheAble,
		}
		item.listElement = c.list.PushBack(item)
		c.items[CacheAble.Key()] = item

		c.capacity -= CacheAble.Size()
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
		delete(c.items, item.CacheAble.Key())
		c.capacity += item.CacheAble.Size()
	}
}

func main() {

}
