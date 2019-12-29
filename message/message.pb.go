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

type StopMsg struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopMsg) Reset()         { *m = StopMsg{} }
func (m *StopMsg) String() string { return proto.CompactTextString(m) }
func (*StopMsg) ProtoMessage()    {}
func (*StopMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{5}
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
	return fileDescriptor_33c57e4bae7b9afd, []int{6}
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
	proto.RegisterType((*StopMsg)(nil), "BCDns_daemon.StopMsg")
	proto.RegisterType((*OrderRep)(nil), "BCDns_daemon.OrderRep")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd) }

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x6d, 0xc4, 0xb6, 0x19, 0x6d, 0x91, 0x05, 0x25, 0x86, 0x2a, 0xb2, 0x27, 0x0f, 0x92,
	0x83, 0x9e, 0x3c, 0x6a, 0x72, 0x12, 0x83, 0x90, 0x3c, 0x80, 0x6c, 0xcd, 0x98, 0x04, 0xda, 0xdd,
	0xed, 0xee, 0x56, 0xa9, 0xcf, 0xe3, 0x83, 0x4a, 0xd6, 0xa4, 0x26, 0x41, 0x7b, 0xcb, 0xcc, 0xff,
	0xf3, 0x31, 0xf9, 0x58, 0x98, 0x2c, 0x51, 0x6b, 0x96, 0x63, 0x20, 0x95, 0x30, 0x82, 0x1c, 0x3d,
	0x84, 0x11, 0xd7, 0x2f, 0x19, 0xc3, 0xa5, 0xe0, 0xf4, 0x1c, 0x0e, 0xd3, 0x0f, 0x26, 0x43, 0x54,
	0x26, 0xd6, 0x39, 0x99, 0x82, 0x53, 0x4a, 0x6f, 0x70, 0x39, 0xb8, 0x72, 0x13, 0xa7, 0x94, 0x34,
	0x80, 0x69, 0x6a, 0x98, 0x32, 0x29, 0xaa, 0x77, 0x54, 0x09, 0xae, 0xc8, 0x0c, 0xdc, 0xf9, 0xe6,
	0x93, 0x71, 0x53, 0x72, 0xb4, 0xc5, 0x71, 0xf2, 0xbb, 0xa0, 0xd7, 0xbd, 0xbe, 0x24, 0x3e, 0x8c,
	0x4b, 0xfd, 0x84, 0x2c, 0x43, 0x55, 0xd7, 0xb7, 0x33, 0xa5, 0x75, 0x3b, 0x5c, 0x94, 0xc8, 0x4d,
	0x45, 0x3f, 0x86, 0xfd, 0x37, 0xb5, 0xb2, 0xc5, 0x83, 0xa4, 0xfa, 0xa4, 0x8f, 0xbd, 0x8e, 0x24,
	0x1e, 0x8c, 0x16, 0xcc, 0x20, 0x7f, 0xdd, 0xd4, 0x87, 0x36, 0x23, 0xb9, 0x00, 0x30, 0x85, 0x12,
	0xeb, 0xbc, 0x10, 0x6b, 0xe3, 0x39, 0x36, 0x6c, 0x6d, 0xa8, 0x0b, 0xa3, 0xd4, 0x08, 0x19, 0xeb,
	0x9c, 0x02, 0x8c, 0x9f, 0x55, 0x66, 0x4f, 0xbc, 0xf9, 0x72, 0x60, 0x18, 0xa3, 0x29, 0x44, 0x46,
	0xee, 0x01, 0x22, 0xd1, 0x08, 0x21, 0x67, 0x41, 0xdb, 0x55, 0xd0, 0x12, 0xe5, 0x9f, 0x76, 0xa3,
	0x86, 0x45, 0xf7, 0x48, 0x0c, 0x93, 0x48, 0xb4, 0x24, 0x90, 0x59, 0x8f, 0xd2, 0xf1, 0xe9, 0xef,
	0x4a, 0xbb, 0xb8, 0x1f, 0x03, 0x7f, 0xe2, 0xb6, 0x02, 0xfd, 0x5d, 0x69, 0x85, 0xbb, 0x83, 0x61,
	0x85, 0x13, 0x92, 0x9c, 0xf4, 0x9b, 0x56, 0xcc, 0xff, 0x3f, 0x36, 0x1f, 0xda, 0xf7, 0x73, 0xfb,
	0x1d, 0x00, 0x00, 0xff, 0xff, 0x6d, 0xaa, 0x41, 0xe9, 0x50, 0x02, 0x00, 0x00,
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

// MethodServer is the server API for Method service.
type MethodServer interface {
	DoSwapCert(context.Context, *SwapCertMsg) (*OrderRep, error)
	DoStartServer(context.Context, *StartServerReq) (*StartServerRep, error)
	DoStartClient(context.Context, *StartClientReq) (*StartClientRep, error)
	DoStop(context.Context, *StopMsg) (*OrderRep, error)
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
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message.proto",
}
