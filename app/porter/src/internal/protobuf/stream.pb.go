package protobuf

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

type ID struct {
	// net_key of the peer (we no longer use the public key as the peer ID, but use it to verify messages)
	NetKey []byte `protobuf:"bytes,1,opt,name=net_key,json=netKey,proto3" json:"net_key,omitempty"`
	// address is the network address of the peer
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// id is the computed hash of the net key
	Id []byte `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	//used in communication with proxy server
	ConnectionId         []byte   `protobuf:"bytes,4,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ID) Reset()         { *m = ID{} }
func (m *ID) String() string { return proto.CompactTextString(m) }
func (*ID) ProtoMessage()    {}
func (*ID) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f7af84f65825b3c, []int{0}
}

func (m *ID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ID.Unmarshal(m, b)
}
func (m *ID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ID.Marshal(b, m, deterministic)
}
func (m *ID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ID.Merge(m, src)
}
func (m *ID) XXX_Size() int {
	return xxx_messageInfo_ID.Size(m)
}
func (m *ID) XXX_DiscardUnknown() {
	xxx_messageInfo_ID.DiscardUnknown(m)
}

var xxx_messageInfo_ID proto.InternalMessageInfo

func (m *ID) GetNetKey() []byte {
	if m != nil {
		return m.NetKey
	}
	return nil
}

func (m *ID) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ID) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *ID) GetConnectionId() []byte {
	if m != nil {
		return m.ConnectionId
	}
	return nil
}

type Message struct {
	Message []byte `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	// Sender's address and net key.
	Sender *ID `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	// Sender's signature of message.
	Signature []byte `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	// request_nonce is the request/response ID. Null if ID associated to a message is not a request/response.
	RequestNonce uint64 `protobuf:"varint,4,opt,name=request_nonce,json=requestNonce,proto3" json:"request_nonce,omitempty"`
	// message_nonce is the sequence ID.
	MessageNonce uint64 `protobuf:"varint,5,opt,name=message_nonce,json=messageNonce,proto3" json:"message_nonce,omitempty"`
	// reply_flag indicates this is a reply to a request
	ReplyFlag bool `protobuf:"varint,6,opt,name=reply_flag,json=replyFlag,proto3" json:"reply_flag,omitempty"`
	// opcode specifies the message type
	Opcode               uint32   `protobuf:"varint,7,opt,name=opcode,proto3" json:"opcode,omitempty"`
	NetID                uint32   `protobuf:"varint,8,opt,name=netID,proto3" json:"netID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f7af84f65825b3c, []int{1}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *Message) GetSender() *ID {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *Message) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *Message) GetRequestNonce() uint64 {
	if m != nil {
		return m.RequestNonce
	}
	return 0
}

func (m *Message) GetMessageNonce() uint64 {
	if m != nil {
		return m.MessageNonce
	}
	return 0
}

func (m *Message) GetReplyFlag() bool {
	if m != nil {
		return m.ReplyFlag
	}
	return false
}

func (m *Message) GetOpcode() uint32 {
	if m != nil {
		return m.Opcode
	}
	return 0
}

func (m *Message) GetNetID() uint32 {
	if m != nil {
		return m.NetID
	}
	return 0
}

type Ping struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ping) Reset()         { *m = Ping{} }
func (m *Ping) String() string { return proto.CompactTextString(m) }
func (*Ping) ProtoMessage()    {}
func (*Ping) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f7af84f65825b3c, []int{2}
}

func (m *Ping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ping.Unmarshal(m, b)
}
func (m *Ping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ping.Marshal(b, m, deterministic)
}
func (m *Ping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ping.Merge(m, src)
}
func (m *Ping) XXX_Size() int {
	return xxx_messageInfo_Ping.Size(m)
}
func (m *Ping) XXX_DiscardUnknown() {
	xxx_messageInfo_Ping.DiscardUnknown(m)
}

var xxx_messageInfo_Ping proto.InternalMessageInfo

type Pong struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pong) Reset()         { *m = Pong{} }
func (m *Pong) String() string { return proto.CompactTextString(m) }
func (*Pong) ProtoMessage()    {}
func (*Pong) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f7af84f65825b3c, []int{3}
}

func (m *Pong) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pong.Unmarshal(m, b)
}
func (m *Pong) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pong.Marshal(b, m, deterministic)
}
func (m *Pong) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pong.Merge(m, src)
}
func (m *Pong) XXX_Size() int {
	return xxx_messageInfo_Pong.Size(m)
}
func (m *Pong) XXX_DiscardUnknown() {
	xxx_messageInfo_Pong.DiscardUnknown(m)
}

var xxx_messageInfo_Pong proto.InternalMessageInfo

type LookupNodeRequest struct {
	Target               *ID      `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LookupNodeRequest) Reset()         { *m = LookupNodeRequest{} }
func (m *LookupNodeRequest) String() string { return proto.CompactTextString(m) }
func (*LookupNodeRequest) ProtoMessage()    {}
func (*LookupNodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f7af84f65825b3c, []int{4}
}

func (m *LookupNodeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LookupNodeRequest.Unmarshal(m, b)
}
func (m *LookupNodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LookupNodeRequest.Marshal(b, m, deterministic)
}
func (m *LookupNodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LookupNodeRequest.Merge(m, src)
}
func (m *LookupNodeRequest) XXX_Size() int {
	return xxx_messageInfo_LookupNodeRequest.Size(m)
}
func (m *LookupNodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LookupNodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LookupNodeRequest proto.InternalMessageInfo

func (m *LookupNodeRequest) GetTarget() *ID {
	if m != nil {
		return m.Target
	}
	return nil
}

type LookupNodeResponse struct {
	Peers                []*ID    `protobuf:"bytes,1,rep,name=peers,proto3" json:"peers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LookupNodeResponse) Reset()         { *m = LookupNodeResponse{} }
func (m *LookupNodeResponse) String() string { return proto.CompactTextString(m) }
func (*LookupNodeResponse) ProtoMessage()    {}
func (*LookupNodeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f7af84f65825b3c, []int{5}
}

func (m *LookupNodeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LookupNodeResponse.Unmarshal(m, b)
}
func (m *LookupNodeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LookupNodeResponse.Marshal(b, m, deterministic)
}
func (m *LookupNodeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LookupNodeResponse.Merge(m, src)
}
func (m *LookupNodeResponse) XXX_Size() int {
	return xxx_messageInfo_LookupNodeResponse.Size(m)
}
func (m *LookupNodeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LookupNodeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LookupNodeResponse proto.InternalMessageInfo

func (m *LookupNodeResponse) GetPeers() []*ID {
	if m != nil {
		return m.Peers
	}
	return nil
}

type Bytes struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Bytes) Reset()         { *m = Bytes{} }
func (m *Bytes) String() string { return proto.CompactTextString(m) }
func (*Bytes) ProtoMessage()    {}
func (*Bytes) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f7af84f65825b3c, []int{6}
}

func (m *Bytes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bytes.Unmarshal(m, b)
}
func (m *Bytes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bytes.Marshal(b, m, deterministic)
}
func (m *Bytes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bytes.Merge(m, src)
}
func (m *Bytes) XXX_Size() int {
	return xxx_messageInfo_Bytes.Size(m)
}
func (m *Bytes) XXX_DiscardUnknown() {
	xxx_messageInfo_Bytes.DiscardUnknown(m)
}

var xxx_messageInfo_Bytes proto.InternalMessageInfo

func (m *Bytes) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type Keepalive struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Keepalive) Reset()         { *m = Keepalive{} }
func (m *Keepalive) String() string { return proto.CompactTextString(m) }
func (*Keepalive) ProtoMessage()    {}
func (*Keepalive) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f7af84f65825b3c, []int{7}
}

func (m *Keepalive) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Keepalive.Unmarshal(m, b)
}
func (m *Keepalive) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Keepalive.Marshal(b, m, deterministic)
}
func (m *Keepalive) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Keepalive.Merge(m, src)
}
func (m *Keepalive) XXX_Size() int {
	return xxx_messageInfo_Keepalive.Size(m)
}
func (m *Keepalive) XXX_DiscardUnknown() {
	xxx_messageInfo_Keepalive.DiscardUnknown(m)
}

var xxx_messageInfo_Keepalive proto.InternalMessageInfo

type KeepaliveResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeepaliveResponse) Reset()         { *m = KeepaliveResponse{} }
func (m *KeepaliveResponse) String() string { return proto.CompactTextString(m) }
func (*KeepaliveResponse) ProtoMessage()    {}
func (*KeepaliveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f7af84f65825b3c, []int{8}
}

func (m *KeepaliveResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeepaliveResponse.Unmarshal(m, b)
}
func (m *KeepaliveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeepaliveResponse.Marshal(b, m, deterministic)
}
func (m *KeepaliveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeepaliveResponse.Merge(m, src)
}
func (m *KeepaliveResponse) XXX_Size() int {
	return xxx_messageInfo_KeepaliveResponse.Size(m)
}
func (m *KeepaliveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_KeepaliveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_KeepaliveResponse proto.InternalMessageInfo

type Disconnect struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Disconnect) Reset()         { *m = Disconnect{} }
func (m *Disconnect) String() string { return proto.CompactTextString(m) }
func (*Disconnect) ProtoMessage()    {}
func (*Disconnect) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f7af84f65825b3c, []int{9}
}

func (m *Disconnect) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Disconnect.Unmarshal(m, b)
}
func (m *Disconnect) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Disconnect.Marshal(b, m, deterministic)
}
func (m *Disconnect) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Disconnect.Merge(m, src)
}
func (m *Disconnect) XXX_Size() int {
	return xxx_messageInfo_Disconnect.Size(m)
}
func (m *Disconnect) XXX_DiscardUnknown() {
	xxx_messageInfo_Disconnect.DiscardUnknown(m)
}

var xxx_messageInfo_Disconnect proto.InternalMessageInfo

type ProxyRequest struct {
	OriginAddress        string   `protobuf:"bytes,1,opt,name=origin_address,json=originAddress,proto3" json:"origin_address,omitempty"`
	ProxyAddress         string   `protobuf:"bytes,2,opt,name=proxy_address,json=proxyAddress,proto3" json:"proxy_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProxyRequest) Reset()         { *m = ProxyRequest{} }
func (m *ProxyRequest) String() string { return proto.CompactTextString(m) }
func (*ProxyRequest) ProtoMessage()    {}
func (*ProxyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f7af84f65825b3c, []int{10}
}

func (m *ProxyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProxyRequest.Unmarshal(m, b)
}
func (m *ProxyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProxyRequest.Marshal(b, m, deterministic)
}
func (m *ProxyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProxyRequest.Merge(m, src)
}
func (m *ProxyRequest) XXX_Size() int {
	return xxx_messageInfo_ProxyRequest.Size(m)
}
func (m *ProxyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProxyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProxyRequest proto.InternalMessageInfo

func (m *ProxyRequest) GetOriginAddress() string {
	if m != nil {
		return m.OriginAddress
	}
	return ""
}

func (m *ProxyRequest) GetProxyAddress() string {
	if m != nil {
		return m.ProxyAddress
	}
	return ""
}

type ProxyResponse struct {
	OriginAddress        string   `protobuf:"bytes,1,opt,name=origin_address,json=originAddress,proto3" json:"origin_address,omitempty"`
	ProxyAddress         string   `protobuf:"bytes,2,opt,name=proxy_address,json=proxyAddress,proto3" json:"proxy_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProxyResponse) Reset()         { *m = ProxyResponse{} }
func (m *ProxyResponse) String() string { return proto.CompactTextString(m) }
func (*ProxyResponse) ProtoMessage()    {}
func (*ProxyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f7af84f65825b3c, []int{11}
}

func (m *ProxyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProxyResponse.Unmarshal(m, b)
}
func (m *ProxyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProxyResponse.Marshal(b, m, deterministic)
}
func (m *ProxyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProxyResponse.Merge(m, src)
}
func (m *ProxyResponse) XXX_Size() int {
	return xxx_messageInfo_ProxyResponse.Size(m)
}
func (m *ProxyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProxyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProxyResponse proto.InternalMessageInfo

func (m *ProxyResponse) GetOriginAddress() string {
	if m != nil {
		return m.OriginAddress
	}
	return ""
}

func (m *ProxyResponse) GetProxyAddress() string {
	if m != nil {
		return m.ProxyAddress
	}
	return ""
}

type MetricRequest struct {
	SendTimestamp        int64    `protobuf:"varint,1,opt,name=send_timestamp,json=sendTimestamp,proto3" json:"send_timestamp,omitempty"`
	RawData              []byte   `protobuf:"bytes,2,opt,name=raw_data,json=rawData,proto3" json:"raw_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetricRequest) Reset()         { *m = MetricRequest{} }
func (m *MetricRequest) String() string { return proto.CompactTextString(m) }
func (*MetricRequest) ProtoMessage()    {}
func (*MetricRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f7af84f65825b3c, []int{12}
}

func (m *MetricRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricRequest.Unmarshal(m, b)
}
func (m *MetricRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricRequest.Marshal(b, m, deterministic)
}
func (m *MetricRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricRequest.Merge(m, src)
}
func (m *MetricRequest) XXX_Size() int {
	return xxx_messageInfo_MetricRequest.Size(m)
}
func (m *MetricRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MetricRequest proto.InternalMessageInfo

func (m *MetricRequest) GetSendTimestamp() int64 {
	if m != nil {
		return m.SendTimestamp
	}
	return 0
}

func (m *MetricRequest) GetRawData() []byte {
	if m != nil {
		return m.RawData
	}
	return nil
}

type MetricResponse struct {
	SendTimestamp        int64    `protobuf:"varint,1,opt,name=send_timestamp,json=sendTimestamp,proto3" json:"send_timestamp,omitempty"`
	RawData              []byte   `protobuf:"bytes,2,opt,name=raw_data,json=rawData,proto3" json:"raw_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetricResponse) Reset()         { *m = MetricResponse{} }
func (m *MetricResponse) String() string { return proto.CompactTextString(m) }
func (*MetricResponse) ProtoMessage()    {}
func (*MetricResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f7af84f65825b3c, []int{13}
}

func (m *MetricResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricResponse.Unmarshal(m, b)
}
func (m *MetricResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricResponse.Marshal(b, m, deterministic)
}
func (m *MetricResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricResponse.Merge(m, src)
}
func (m *MetricResponse) XXX_Size() int {
	return xxx_messageInfo_MetricResponse.Size(m)
}
func (m *MetricResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MetricResponse proto.InternalMessageInfo

func (m *MetricResponse) GetSendTimestamp() int64 {
	if m != nil {
		return m.SendTimestamp
	}
	return 0
}

func (m *MetricResponse) GetRawData() []byte {
	if m != nil {
		return m.RawData
	}
	return nil
}

func init() {
	proto.RegisterType((*ID)(nil), "protobuf.ID")
	proto.RegisterType((*Message)(nil), "protobuf.Message")
	proto.RegisterType((*Ping)(nil), "protobuf.Ping")
	proto.RegisterType((*Pong)(nil), "protobuf.Pong")
	proto.RegisterType((*LookupNodeRequest)(nil), "protobuf.LookupNodeRequest")
	proto.RegisterType((*LookupNodeResponse)(nil), "protobuf.LookupNodeResponse")
	proto.RegisterType((*Bytes)(nil), "protobuf.Bytes")
	proto.RegisterType((*Keepalive)(nil), "protobuf.Keepalive")
	proto.RegisterType((*KeepaliveResponse)(nil), "protobuf.KeepaliveResponse")
	proto.RegisterType((*Disconnect)(nil), "protobuf.Disconnect")
	proto.RegisterType((*ProxyRequest)(nil), "protobuf.ProxyRequest")
	proto.RegisterType((*ProxyResponse)(nil), "protobuf.ProxyResponse")
	proto.RegisterType((*MetricRequest)(nil), "protobuf.MetricRequest")
	proto.RegisterType((*MetricResponse)(nil), "protobuf.MetricResponse")
}

func init() {
	proto.RegisterFile("v2ray.com/core/app/porter/src/internal/protobuf/stream.proto", fileDescriptor_0f7af84f65825b3c)
}

var fileDescriptor_0f7af84f65825b3c = []byte{
	// 517 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xdf, 0x6b, 0xd4, 0x40,
	0x10, 0xc7, 0xc9, 0xfd, 0xc8, 0xdd, 0x4d, 0x93, 0x83, 0xae, 0xa2, 0x11, 0x15, 0x8e, 0xad, 0xc2,
	0x3d, 0x5d, 0xa0, 0xbe, 0x28, 0xf8, 0xa2, 0x1c, 0xc2, 0x51, 0x5b, 0x6a, 0xf0, 0x49, 0x1f, 0xc2,
	0x36, 0x99, 0x86, 0xa5, 0xc9, 0xee, 0xba, 0xbb, 0xd7, 0x9a, 0x7f, 0x5e, 0x24, 0x9b, 0x8d, 0x57,
	0xe9, 0xa3, 0x3e, 0x25, 0xf3, 0x99, 0xd9, 0x99, 0xef, 0xce, 0x7e, 0xe1, 0xfd, 0xed, 0xa9, 0x66,
	0xed, 0xa6, 0x90, 0x4d, 0x5a, 0x48, 0x8d, 0x29, 0x53, 0x2a, 0x55, 0x52, 0x5b, 0xd4, 0xa9, 0xd1,
	0x45, 0xca, 0x85, 0x45, 0x2d, 0x58, 0x9d, 0x2a, 0x2d, 0xad, 0xbc, 0xda, 0x5f, 0xa7, 0xc6, 0x6a,
	0x64, 0xcd, 0xc6, 0xc5, 0x64, 0x3e, 0x60, 0x5a, 0xc3, 0x68, 0xb7, 0x25, 0x4f, 0x61, 0x26, 0xd0,
	0xe6, 0x37, 0xd8, 0x26, 0xc1, 0x2a, 0x58, 0x47, 0x59, 0x28, 0xd0, 0x9e, 0x61, 0x4b, 0x12, 0x98,
	0xb1, 0xb2, 0xd4, 0x68, 0x4c, 0x32, 0x5a, 0x05, 0xeb, 0x45, 0x36, 0x84, 0x64, 0x09, 0x23, 0x5e,
	0x26, 0x63, 0x57, 0x3d, 0xe2, 0x25, 0x39, 0x81, 0xb8, 0x90, 0x42, 0x60, 0x61, 0xb9, 0x14, 0x39,
	0x2f, 0x93, 0x89, 0x4b, 0x45, 0x07, 0xb8, 0x2b, 0xe9, 0xaf, 0x00, 0x66, 0xe7, 0x68, 0x0c, 0xab,
	0xb0, 0x6b, 0xdd, 0xf4, 0xbf, 0x7e, 0xe6, 0x10, 0x92, 0x57, 0x10, 0x1a, 0x14, 0x25, 0x6a, 0x37,
	0xf3, 0xe8, 0x34, 0xda, 0x0c, 0x72, 0x37, 0xbb, 0x6d, 0xe6, 0x73, 0xe4, 0x05, 0x2c, 0x0c, 0xaf,
	0x04, 0xb3, 0x7b, 0x8d, 0x5e, 0xc7, 0x01, 0x74, 0x72, 0x34, 0xfe, 0xd8, 0xa3, 0xb1, 0xb9, 0x90,
	0xa2, 0x40, 0x27, 0x67, 0x92, 0x45, 0x1e, 0x5e, 0x74, 0xac, 0x2b, 0xf2, 0x33, 0x7d, 0xd1, 0xb4,
	0x2f, 0xf2, 0xb0, 0x2f, 0x7a, 0x09, 0xa0, 0x51, 0xd5, 0x6d, 0x7e, 0x5d, 0xb3, 0x2a, 0x09, 0x57,
	0xc1, 0x7a, 0x9e, 0x2d, 0x1c, 0xf9, 0x54, 0xb3, 0x8a, 0x3c, 0x81, 0x50, 0xaa, 0x42, 0x96, 0x98,
	0xcc, 0x56, 0xc1, 0x3a, 0xce, 0x7c, 0x44, 0x1e, 0xc3, 0x54, 0xa0, 0xdd, 0x6d, 0x93, 0xb9, 0xc3,
	0x7d, 0x40, 0x43, 0x98, 0x5c, 0x72, 0x51, 0xb9, 0xaf, 0x14, 0x15, 0x7d, 0x07, 0xc7, 0x9f, 0xa5,
	0xbc, 0xd9, 0xab, 0x0b, 0x59, 0x62, 0xd6, 0x6b, 0xeb, 0xee, 0x6f, 0x99, 0xae, 0xd0, 0xba, 0xc5,
	0x3c, 0xb8, 0x7f, 0x9f, 0xa3, 0x6f, 0x81, 0xdc, 0x3f, 0x6a, 0x94, 0x14, 0x06, 0x09, 0x85, 0xa9,
	0x42, 0xd4, 0x26, 0x09, 0x56, 0xe3, 0x07, 0x47, 0xfb, 0x14, 0x7d, 0x0e, 0xd3, 0x8f, 0xad, 0x45,
	0x43, 0x08, 0x4c, 0x4a, 0x66, 0x99, 0xdf, 0xbf, 0xfb, 0xa7, 0x47, 0xb0, 0x38, 0x43, 0x54, 0xac,
	0xe6, 0xb7, 0x48, 0x1f, 0xc1, 0xf1, 0x9f, 0x60, 0x18, 0x41, 0x23, 0x80, 0x2d, 0x37, 0xfe, 0x5d,
	0xe9, 0x37, 0x88, 0x2e, 0xb5, 0xfc, 0xd9, 0x0e, 0xe2, 0x5f, 0xc3, 0x52, 0x6a, 0x5e, 0x71, 0x91,
	0x0f, 0xc6, 0x09, 0x9c, 0x71, 0xe2, 0x9e, 0x7e, 0xf0, 0xf6, 0x39, 0x81, 0x58, 0x75, 0xc7, 0xf2,
	0xbf, 0xed, 0x15, 0x39, 0xe8, 0x8b, 0xe8, 0x77, 0x88, 0x7d, 0x6f, 0x7f, 0xbb, 0xff, 0xd9, 0xfc,
	0x0b, 0xc4, 0xe7, 0x68, 0x35, 0x2f, 0xee, 0x29, 0xef, 0xac, 0x95, 0x5b, 0xde, 0xa0, 0xb1, 0xac,
	0x51, 0xae, 0xf9, 0x38, 0x8b, 0x3b, 0xfa, 0x75, 0x80, 0xe4, 0x19, 0xcc, 0x35, 0xbb, 0xcb, 0xdd,
	0xe2, 0x46, 0xbd, 0x71, 0x35, 0xbb, 0xdb, 0x76, 0xbb, 0xcb, 0x60, 0x39, 0xb4, 0x3c, 0x08, 0xfe,
	0xb7, 0x9e, 0x57, 0xa1, 0x7b, 0xc0, 0x37, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x84, 0xb8, 0xe1,
	0xcc, 0xf1, 0x03, 0x00, 0x00,
}
