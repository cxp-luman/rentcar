package main

import (
	authpb "coolcar/auth/api/gen/v1"
	"coolcar/auth/auth"
	"coolcar/auth/wechat"
	"log"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("cannot creat logger: %v", err)
	}
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		logger.Fatal("cannot listen", zap.Error(err))
	}
	s := grpc.NewServer()
	authpb.RegisterAuthServiceServer(s, &auth.Service{
		OpenIDResolver: &wechat.Service{
			AppID: "wx32db26324f2c6b66",
			AppSecret: "82ca641ba23fe0c26dc4d08fb5bfaa3c",
		},
		Logger: logger,
	})
	err = s.Serve(lis)
	if err != nil {
		logger.Fatal("auth cannot serve", zap.Error(err))
	}

 
}