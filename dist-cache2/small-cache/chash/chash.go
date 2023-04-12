package chash

import (
	"errors"
	"hash/crc32"
	"sort"
	"strconv"
)

// 一致性哈希实现

// CHash 一致性哈希核心结构
type CHash struct {
	replicas      int            // 每个物理节点的分片数量
	keys          []int          // 哈希环  有序的
	virtualToReal map[int]string // 虚拟节点到物理节点的映射
}

func New(replicas int) *CHash {
	return &CHash{
		replicas:      replicas,
		virtualToReal: make(map[int]string),
	}
}

// AddNode 添加物理节点   物理节点的标识是【节点的名称、编号和 IP 地址】组成的字符串
func (ch *CHash) AddNode(nodes ...string) {
	if len(nodes) == 0 {
		return
	}
	for _, node := range nodes {
		for i := 0; i < ch.replicas; i++ { // 每个物理节点有replicas个虚拟节点
			hval := int(crc32.ChecksumIEEE([]byte(node + strconv.Itoa(i))))
			ch.keys = append(ch.keys, hval)
			ch.virtualToReal[hval] = node
		}
	}
	sort.Ints(ch.keys)
}

// GetNode 对于一个key，找到它对应的物理节点
func (ch *CHash) GetNode(key string) (string, error) {
	if len(key) == 0 {
		return "", errors.New("key is invalid.")
	}
	hval := int(crc32.ChecksumIEEE([]byte(key)))
	// 右时针找到最近的虚拟节点
	idx := sort.Search(len(ch.keys), func(i int) bool {
		return ch.keys[i] >= hval
	})
	// 此时说明没有对应的
	// todo：这里可能有问题
	if idx == len(ch.keys) {
		return ch.virtualToReal[ch.keys[0]], nil
	}
	return ch.virtualToReal[ch.keys[idx]], nil
}

func (ch *CHash) VirtualNodeNums() int {
	return len(ch.keys)
}
