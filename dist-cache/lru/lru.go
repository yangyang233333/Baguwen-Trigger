package lru

import (
	"container/list"
	"sync"
)

//  实现一个支持并发的LRU缓存

type Cache struct {
	cap_    int64
	hmap_   map[string]*list.Element
	list_   *list.List
	locker_ *sync.RWMutex
}

// Entry 键值对 entry 是双向链表节点的数据类型
type Entry struct {
	key   string
	value Value
}

// Value 为了通用性，我们允许值是实现了 Value 接口的任意类型，该接口只包含了一个方法
// Len() int，用于返回值所占用的内存大小。
type Value interface {
	Len() int
}

func New(cap int64) *Cache {
	return &Cache{
		cap_:    cap,
		hmap_:   make(map[string]*list.Element),
		list_:   list.New(),
		locker_: &sync.RWMutex{},
	}
}

func (c *Cache) Get(key string) (Value, bool) {
	c.locker_.Lock()
	defer c.locker_.Unlock()
	if element, ok := c.hmap_[key]; ok {
		c.list_.MoveToFront(element)
		entry := element.Value.(*Entry) // 一个迷惑操作：实际上链表中存的是*Entry而不是Value
		return entry.value, true
	}
	return nil, false
}

func (c *Cache) Add(key string, value Value) {
	c.locker_.Lock()
	defer c.locker_.Unlock()
	if element, ok := c.hmap_[key]; ok { // Update
		c.list_.MoveToFront(element)
		entry := element.Value.(*Entry)
		entry.value = value
	} else { // Insert
		entry := &Entry{
			key:   key,
			value: value,
		}
		element := c.list_.PushFront(entry)
		c.hmap_[key] = element
		if int64(c.list_.Len()) > c.cap_ {
			delete(c.hmap_, c.list_.Back().Value.(*Entry).key)
			c.list_.Remove(c.list_.Back())
		}
	}
}

func (c *Cache) Len() int64 {
	c.locker_.RLock()
	defer c.locker_.RUnlock()
	return int64(c.list_.Len())
}
