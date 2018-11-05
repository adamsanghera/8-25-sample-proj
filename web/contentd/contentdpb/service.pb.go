// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package contentdpb

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

type NewTweetRequest struct {
	Content   string `protobuf:"bytes,1,opt,name=Content" json:"Content,omitempty"`
	PosterUID string `protobuf:"bytes,2,opt,name=PosterUID" json:"PosterUID,omitempty"`
}

func (m *NewTweetRequest) Reset()                    { *m = NewTweetRequest{} }
func (m *NewTweetRequest) String() string            { return proto.CompactTextString(m) }
func (*NewTweetRequest) ProtoMessage()               {}
func (*NewTweetRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *NewTweetRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *NewTweetRequest) GetPosterUID() string {
	if m != nil {
		return m.PosterUID
	}
	return ""
}

type NewTweetResponse struct {
}

func (m *NewTweetResponse) Reset()                    { *m = NewTweetResponse{} }
func (m *NewTweetResponse) String() string            { return proto.CompactTextString(m) }
func (*NewTweetResponse) ProtoMessage()               {}
func (*NewTweetResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

type GetTweetRequest struct {
	TID string `protobuf:"bytes,1,opt,name=TID" json:"TID,omitempty"`
}

func (m *GetTweetRequest) Reset()                    { *m = GetTweetRequest{} }
func (m *GetTweetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetTweetRequest) ProtoMessage()               {}
func (*GetTweetRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *GetTweetRequest) GetTID() string {
	if m != nil {
		return m.TID
	}
	return ""
}

type GetTweetResponse struct {
	Tweet *Tweet `protobuf:"bytes,1,opt,name=Tweet" json:"Tweet,omitempty"`
}

func (m *GetTweetResponse) Reset()                    { *m = GetTweetResponse{} }
func (m *GetTweetResponse) String() string            { return proto.CompactTextString(m) }
func (*GetTweetResponse) ProtoMessage()               {}
func (*GetTweetResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *GetTweetResponse) GetTweet() *Tweet {
	if m != nil {
		return m.Tweet
	}
	return nil
}

func init() {
	proto.RegisterType((*NewTweetRequest)(nil), "NewTweetRequest")
	proto.RegisterType((*NewTweetResponse)(nil), "NewTweetResponse")
	proto.RegisterType((*GetTweetRequest)(nil), "GetTweetRequest")
	proto.RegisterType((*GetTweetResponse)(nil), "GetTweetResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Contentd service

type ContentdClient interface {
	NewTweet(ctx context.Context, in *NewTweetRequest, opts ...grpc.CallOption) (*NewTweetResponse, error)
	GetTweet(ctx context.Context, in *GetTweetRequest, opts ...grpc.CallOption) (*GetTweetResponse, error)
}

type contentdClient struct {
	cc *grpc.ClientConn
}

func NewContentdClient(cc *grpc.ClientConn) ContentdClient {
	return &contentdClient{cc}
}

func (c *contentdClient) NewTweet(ctx context.Context, in *NewTweetRequest, opts ...grpc.CallOption) (*NewTweetResponse, error) {
	out := new(NewTweetResponse)
	err := grpc.Invoke(ctx, "/Contentd/NewTweet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentdClient) GetTweet(ctx context.Context, in *GetTweetRequest, opts ...grpc.CallOption) (*GetTweetResponse, error) {
	out := new(GetTweetResponse)
	err := grpc.Invoke(ctx, "/Contentd/GetTweet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Contentd service

type ContentdServer interface {
	NewTweet(context.Context, *NewTweetRequest) (*NewTweetResponse, error)
	GetTweet(context.Context, *GetTweetRequest) (*GetTweetResponse, error)
}

func RegisterContentdServer(s *grpc.Server, srv ContentdServer) {
	s.RegisterService(&_Contentd_serviceDesc, srv)
}

func _Contentd_NewTweet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewTweetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentdServer).NewTweet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Contentd/NewTweet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentdServer).NewTweet(ctx, req.(*NewTweetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Contentd_GetTweet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTweetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentdServer).GetTweet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Contentd/GetTweet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentdServer).GetTweet(ctx, req.(*GetTweetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Contentd_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Contentd",
	HandlerType: (*ContentdServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewTweet",
			Handler:    _Contentd_NewTweet_Handler,
		},
		{
			MethodName: "GetTweet",
			Handler:    _Contentd_GetTweet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

func init() { proto.RegisterFile("service.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 206 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x97, 0xe2, 0x4b, 0xce, 0xcf, 0x2b, 0x49,
	0xcd, 0x2b, 0x49, 0x81, 0xf0, 0x95, 0x3c, 0xb9, 0xf8, 0xfd, 0x52, 0xcb, 0x43, 0xca, 0x53, 0x53,
	0x4b, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x24, 0xb8, 0xd8, 0x9d, 0x21, 0x8a, 0x24,
	0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x60, 0x5c, 0x21, 0x19, 0x2e, 0xce, 0x80, 0xfc, 0xe2, 0x92,
	0xd4, 0xa2, 0x50, 0x4f, 0x17, 0x09, 0x26, 0xb0, 0x1c, 0x42, 0x40, 0x49, 0x88, 0x4b, 0x00, 0x61,
	0x54, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x92, 0x32, 0x17, 0xbf, 0x7b, 0x6a, 0x09, 0x8a, 0xf1,
	0x02, 0x5c, 0xcc, 0x21, 0x9e, 0x2e, 0x50, 0xa3, 0x41, 0x4c, 0x25, 0x4b, 0x2e, 0x01, 0x84, 0x22,
	0x88, 0x46, 0x21, 0x55, 0x2e, 0x56, 0xb0, 0x00, 0x58, 0x1d, 0xb7, 0x11, 0xbf, 0x1e, 0xdc, 0xdd,
	0x10, 0x75, 0x10, 0x59, 0xa3, 0x1c, 0x2e, 0x0e, 0xa8, 0xe3, 0x52, 0x84, 0xf4, 0xb9, 0x38, 0x60,
	0xf6, 0x0b, 0x09, 0xe8, 0xa1, 0xf9, 0x4a, 0x4a, 0x50, 0x0f, 0xdd, 0x71, 0x20, 0x0d, 0x30, 0x7b,
	0x85, 0x04, 0xf4, 0xd0, 0xdc, 0x29, 0x25, 0xa8, 0x87, 0xee, 0xa8, 0x24, 0x36, 0x70, 0x98, 0x19,
	0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xdf, 0x8a, 0xfd, 0xb6, 0x54, 0x01, 0x00, 0x00,
}
