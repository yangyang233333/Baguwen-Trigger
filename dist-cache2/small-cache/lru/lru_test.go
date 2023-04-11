package lru

import (
	"testing"
)

func TestCache_Add(t *testing.T) {
	lruCache := New(2)
	lruCache.Add("key_1", "value_1")
	lruCache.Add("key_2", "value_2")
	lruCache.Add("key_3", "value_3")
	if lruCache.Len() != 2 {
		t.Fatalf("lruCache.Len() != 2")
	}
	_, ok := lruCache.Get("key_1")
	if ok {
		t.Fatalf("")
	}
	_, ok = lruCache.Get("key_2")
	if !ok {
		t.Fatalf("")
	}
}
