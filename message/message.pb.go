// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package BCDns_daemon

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type SwapCertMsg struct {
	Ip                   string   `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SwapCertMsg) Reset()         { *m = SwapCertMsg{} }
func (m *SwapCertMsg) String() string { return proto.CompactTextString(m) }
func (*SwapCertMsg) ProtoMessage()    {}
func (*SwapCertMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}

func (m *SwapCertMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SwapCertMsg.Unmarshal(m, b)
}
func (m *SwapCertMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SwapCertMsg.Marshal(b, m, deterministic)
}
func (m *SwapCertMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SwapCertMsg.Merge(m, src)
}
func (m *SwapCertMsg) XXX_Size() int {
	return xxx_messageInfo_SwapCertMsg.Size(m)
}
func (m *SwapCertMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_SwapCertMsg.DiscardUnknown(m)
}

var xxx_messageInfo_SwapCertMsg proto.InternalMessageInfo

func (m *SwapCertMsg) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

type StartServerReq struct {
	Byzantine            bool     `protobuf:"varint,1,opt,name=byzantine,proto3" json:"byzantine,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartServerReq) Reset()         { *m = StartServerReq{} }
func (m *StartServerReq) String() string { return proto.CompactTextString(m) }
func (*StartServerReq) ProtoMessage()    {}
func (*StartServerReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{1}
}

func (m *StartServerReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartServerReq.Unmarshal(m, b)
}
func (m *StartServerReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartServerReq.Marshal(b, m, deterministic)
}
func (m *StartServerReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartServerReq.Merge(m, src)
}
func (m *StartServerReq) XXX_Size() int {
	return xxx_messageInfo_StartServerReq.Size(m)
}
func (m *StartServerReq) XXX_DiscardUnknown() {
	xxx_messageInfo_StartServerReq.DiscardUnknown(m)
}

var xxx_messageInfo_StartServerReq proto.InternalMessageInfo

func (m *StartServerReq) GetByzantine() bool {
	if m != nil {
		return m.Byzantine
	}
	return false
}

type StartServerRep struct {
	IsLeader             bool     `protobuf:"varint,1,opt,name=isLeader,proto3" json:"isLeader,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartServerRep) Reset()         { *m = StartServerRep{} }
func (m *StartServerRep) String() string { return proto.CompactTextString(m) }
func (*StartServerRep) ProtoMessage()    {}
func (*StartServerRep) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{2}
}

func (m *StartServerRep) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartServerRep.Unmarshal(m, b)
}
func (m *StartServerRep) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartServerRep.Marshal(b, m, deterministic)
}
func (m *StartServerRep) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartServerRep.Merge(m, src)
}
func (m *StartServerRep) XXX_Size() int {
	return xxx_messageInfo_StartServerRep.Size(m)
}
func (m *StartServerRep) XXX_DiscardUnknown() {
	xxx_messageInfo_StartServerRep.DiscardUnknown(m)
}

var xxx_messageInfo_StartServerRep proto.InternalMessageInfo

func (m *StartServerRep) GetIsLeader() bool {
	if m != nil {
		return m.IsLeader
	}
	return false
}

type StartClientReq struct {
	Frq                  int32    `protobuf:"varint,1,opt,name=frq,proto3" json:"frq,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartClientReq) Reset()         { *m = StartClientReq{} }
func (m *StartClientReq) String() string { return proto.CompactTextString(m) }
func (*StartClientReq) ProtoMessage()    {}
func (*StartClientReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{3}
}

func (m *StartClientReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartClientReq.Unmarshal(m, b)
}
func (m *StartClientReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartClientReq.Marshal(b, m, deterministic)
}
func (m *StartClientReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartClientReq.Merge(m, src)
}
func (m *StartClientReq) XXX_Size() int {
	return xxx_messageInfo_StartClientReq.Size(m)
}
func (m *StartClientReq) XXX_DiscardUnknown() {
	xxx_messageInfo_StartClientReq.DiscardUnknown(m)
}

var xxx_messageInfo_StartClientReq proto.InternalMessageInfo

func (m *StartClientReq) GetFrq() int32 {
	if m != nil {
		return m.Frq
	}
	return 0
}

type StartClientRep struct {
	Latency              string   `protobuf:"bytes,1,opt,name=latency,proto3" json:"latency,omitempty"`
	Throughout           string   `protobuf:"bytes,2,opt,name=throughout,proto3" json:"throughout,omitempty"`
	SendRate             string   `protobuf:"bytes,3,opt,name=sendRate,proto3" json:"sendRate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartClientRep) Reset()         { *m = StartClientRep{} }
func (m *StartClientRep) String() string { return proto.CompactTextString(m) }
func (*StartClientRep) ProtoMessage()    {}
func (*StartClientRep) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{4}
}

func (m *StartClientRep) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartClientRep.Unmarshal(m, b)
}
func (m *StartClientRep) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartClientRep.Marshal(b, m, deterministic)
}
func (m *StartClientRep) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartClientRep.Merge(m, src)
}
func (m *StartClientRep) XXX_Size() int {
	return xxx_messageInfo_StartClientRep.Size(m)
}
func (m *StartClientRep) XXX_DiscardUnknown() {
	xxx_messageInfo_StartClientRep.DiscardUnknown(m)
}

var xxx_messageInfo_StartClientRep proto.InternalMessageInfo

func (m *StartClientRep) GetLatency() string {
	if m != nil {
		return m.Latency
	}
	return ""
}

func (m *StartClientRep) GetThroughout() string {
	if m != nil {
		return m.Throughout
	}
	return ""
}

func (m *StartClientRep) GetSendRate() string {
	if m != nil {
		return m.SendRate
	}
	return ""
}

type SwitchReq struct {
	Mode                 int32    `protobuf:"varint,1,opt,name=mode,proto3" json:"mode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SwitchReq) Reset()         { *m = SwitchReq{} }
func (m *SwitchReq) String() string { return proto.CompactTextString(m) }
func (*SwitchReq) ProtoMessage()    {}
func (*SwitchReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{5}
}

func (m *SwitchReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SwitchReq.Unmarshal(m, b)
}
func (m *SwitchReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SwitchReq.Marshal(b, m, deterministic)
}
func (m *SwitchReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SwitchReq.Merge(m, src)
}
func (m *SwitchReq) XXX_Size() int {
	return xxx_messageInfo_SwitchReq.Size(m)
}
func (m *SwitchReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SwitchReq.DiscardUnknown(m)
}

var xxx_messageInfo_SwitchReq proto.InternalMessageInfo

func (m *SwitchReq) GetMode() int32 {
	if m != nil {
		return m.Mode
	}
	return 0
}

type SwitchRep struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SwitchRep) Reset()         { *m = SwitchRep{} }
func (m *SwitchRep) String() string { return proto.CompactTextString(m) }
func (*SwitchRep) ProtoMessage()    {}
func (*SwitchRep) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{6}
}

func (m *SwitchRep) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SwitchRep.Unmarshal(m, b)
}
func (m *SwitchRep) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SwitchRep.Marshal(b, m, deterministic)
}
func (m *SwitchRep) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SwitchRep.Merge(m, src)
}
func (m *SwitchRep) XXX_Size() int {
	return xxx_messageInfo_SwitchRep.Size(m)
}
func (m *SwitchRep) XXX_DiscardUnknown() {
	xxx_messageInfo_SwitchRep.DiscardUnknown(m)
}

var xxx_messageInfo_SwitchRep proto.InternalMessageInfo

type StopMsg struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopMsg) Reset()         { *m = StopMsg{} }
func (m *StopMsg) String() string { return proto.CompactTextString(m) }
func (*StopMsg) ProtoMessage()    {}
func (*StopMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{7}
}

func (m *StopMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopMsg.Unmarshal(m, b)
}
func (m *StopMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopMsg.Marshal(b, m, deterministic)
}
func (m *StopMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopMsg.Merge(m, src)
}
func (m *StopMsg) XXX_Size() int {
	return xxx_messageInfo_StopMsg.Size(m)
}
func (m *StopMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_StopMsg.DiscardUnknown(m)
}

var xxx_messageInfo_StopMsg proto.InternalMessageInfo

type OrderRep struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderRep) Reset()         { *m = OrderRep{} }
func (m *OrderRep) String() string { return proto.CompactTextString(m) }
func (*OrderRep) ProtoMessage()    {}
func (*OrderRep) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{8}
}

func (m *OrderRep) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderRep.Unmarshal(m, b)
}
func (m *OrderRep) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderRep.Marshal(b, m, deterministic)
}
func (m *OrderRep) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderRep.Merge(m, src)
}
func (m *OrderRep) XXX_Size() int {
	return xxx_messageInfo_OrderRep.Size(m)
}
func (m *OrderRep) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderRep.DiscardUnknown(m)
}

var xxx_messageInfo_OrderRep proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SwapCertMsg)(nil), "BCDns_daemon.SwapCertMsg")
	proto.RegisterType((*StartServerReq)(nil), "BCDns_daemon.StartServerReq")
	proto.RegisterType((*StartServerRep)(nil), "BCDns_daemon.StartServerRep")
	proto.RegisterType((*StartClientReq)(nil), "BCDns_daemon.StartClientReq")
	proto.RegisterType((*StartClientRep)(nil), "BCDns_daemon.StartClientRep")
	proto.RegisterType((*SwitchReq)(nil), "BCDns_daemon.SwitchReq")
	proto.RegisterType((*SwitchRep)(nil), "BCDns_daemon.SwitchRep")
	proto.RegisterType((*StopMsg)(nil), "BCDns_daemon.StopMsg")
	proto.RegisterType((*OrderRep)(nil), "BCDns_daemon.OrderRep")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd) }

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 358 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcd, 0x6e, 0xe2, 0x30,
	0x14, 0x85, 0x21, 0xcc, 0x00, 0xb9, 0xfc, 0x68, 0x64, 0x69, 0x66, 0xd2, 0x88, 0xfe, 0xc8, 0xab,
	0x2e, 0xaa, 0x2c, 0xda, 0x55, 0x97, 0x05, 0x96, 0x8d, 0x2a, 0x25, 0x0f, 0x50, 0x19, 0x72, 0x49,
	0x22, 0x41, 0x6c, 0x6c, 0x53, 0x44, 0x9f, 0xb8, 0x8f, 0x51, 0xd9, 0x0d, 0x10, 0xa2, 0x96, 0x9d,
	0xaf, 0xcf, 0xd1, 0xd1, 0xf1, 0x77, 0x0d, 0x83, 0x15, 0x2a, 0xc5, 0x52, 0x0c, 0x84, 0xe4, 0x9a,
	0x93, 0xfe, 0x78, 0x32, 0x2d, 0xd4, 0x6b, 0xc2, 0x70, 0xc5, 0x0b, 0x7a, 0x09, 0xbd, 0x78, 0xcb,
	0xc4, 0x04, 0xa5, 0x0e, 0x55, 0x4a, 0x86, 0xe0, 0xe4, 0xc2, 0x6b, 0xde, 0x34, 0x6f, 0xdd, 0xc8,
	0xc9, 0x05, 0x0d, 0x60, 0x18, 0x6b, 0x26, 0x75, 0x8c, 0xf2, 0x0d, 0x65, 0x84, 0x6b, 0x32, 0x02,
	0x77, 0xb6, 0x7b, 0x67, 0x85, 0xce, 0x0b, 0xb4, 0xc6, 0x6e, 0x74, 0xbc, 0xa0, 0x77, 0x35, 0xbf,
	0x20, 0x3e, 0x74, 0x73, 0xf5, 0x8c, 0x2c, 0x41, 0x59, 0xda, 0x0f, 0x33, 0xa5, 0xa5, 0x7b, 0xb2,
	0xcc, 0xb1, 0xd0, 0x26, 0xfd, 0x0f, 0xb4, 0x16, 0x72, 0x6d, 0x8d, 0xbf, 0x23, 0x73, 0xa4, 0x8b,
	0x9a, 0x47, 0x10, 0x0f, 0x3a, 0x4b, 0xa6, 0xb1, 0x98, 0xef, 0xca, 0xa2, 0xfb, 0x91, 0x5c, 0x01,
	0xe8, 0x4c, 0xf2, 0x4d, 0x9a, 0xf1, 0x8d, 0xf6, 0x1c, 0x2b, 0x56, 0x6e, 0x4c, 0x17, 0x85, 0x45,
	0x12, 0x31, 0x8d, 0x5e, 0xcb, 0xaa, 0x87, 0x99, 0x5e, 0x83, 0x1b, 0x6f, 0x73, 0x3d, 0xcf, 0x4c,
	0x0d, 0x02, 0xbf, 0x56, 0x3c, 0xc1, 0xb2, 0x87, 0x3d, 0xd3, 0xde, 0xd1, 0x20, 0xa8, 0x0b, 0x9d,
	0x58, 0x73, 0x11, 0xaa, 0x94, 0x02, 0x74, 0x5f, 0x64, 0x62, 0x1f, 0x7b, 0xff, 0xe1, 0x40, 0x3b,
	0x44, 0x9d, 0xf1, 0x84, 0x3c, 0x01, 0x4c, 0xf9, 0x1e, 0x2d, 0xb9, 0x08, 0xaa, 0xd4, 0x83, 0x0a,
	0x72, 0xff, 0xdf, 0xa9, 0xb4, 0xcf, 0xa2, 0x0d, 0x12, 0xc2, 0x60, 0xca, 0x2b, 0x38, 0xc9, 0xa8,
	0x96, 0x72, 0xb2, 0x19, 0xff, 0x9c, 0x7a, 0x1a, 0xf7, 0xc5, 0xf2, 0xdb, 0xb8, 0xc3, 0x2a, 0xfc,
	0x73, 0xaa, 0x89, 0x7b, 0x84, 0xb6, 0x89, 0xe3, 0x82, 0xfc, 0xad, 0x3b, 0x2d, 0x98, 0x33, 0x0f,
	0x1b, 0x43, 0xdf, 0xb0, 0x31, 0x30, 0x43, 0x9e, 0x20, 0xf9, 0x5f, 0xa7, 0x53, 0xee, 0xc1, 0xff,
	0x49, 0xa0, 0x8d, 0x59, 0xdb, 0xfe, 0xe6, 0x87, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x63, 0xb1,
	0x6a, 0x91, 0xde, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MethodClient is the client API for Method service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MethodClient interface {
	DoSwapCert(ctx context.Context, in *SwapCertMsg, opts ...grpc.CallOption) (*OrderRep, error)
	DoStartServer(ctx context.Context, in *StartServerReq, opts ...grpc.CallOption) (*StartServerRep, error)
	DoStartClient(ctx context.Context, in *StartClientReq, opts ...grpc.CallOption) (*StartClientRep, error)
	DoStop(ctx context.Context, in *StopMsg, opts ...grpc.CallOption) (*OrderRep, error)
	DoSwitchMode(ctx context.Context, in *SwitchReq, opts ...grpc.CallOption) (*SwitchReq, error)
}

type methodClient struct {
	cc *grpc.ClientConn
}

func NewMethodClient(cc *grpc.ClientConn) MethodClient {
	return &methodClient{cc}
}

func (c *methodClient) DoSwapCert(ctx context.Context, in *SwapCertMsg, opts ...grpc.CallOption) (*OrderRep, error) {
	out := new(OrderRep)
	err := c.cc.Invoke(ctx, "/BCDns_daemon.Method/DoSwapCert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *methodClient) DoStartServer(ctx context.Context, in *StartServerReq, opts ...grpc.CallOption) (*StartServerRep, error) {
	out := new(StartServerRep)
	err := c.cc.Invoke(ctx, "/BCDns_daemon.Method/DoStartServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *methodClient) DoStartClient(ctx context.Context, in *StartClientReq, opts ...grpc.CallOption) (*StartClientRep, error) {
	out := new(StartClientRep)
	err := c.cc.Invoke(ctx, "/BCDns_daemon.Method/DoStartClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *methodClient) DoStop(ctx context.Context, in *StopMsg, opts ...grpc.CallOption) (*OrderRep, error) {
	out := new(OrderRep)
	err := c.cc.Invoke(ctx, "/BCDns_daemon.Method/DoStop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *methodClient) DoSwitchMode(ctx context.Context, in *SwitchReq, opts ...grpc.CallOption) (*SwitchReq, error) {
	out := new(SwitchReq)
	err := c.cc.Invoke(ctx, "/BCDns_daemon.Method/DoSwitchMode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MethodServer is the server API for Method service.
type MethodServer interface {
	DoSwapCert(context.Context, *SwapCertMsg) (*OrderRep, error)
	DoStartServer(context.Context, *StartServerReq) (*StartServerRep, error)
	DoStartClient(context.Context, *StartClientReq) (*StartClientRep, error)
	DoStop(context.Context, *StopMsg) (*OrderRep, error)
	DoSwitchMode(context.Context, *SwitchReq) (*SwitchReq, error)
}

// UnimplementedMethodServer can be embedded to have forward compatible implementations.
type UnimplementedMethodServer struct {
}

func (*UnimplementedMethodServer) DoSwapCert(ctx context.Context, req *SwapCertMsg) (*OrderRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoSwapCert not implemented")
}
func (*UnimplementedMethodServer) DoStartServer(ctx context.Context, req *StartServerReq) (*StartServerRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoStartServer not implemented")
}
func (*UnimplementedMethodServer) DoStartClient(ctx context.Context, req *StartClientReq) (*StartClientRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoStartClient not implemented")
}
func (*UnimplementedMethodServer) DoStop(ctx context.Context, req *StopMsg) (*OrderRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoStop not implemented")
}
func (*UnimplementedMethodServer) DoSwitchMode(ctx context.Context, req *SwitchReq) (*SwitchReq, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoSwitchMode not implemented")
}

func RegisterMethodServer(s *grpc.Server, srv MethodServer) {
	s.RegisterService(&_Method_serviceDesc, srv)
}

func _Method_DoSwapCert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SwapCertMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MethodServer).DoSwapCert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BCDns_daemon.Method/DoSwapCert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MethodServer).DoSwapCert(ctx, req.(*SwapCertMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Method_DoStartServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartServerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MethodServer).DoStartServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BCDns_daemon.Method/DoStartServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MethodServer).DoStartServer(ctx, req.(*StartServerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Method_DoStartClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartClientReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MethodServer).DoStartClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BCDns_daemon.Method/DoStartClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MethodServer).DoStartClient(ctx, req.(*StartClientReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Method_DoStop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MethodServer).DoStop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BCDns_daemon.Method/DoStop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MethodServer).DoStop(ctx, req.(*StopMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Method_DoSwitchMode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SwitchReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MethodServer).DoSwitchMode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BCDns_daemon.Method/DoSwitchMode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MethodServer).DoSwitchMode(ctx, req.(*SwitchReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Method_serviceDesc = grpc.ServiceDesc{
	ServiceName: "BCDns_daemon.Method",
	HandlerType: (*MethodServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DoSwapCert",
			Handler:    _Method_DoSwapCert_Handler,
		},
		{
			MethodName: "DoStartServer",
			Handler:    _Method_DoStartServer_Handler,
		},
		{
			MethodName: "DoStartClient",
			Handler:    _Method_DoStartClient_Handler,
		},
		{
			MethodName: "DoStop",
			Handler:    _Method_DoStop_Handler,
		},
		{
			MethodName: "DoSwitchMode",
			Handler:    _Method_DoSwitchMode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message.proto",
}
