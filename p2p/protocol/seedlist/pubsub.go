/**
 * Description:
 * Author: Yihen.Liu
 * Create: 2019-08-01
 */
package seedlist

import (
	"context"
	"fmt"
	"os"

	"v2ray.com/core/p2p/channels"
	"v2ray.com/core/p2p/grpc/p2psocks/p2p"
	"v2ray.com/core/p2p/wire/pb/seedlist/types"
	"github.com/libp2p/go-libp2p-core"
	"github.com/libp2p/go-libp2p-pubsub"
)

const seedlistTopic = "/seedlist/pubsub/0.0.1/notice"

func seedlistSubscribeHandle(ctx context.Context, sub *pubsub.Subscription) {
	for {
		select {
		default:
			msg, err := sub.Next(ctx)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}

			seed := &types.HelloSeedList{}
			err = seed.Unmarshal(msg.Data)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
			grpc_p2p.GRpcClientStart(seed)
			fmt.Println("subscribe receive:", seed.GetAction(), ", data:", seed.Seed, "from:", fmt.Sprintf("%s", msg.GetFrom()))

		case <-ctx.Done():
			return
		}
	}
}

func seedlistPublishHandle(ctx context.Context, ps *pubsub.PubSub) {
	for {
		select {
		case seedInfo := <-channels.SeedlistNotice:
			msgBytes, err := seedInfo.Marshal()
			if err != nil {
				fmt.Printf("seed Marshal err:%s", err.Error())
			}
			ps.Publish(seedlistTopic, msgBytes)
		case <-ctx.Done():
			return
		}
	}
}

func (*SeedListProtocol) StartSeedlistGossipPubSub(ctx context.Context, host core.Host) {
	ps, err := pubsub.NewGossipSub(ctx, host)
	if err != nil {
		panic(err)
	}
	sub, err := ps.Subscribe(seedlistTopic)
	if err != nil {
		panic(err)
	}
	go seedlistSubscribeHandle(ctx, sub)
	go seedlistPublishHandle(ctx, ps)
}
