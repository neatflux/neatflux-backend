// Code generated by protoc-gen-go. DO NOT EDIT.
// source: content-service.proto

package contentservice

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

type VideoInfo struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VideoInfo) Reset()         { *m = VideoInfo{} }
func (m *VideoInfo) String() string { return proto.CompactTextString(m) }
func (*VideoInfo) ProtoMessage()    {}
func (*VideoInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0ada9bf3eb20d4d, []int{0}
}

func (m *VideoInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VideoInfo.Unmarshal(m, b)
}
func (m *VideoInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VideoInfo.Marshal(b, m, deterministic)
}
func (m *VideoInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VideoInfo.Merge(m, src)
}
func (m *VideoInfo) XXX_Size() int {
	return xxx_messageInfo_VideoInfo.Size(m)
}
func (m *VideoInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_VideoInfo.DiscardUnknown(m)
}

var xxx_messageInfo_VideoInfo proto.InternalMessageInfo

func (m *VideoInfo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *VideoInfo) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

type VideoInfoRequest struct {
	ApiVersion           string   `protobuf:"bytes,1,opt,name=apiVersion,proto3" json:"apiVersion,omitempty"`
	Id                   int64    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VideoInfoRequest) Reset()         { *m = VideoInfoRequest{} }
func (m *VideoInfoRequest) String() string { return proto.CompactTextString(m) }
func (*VideoInfoRequest) ProtoMessage()    {}
func (*VideoInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0ada9bf3eb20d4d, []int{1}
}

func (m *VideoInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VideoInfoRequest.Unmarshal(m, b)
}
func (m *VideoInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VideoInfoRequest.Marshal(b, m, deterministic)
}
func (m *VideoInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VideoInfoRequest.Merge(m, src)
}
func (m *VideoInfoRequest) XXX_Size() int {
	return xxx_messageInfo_VideoInfoRequest.Size(m)
}
func (m *VideoInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VideoInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VideoInfoRequest proto.InternalMessageInfo

func (m *VideoInfoRequest) GetApiVersion() string {
	if m != nil {
		return m.ApiVersion
	}
	return ""
}

func (m *VideoInfoRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type VideoInfoReply struct {
	ApiVersion           string     `protobuf:"bytes,1,opt,name=apiVersion,proto3" json:"apiVersion,omitempty"`
	VideoInfo            *VideoInfo `protobuf:"bytes,2,opt,name=videoInfo,proto3" json:"videoInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *VideoInfoReply) Reset()         { *m = VideoInfoReply{} }
func (m *VideoInfoReply) String() string { return proto.CompactTextString(m) }
func (*VideoInfoReply) ProtoMessage()    {}
func (*VideoInfoReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0ada9bf3eb20d4d, []int{2}
}

func (m *VideoInfoReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VideoInfoReply.Unmarshal(m, b)
}
func (m *VideoInfoReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VideoInfoReply.Marshal(b, m, deterministic)
}
func (m *VideoInfoReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VideoInfoReply.Merge(m, src)
}
func (m *VideoInfoReply) XXX_Size() int {
	return xxx_messageInfo_VideoInfoReply.Size(m)
}
func (m *VideoInfoReply) XXX_DiscardUnknown() {
	xxx_messageInfo_VideoInfoReply.DiscardUnknown(m)
}

var xxx_messageInfo_VideoInfoReply proto.InternalMessageInfo

func (m *VideoInfoReply) GetApiVersion() string {
	if m != nil {
		return m.ApiVersion
	}
	return ""
}

func (m *VideoInfoReply) GetVideoInfo() *VideoInfo {
	if m != nil {
		return m.VideoInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*VideoInfo)(nil), "contentservice.VideoInfo")
	proto.RegisterType((*VideoInfoRequest)(nil), "contentservice.VideoInfoRequest")
	proto.RegisterType((*VideoInfoReply)(nil), "contentservice.VideoInfoReply")
}

func init() { proto.RegisterFile("content-service.proto", fileDescriptor_e0ada9bf3eb20d4d) }

var fileDescriptor_e0ada9bf3eb20d4d = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4d, 0xce, 0xcf, 0x2b,
	0x49, 0xcd, 0x2b, 0xd1, 0x2d, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0xe2, 0x83, 0x0a, 0x43, 0x45, 0x95, 0x0c, 0xb9, 0x38, 0xc3, 0x32, 0x53, 0x52, 0xf3,
	0x3d, 0xf3, 0xd2, 0xf2, 0x85, 0xf8, 0xb8, 0x98, 0x32, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x98,
	0x83, 0x98, 0x32, 0x53, 0x84, 0x44, 0xb8, 0x58, 0x4b, 0x32, 0x4b, 0x72, 0x52, 0x25, 0x98, 0x14,
	0x18, 0x35, 0x38, 0x83, 0x20, 0x1c, 0x25, 0x27, 0x2e, 0x01, 0xb8, 0x96, 0xa0, 0xd4, 0xc2, 0xd2,
	0xd4, 0xe2, 0x12, 0x21, 0x39, 0x2e, 0xae, 0xc4, 0x82, 0xcc, 0xb0, 0xd4, 0xa2, 0xe2, 0xcc, 0xfc,
	0x3c, 0xb0, 0x09, 0x9c, 0x41, 0x48, 0x22, 0x50, 0x93, 0x99, 0x60, 0x26, 0x2b, 0x65, 0x72, 0xf1,
	0x21, 0x99, 0x51, 0x90, 0x53, 0x49, 0xd0, 0x04, 0x73, 0x2e, 0xce, 0x32, 0x98, 0x0e, 0xb0, 0x41,
	0xdc, 0x46, 0x92, 0x7a, 0xa8, 0x9e, 0xd1, 0x43, 0x18, 0x89, 0x50, 0x6b, 0x94, 0xc2, 0x85, 0xe6,
	0x67, 0xa1, 0x20, 0x2e, 0x1e, 0xf7, 0xd4, 0x12, 0x84, 0xb7, 0x15, 0x70, 0x9b, 0x03, 0xf1, 0x9e,
	0x94, 0x1c, 0x1e, 0x15, 0x05, 0x39, 0x95, 0x4a, 0x0c, 0x49, 0x6c, 0xe0, 0xe0, 0x35, 0x06, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x53, 0x8e, 0x59, 0x47, 0x77, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ContentserviceClient is the client API for Contentservice service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ContentserviceClient interface {
	GetVideoInfo(ctx context.Context, in *VideoInfoRequest, opts ...grpc.CallOption) (*VideoInfoReply, error)
}

type contentserviceClient struct {
	cc *grpc.ClientConn
}

func NewContentserviceClient(cc *grpc.ClientConn) ContentserviceClient {
	return &contentserviceClient{cc}
}

func (c *contentserviceClient) GetVideoInfo(ctx context.Context, in *VideoInfoRequest, opts ...grpc.CallOption) (*VideoInfoReply, error) {
	out := new(VideoInfoReply)
	err := c.cc.Invoke(ctx, "/contentservice.contentservice/GetVideoInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContentserviceServer is the server API for Contentservice service.
type ContentserviceServer interface {
	GetVideoInfo(context.Context, *VideoInfoRequest) (*VideoInfoReply, error)
}

// UnimplementedContentserviceServer can be embedded to have forward compatible implementations.
type UnimplementedContentserviceServer struct {
}

func (*UnimplementedContentserviceServer) GetVideoInfo(ctx context.Context, req *VideoInfoRequest) (*VideoInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVideoInfo not implemented")
}

func RegisterContentserviceServer(s *grpc.Server, srv ContentserviceServer) {
	s.RegisterService(&_Contentservice_serviceDesc, srv)
}

func _Contentservice_GetVideoInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VideoInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentserviceServer).GetVideoInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contentservice.contentservice/GetVideoInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentserviceServer).GetVideoInfo(ctx, req.(*VideoInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Contentservice_serviceDesc = grpc.ServiceDesc{
	ServiceName: "contentservice.contentservice",
	HandlerType: (*ContentserviceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVideoInfo",
			Handler:    _Contentservice_GetVideoInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "content-service.proto",
}
