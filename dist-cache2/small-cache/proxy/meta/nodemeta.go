package meta

import "sync"

// NodeMeta 在线的物理节点的信息
type NodeMeta struct {
	NodeAddr   string // IP:Port
	NodeNumber int    // 节点编号  Proxy中全局唯一
}

var (
	nodeNumber = -1 // 节点编号
	mu         = sync.Mutex{}
)

func NewNodeMeta(NodeAddr string) *NodeMeta {
	mu.Lock()
	defer mu.Unlock()
	nodeNumber++
	return &NodeMeta{
		NodeAddr:   NodeAddr,
		NodeNumber: nodeNumber,
	}
}
