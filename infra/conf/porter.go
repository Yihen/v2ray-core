/**
 * Description:
 * Author: Yihen.Liu
 * Create: 2020-03-16
 */
package conf

import (
	"v2ray.com/core/app/porter/config"
)

type Compression struct {
	Enable       bool  `json:"Enable"`
	CompressAlgo int32 `json:"CompressAlgo"`
	FileSize     int32 `json:"FileSize"`
}

type InfluxDB struct {
	URL      string `json:"URL"`
	DBName   string `json:"DBName"`
	UserName string `json:"UserName"`
	Password string `json:"Password"`
	Interval int32  `json:"Interval"`
}

type PorterConfig struct {
	Protocol        string      `json:"Protocol"`
	InterfaceName   string      `json:"InterfaceName"`
	InnerIP         string      `json:"InnerIP"`
	PublicIP        string      `json:"PublicIP"`
	RandomPortBegin int32       `json:"RandomPortBegin"`
	RandomPortRange int32       `json:"RandomPortRange"`
	UPort           int32       `json:"UPort"`
	KPort           int32       `json:"KPort"`
	QPort           int32       `json:"QPort"`
	TPort           int32       `json:"TPort"`
	PortTimeout     int32       `json:"PortTimeout"`
	NetworkID       uint32      `json:"NetworkID"`
	PubKey          string      `json:"PubKey"`
	WriteBufferSize int32       `json:"WriteBufferSize"`
	LogDir          string      `json:"LogDir"`
	LogLevel        int32       `json:"LogLevel"`
	PorterDBPath    string      `json:"PorterDBPath"`
	Compression     Compression `json:"Compression"`
	InfluxDB        InfluxDB    `json:"InfluxDB"`
}

func (v *PorterConfig) Build() *config.Config {
	return &config.Config{
		Protocol:        v.Protocol,
		InterfaceName:   v.InterfaceName,
		InnerIp:         v.InnerIP,
		PublicIp:        v.PublicIP,
		RandomPortBegin: v.RandomPortBegin,
		RandomPortRange: v.RandomPortRange,
		UPort:           v.UPort,
		KPort:           v.KPort,
		QPort:           v.QPort,
		TPort:           v.TPort,
		PortTimeout:     v.PortTimeout,
		NetworkId:       v.NetworkID,
		PubKey:          v.PubKey,
		WriteBufferSize: v.WriteBufferSize,
		LogDir:          v.LogDir,
		LogLevel:        v.LogLevel,
		PorterDbPath:    v.PorterDBPath,
		Compression: &config.Compression{
			Enable:       v.Compression.Enable,
			CompressAlgo: v.Compression.CompressAlgo,
			FileSize:     v.Compression.FileSize,
		},
	}
}
