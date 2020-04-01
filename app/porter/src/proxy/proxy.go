/**
 * Description:
 * Author: Yihen.Liu
 * Create: 2020-03-16
 */
package proxy

import (
	"v2ray.com/core/app/porter/config"
	"v2ray.com/core/app/porter/src/proxy/kcp"
	"v2ray.com/core/app/porter/src/proxy/quic"
	"v2ray.com/core/app/porter/src/proxy/tcp"
	"v2ray.com/core/app/porter/src/proxy/udp"
)

func Run() {

	switch config.Parameters.Protocol {
	case "udp":
		udp.Init().StartUDPServer(uint16(config.Parameters.UPort))
	case "kcp":
		kcp.Init().StartKCPServer(uint16(config.Parameters.KPort))
	case "quic":
		quic.Init().StartQuicServer(uint16(config.Parameters.QPort))
	case "tcp":
		tcp.Init().StartTCPServer(uint16(config.Parameters.TPort))
	default:
		udp.Init().StartUDPServer(uint16(config.Parameters.UPort))
		kcp.Init().StartKCPServer(uint16(config.Parameters.KPort))
		quic.Init().StartQuicServer(uint16(config.Parameters.QPort))
		tcp.Init().StartTCPServer(uint16(config.Parameters.TPort))
	}

}
