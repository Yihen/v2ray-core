package rpc

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

type Config struct {
	Protocol             string   `protobuf:"bytes,1,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Ip                   string   `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
	Port                 int32    `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba26b287be6b8d69, []int{0}
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

func (m *Config) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *Config) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *Config) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func init() {
	proto.RegisterType((*Config)(nil), "v2ray.core.app.rpc.config")
}

func init() {
	proto.RegisterFile("v2ray.com/core/app/rpc/config.proto", fileDescriptor_ba26b287be6b8d69)
}

var fileDescriptor_ba26b287be6b8d69 = []byte{
	// 170 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2e, 0x33, 0x2a, 0x4a,
	0xac, 0xd4, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xce, 0x2f, 0x4a, 0xd5, 0x4f, 0x2c, 0x28, 0xd0, 0x2f,
	0x2a, 0x48, 0xd6, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x12, 0x82, 0x29, 0x2a, 0x4a, 0xd5, 0x4b, 0x2c, 0x28, 0xd0, 0x2b, 0x2a, 0x48, 0x56, 0xf2, 0xe0,
	0x62, 0x83, 0xa8, 0x11, 0x92, 0xe2, 0xe2, 0x00, 0x2b, 0x4b, 0xce, 0xcf, 0x91, 0x60, 0x54, 0x60,
	0xd4, 0xe0, 0x0c, 0x82, 0xf3, 0x85, 0xf8, 0xb8, 0x98, 0x32, 0x0b, 0x24, 0x98, 0xc0, 0xa2, 0x4c,
	0x99, 0x05, 0x42, 0x42, 0x5c, 0x2c, 0x05, 0xf9, 0x45, 0x25, 0x12, 0xcc, 0x0a, 0x8c, 0x1a, 0xac,
	0x41, 0x60, 0xb6, 0x93, 0x09, 0x97, 0x58, 0x72, 0x7e, 0xae, 0x1e, 0xa6, 0x1d, 0x01, 0x8c, 0x51,
	0xcc, 0x45, 0x05, 0xc9, 0xab, 0x98, 0x84, 0xc2, 0x8c, 0x82, 0x12, 0x2b, 0xf5, 0x9c, 0x41, 0x72,
	0x8e, 0x05, 0x05, 0x7a, 0x41, 0x05, 0xc9, 0x49, 0x6c, 0x60, 0x3b, 0x8c, 0x01, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x7b, 0xbc, 0xa5, 0xf6, 0xc1, 0x00, 0x00, 0x00,
}
