package main

import (
	"common"
	pb "common/proto"
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// 三个kvholder的地址，这里方其实应该通过注册中心来获取底层的holder的地址
// 简单起见，省略了注册中心，直接采用硬编码的方法
var holderAddrs = []string{"127.0.0.1:10051", "127.0.0.1:10052", "127.0.0.1:10053"}

var addr2Conn = make(map[string]pb.KVCacheHolderClient)

// MakeConn 链接底层的Holder节点
func MakeConn() {
	for _, addr := range holderAddrs {
		// 先添加节点
		chashOBJ.AddRealNode(addr)

		// 建立链接
		conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			common.LogInstance().Fatal("连接失败", zap.Error(err))
			return
		}
		cli := pb.NewKVCacheHolderClient(conn)
		addr2Conn[addr] = cli
	}
}

// Ping 向所有holder定时（周期1秒）发送心跳包，如果在规定时间内没得到回复，就认为该节点已经故障下线
// 这内部是一个异步的线程调用
func Ping() {
	go func() {
		for true {
			for _, cli := range addr2Conn {
				// todo： 发送心跳包检测  如果规定时间没有得到结果，则认为该物理节点下线，应该删除对应的虚拟节点
				cli.Get(context.Background(), &pb.GetRequest{
					GroupName: "",
					Key:       "",
				})
			}
		}
	}()
}
