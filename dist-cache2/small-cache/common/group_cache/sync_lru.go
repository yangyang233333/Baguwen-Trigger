package group_cache

import (
	"sync"
)

// SyncLRU LRU的并发版本
type SyncLRU struct {
	mu       *sync.Mutex
	lru      *Cache
	maxItems int
}

func NewSyncLRU(maxItems int) *SyncLRU {
	return &SyncLRU{
		mu:       &sync.Mutex{},
		lru:      New(maxItems),
		maxItems: maxItems,
	}
}

func (sl *SyncLRU) Add(key, value string) {
	sl.mu.Lock()
	defer sl.mu.Unlock()
	sl.lru.Add(key, value)
}

func (sl *SyncLRU) Get(key string) (value string, ok bool) {
	sl.mu.Lock()
	defer sl.mu.Unlock()
	return sl.lru.Get(key)
}
