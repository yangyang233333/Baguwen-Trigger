package lru

import (
	"testing"
)

type String string

func (s String) Len() int {
	return len(s)
}

func TestCache_ALL(t *testing.T) {
	lru := New(2)
	lru.Add("key_1", String("1234"))
	if _, ok := lru.Get("key_1"); !ok {
		t.Fatalf("++++++")
	}
	if lru.Len() != 1 {
		t.Fatalf("++++++")
	}
	lru.Add("key_2", String("1234"))
	lru.Add("key_3", String("1234"))
	if lru.Len() != 2 {
		t.Fatalf("++++++")
	}
	if _, ok := lru.Get("key_1"); ok {
		t.Fatalf("-----")
	}
}
