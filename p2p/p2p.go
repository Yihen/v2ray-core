/**
 * Description:
 * Author: Yihen.Liu
 * Create: 2019-07-22
 */
package p2p

import (
	"context"
	"fmt"

	"time"

	"v2ray.com/core/p2p/channels"
	"v2ray.com/core/p2p/config"
	"v2ray.com/core/p2p/account"
	"v2ray.com/core/p2p/protocol"
	"v2ray.com/core/p2p/protocol/seedlist"
	"v2ray.com/core/p2p/wire/pb/seedlist/types"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core"
	"github.com/libp2p/go-libp2p-core/peer"
	mplex "github.com/libp2p/go-libp2p-mplex"
	"github.com/libp2p/go-libp2p-secio"
	yamux "github.com/libp2p/go-libp2p-yamux"
	"github.com/libp2p/go-tcp-transport"
	ws "github.com/libp2p/go-ws-transport"
	"github.com/multiformats/go-multiaddr"
	"v2ray.com/core/p2p/grpc"
)

const DefaultSleepInterval = 100 * time.Millisecond

type gRpcService struct{}

// SayHello implements helloworld.GreeterServer
func (s *gRpcService) SayHello(ctx context.Context, in *types.HelloSeedList) (*types.HelloReply, error) {
	fmt.Printf("Received: %v", in.Action)
	return &types.HelloReply{}, nil
}

type P2PNode struct {
	gRpcService
	Host      core.Host
	Protocols protocol.Protocol
	ctx       context.Context
}

func (this *P2PNode) DoSeedListRequest(pid peer.ID) {
	this.Protocols.RequestSeedList(pid)
}

func (this *P2PNode) onceBootstrap(seed config.Address) peer.ID {
	targetAddr, err := multiaddr.NewMultiaddr(fmt.Sprintf("/ip4/%s/%s/%d/p2p/%s", seed.IP, seed.Protocol, seed.Port, seed.PubID))
	if err != nil {
		panic(err)
	}

	targetInfo, err := peer.AddrInfoFromP2pAddr(targetAddr)
	if err != nil {
		panic(err)
	}

	err = this.Host.Connect(this.ctx, *targetInfo)
	if err != nil {
		panic(err)
	}
	time.Sleep(DefaultSleepInterval)
	return targetInfo.ID
}

func (this *P2PNode) Bootstrap() {
	for _, seed := range config.Parameters.P2P.SeedList {
		fmt.Println("bootstrap once start:", seed.PubID, seed.Port, seed.Protocol, seed.IP)
		this.onceBootstrap(seed)

		channels.SeedlistNotice <- types.HelloSeedList{
			Action: types.ActionType_SEED_ONLINE,
			Seed: &types.SeedInfo{
				Protocol: config.Parameters.P2P.Address.Protocol,
				Ip:       config.Parameters.P2P.Address.IP,
				Port:     int32(config.Parameters.P2P.Address.Port),
				HostID:   this.Host.ID().Pretty(),
			}}
	}
}

func (this *P2PNode) StartListen(protocol string, ip string, port int, done chan struct{}) {
	p := account.GetAccount().GetPrivKey()

	transports := libp2p.ChainOptions(
		libp2p.Transport(tcp.NewTCPTransport),
		libp2p.Transport(ws.New),
	)

	listenAddrs := libp2p.ListenAddrStrings(
		fmt.Sprintf("/ip4/%s/%s/%d", ip, protocol, port),
	)

	muxers := libp2p.ChainOptions(
		libp2p.Muxer("/yamux/1.0.0", yamux.DefaultTransport),
		libp2p.Muxer("/mplex/6.7.0", mplex.DefaultTransport),
	)

	security := libp2p.Security(secio.ID, secio.New)
	var err error
	this.Host, err = libp2p.New(
		context.Background(),
		transports,
		muxers,
		security,
		listenAddrs,
		libp2p.Identity(p),
	)

	if err != nil {
		fmt.Printf("libp2p.New in StartListen, err:%s", err.Error())
	}
	seedNode := seedlist.NewSeedNode(this.Host, done)
	this.Protocols.SeedListProtocol = seedlist.NewSeedListProtocol(seedNode, done)
	fmt.Println("host.ID:", this.Host.ID().Pretty())
}

func NewP2PNode() *P2PNode {
	return &P2PNode{ctx: context.Background()}
}

func (this *P2PNode)StartGrpcServer()  {
	go grpc.GRpcServiceStart(this)
}

func (this *P2PNode) StartService() {
	this.StartListen(config.Parameters.P2P.Address.Protocol, config.Parameters.P2P.Address.IP, config.Parameters.P2P.Address.Port, make(chan struct{}))
	this.Protocols.StartSeedlistGossipPubSub(this.ctx, this.Host)
	this.Bootstrap()
	this.StartGrpcServer()
}
