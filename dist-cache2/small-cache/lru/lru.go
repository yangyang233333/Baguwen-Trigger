package lru

import (
	"container/list"
	"log"
)

type Cache struct {
	curItem_ int // 当前键值对的数量
	maxItem_ int // 键值对最大数量
	list_    *list.List
	cache    map[string]*list.Element
}

type Entry struct {
	key   string
	value string
}

func New(maxItem int) *Cache {
	if maxItem <= 0 {
		log.Fatalf("[ERROR] maxItem <= 0")
	}
	return &Cache{
		curItem_: 0,
		maxItem_: maxItem,
		list_:    list.New(),
		cache:    make(map[string]*list.Element),
	}
}

// Get 查找
func (c *Cache) Get(key string) (value string, ok bool) {
	if elementPtr, ok := c.cache[key]; ok {
		c.list_.MoveToFront(elementPtr)
		entry := elementPtr.Value.(*Entry)
		return entry.value, true
	}
	return "", false
}

// Add 插入
func (c *Cache) Add(key, value string) {
	elementPtr, ok := c.cache[key]
	if ok { // 更新旧值
		c.list_.MoveToFront(elementPtr)
		entry := elementPtr.Value.(*Entry)
		entry.value = value
	} else { // 插入新值
		entry := &Entry{
			key:   key,
			value: value,
		}
		elementPtr = c.list_.PushFront(entry)
		c.cache[key] = elementPtr
		c.curItem_++
		if c.curItem_ > c.maxItem_ {
			backPtr := c.list_.Back()
			backKey := backPtr.Value.(*Entry).key
			delete(c.cache, backKey)
			c.list_.Remove(c.list_.Back())
			c.curItem_--
		}
	}
}

// Len 返回当前缓存的键值对数量
func (c *Cache) Len() int {
	return c.curItem_
}
