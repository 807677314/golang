// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user_server.proto

package grpc_server

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

type UserRequest struct {
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserName             string   `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRequest) Reset()         { *m = UserRequest{} }
func (m *UserRequest) String() string { return proto.CompactTextString(m) }
func (*UserRequest) ProtoMessage()    {}
func (*UserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_44f3f4a2a039ae28, []int{0}
}

func (m *UserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRequest.Unmarshal(m, b)
}
func (m *UserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRequest.Marshal(b, m, deterministic)
}
func (m *UserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRequest.Merge(m, src)
}
func (m *UserRequest) XXX_Size() int {
	return xxx_messageInfo_UserRequest.Size(m)
}
func (m *UserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserRequest proto.InternalMessageInfo

func (m *UserRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *UserRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

type UserIDResponse struct {
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserIDResponse) Reset()         { *m = UserIDResponse{} }
func (m *UserIDResponse) String() string { return proto.CompactTextString(m) }
func (*UserIDResponse) ProtoMessage()    {}
func (*UserIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_44f3f4a2a039ae28, []int{1}
}

func (m *UserIDResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserIDResponse.Unmarshal(m, b)
}
func (m *UserIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserIDResponse.Marshal(b, m, deterministic)
}
func (m *UserIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserIDResponse.Merge(m, src)
}
func (m *UserIDResponse) XXX_Size() int {
	return xxx_messageInfo_UserIDResponse.Size(m)
}
func (m *UserIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserIDResponse proto.InternalMessageInfo

func (m *UserIDResponse) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type UserNameResponse struct {
	UserName             string   `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserNameResponse) Reset()         { *m = UserNameResponse{} }
func (m *UserNameResponse) String() string { return proto.CompactTextString(m) }
func (*UserNameResponse) ProtoMessage()    {}
func (*UserNameResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_44f3f4a2a039ae28, []int{2}
}

func (m *UserNameResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserNameResponse.Unmarshal(m, b)
}
func (m *UserNameResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserNameResponse.Marshal(b, m, deterministic)
}
func (m *UserNameResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserNameResponse.Merge(m, src)
}
func (m *UserNameResponse) XXX_Size() int {
	return xxx_messageInfo_UserNameResponse.Size(m)
}
func (m *UserNameResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserNameResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserNameResponse proto.InternalMessageInfo

func (m *UserNameResponse) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func init() {
	proto.RegisterType((*UserRequest)(nil), "controllers.UserRequest")
	proto.RegisterType((*UserIDResponse)(nil), "controllers.UserIDResponse")
	proto.RegisterType((*UserNameResponse)(nil), "controllers.UserNameResponse")
}

func init() { proto.RegisterFile("user_server.proto", fileDescriptor_44f3f4a2a039ae28) }

var fileDescriptor_44f3f4a2a039ae28 = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x2d, 0x4e, 0x2d,
	0x8a, 0x2f, 0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4e,
	0xce, 0xcf, 0x2b, 0x29, 0xca, 0xcf, 0xc9, 0x49, 0x2d, 0x2a, 0x56, 0x72, 0xe6, 0xe2, 0x0e, 0x2d,
	0x4e, 0x2d, 0x0a, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe7, 0x62, 0x07, 0x6b, 0xc8,
	0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0e, 0x62, 0x03, 0x71, 0x3d, 0x53, 0x84, 0xa4, 0xb9,
	0x38, 0xc1, 0x12, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x1c, 0x20,
	0x01, 0xbf, 0xc4, 0xdc, 0x54, 0x25, 0x4d, 0x2e, 0x3e, 0x90, 0x21, 0x9e, 0x2e, 0x41, 0xa9, 0xc5,
	0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x38, 0xcd, 0x51, 0xd2, 0xe7, 0x12, 0x08, 0x85, 0x6a, 0x83, 0x2b,
	0x46, 0x31, 0x9b, 0x11, 0xd5, 0x6c, 0xa3, 0x59, 0x8c, 0x5c, 0x5c, 0x20, 0x1d, 0xc1, 0x60, 0x2f,
	0x08, 0xb9, 0x70, 0x71, 0xbb, 0xa7, 0x96, 0x80, 0x04, 0x9c, 0x2a, 0x3d, 0x53, 0x84, 0x24, 0xf4,
	0x90, 0x3c, 0xa3, 0x87, 0xe4, 0x13, 0x29, 0x69, 0x0c, 0x19, 0x24, 0xe7, 0xb9, 0xc1, 0x4d, 0x01,
	0xd9, 0x81, 0xc7, 0x14, 0x59, 0x0c, 0x19, 0x64, 0x97, 0x27, 0xb1, 0x81, 0x43, 0xd4, 0x18, 0x10,
	0x00, 0x00, 0xff, 0xff, 0x3b, 0xaf, 0xcd, 0x90, 0x66, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServerClient is the client API for UserServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServerClient interface {
	GetUserById(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserIDResponse, error)
	GetUserName(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserNameResponse, error)
}

type userServerClient struct {
	cc *grpc.ClientConn
}

func NewUserServerClient(cc *grpc.ClientConn) UserServerClient {
	return &userServerClient{cc}
}

func (c *userServerClient) GetUserById(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserIDResponse, error) {
	out := new(UserIDResponse)
	err := c.cc.Invoke(ctx, "/controllers.UserServer/GetUserById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServerClient) GetUserName(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserNameResponse, error) {
	out := new(UserNameResponse)
	err := c.cc.Invoke(ctx, "/controllers.UserServer/GetUserName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServerServer is the server API for UserServer service.
type UserServerServer interface {
	GetUserById(context.Context, *UserRequest) (*UserIDResponse, error)
	GetUserName(context.Context, *UserRequest) (*UserNameResponse, error)
}

// UnimplementedUserServerServer can be embedded to have forward compatible implementations.
type UnimplementedUserServerServer struct {
}

func (*UnimplementedUserServerServer) GetUserById(ctx context.Context, req *UserRequest) (*UserIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserById not implemented")
}
func (*UnimplementedUserServerServer) GetUserName(ctx context.Context, req *UserRequest) (*UserNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserName not implemented")
}

func RegisterUserServerServer(s *grpc.Server, srv UserServerServer) {
	s.RegisterService(&_UserServer_serviceDesc, srv)
}

func _UserServer_GetUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServerServer).GetUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/controllers.UserServer/GetUserById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServerServer).GetUserById(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServer_GetUserName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServerServer).GetUserName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/controllers.UserServer/GetUserName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServerServer).GetUserName(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "controllers.UserServer",
	HandlerType: (*UserServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserById",
			Handler:    _UserServer_GetUserById_Handler,
		},
		{
			MethodName: "GetUserName",
			Handler:    _UserServer_GetUserName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_server.proto",
}