package main

import (
	"common"
	pb "common/proto"
	"flag"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

var (
	port = flag.Int("port", 10050, "端口")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		common.LogInstance().Fatal("", zap.Error(err))
	}

	// 启动一个grpc服务器
	s := grpc.NewServer()
	pb.RegisterKVCacheHolderServer(s, &KVCacheHolder{})
	common.LogInstance().Info(fmt.Sprintf("server listening at %v", lis.Addr()))
	if err := s.Serve(lis); err != nil {
		common.LogInstance().Info(fmt.Sprintf("failed to serve: %v", err))
	}
}
