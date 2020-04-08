/**
 * Description:
 * Author: Yihen.Liu
 * Create: 2020-04-08
 */
package conf

import (
	"v2ray.com/core/app/rpc"
)

func DefaultRPCConfig() *rpc.Config {
	return &rpc.Config{
		Ip:       "127.0.0.1",
		Port:     10080,
		Protocol: "tcp",
	}
}

type RPCConfig struct {
	Protocol string `json:"protocol"`
	IP       string `json:"ip"`
	Port     int32  `json:"port"`
}

func (v *RPCConfig) Build() *rpc.Config {
	return &rpc.Config{
		Ip:       v.IP,
		Port:     v.Port,
		Protocol: v.Protocol,
	}
}
