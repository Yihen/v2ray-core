/**
 * Description:
 * Author: Yihen.Liu
 * Create: 2019-07-22
 */
package config

var defaultSocksClientConfig = ClientConfiguration{
	Socks:          ":1080",
	UDPSocksEnable: false,
	RedirTCP:       "",
	RedirTCP6:      "",
	TCPTun:         []string{":8053=8.8.8.8:53", ":8054=8.8.4.4:53"},
	UDPTun:         []string{":8053=8.8.8.8:53", ":8054=8.8.4.4:53"},
}

var defaultSocksConfig = SocksConfiguration{
	Mode:            SERVER_MODE,
	NodeShareEnable: true,
	Verbose:         true,
	Cipher:          "AEAD_CHACHA20_POLY1305",
	Key:             "",
	Keygen:          0,
	Password:        "",
	Server:          "ss://AEAD_CHACHA20_POLY1305:your-password@:8488",
	Client:          "ss://AEAD_CHACHA20_POLY1305:your-password@127.0.0.1:8488",
	UDPTimeout:      5000,
	LogLevel:        1,
	LogDir:          "./log/default/",
	ClientConfig:    defaultSocksClientConfig,
}

var defaultP2PConfig = P2PConfiguration{
	Address: Address{
		Protocol: "tcp",
		IP:       "",
		Port:     1,
		PubID:    ""},
	GRPCPort:   55051,
	NetworkID:  0,
	SocksDBDir: "./socks.db",
}

var Parameters = &Configuration{
	P2P:   defaultP2PConfig,
	Socks: defaultSocksConfig,
}
