// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/user.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type User struct {
	ID                   int64                `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email                string               `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phone                string               `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,6,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4f917211a932270, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *User) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *User) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type GetUsersRequest struct {
	SalesOrderID         int64    `protobuf:"varint,1,opt,name=salesOrderID,proto3" json:"salesOrderID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUsersRequest) Reset()         { *m = GetUsersRequest{} }
func (m *GetUsersRequest) String() string { return proto.CompactTextString(m) }
func (*GetUsersRequest) ProtoMessage()    {}
func (*GetUsersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4f917211a932270, []int{1}
}

func (m *GetUsersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUsersRequest.Unmarshal(m, b)
}
func (m *GetUsersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUsersRequest.Marshal(b, m, deterministic)
}
func (m *GetUsersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUsersRequest.Merge(m, src)
}
func (m *GetUsersRequest) XXX_Size() int {
	return xxx_messageInfo_GetUsersRequest.Size(m)
}
func (m *GetUsersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUsersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUsersRequest proto.InternalMessageInfo

func (m *GetUsersRequest) GetSalesOrderID() int64 {
	if m != nil {
		return m.SalesOrderID
	}
	return 0
}

type GetUsersResponse struct {
	Users                []*User  `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUsersResponse) Reset()         { *m = GetUsersResponse{} }
func (m *GetUsersResponse) String() string { return proto.CompactTextString(m) }
func (*GetUsersResponse) ProtoMessage()    {}
func (*GetUsersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4f917211a932270, []int{2}
}

func (m *GetUsersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUsersResponse.Unmarshal(m, b)
}
func (m *GetUsersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUsersResponse.Marshal(b, m, deterministic)
}
func (m *GetUsersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUsersResponse.Merge(m, src)
}
func (m *GetUsersResponse) XXX_Size() int {
	return xxx_messageInfo_GetUsersResponse.Size(m)
}
func (m *GetUsersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUsersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUsersResponse proto.InternalMessageInfo

func (m *GetUsersResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "pb.User")
	proto.RegisterType((*GetUsersRequest)(nil), "pb.GetUsersRequest")
	proto.RegisterType((*GetUsersResponse)(nil), "pb.GetUsersResponse")
}

func init() {
	proto.RegisterFile("pb/user.proto", fileDescriptor_c4f917211a932270)
}

var fileDescriptor_c4f917211a932270 = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x4f, 0x4f, 0xb4, 0x30,
	0x18, 0xc4, 0xdf, 0xb2, 0xb0, 0xd9, 0x7d, 0x78, 0xfd, 0x93, 0xba, 0x87, 0x86, 0x83, 0x12, 0x4e,
	0x9c, 0x4a, 0x82, 0x31, 0x7a, 0x35, 0xd9, 0x68, 0xf6, 0x64, 0x82, 0xfa, 0x01, 0x60, 0x79, 0x5c,
	0x49, 0x80, 0xd6, 0xb6, 0xf8, 0x21, 0xfd, 0x54, 0xa6, 0x54, 0x76, 0xb3, 0x5e, 0xbc, 0x75, 0x7e,
	0x99, 0x69, 0x67, 0x0a, 0x27, 0xb2, 0xca, 0x06, 0x8d, 0x8a, 0x4b, 0x25, 0x8c, 0xa0, 0x9e, 0xac,
	0xa2, 0xab, 0x9d, 0x10, 0xbb, 0x16, 0xb3, 0x91, 0x54, 0xc3, 0x5b, 0x66, 0x9a, 0x0e, 0xb5, 0x29,
	0x3b, 0xe9, 0x4c, 0xc9, 0x17, 0x01, 0xff, 0x55, 0xa3, 0xa2, 0xa7, 0xe0, 0x6d, 0xd6, 0x8c, 0xc4,
	0x24, 0x9d, 0x15, 0xde, 0x66, 0x4d, 0x29, 0xf8, 0x7d, 0xd9, 0x21, 0xf3, 0x62, 0x92, 0x2e, 0x8b,
	0xf1, 0x4c, 0x57, 0x10, 0x60, 0x57, 0x36, 0x2d, 0x9b, 0x8d, 0xd0, 0x09, 0x4b, 0xe5, 0xbb, 0xe8,
	0x91, 0xf9, 0x8e, 0x8e, 0x82, 0xde, 0xc1, 0x72, 0xab, 0xb0, 0x34, 0x58, 0xdf, 0x1b, 0x16, 0xc4,
	0x24, 0x0d, 0xf3, 0x88, 0xbb, 0x36, 0x7c, 0x6a, 0xc3, 0x5f, 0xa6, 0x36, 0xc5, 0xc1, 0x6c, 0x93,
	0x83, 0xac, 0x7f, 0x92, 0xf3, 0xbf, 0x93, 0x7b, 0x73, 0x72, 0x03, 0x67, 0x8f, 0x68, 0xec, 0x1c,
	0x5d, 0xe0, 0xc7, 0x80, 0xda, 0xd0, 0x04, 0xfe, 0xeb, 0xb2, 0x45, 0xfd, 0xa4, 0x6a, 0x54, 0xfb,
	0x81, 0x47, 0x2c, 0xc9, 0xe1, 0xfc, 0x10, 0xd3, 0x52, 0xf4, 0x1a, 0xe9, 0x25, 0x04, 0xf6, 0x2b,
	0x35, 0x23, 0xf1, 0x2c, 0x0d, 0xf3, 0x05, 0x97, 0x15, 0xb7, 0x8e, 0xc2, 0xe1, 0xfc, 0x01, 0x42,
	0x2b, 0x9f, 0x51, 0x7d, 0x36, 0x5b, 0xa4, 0xb7, 0xb0, 0x98, 0xae, 0xa0, 0x17, 0xd6, 0xfb, 0xab,
	0x47, 0xb4, 0x3a, 0x86, 0xee, 0x95, 0xe4, 0x5f, 0x35, 0x1f, 0x17, 0x5d, 0x7f, 0x07, 0x00, 0x00,
	0xff, 0xff, 0x9a, 0x93, 0xd0, 0x8b, 0xbc, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*GetUsersResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*GetUsersResponse, error) {
	out := new(GetUsersResponse)
	err := c.cc.Invoke(ctx, "/pb.UserService/GetUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	GetUsers(context.Context, *GetUsersRequest) (*GetUsersResponse, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) GetUsers(ctx context.Context, req *GetUsersRequest) (*GetUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/GetUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUsers(ctx, req.(*GetUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUsers",
			Handler:    _UserService_GetUsers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/user.proto",
}
