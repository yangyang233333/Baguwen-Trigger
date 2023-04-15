package consistenhash

import (
	"common"
	"hash/crc32"
	"proxy/meta"
	"sort"
	"strconv"
)

// ConsistentHash 一致性哈希实现
type ConsistentHash struct {
	replicas     int                    // 每个物理节点的分片数量，也就是对应几个虚拟节点
	virtualNodes []int                  // 哈希环上虚拟节点的列表
	virtual2Real map[int]*meta.NodeMeta // 虚拟节点到物理节点的映射
}

func NewConsistentHash() *ConsistentHash {
	return &ConsistentHash{
		replicas:     DefaultReplicas,
		virtualNodes: make([]int, 0),
		virtual2Real: make(map[int]*meta.NodeMeta),
	}
}

// AddRealNode 添加一个新的物理节点
func (ch *ConsistentHash) AddRealNode(addr string) {
	nodeMeta := meta.NewNodeMeta(addr)
	for i := 0; i < ch.replicas; i++ {
		// 根据物理节点的地址、物理节点的编号、虚拟节点的序号i计算虚拟节点的哈希值
		hashVal := int(crc32.ChecksumIEEE([]byte(nodeMeta.NodeAddr + "--" +
			strconv.Itoa(nodeMeta.NodeNumber) + "--" + strconv.Itoa(i)))) // 算一下哈希值
		ch.virtualNodes = append(ch.virtualNodes, hashVal)
		ch.virtual2Real[hashVal] = nodeMeta // 保存对应的物理节点的映射
	}
	sort.Ints(ch.virtualNodes)
}

// GetRealNode 对于一个key 找到它对应的物理节点
func (ch *ConsistentHash) GetRealNode(key string) *meta.NodeMeta {
	if len(key) == 0 {
		return nil
	}
	if len(ch.virtualNodes) == 0 {
		common.LogInstance().Error("虚拟节点数量为0")
		return nil
	}
	// key -> 虚拟节点 -> 物理节点
	hashVal := int(crc32.ChecksumIEEE([]byte(key)))
	idx := sort.Search(len(ch.virtualNodes), func(i int) bool {
		return hashVal <= ch.virtualNodes[i]
	})
	// 因为是环，所以如果 hashVal 比所有节点都大，应该对应第一个节点  所以取余一下就行
	idx = idx % len(ch.virtualNodes)
	return ch.virtual2Real[ch.virtualNodes[idx]]
}
