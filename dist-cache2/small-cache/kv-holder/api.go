package main

import (
	gcache "common/group_cache"
	"common/proto"
	"context"
)

type KVCacheHolder struct {
	proto.UnimplementedKVCacheHolderServer
}

func (kvch *KVCacheHolder) Add(ctx context.Context, rst *proto.AddRequest) (*proto.AddReply, error) {
	if gcache.GetGroup(rst.GroupName) == nil { // 需要新建一个组
		g := gcache.NewGroup(rst.GroupName, DefaultMaxItems)
		g.Add(rst.Key, rst.Value)
	} else { // 在已有的组里面插入
		gcache.GetGroup(rst.GroupName).Add(rst.Key, rst.Value)
	}
	return &proto.AddReply{Success: true}, nil
}

func (kvch *KVCacheHolder) Get(ctx context.Context, rst *proto.GetRequest) (*proto.GetReply, error) {
	ok := true
	val := ""
	if gcache.GetGroup(rst.GroupName) == nil {
		ok = false
		val = ""
	} else {
		val, ok = gcache.GetGroup(rst.GroupName).Get(rst.Key)
	}

	return &proto.GetReply{
		Exists: ok,
		Value:  val,
	}, nil
}
