// Package main implements a client for Greeter service.
package grpc_p2p

import (
	"context"
	"log"
	"os"
	"time"

	"fmt"

	"v2ray.com/core/p2p/config"
	"v2ray.com/core/p2p/wire/pb/seedlist"
	"v2ray.com/core/p2p/wire/pb/seedlist/types"
	"google.golang.org/grpc"
)

const (
	defaultName = "sparrow-client"
)

func GRpcClientStart(seed *types.HelloSeedList) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", config.Parameters.P2P.GRPCPort), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := seedlist.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	_, err = c.SayHello(ctx, seed)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: ,relpy:%s", name)
	cancel()
}
