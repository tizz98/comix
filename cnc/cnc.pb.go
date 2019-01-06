// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cnc.proto

package cnc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Response struct {
	HasUpdate            bool     `protobuf:"varint,1,opt,name=has_update,json=hasUpdate,proto3" json:"has_update,omitempty"`
	Url                  string   `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Checksum             string   `protobuf:"bytes,3,opt,name=checksum,proto3" json:"checksum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_cnc_71345804cefc447c, []int{0}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetHasUpdate() bool {
	if m != nil {
		return m.HasUpdate
	}
	return false
}

func (m *Response) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Response) GetChecksum() string {
	if m != nil {
		return m.Checksum
	}
	return ""
}

type UpdateMsg struct {
	UpdateComplete       bool     `protobuf:"varint,1,opt,name=updateComplete,proto3" json:"updateComplete,omitempty"`
	UpdateMessage        string   `protobuf:"bytes,2,opt,name=updateMessage,proto3" json:"updateMessage,omitempty"`
	ClientId             string   `protobuf:"bytes,3,opt,name=clientId,proto3" json:"clientId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateMsg) Reset()         { *m = UpdateMsg{} }
func (m *UpdateMsg) String() string { return proto.CompactTextString(m) }
func (*UpdateMsg) ProtoMessage()    {}
func (*UpdateMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_cnc_71345804cefc447c, []int{1}
}
func (m *UpdateMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateMsg.Unmarshal(m, b)
}
func (m *UpdateMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateMsg.Marshal(b, m, deterministic)
}
func (dst *UpdateMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateMsg.Merge(dst, src)
}
func (m *UpdateMsg) XXX_Size() int {
	return xxx_messageInfo_UpdateMsg.Size(m)
}
func (m *UpdateMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateMsg.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateMsg proto.InternalMessageInfo

func (m *UpdateMsg) GetUpdateComplete() bool {
	if m != nil {
		return m.UpdateComplete
	}
	return false
}

func (m *UpdateMsg) GetUpdateMessage() string {
	if m != nil {
		return m.UpdateMessage
	}
	return ""
}

func (m *UpdateMsg) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

type Status struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_cnc_71345804cefc447c, []int{2}
}
func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (dst *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(dst, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

type PingMsg struct {
	Ok                   bool     `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	StatusMessage        string   `protobuf:"bytes,2,opt,name=statusMessage,proto3" json:"statusMessage,omitempty"`
	ClientId             string   `protobuf:"bytes,3,opt,name=clientId,proto3" json:"clientId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingMsg) Reset()         { *m = PingMsg{} }
func (m *PingMsg) String() string { return proto.CompactTextString(m) }
func (*PingMsg) ProtoMessage()    {}
func (*PingMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_cnc_71345804cefc447c, []int{3}
}
func (m *PingMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingMsg.Unmarshal(m, b)
}
func (m *PingMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingMsg.Marshal(b, m, deterministic)
}
func (dst *PingMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingMsg.Merge(dst, src)
}
func (m *PingMsg) XXX_Size() int {
	return xxx_messageInfo_PingMsg.Size(m)
}
func (m *PingMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_PingMsg.DiscardUnknown(m)
}

var xxx_messageInfo_PingMsg proto.InternalMessageInfo

func (m *PingMsg) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *PingMsg) GetStatusMessage() string {
	if m != nil {
		return m.StatusMessage
	}
	return ""
}

func (m *PingMsg) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func init() {
	proto.RegisterType((*Response)(nil), "cnc.Response")
	proto.RegisterType((*UpdateMsg)(nil), "cnc.UpdateMsg")
	proto.RegisterType((*Status)(nil), "cnc.Status")
	proto.RegisterType((*PingMsg)(nil), "cnc.PingMsg")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CnCClient is the client API for CnC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CnCClient interface {
	Ping(ctx context.Context, in *PingMsg, opts ...grpc.CallOption) (*Response, error)
	UpdateStatus(ctx context.Context, in *UpdateMsg, opts ...grpc.CallOption) (*Status, error)
}

type cnCClient struct {
	cc *grpc.ClientConn
}

func NewCnCClient(cc *grpc.ClientConn) CnCClient {
	return &cnCClient{cc}
}

func (c *cnCClient) Ping(ctx context.Context, in *PingMsg, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/cnc.CnC/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cnCClient) UpdateStatus(ctx context.Context, in *UpdateMsg, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/cnc.CnC/UpdateStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CnCServer is the server API for CnC service.
type CnCServer interface {
	Ping(context.Context, *PingMsg) (*Response, error)
	UpdateStatus(context.Context, *UpdateMsg) (*Status, error)
}

func RegisterCnCServer(s *grpc.Server, srv CnCServer) {
	s.RegisterService(&_CnC_serviceDesc, srv)
}

func _CnC_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CnCServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cnc.CnC/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CnCServer).Ping(ctx, req.(*PingMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _CnC_UpdateStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CnCServer).UpdateStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cnc.CnC/UpdateStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CnCServer).UpdateStatus(ctx, req.(*UpdateMsg))
	}
	return interceptor(ctx, in, info, handler)
}

var _CnC_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cnc.CnC",
	HandlerType: (*CnCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _CnC_Ping_Handler,
		},
		{
			MethodName: "UpdateStatus",
			Handler:    _CnC_UpdateStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cnc.proto",
}

func init() { proto.RegisterFile("cnc.proto", fileDescriptor_cnc_71345804cefc447c) }

var fileDescriptor_cnc_71345804cefc447c = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x86, 0x9b, 0x44, 0x6a, 0x32, 0xb6, 0x41, 0xe6, 0x14, 0x02, 0x42, 0x09, 0x2a, 0xbd, 0xd8,
	0x83, 0xfe, 0x84, 0x9c, 0x3c, 0x14, 0x24, 0x22, 0x1e, 0x7a, 0x90, 0x75, 0xbb, 0x24, 0x25, 0xe9,
	0x6e, 0xe8, 0xec, 0xfe, 0x7f, 0xd9, 0x0f, 0x03, 0xed, 0xa1, 0xb7, 0xbc, 0x0f, 0x93, 0x79, 0x86,
	0x77, 0x21, 0xe3, 0x92, 0x6f, 0xc6, 0x93, 0xd2, 0x0a, 0x13, 0x2e, 0x79, 0xf5, 0x0d, 0x69, 0x23,
	0x68, 0x54, 0x92, 0x04, 0x3e, 0x00, 0x74, 0x8c, 0x7e, 0xcc, 0xb8, 0x67, 0x5a, 0x14, 0xd1, 0x2a,
	0x5a, 0xa7, 0x4d, 0xd6, 0x31, 0xfa, 0x72, 0x00, 0xef, 0x21, 0x31, 0xa7, 0xa1, 0x88, 0x57, 0xd1,
	0x3a, 0x6b, 0xec, 0x27, 0x96, 0x90, 0xf2, 0x4e, 0xf0, 0x9e, 0xcc, 0xb1, 0x48, 0x1c, 0x9e, 0x72,
	0x65, 0x20, 0xf3, 0xff, 0x6d, 0xa9, 0xc5, 0x67, 0xc8, 0xfd, 0xd6, 0x5a, 0x1d, 0xc7, 0x41, 0x4c,
	0xdb, 0x2f, 0x28, 0x3e, 0xc2, 0xd2, 0x93, 0xad, 0x20, 0x62, 0xad, 0x08, 0xb2, 0x73, 0xe8, 0xb4,
	0xc3, 0x41, 0x48, 0xfd, 0xbe, 0x9f, 0xb4, 0x21, 0x57, 0x29, 0xcc, 0x3f, 0x35, 0xd3, 0x86, 0xaa,
	0x1d, 0xdc, 0x7e, 0x1c, 0x64, 0x6b, 0xf5, 0x39, 0xc4, 0xaa, 0x0f, 0xca, 0x58, 0xf5, 0x56, 0x43,
	0x6e, 0xe8, 0x42, 0x73, 0x06, 0xaf, 0x69, 0x5e, 0x77, 0x90, 0xd4, 0xb2, 0xc6, 0x27, 0xb8, 0xb1,
	0x0e, 0x5c, 0x6c, 0x6c, 0xad, 0x41, 0x57, 0x2e, 0x5d, 0xfa, 0xaf, 0xb5, 0x9a, 0xe1, 0x0b, 0x2c,
	0x7c, 0x17, 0xfe, 0x34, 0xcc, 0xdd, 0xc0, 0x54, 0x4f, 0x79, 0xe7, 0x72, 0xb8, 0x7b, 0xf6, 0x3b,
	0x77, 0xef, 0xf3, 0xf6, 0x17, 0x00, 0x00, 0xff, 0xff, 0x14, 0x77, 0xf3, 0x14, 0xac, 0x01, 0x00,
	0x00,
}
