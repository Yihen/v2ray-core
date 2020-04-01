package config

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Compression struct {
	Enable               bool     `protobuf:"varint,1,opt,name=enable,proto3" json:"enable,omitempty"`
	CompressAlgo         int32    `protobuf:"varint,2,opt,name=compress_algo,json=compressAlgo,proto3" json:"compress_algo,omitempty"`
	FileSize             int32    `protobuf:"varint,3,opt,name=file_size,json=fileSize,proto3" json:"file_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Compression) Reset()         { *m = Compression{} }
func (m *Compression) String() string { return proto.CompactTextString(m) }
func (*Compression) ProtoMessage()    {}
func (*Compression) Descriptor() ([]byte, []int) {
	return fileDescriptor_16e417d9f6cf30bb, []int{0}
}

func (m *Compression) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Compression.Unmarshal(m, b)
}
func (m *Compression) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Compression.Marshal(b, m, deterministic)
}
func (m *Compression) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Compression.Merge(m, src)
}
func (m *Compression) XXX_Size() int {
	return xxx_messageInfo_Compression.Size(m)
}
func (m *Compression) XXX_DiscardUnknown() {
	xxx_messageInfo_Compression.DiscardUnknown(m)
}

var xxx_messageInfo_Compression proto.InternalMessageInfo

func (m *Compression) GetEnable() bool {
	if m != nil {
		return m.Enable
	}
	return false
}

func (m *Compression) GetCompressAlgo() int32 {
	if m != nil {
		return m.CompressAlgo
	}
	return 0
}

func (m *Compression) GetFileSize() int32 {
	if m != nil {
		return m.FileSize
	}
	return 0
}

type InfluxDb struct {
	URL                  string   `protobuf:"bytes,1,opt,name=URL,proto3" json:"URL,omitempty"`
	DBName               string   `protobuf:"bytes,2,opt,name=DBName,proto3" json:"DBName,omitempty"`
	UserName             string   `protobuf:"bytes,3,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	Password             string   `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Interval             int32    `protobuf:"varint,5,opt,name=interval,proto3" json:"interval,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InfluxDb) Reset()         { *m = InfluxDb{} }
func (m *InfluxDb) String() string { return proto.CompactTextString(m) }
func (*InfluxDb) ProtoMessage()    {}
func (*InfluxDb) Descriptor() ([]byte, []int) {
	return fileDescriptor_16e417d9f6cf30bb, []int{1}
}

func (m *InfluxDb) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InfluxDb.Unmarshal(m, b)
}
func (m *InfluxDb) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InfluxDb.Marshal(b, m, deterministic)
}
func (m *InfluxDb) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InfluxDb.Merge(m, src)
}
func (m *InfluxDb) XXX_Size() int {
	return xxx_messageInfo_InfluxDb.Size(m)
}
func (m *InfluxDb) XXX_DiscardUnknown() {
	xxx_messageInfo_InfluxDb.DiscardUnknown(m)
}

var xxx_messageInfo_InfluxDb proto.InternalMessageInfo

func (m *InfluxDb) GetURL() string {
	if m != nil {
		return m.URL
	}
	return ""
}

func (m *InfluxDb) GetDBName() string {
	if m != nil {
		return m.DBName
	}
	return ""
}

func (m *InfluxDb) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *InfluxDb) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *InfluxDb) GetInterval() int32 {
	if m != nil {
		return m.Interval
	}
	return 0
}

type Config struct {
	InterfaceName        string       `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	InnerIp              string       `protobuf:"bytes,2,opt,name=inner_ip,json=innerIp,proto3" json:"inner_ip,omitempty"`
	PublicIp             string       `protobuf:"bytes,3,opt,name=public_ip,json=publicIp,proto3" json:"public_ip,omitempty"`
	RandomPortBegin      int32        `protobuf:"varint,4,opt,name=random_port_begin,json=randomPortBegin,proto3" json:"random_port_begin,omitempty"`
	RandomPortRange      int32        `protobuf:"varint,5,opt,name=random_port_range,json=randomPortRange,proto3" json:"random_port_range,omitempty"`
	UPort                int32        `protobuf:"varint,6,opt,name=u_port,json=uPort,proto3" json:"u_port,omitempty"`
	KPort                int32        `protobuf:"varint,7,opt,name=k_port,json=kPort,proto3" json:"k_port,omitempty"`
	QPort                int32        `protobuf:"varint,8,opt,name=q_port,json=qPort,proto3" json:"q_port,omitempty"`
	TPort                int32        `protobuf:"varint,9,opt,name=t_port,json=tPort,proto3" json:"t_port,omitempty"`
	PortTimeout          int32        `protobuf:"varint,10,opt,name=port_timeout,json=portTimeout,proto3" json:"port_timeout,omitempty"`
	NetworkId            uint32       `protobuf:"varint,11,opt,name=network_id,json=networkId,proto3" json:"network_id,omitempty"`
	PubKey               string       `protobuf:"bytes,12,opt,name=pub_key,json=pubKey,proto3" json:"pub_key,omitempty"`
	WriteBufferSize      int32        `protobuf:"varint,13,opt,name=write_buffer_size,json=writeBufferSize,proto3" json:"write_buffer_size,omitempty"`
	LogDir               string       `protobuf:"bytes,14,opt,name=log_dir,json=logDir,proto3" json:"log_dir,omitempty"`
	LogLevel             int32        `protobuf:"varint,15,opt,name=log_level,json=logLevel,proto3" json:"log_level,omitempty"`
	PorterDbPath         string       `protobuf:"bytes,16,opt,name=porter_db_path,json=porterDbPath,proto3" json:"porter_db_path,omitempty"`
	Compression          *Compression `protobuf:"bytes,17,opt,name=compression,proto3" json:"compression,omitempty"`
	InfluxDb             *InfluxDb    `protobuf:"bytes,18,opt,name=influx_db,json=influxDb,proto3" json:"influx_db,omitempty"`
	Protocol             string       `protobuf:"bytes,19,opt,name=protocol,proto3" json:"protocol,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_16e417d9f6cf30bb, []int{2}
}

func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *Config) GetInnerIp() string {
	if m != nil {
		return m.InnerIp
	}
	return ""
}

func (m *Config) GetPublicIp() string {
	if m != nil {
		return m.PublicIp
	}
	return ""
}

func (m *Config) GetRandomPortBegin() int32 {
	if m != nil {
		return m.RandomPortBegin
	}
	return 0
}

func (m *Config) GetRandomPortRange() int32 {
	if m != nil {
		return m.RandomPortRange
	}
	return 0
}

func (m *Config) GetUPort() int32 {
	if m != nil {
		return m.UPort
	}
	return 0
}

func (m *Config) GetKPort() int32 {
	if m != nil {
		return m.KPort
	}
	return 0
}

func (m *Config) GetQPort() int32 {
	if m != nil {
		return m.QPort
	}
	return 0
}

func (m *Config) GetTPort() int32 {
	if m != nil {
		return m.TPort
	}
	return 0
}

func (m *Config) GetPortTimeout() int32 {
	if m != nil {
		return m.PortTimeout
	}
	return 0
}

func (m *Config) GetNetworkId() uint32 {
	if m != nil {
		return m.NetworkId
	}
	return 0
}

func (m *Config) GetPubKey() string {
	if m != nil {
		return m.PubKey
	}
	return ""
}

func (m *Config) GetWriteBufferSize() int32 {
	if m != nil {
		return m.WriteBufferSize
	}
	return 0
}

func (m *Config) GetLogDir() string {
	if m != nil {
		return m.LogDir
	}
	return ""
}

func (m *Config) GetLogLevel() int32 {
	if m != nil {
		return m.LogLevel
	}
	return 0
}

func (m *Config) GetPorterDbPath() string {
	if m != nil {
		return m.PorterDbPath
	}
	return ""
}

func (m *Config) GetCompression() *Compression {
	if m != nil {
		return m.Compression
	}
	return nil
}

func (m *Config) GetInfluxDb() *InfluxDb {
	if m != nil {
		return m.InfluxDb
	}
	return nil
}

func (m *Config) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func init() {
	proto.RegisterType((*Compression)(nil), "v2ray.core.app.porter.config.Compression")
	proto.RegisterType((*InfluxDb)(nil), "v2ray.core.app.porter.config.Influx_db")
	proto.RegisterType((*Config)(nil), "v2ray.core.app.porter.config.Config")
}

func init() {
	proto.RegisterFile("v2ray.com/core/app/porter/config/config.proto", fileDescriptor_16e417d9f6cf30bb)
}

var fileDescriptor_16e417d9f6cf30bb = []byte{
	// 598 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0x5d, 0x6f, 0xd3, 0x3e,
	0x14, 0xc6, 0x95, 0xed, 0xdf, 0xae, 0x71, 0xdb, 0xbd, 0xf8, 0x2f, 0xc0, 0xc0, 0x90, 0xca, 0x00,
	0x51, 0x90, 0x48, 0xa5, 0xf1, 0x09, 0xd6, 0xf5, 0xa6, 0xda, 0x84, 0xaa, 0xf0, 0x72, 0xc1, 0x8d,
	0xe5, 0x24, 0x6e, 0x66, 0xd5, 0x8d, 0x3d, 0x27, 0xd9, 0xe8, 0x3e, 0x02, 0x1f, 0x05, 0x89, 0xef,
	0x88, 0x8e, 0x4f, 0x32, 0x26, 0x2e, 0x7a, 0x15, 0x9f, 0xdf, 0xf3, 0xe4, 0xf8, 0xc4, 0x79, 0x4c,
	0x3e, 0xdc, 0x9c, 0x3a, 0xb1, 0x89, 0x52, 0xb3, 0x9e, 0xa4, 0xc6, 0xc9, 0x89, 0xb0, 0x76, 0x62,
	0x8d, 0xab, 0xa4, 0x9b, 0xa4, 0xa6, 0x58, 0xaa, 0xbc, 0x79, 0x44, 0xd6, 0x99, 0xca, 0xd0, 0xe3,
	0xd6, 0xee, 0x64, 0x24, 0xac, 0x8d, 0xd0, 0x1a, 0xa1, 0xe7, 0x24, 0x27, 0xfd, 0x73, 0xb3, 0xb6,
	0x4e, 0x96, 0xa5, 0x32, 0x05, 0x7d, 0x4c, 0xba, 0xb2, 0x10, 0x89, 0x96, 0x2c, 0x18, 0x05, 0xe3,
	0x5e, 0xdc, 0x54, 0xf4, 0x15, 0x19, 0xa6, 0x8d, 0x8d, 0x0b, 0x9d, 0x1b, 0xb6, 0x33, 0x0a, 0xc6,
	0x9d, 0x78, 0xd0, 0xc2, 0x33, 0x9d, 0x1b, 0xfa, 0x9c, 0x84, 0x4b, 0xa5, 0x25, 0x2f, 0xd5, 0x9d,
	0x64, 0xbb, 0xde, 0xd0, 0x03, 0xf0, 0x59, 0xdd, 0xc9, 0x93, 0x9f, 0x01, 0x09, 0xe7, 0xc5, 0x52,
	0xd7, 0x3f, 0x78, 0x96, 0xd0, 0x43, 0xb2, 0xfb, 0x35, 0xbe, 0xf4, 0x9b, 0x84, 0x31, 0x2c, 0x61,
	0xe7, 0xd9, 0xf4, 0x93, 0x58, 0x4b, 0xdf, 0x3a, 0x8c, 0x9b, 0x0a, 0x9a, 0xd6, 0xa5, 0x74, 0xbc,
	0x00, 0x69, 0xd7, 0x4b, 0x3d, 0x00, 0x5e, 0x7c, 0x46, 0x7a, 0x56, 0x94, 0xe5, 0xad, 0x71, 0x19,
	0xfb, 0x0f, 0xb5, 0xb6, 0x06, 0x4d, 0x15, 0x95, 0x74, 0x37, 0x42, 0xb3, 0x0e, 0x0e, 0xd3, 0xd6,
	0x27, 0xbf, 0x3b, 0xa4, 0x7b, 0xee, 0x0f, 0x80, 0xbe, 0x21, 0xfb, 0x1e, 0x2f, 0x45, 0x2a, 0x71,
	0x13, 0x1c, 0x6a, 0x78, 0x4f, 0xfd, 0x4e, 0x4f, 0xa1, 0x5b, 0x21, 0x1d, 0x57, 0xb6, 0x19, 0x70,
	0xcf, 0xd7, 0x73, 0x0b, 0x13, 0xda, 0x3a, 0xd1, 0x2a, 0x05, 0xad, 0x99, 0x10, 0xc1, 0xdc, 0xd2,
	0xf7, 0xe4, 0xc8, 0x89, 0x22, 0x33, 0x6b, 0x0e, 0xe7, 0xce, 0x13, 0x99, 0xab, 0xc2, 0x8f, 0xda,
	0x89, 0x0f, 0x50, 0x58, 0x18, 0x57, 0x4d, 0x01, 0xff, 0xeb, 0x75, 0xa2, 0xc8, 0x65, 0x33, 0xfa,
	0x03, 0x6f, 0x0c, 0x98, 0x3e, 0x22, 0xdd, 0xda, 0xdb, 0x58, 0xd7, 0x1b, 0x3a, 0x35, 0x68, 0x80,
	0x57, 0x88, 0xf7, 0x10, 0xaf, 0x5a, 0x7c, 0x8d, 0xb8, 0x87, 0xf8, 0xba, 0xc5, 0x15, 0xe2, 0x10,
	0x71, 0xe5, 0xf1, 0x4b, 0x32, 0xf0, 0x03, 0x54, 0x6a, 0x2d, 0x4d, 0x5d, 0x31, 0xe2, 0xc5, 0x3e,
	0xb0, 0x2f, 0x88, 0xe8, 0x0b, 0x42, 0x0a, 0x59, 0xdd, 0x1a, 0xb7, 0xe2, 0x2a, 0x63, 0xfd, 0x51,
	0x30, 0x1e, 0xc6, 0x61, 0x43, 0xe6, 0x19, 0x7d, 0x42, 0xf6, 0x6c, 0x9d, 0xf0, 0x95, 0xdc, 0xb0,
	0x01, 0xfe, 0x4d, 0x5b, 0x27, 0x17, 0x72, 0x03, 0x9f, 0x78, 0xeb, 0x54, 0x25, 0x79, 0x52, 0x2f,
	0x97, 0xd2, 0x61, 0x54, 0x86, 0xf8, 0x89, 0x5e, 0x98, 0x7a, 0x0e, 0x89, 0x81, 0x26, 0xda, 0xe4,
	0x3c, 0x53, 0x8e, 0xed, 0x63, 0x13, 0x6d, 0xf2, 0x99, 0x72, 0x70, 0xe0, 0x20, 0x68, 0x79, 0x23,
	0x35, 0x3b, 0xc0, 0x5f, 0xab, 0x4d, 0x7e, 0x09, 0x35, 0x7d, 0x4d, 0xf6, 0x31, 0xe1, 0x3c, 0x4b,
	0xb8, 0x15, 0xd5, 0x15, 0x3b, 0xf4, 0x2f, 0x0f, 0x90, 0xce, 0x92, 0x85, 0xa8, 0xae, 0xe8, 0x05,
	0xe9, 0xa7, 0x7f, 0x63, 0xcf, 0x8e, 0x46, 0xc1, 0xb8, 0x7f, 0xfa, 0x2e, 0xda, 0x76, 0x55, 0xa2,
	0x07, 0xf7, 0x24, 0x7e, 0xf8, 0x36, 0x9d, 0x91, 0x50, 0xb5, 0xc9, 0x66, 0xd4, 0xb7, 0x7a, 0xbb,
	0xbd, 0xd5, 0xfd, 0x45, 0x80, 0x4c, 0xc2, 0x72, 0x96, 0xf8, 0x2c, 0xc3, 0x85, 0x4d, 0x8d, 0x66,
	0xff, 0x37, 0x29, 0x6a, 0xea, 0xe9, 0x05, 0x19, 0xa5, 0x66, 0xbd, 0xb5, 0xe7, 0x22, 0xf8, 0xde,
	0xc5, 0xd5, 0xaf, 0x9d, 0xe3, 0x6f, 0xa7, 0xb1, 0xd8, 0x44, 0xe7, 0x60, 0x3c, 0xb3, 0x36, 0x5a,
	0xa0, 0x11, 0x13, 0x9f, 0x74, 0x7d, 0xdb, 0x8f, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x70, 0x9e,
	0xa4, 0x02, 0x48, 0x04, 0x00, 0x00,
}
