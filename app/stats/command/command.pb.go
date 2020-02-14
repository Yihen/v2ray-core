package command

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type GetStatsRequest struct {
	// Name of the stat counter.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Whether or not to reset the counter to fetching its value.
	Reset_               bool     `protobuf:"varint,2,opt,name=reset,proto3" json:"reset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetStatsRequest) Reset()         { *m = GetStatsRequest{} }
func (m *GetStatsRequest) String() string { return proto.CompactTextString(m) }
func (*GetStatsRequest) ProtoMessage()    {}
func (*GetStatsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c902411c4948f26b, []int{0}
}

func (m *GetStatsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStatsRequest.Unmarshal(m, b)
}
func (m *GetStatsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStatsRequest.Marshal(b, m, deterministic)
}
func (m *GetStatsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStatsRequest.Merge(m, src)
}
func (m *GetStatsRequest) XXX_Size() int {
	return xxx_messageInfo_GetStatsRequest.Size(m)
}
func (m *GetStatsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStatsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetStatsRequest proto.InternalMessageInfo

func (m *GetStatsRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetStatsRequest) GetReset_() bool {
	if m != nil {
		return m.Reset_
	}
	return false
}

type Stat struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value                int64    `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Stat) Reset()         { *m = Stat{} }
func (m *Stat) String() string { return proto.CompactTextString(m) }
func (*Stat) ProtoMessage()    {}
func (*Stat) Descriptor() ([]byte, []int) {
	return fileDescriptor_c902411c4948f26b, []int{1}
}

func (m *Stat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Stat.Unmarshal(m, b)
}
func (m *Stat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Stat.Marshal(b, m, deterministic)
}
func (m *Stat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Stat.Merge(m, src)
}
func (m *Stat) XXX_Size() int {
	return xxx_messageInfo_Stat.Size(m)
}
func (m *Stat) XXX_DiscardUnknown() {
	xxx_messageInfo_Stat.DiscardUnknown(m)
}

var xxx_messageInfo_Stat proto.InternalMessageInfo

func (m *Stat) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Stat) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type GetStatsResponse struct {
	Stat                 *Stat    `protobuf:"bytes,1,opt,name=stat,proto3" json:"stat,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetStatsResponse) Reset()         { *m = GetStatsResponse{} }
func (m *GetStatsResponse) String() string { return proto.CompactTextString(m) }
func (*GetStatsResponse) ProtoMessage()    {}
func (*GetStatsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c902411c4948f26b, []int{2}
}

func (m *GetStatsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStatsResponse.Unmarshal(m, b)
}
func (m *GetStatsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStatsResponse.Marshal(b, m, deterministic)
}
func (m *GetStatsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStatsResponse.Merge(m, src)
}
func (m *GetStatsResponse) XXX_Size() int {
	return xxx_messageInfo_GetStatsResponse.Size(m)
}
func (m *GetStatsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStatsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetStatsResponse proto.InternalMessageInfo

func (m *GetStatsResponse) GetStat() *Stat {
	if m != nil {
		return m.Stat
	}
	return nil
}

type QueryStatsRequest struct {
	Pattern              string   `protobuf:"bytes,1,opt,name=pattern,proto3" json:"pattern,omitempty"`
	Reset_               bool     `protobuf:"varint,2,opt,name=reset,proto3" json:"reset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryStatsRequest) Reset()         { *m = QueryStatsRequest{} }
func (m *QueryStatsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryStatsRequest) ProtoMessage()    {}
func (*QueryStatsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c902411c4948f26b, []int{3}
}

func (m *QueryStatsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryStatsRequest.Unmarshal(m, b)
}
func (m *QueryStatsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryStatsRequest.Marshal(b, m, deterministic)
}
func (m *QueryStatsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryStatsRequest.Merge(m, src)
}
func (m *QueryStatsRequest) XXX_Size() int {
	return xxx_messageInfo_QueryStatsRequest.Size(m)
}
func (m *QueryStatsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryStatsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryStatsRequest proto.InternalMessageInfo

func (m *QueryStatsRequest) GetPattern() string {
	if m != nil {
		return m.Pattern
	}
	return ""
}

func (m *QueryStatsRequest) GetReset_() bool {
	if m != nil {
		return m.Reset_
	}
	return false
}

type QueryStatsResponse struct {
	Stat                 []*Stat  `protobuf:"bytes,1,rep,name=stat,proto3" json:"stat,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryStatsResponse) Reset()         { *m = QueryStatsResponse{} }
func (m *QueryStatsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryStatsResponse) ProtoMessage()    {}
func (*QueryStatsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c902411c4948f26b, []int{4}
}

func (m *QueryStatsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryStatsResponse.Unmarshal(m, b)
}
func (m *QueryStatsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryStatsResponse.Marshal(b, m, deterministic)
}
func (m *QueryStatsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryStatsResponse.Merge(m, src)
}
func (m *QueryStatsResponse) XXX_Size() int {
	return xxx_messageInfo_QueryStatsResponse.Size(m)
}
func (m *QueryStatsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryStatsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryStatsResponse proto.InternalMessageInfo

func (m *QueryStatsResponse) GetStat() []*Stat {
	if m != nil {
		return m.Stat
	}
	return nil
}

type SysStatsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SysStatsRequest) Reset()         { *m = SysStatsRequest{} }
func (m *SysStatsRequest) String() string { return proto.CompactTextString(m) }
func (*SysStatsRequest) ProtoMessage()    {}
func (*SysStatsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c902411c4948f26b, []int{5}
}

func (m *SysStatsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SysStatsRequest.Unmarshal(m, b)
}
func (m *SysStatsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SysStatsRequest.Marshal(b, m, deterministic)
}
func (m *SysStatsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SysStatsRequest.Merge(m, src)
}
func (m *SysStatsRequest) XXX_Size() int {
	return xxx_messageInfo_SysStatsRequest.Size(m)
}
func (m *SysStatsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SysStatsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SysStatsRequest proto.InternalMessageInfo

type SysStatsResponse struct {
	NumGoroutine         uint32   `protobuf:"varint,1,opt,name=NumGoroutine,proto3" json:"NumGoroutine,omitempty"`
	NumGC                uint32   `protobuf:"varint,2,opt,name=NumGC,proto3" json:"NumGC,omitempty"`
	Alloc                uint64   `protobuf:"varint,3,opt,name=Alloc,proto3" json:"Alloc,omitempty"`
	TotalAlloc           uint64   `protobuf:"varint,4,opt,name=TotalAlloc,proto3" json:"TotalAlloc,omitempty"`
	Sys                  uint64   `protobuf:"varint,5,opt,name=Sys,proto3" json:"Sys,omitempty"`
	Mallocs              uint64   `protobuf:"varint,6,opt,name=Mallocs,proto3" json:"Mallocs,omitempty"`
	Frees                uint64   `protobuf:"varint,7,opt,name=Frees,proto3" json:"Frees,omitempty"`
	LiveObjects          uint64   `protobuf:"varint,8,opt,name=LiveObjects,proto3" json:"LiveObjects,omitempty"`
	PauseTotalNs         uint64   `protobuf:"varint,9,opt,name=PauseTotalNs,proto3" json:"PauseTotalNs,omitempty"`
	Uptime               uint32   `protobuf:"varint,10,opt,name=Uptime,proto3" json:"Uptime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SysStatsResponse) Reset()         { *m = SysStatsResponse{} }
func (m *SysStatsResponse) String() string { return proto.CompactTextString(m) }
func (*SysStatsResponse) ProtoMessage()    {}
func (*SysStatsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c902411c4948f26b, []int{6}
}

func (m *SysStatsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SysStatsResponse.Unmarshal(m, b)
}
func (m *SysStatsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SysStatsResponse.Marshal(b, m, deterministic)
}
func (m *SysStatsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SysStatsResponse.Merge(m, src)
}
func (m *SysStatsResponse) XXX_Size() int {
	return xxx_messageInfo_SysStatsResponse.Size(m)
}
func (m *SysStatsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SysStatsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SysStatsResponse proto.InternalMessageInfo

func (m *SysStatsResponse) GetNumGoroutine() uint32 {
	if m != nil {
		return m.NumGoroutine
	}
	return 0
}

func (m *SysStatsResponse) GetNumGC() uint32 {
	if m != nil {
		return m.NumGC
	}
	return 0
}

func (m *SysStatsResponse) GetAlloc() uint64 {
	if m != nil {
		return m.Alloc
	}
	return 0
}

func (m *SysStatsResponse) GetTotalAlloc() uint64 {
	if m != nil {
		return m.TotalAlloc
	}
	return 0
}

func (m *SysStatsResponse) GetSys() uint64 {
	if m != nil {
		return m.Sys
	}
	return 0
}

func (m *SysStatsResponse) GetMallocs() uint64 {
	if m != nil {
		return m.Mallocs
	}
	return 0
}

func (m *SysStatsResponse) GetFrees() uint64 {
	if m != nil {
		return m.Frees
	}
	return 0
}

func (m *SysStatsResponse) GetLiveObjects() uint64 {
	if m != nil {
		return m.LiveObjects
	}
	return 0
}

func (m *SysStatsResponse) GetPauseTotalNs() uint64 {
	if m != nil {
		return m.PauseTotalNs
	}
	return 0
}

func (m *SysStatsResponse) GetUptime() uint32 {
	if m != nil {
		return m.Uptime
	}
	return 0
}

type Config struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_c902411c4948f26b, []int{7}
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

func init() {
	proto.RegisterType((*GetStatsRequest)(nil), "v2ray.core.app.stats.command.GetStatsRequest")
	proto.RegisterType((*Stat)(nil), "v2ray.core.app.stats.command.Stat")
	proto.RegisterType((*GetStatsResponse)(nil), "v2ray.core.app.stats.command.GetStatsResponse")
	proto.RegisterType((*QueryStatsRequest)(nil), "v2ray.core.app.stats.command.QueryStatsRequest")
	proto.RegisterType((*QueryStatsResponse)(nil), "v2ray.core.app.stats.command.QueryStatsResponse")
	proto.RegisterType((*SysStatsRequest)(nil), "v2ray.core.app.stats.command.SysStatsRequest")
	proto.RegisterType((*SysStatsResponse)(nil), "v2ray.core.app.stats.command.SysStatsResponse")
	proto.RegisterType((*Config)(nil), "v2ray.core.app.stats.command.Config")
}

func init() {
	proto.RegisterFile("v2ray.com/core/app/stats/command/command.proto", fileDescriptor_c902411c4948f26b)
}

var fileDescriptor_c902411c4948f26b = []byte{
	// 488 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x5d, 0x6f, 0xd3, 0x30,
	0x14, 0x5d, 0xdb, 0xac, 0xed, 0x6e, 0x37, 0xad, 0xb3, 0x10, 0xb2, 0xa6, 0x09, 0x45, 0x7e, 0xea,
	0x0b, 0xce, 0x54, 0x24, 0x5e, 0x78, 0x1a, 0x91, 0x98, 0x84, 0xca, 0x18, 0x29, 0xf0, 0xc0, 0x9b,
	0x17, 0x2e, 0x28, 0xd0, 0xc4, 0x9e, 0xed, 0x54, 0xca, 0xdf, 0xe1, 0x91, 0x1f, 0xc5, 0x6f, 0x41,
	0x76, 0x12, 0xfa, 0x31, 0xad, 0x1b, 0x4f, 0xf5, 0x39, 0xf7, 0x9e, 0xdc, 0x73, 0xed, 0xa3, 0x02,
	0x5f, 0x4e, 0xb5, 0xa8, 0x78, 0x2a, 0xf3, 0x28, 0x95, 0x1a, 0x23, 0xa1, 0x54, 0x64, 0xac, 0xb0,
	0x26, 0x4a, 0x65, 0x9e, 0x8b, 0xe2, 0x6b, 0xfb, 0xcb, 0x95, 0x96, 0x56, 0x92, 0xb3, 0xb6, 0x5f,
	0x23, 0x17, 0x4a, 0x71, 0xdf, 0xcb, 0x9b, 0x1e, 0xf6, 0x0a, 0x8e, 0x2f, 0xd1, 0xce, 0x1d, 0x97,
	0xe0, 0x6d, 0x89, 0xc6, 0x12, 0x02, 0x41, 0x21, 0x72, 0xa4, 0x9d, 0xb0, 0x33, 0x39, 0x48, 0xfc,
	0x99, 0x3c, 0x81, 0x7d, 0x8d, 0x06, 0x2d, 0xed, 0x86, 0x9d, 0xc9, 0x30, 0xa9, 0x01, 0x3b, 0x87,
	0xc0, 0x29, 0xef, 0x53, 0x2c, 0xc5, 0xa2, 0x44, 0xaf, 0xe8, 0x25, 0x35, 0x60, 0x6f, 0x61, 0xbc,
	0x1a, 0x67, 0x94, 0x2c, 0x0c, 0x92, 0x97, 0x10, 0x38, 0x4f, 0x5e, 0x3d, 0x9a, 0x32, 0xbe, 0xcb,
	0x2f, 0x77, 0xd2, 0xc4, 0xf7, 0xb3, 0x18, 0x4e, 0x3e, 0x94, 0xa8, 0xab, 0x0d, 0xf3, 0x14, 0x06,
	0x4a, 0x58, 0x8b, 0xba, 0x68, 0xdc, 0xb4, 0xf0, 0x9e, 0x15, 0x66, 0x40, 0xd6, 0x3f, 0x72, 0xc7,
	0x52, 0xef, 0xbf, 0x2c, 0x9d, 0xc0, 0xf1, 0xbc, 0x32, 0xeb, 0x86, 0xd8, 0xaf, 0x2e, 0x8c, 0x57,
	0x5c, 0xf3, 0x7d, 0x06, 0x87, 0x57, 0x65, 0x7e, 0x29, 0xb5, 0x2c, 0x6d, 0x56, 0xd4, 0x17, 0x77,
	0x94, 0x6c, 0x70, 0xce, 0xaf, 0xc3, 0xb1, 0xf7, 0x7b, 0x94, 0xd4, 0xc0, 0xb1, 0x17, 0x8b, 0x85,
	0x4c, 0x69, 0x2f, 0xec, 0x4c, 0x82, 0xa4, 0x06, 0xe4, 0x19, 0xc0, 0x47, 0x69, 0xc5, 0xa2, 0x2e,
	0x05, 0xbe, 0xb4, 0xc6, 0x90, 0x31, 0xf4, 0xe6, 0x95, 0xa1, 0xfb, 0xbe, 0xe0, 0x8e, 0xee, 0x9e,
	0xde, 0x09, 0x57, 0x33, 0xb4, 0xef, 0xd9, 0x16, 0xba, 0x09, 0x6f, 0x34, 0xa2, 0xa1, 0x83, 0x7a,
	0x82, 0x07, 0x24, 0x84, 0xd1, 0x2c, 0x5b, 0xe2, 0xfb, 0x9b, 0x1f, 0x98, 0x5a, 0x43, 0x87, 0xbe,
	0xb6, 0x4e, 0xb9, 0x9d, 0xae, 0x45, 0x69, 0xd0, 0x8f, 0xbd, 0x32, 0xf4, 0xc0, 0xb7, 0x6c, 0x70,
	0xe4, 0x29, 0xf4, 0x3f, 0x29, 0x9b, 0xe5, 0x48, 0xc1, 0x2f, 0xd5, 0x20, 0x36, 0x84, 0x7e, 0x2c,
	0x8b, 0x6f, 0xd9, 0xf7, 0xe9, 0x9f, 0x2e, 0x1c, 0xfa, 0xbb, 0x9a, 0xa3, 0x5e, 0x66, 0x29, 0x92,
	0x9f, 0x30, 0x6c, 0x13, 0x43, 0x9e, 0xef, 0x7e, 0x88, 0xad, 0x20, 0x9f, 0xf2, 0xc7, 0xb6, 0xd7,
	0xaf, 0xc2, 0xf6, 0xc8, 0x2d, 0xc0, 0x2a, 0x0d, 0x24, 0xda, 0xad, 0xbf, 0x13, 0xbe, 0xd3, 0xf3,
	0xc7, 0x0b, 0xfe, 0x8d, 0x2c, 0x60, 0xe4, 0x8c, 0x34, 0x09, 0x79, 0x68, 0xc5, 0xad, 0x74, 0x3d,
	0xb4, 0xe2, 0x76, 0xf0, 0xd8, 0xde, 0xeb, 0x19, 0x84, 0xa9, 0xcc, 0x77, 0xca, 0xae, 0x3b, 0x5f,
	0x06, 0xcd, 0xf1, 0x77, 0xf7, 0xec, 0xf3, 0x34, 0x11, 0x15, 0x8f, 0x5d, 0xe7, 0x85, 0x52, 0x3e,
	0xed, 0x86, 0xc7, 0x75, 0xf9, 0xa6, 0xef, 0xff, 0x63, 0x5e, 0xfc, 0x0d, 0x00, 0x00, 0xff, 0xff,
	0x25, 0xa8, 0x5c, 0x1b, 0x95, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StatsServiceClient is the client API for StatsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StatsServiceClient interface {
	GetStats(ctx context.Context, in *GetStatsRequest, opts ...grpc.CallOption) (*GetStatsResponse, error)
	QueryStats(ctx context.Context, in *QueryStatsRequest, opts ...grpc.CallOption) (*QueryStatsResponse, error)
	GetSysStats(ctx context.Context, in *SysStatsRequest, opts ...grpc.CallOption) (*SysStatsResponse, error)
}

type statsServiceClient struct {
	cc *grpc.ClientConn
}

func NewStatsServiceClient(cc *grpc.ClientConn) StatsServiceClient {
	return &statsServiceClient{cc}
}

func (c *statsServiceClient) GetStats(ctx context.Context, in *GetStatsRequest, opts ...grpc.CallOption) (*GetStatsResponse, error) {
	out := new(GetStatsResponse)
	err := c.cc.Invoke(ctx, "/v2ray.core.app.stats.command.StatsService/GetStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statsServiceClient) QueryStats(ctx context.Context, in *QueryStatsRequest, opts ...grpc.CallOption) (*QueryStatsResponse, error) {
	out := new(QueryStatsResponse)
	err := c.cc.Invoke(ctx, "/v2ray.core.app.stats.command.StatsService/QueryStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statsServiceClient) GetSysStats(ctx context.Context, in *SysStatsRequest, opts ...grpc.CallOption) (*SysStatsResponse, error) {
	out := new(SysStatsResponse)
	err := c.cc.Invoke(ctx, "/v2ray.core.app.stats.command.StatsService/GetSysStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StatsServiceServer is the server API for StatsService service.
type StatsServiceServer interface {
	GetStats(context.Context, *GetStatsRequest) (*GetStatsResponse, error)
	QueryStats(context.Context, *QueryStatsRequest) (*QueryStatsResponse, error)
	GetSysStats(context.Context, *SysStatsRequest) (*SysStatsResponse, error)
}

func RegisterStatsServiceServer(s *grpc.Server, srv StatsServiceServer) {
	s.RegisterService(&_StatsService_serviceDesc, srv)
}

func _StatsService_GetStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsServiceServer).GetStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v2ray.core.app.stats.command.StatsService/GetStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsServiceServer).GetStats(ctx, req.(*GetStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatsService_QueryStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsServiceServer).QueryStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v2ray.core.app.stats.command.StatsService/QueryStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsServiceServer).QueryStats(ctx, req.(*QueryStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatsService_GetSysStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsServiceServer).GetSysStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v2ray.core.app.stats.command.StatsService/GetSysStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsServiceServer).GetSysStats(ctx, req.(*SysStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StatsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v2ray.core.app.stats.command.StatsService",
	HandlerType: (*StatsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStats",
			Handler:    _StatsService_GetStats_Handler,
		},
		{
			MethodName: "QueryStats",
			Handler:    _StatsService_QueryStats_Handler,
		},
		{
			MethodName: "GetSysStats",
			Handler:    _StatsService_GetSysStats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v2ray.com/core/app/stats/command/command.proto",
}
