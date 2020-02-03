/**
 * Description:
 * Author: Yihen.Liu
 * Create: 2019-07-22
 */

package config

import (
	"net"
	"net/url"
	"strconv"
	"strings"
)

const (
	DefaultConfigFilename = "./config.json"
	SERVER_MODE           = "server"
	CLIENT_MODE           = "client"
)

var Version string

type Address struct {
	Protocol string `json:"Protocol"`
	IP       string `json:"IP"`
	Port     int    `json:"Port"`
	PubID    string `json:"PubID"`
}

type P2PConfiguration struct {
	Address    Address   `json:"Address"`
	SeedList   []Address `json:"SeedList"`
	GRPCPort   uint16    `json:"GRPCPort"`
	NetworkID  uint32    `json:"NetworkID"`
	SocksDBDir string    `json:"SocksDBDir"`
}

type ClientConfiguration struct {
	Socks          string   `json:"Socks"`           //SOCKS listen address
	UDPSocksEnable bool     `json:"UDPSocksEnable"`  //Enable UDP support for SOCKS
	RedirTCP       string   `json:"RedirTCPEnable"`  //redirect TCP from this address
	RedirTCP6      string   `json:"RedirTCP6Enable"` //redirect TCP IPv6 from this address
	TCPTun         []string `json:"TCPTun"`          //TCP tunnel (laddr1=raddr1,laddr2=raddr2,...)
	UDPTun         []string `json:"UDPTun"`          //UDP tunnel (laddr1=raddr1,laddr2=raddr2,...)
}

type SocksConfiguration struct {
	Mode            string              `json:"Mode"`            // mode client or server
	NodeShareEnable bool                `json:"NodeShareEnable"` // if true, enable share data flow by dht network
	Verbose         bool                `json:"Verbose"`         // verbose mode
	Cipher          string              `json:"Cipher"`          // available cipher used
	Key             string              `json:"Key"`             // base64url-encoded key (derive from password if empty)
	Keygen          int                 `json:"Keygen"`          // generate a base64url-encoded random key of given length in byte
	Password        string              `json:"Password"`        // password
	Server          string              `json:"Server"`          // server listen address or url
	Client          string              `json:"Client"`          //server listen address or url
	UDPTimeout      int                 `json:"UDPTimeout"`      //UDP tunnel timeout
	LogLevel        int                 `json:"LogLevel"`
	LogDir          string              `json:"LogDir"`
	ClientConfig    ClientConfiguration `json:"ClientConfig"`
}

type Configuration struct {
	P2P   P2PConfiguration   `json:"P2P"`
	Socks SocksConfiguration `json:"Socks"`
}

func ParseAddress(address string) (*Address, error) {
	arr := strings.Split(address, "-")
	if len(arr) < 2 {
		return nil, nil
	}

	urlInfo, err := url.Parse(arr[0])
	if err != nil {
		return nil, err
	}

	host, rawPort, err := net.SplitHostPort(urlInfo.Host)
	if err != nil {
		return nil, err
	}

	port, err := strconv.ParseUint(rawPort, 10, 16)
	if err != nil {
		return nil, err
	}

	return &Address{
		Protocol: urlInfo.Scheme,
		IP:       host,
		Port:     int(port),
		PubID:    arr[1],
	}, nil
}
