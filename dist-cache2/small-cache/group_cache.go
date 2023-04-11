package small_cache

import (
	"errors"
	"sync"
)

//一个 Group 可以认为是一个缓存的命名空间，每个 Group 拥有一个唯一的名称 name。
//比如可以创建三个 Group，缓存学生的成绩命名为 scores，缓存学生信息的命名为 info，
//缓存学生课程的命名为 courses。

type Group struct {
	name       string // Group的唯一名字
	localCache *SyncLRU
}

var (
	mu     sync.Mutex
	groups = make(map[string]*Group)
)

func NewGroup(name string, maxItems int) *Group {
	mu.Lock()
	defer mu.Unlock()
	g := &Group{
		name:       name,
		localCache: NewSyncLRU(maxItems),
	}
	groups[name] = g
	return g
}

func GetGroup(name string) *Group {
	mu.Lock()
	defer mu.Unlock()
	g, ok := groups[name]
	if ok {
		return g
	}
	return nil
}

func (g *Group) Get(key string) (string, error) {
	if len(key) == 0 {
		return "", errors.New("key is null")
	}
	if val, ok := g.localCache.Get(key); ok { // 存在于本地缓存，直接返回即可
		return val, nil
	}

	// todo 本地缓存不不存在，需要从peer节点查询
	return "", errors.New("todo")
}
