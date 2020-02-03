//go:generate protoc -I ../helloworld --go_out=plugins=grpc:../helloworld ../helloworld/helloworld.proto

// Package main implements a server for Greeter service.
package grpc

import (
	"log"
	"net"

	"fmt"

	"v2ray.com/core/p2p/config"
	"v2ray.com/core/p2p/wire/pb/seedlist"
	"google.golang.org/grpc"
)

func GRpcServiceStart(srv seedlist.GreeterServer) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Parameters.P2P.GRPCPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	seedlist.RegisterGreeterServer(s, srv)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
