package main

import (
	"common"
	pb "common/proto"
	"context"
	"go.uber.org/zap"
	"proxy/consistenhash"
)

// proxy 对外的 API，后续可以套一个 grpc 或者 http 接口

// 单例 只此一份
var chashOBJ *consistenhash.ConsistentHash

func init() {
	chashOBJ = consistenhash.NewConsistentHash()
}

// Add 向 Group: gName 中写入 <key, val>
func Add(gName, key, val string) bool {
	/*
	   逻辑：
	   	key --> 物理节点 --> rpc --> 写入

	*/
	realNode := chashOBJ.GetRealNode(key) // 计算出对应的物理节点
	cli := addr2Conn[realNode.NodeAddr]
	ctx := context.Background()
	rst := &pb.AddRequest{
		GroupName: gName,
		Key:       key,
		Value:     val,
	}
	rsp, err := cli.Add(ctx, rst)
	if err != nil {
		common.LogInstance().Error("", zap.Error(err))
		return false
	}
	return rsp.Success
}

func Get(gName, key string) (string, bool) {
	/*
	   逻辑：
	   	key --> 物理节点 --> rpc --> 读取
	*/
	realNode := chashOBJ.GetRealNode(key) // 计算出对应的物理节点
	if realNode == nil {
		common.LogInstance().Error(key + "对应的物理节点不存在")
		return "", false
	}
	cli, ok := addr2Conn[realNode.NodeAddr]
	if !ok {
		common.LogInstance().Error(realNode.NodeAddr + "不存在")
		return "", false
	}
	ctx := context.Background()
	rst := &pb.GetRequest{
		GroupName: gName,
		Key:       key,
	}
	rsp, err := cli.Get(ctx, rst)
	if err != nil {
		common.LogInstance().Error("", zap.Error(err))
		return "", false
	}
	return rsp.Value, rsp.Exists
}
