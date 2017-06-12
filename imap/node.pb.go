// Code generated by protoc-gen-go. DO NOT EDIT.
// source: node.proto

/*
Package imap is a generated protocol buffer package.

It is generated from these files:
	node.proto

It has these top-level messages:
	Context
	Confirmation
	Command
	Reply
*/
package imap

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

type Context struct {
	ClientID string `protobuf:"bytes,1,opt,name=clientID" json:"clientID,omitempty"`
	UserName string `protobuf:"bytes,2,opt,name=userName" json:"userName,omitempty"`
}

func (m *Context) Reset()                    { *m = Context{} }
func (m *Context) String() string            { return proto.CompactTextString(m) }
func (*Context) ProtoMessage()               {}
func (*Context) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Context) GetClientID() string {
	if m != nil {
		return m.ClientID
	}
	return ""
}

func (m *Context) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

type Confirmation struct {
	Status uint32 `protobuf:"varint,1,opt,name=status" json:"status,omitempty"`
}

func (m *Confirmation) Reset()                    { *m = Confirmation{} }
func (m *Confirmation) String() string            { return proto.CompactTextString(m) }
func (*Confirmation) ProtoMessage()               {}
func (*Confirmation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Confirmation) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

type Command struct {
	Command string `protobuf:"bytes,1,opt,name=command" json:"command,omitempty"`
}

func (m *Command) Reset()                    { *m = Command{} }
func (m *Command) String() string            { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()               {}
func (*Command) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Command) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

type Reply struct {
	Reply string `protobuf:"bytes,1,opt,name=reply" json:"reply,omitempty"`
}

func (m *Reply) Reset()                    { *m = Reply{} }
func (m *Reply) String() string            { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()               {}
func (*Reply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Reply) GetReply() string {
	if m != nil {
		return m.Reply
	}
	return ""
}

func init() {
	proto.RegisterType((*Context)(nil), "imap.Context")
	proto.RegisterType((*Confirmation)(nil), "imap.Confirmation")
	proto.RegisterType((*Command)(nil), "imap.Command")
	proto.RegisterType((*Reply)(nil), "imap.Reply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Node service

type NodeClient interface {
	SendContext(ctx context.Context, in *Context, opts ...grpc.CallOption) (*Confirmation, error)
	Select(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Reply, error)
	Create(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Reply, error)
	Delete(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Reply, error)
	List(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Reply, error)
	Append(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Reply, error)
	Expunge(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Reply, error)
	Store(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Reply, error)
}

type nodeClient struct {
	cc *grpc.ClientConn
}

func NewNodeClient(cc *grpc.ClientConn) NodeClient {
	return &nodeClient{cc}
}

func (c *nodeClient) SendContext(ctx context.Context, in *Context, opts ...grpc.CallOption) (*Confirmation, error) {
	out := new(Confirmation)
	err := grpc.Invoke(ctx, "/imap.Node/SendContext", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) Select(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := grpc.Invoke(ctx, "/imap.Node/Select", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) Create(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := grpc.Invoke(ctx, "/imap.Node/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) Delete(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := grpc.Invoke(ctx, "/imap.Node/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) List(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := grpc.Invoke(ctx, "/imap.Node/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) Append(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := grpc.Invoke(ctx, "/imap.Node/Append", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) Expunge(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := grpc.Invoke(ctx, "/imap.Node/Expunge", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) Store(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := grpc.Invoke(ctx, "/imap.Node/Store", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Node service

type NodeServer interface {
	SendContext(context.Context, *Context) (*Confirmation, error)
	Select(context.Context, *Command) (*Reply, error)
	Create(context.Context, *Command) (*Reply, error)
	Delete(context.Context, *Command) (*Reply, error)
	List(context.Context, *Command) (*Reply, error)
	Append(context.Context, *Command) (*Reply, error)
	Expunge(context.Context, *Command) (*Reply, error)
	Store(context.Context, *Command) (*Reply, error)
}

func RegisterNodeServer(s *grpc.Server, srv NodeServer) {
	s.RegisterService(&_Node_serviceDesc, srv)
}

func _Node_SendContext_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Context)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).SendContext(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/imap.Node/SendContext",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).SendContext(ctx, req.(*Context))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_Select_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Command)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Select(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/imap.Node/Select",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Select(ctx, req.(*Command))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Command)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/imap.Node/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Create(ctx, req.(*Command))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Command)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/imap.Node/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Delete(ctx, req.(*Command))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Command)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/imap.Node/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).List(ctx, req.(*Command))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_Append_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Command)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Append(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/imap.Node/Append",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Append(ctx, req.(*Command))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_Expunge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Command)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Expunge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/imap.Node/Expunge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Expunge(ctx, req.(*Command))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_Store_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Command)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Store(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/imap.Node/Store",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Store(ctx, req.(*Command))
	}
	return interceptor(ctx, in, info, handler)
}

var _Node_serviceDesc = grpc.ServiceDesc{
	ServiceName: "imap.Node",
	HandlerType: (*NodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendContext",
			Handler:    _Node_SendContext_Handler,
		},
		{
			MethodName: "Select",
			Handler:    _Node_Select_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Node_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Node_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Node_List_Handler,
		},
		{
			MethodName: "Append",
			Handler:    _Node_Append_Handler,
		},
		{
			MethodName: "Expunge",
			Handler:    _Node_Expunge_Handler,
		},
		{
			MethodName: "Store",
			Handler:    _Node_Store_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "node.proto",
}

func init() { proto.RegisterFile("node.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcd, 0x4a, 0xf4, 0x30,
	0x14, 0x40, 0xbf, 0x19, 0xfa, 0xf3, 0x79, 0xc7, 0xd9, 0x04, 0x91, 0x52, 0x10, 0x24, 0xea, 0xe8,
	0xaa, 0x8b, 0xf1, 0x09, 0x86, 0x8e, 0x0b, 0x41, 0x66, 0xd1, 0x3e, 0x41, 0x6c, 0xae, 0x12, 0x68,
	0x7e, 0x48, 0x53, 0x18, 0xdf, 0xd2, 0x47, 0x92, 0x36, 0x6d, 0x71, 0x17, 0x77, 0xf7, 0x70, 0x0f,
	0x3d, 0xb7, 0x10, 0x00, 0xa5, 0x39, 0x16, 0xc6, 0x6a, 0xa7, 0x49, 0x24, 0x24, 0x33, 0xf4, 0x00,
	0x69, 0xa9, 0x95, 0xc3, 0xb3, 0x23, 0x39, 0xfc, 0x6f, 0x5a, 0x81, 0xca, 0xbd, 0x1e, 0xb3, 0xd5,
	0xed, 0xea, 0xe9, 0xa2, 0x5a, 0x78, 0xd8, 0xf5, 0x1d, 0xda, 0x13, 0x93, 0x98, 0xad, 0xfd, 0x6e,
	0x66, 0xba, 0x83, 0xcb, 0x52, 0xab, 0x0f, 0x61, 0x25, 0x73, 0x42, 0x2b, 0x72, 0x0d, 0x49, 0xe7,
	0x98, 0xeb, 0xbb, 0xf1, 0x2b, 0xdb, 0x6a, 0x22, 0x7a, 0x37, 0xa4, 0xa4, 0x64, 0x8a, 0x93, 0x0c,
	0xd2, 0xc6, 0x8f, 0x53, 0x69, 0x46, 0x7a, 0x03, 0x71, 0x85, 0xa6, 0xfd, 0x22, 0x57, 0x10, 0xdb,
	0x61, 0x98, 0x04, 0x0f, 0xfb, 0xef, 0x35, 0x44, 0x27, 0xcd, 0x91, 0xec, 0x61, 0x53, 0xa3, 0xe2,
	0xf3, 0xed, 0xdb, 0x62, 0xf8, 0x9b, 0x62, 0xc2, 0x9c, 0x2c, 0xb8, 0x9c, 0x45, 0xff, 0x91, 0x1d,
	0x24, 0x35, 0xb6, 0xd8, 0xfc, 0xd2, 0xc7, 0x68, 0xbe, 0xf1, 0x38, 0x86, 0xbd, 0x57, 0x5a, 0x64,
	0x0e, 0xc3, 0xde, 0x11, 0x5b, 0x0c, 0x7a, 0xf7, 0x10, 0xbd, 0x89, 0xee, 0x0f, 0xd5, 0x83, 0x31,
	0xa8, 0x78, 0xc0, 0x7b, 0x84, 0xf4, 0xe5, 0x6c, 0x7a, 0xf5, 0x19, 0xca, 0x3e, 0x40, 0x5c, 0x3b,
	0x6d, 0x03, 0xda, 0x7b, 0x32, 0x3e, 0x87, 0xe7, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x06, 0x00,
	0xb0, 0xf7, 0x1c, 0x02, 0x00, 0x00,
}
