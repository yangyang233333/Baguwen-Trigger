package chash

import (
	"testing"
)

func TestCHash(t *testing.T) {
	chashNode := New(3) // 每个节点三个虚拟节点
	chashNode.AddNode("node_1", "node_2", "node_3")

	if chashNode.VirtualNodeNums() != 9 {
		t.Error("节点数量有误")
	}
}
