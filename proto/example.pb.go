// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: example.proto

// GoGo gRPC Example
//
// This example is used to show how to use gRPC and
// gRPC-Gateway with GoGo Protobuf.

package example

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/googleapis/google/api"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	types "github.com/gogo/protobuf/types"
	golang_proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "github.com/mwitkow/go-proto-validators"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Role int32

const (
	Role_GUEST  Role = 0
	Role_MEMBER Role = 1
	Role_ADMIN  Role = 2
)

var Role_name = map[int32]string{
	0: "GUEST",
	1: "MEMBER",
	2: "ADMIN",
}

var Role_value = map[string]int32{
	"GUEST":  0,
	"MEMBER": 1,
	"ADMIN":  2,
}

func (x Role) String() string {
	return proto.EnumName(Role_name, int32(x))
}

func (Role) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{0}
}

type User struct {
	ID                   uint32     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Role                 Role       `protobuf:"varint,2,opt,name=role,proto3,enum=example.Role" json:"role,omitempty"`
	CreateDate           *time.Time `protobuf:"bytes,3,opt,name=create_date,json=createDate,proto3,stdtime" json:"create_date,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{0}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_User.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return m.Size()
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetID() uint32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *User) GetRole() Role {
	if m != nil {
		return m.Role
	}
	return Role_GUEST
}

func (m *User) GetCreateDate() *time.Time {
	if m != nil {
		return m.CreateDate
	}
	return nil
}

func (*User) XXX_MessageName() string {
	return "example.User"
}

type UserRole struct {
	Role                 Role     `protobuf:"varint,1,opt,name=role,proto3,enum=example.Role" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRole) Reset()         { *m = UserRole{} }
func (m *UserRole) String() string { return proto.CompactTextString(m) }
func (*UserRole) ProtoMessage()    {}
func (*UserRole) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{1}
}
func (m *UserRole) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserRole) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserRole.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UserRole) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRole.Merge(m, src)
}
func (m *UserRole) XXX_Size() int {
	return m.Size()
}
func (m *UserRole) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRole.DiscardUnknown(m)
}

var xxx_messageInfo_UserRole proto.InternalMessageInfo

func (m *UserRole) GetRole() Role {
	if m != nil {
		return m.Role
	}
	return Role_GUEST
}

func (*UserRole) XXX_MessageName() string {
	return "example.UserRole"
}

type UpdateUserRequest struct {
	// The user resource which replaces the resource on the server.
	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	// The update mask applies to the resource. For the `FieldMask` definition,
	// see https://developers.google.com/protocol-buffers/docs/reference/google.protobuf#fieldmask
	UpdateMask           *types.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *UpdateUserRequest) Reset()         { *m = UpdateUserRequest{} }
func (m *UpdateUserRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateUserRequest) ProtoMessage()    {}
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{2}
}
func (m *UpdateUserRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpdateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpdateUserRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UpdateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserRequest.Merge(m, src)
}
func (m *UpdateUserRequest) XXX_Size() int {
	return m.Size()
}
func (m *UpdateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserRequest proto.InternalMessageInfo

func (m *UpdateUserRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *UpdateUserRequest) GetUpdateMask() *types.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

func (*UpdateUserRequest) XXX_MessageName() string {
	return "example.UpdateUserRequest"
}

type ListUsersRequest struct {
	// Only list users created after this timestamp
	CreatedSince *time.Time `protobuf:"bytes,1,opt,name=created_since,json=createdSince,proto3,stdtime" json:"created_since,omitempty"`
	// Only list users older than this Duration
	OlderThan            *time.Duration `protobuf:"bytes,2,opt,name=older_than,json=olderThan,proto3,stdduration" json:"older_than,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ListUsersRequest) Reset()         { *m = ListUsersRequest{} }
func (m *ListUsersRequest) String() string { return proto.CompactTextString(m) }
func (*ListUsersRequest) ProtoMessage()    {}
func (*ListUsersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{3}
}
func (m *ListUsersRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListUsersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListUsersRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListUsersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUsersRequest.Merge(m, src)
}
func (m *ListUsersRequest) XXX_Size() int {
	return m.Size()
}
func (m *ListUsersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUsersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListUsersRequest proto.InternalMessageInfo

func (m *ListUsersRequest) GetCreatedSince() *time.Time {
	if m != nil {
		return m.CreatedSince
	}
	return nil
}

func (m *ListUsersRequest) GetOlderThan() *time.Duration {
	if m != nil {
		return m.OlderThan
	}
	return nil
}

func (*ListUsersRequest) XXX_MessageName() string {
	return "example.ListUsersRequest"
}
func init() {
	proto.RegisterEnum("example.Role", Role_name, Role_value)
	golang_proto.RegisterEnum("example.Role", Role_name, Role_value)
	proto.RegisterType((*User)(nil), "example.User")
	golang_proto.RegisterType((*User)(nil), "example.User")
	proto.RegisterType((*UserRole)(nil), "example.UserRole")
	golang_proto.RegisterType((*UserRole)(nil), "example.UserRole")
	proto.RegisterType((*UpdateUserRequest)(nil), "example.UpdateUserRequest")
	golang_proto.RegisterType((*UpdateUserRequest)(nil), "example.UpdateUserRequest")
	proto.RegisterType((*ListUsersRequest)(nil), "example.ListUsersRequest")
	golang_proto.RegisterType((*ListUsersRequest)(nil), "example.ListUsersRequest")
}

func init() { proto.RegisterFile("example.proto", fileDescriptor_15a1dc8d40dadaa6) }
func init() { golang_proto.RegisterFile("example.proto", fileDescriptor_15a1dc8d40dadaa6) }

var fileDescriptor_15a1dc8d40dadaa6 = []byte{
	// 721 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xcf, 0x6e, 0x13, 0x39,
	0x18, 0xaf, 0xa7, 0x69, 0xbb, 0x75, 0x36, 0xdb, 0xd4, 0xda, 0xdd, 0xa6, 0xb3, 0x55, 0x32, 0x3b,
	0xe2, 0x10, 0x45, 0x64, 0xa6, 0x0d, 0x12, 0x42, 0xad, 0x04, 0x4a, 0x48, 0x40, 0x05, 0x52, 0x41,
	0xda, 0x5e, 0xe0, 0x10, 0xb9, 0x19, 0x77, 0x62, 0x75, 0x66, 0x3c, 0x8c, 0x9d, 0x94, 0x0a, 0x71,
	0xe1, 0x01, 0x10, 0x82, 0x0b, 0xbc, 0x00, 0xcf, 0xc1, 0xb1, 0x47, 0x24, 0x4e, 0x5c, 0x28, 0x4a,
	0x79, 0x10, 0x64, 0x67, 0x92, 0x36, 0xd3, 0x02, 0x27, 0xfb, 0xf3, 0xf7, 0xfb, 0xfd, 0x3e, 0x7f,
	0x7f, 0x6c, 0x98, 0x21, 0xcf, 0xb0, 0x1f, 0x7a, 0xc4, 0x0a, 0x23, 0x26, 0x18, 0x9a, 0x8b, 0x4d,
	0xfd, 0x3f, 0x97, 0x31, 0xd7, 0x23, 0xb6, 0x3a, 0xde, 0xeb, 0xed, 0xdb, 0xc4, 0x0f, 0xc5, 0xd1,
	0x10, 0xa5, 0x17, 0x92, 0x4e, 0x41, 0x7d, 0xc2, 0x05, 0xf6, 0xc3, 0x18, 0x90, 0x4f, 0x02, 0x9c,
	0x5e, 0x84, 0x05, 0x65, 0x41, 0xec, 0x37, 0x92, 0xfe, 0x7d, 0x4a, 0x3c, 0xa7, 0xed, 0x63, 0x7e,
	0x10, 0x23, 0x56, 0x62, 0x04, 0x0e, 0xa9, 0x8d, 0x83, 0x80, 0x09, 0x45, 0xe7, 0xb1, 0xf7, 0xaa,
	0x5a, 0x3a, 0x65, 0x97, 0x04, 0x65, 0x7e, 0x88, 0x5d, 0x97, 0x44, 0x36, 0x0b, 0x15, 0xe2, 0x12,
	0x74, 0xd9, 0xa5, 0xa2, 0xdb, 0xdb, 0xb3, 0x3a, 0xcc, 0xb7, 0x5d, 0xe6, 0xb2, 0xb3, 0xb0, 0xd2,
	0x52, 0x86, 0xda, 0xc5, 0xf0, 0xeb, 0xe7, 0xe0, 0xfe, 0x21, 0x15, 0x07, 0xec, 0xd0, 0x76, 0x59,
	0x59, 0x39, 0xcb, 0x7d, 0xec, 0x51, 0x07, 0x0b, 0x16, 0x71, 0x7b, 0xbc, 0x1d, 0xf2, 0xcc, 0x0f,
	0x00, 0xa6, 0x76, 0x39, 0x89, 0xd0, 0x0d, 0xa8, 0x51, 0x27, 0x07, 0x0c, 0x50, 0xcc, 0xd4, 0x8a,
	0x83, 0xaf, 0x05, 0x6d, 0xb3, 0x3e, 0x38, 0x29, 0xe4, 0xb3, 0x53, 0x25, 0x7d, 0xb3, 0x6e, 0xf8,
	0x3d, 0x2e, 0x0c, 0x6c, 0x84, 0x8c, 0x53, 0x41, 0xfb, 0xc4, 0xa0, 0x81, 0x20, 0x2e, 0x89, 0x5a,
	0x1a, 0x75, 0xd0, 0xff, 0x30, 0x15, 0x31, 0x8f, 0xe4, 0x34, 0x03, 0x14, 0xff, 0xaa, 0x64, 0xac,
	0x51, 0x73, 0x5a, 0xcc, 0x23, 0x2d, 0xe5, 0x42, 0x55, 0x98, 0xee, 0x44, 0x04, 0x0b, 0xd2, 0x76,
	0xb0, 0x20, 0xb9, 0x69, 0x03, 0x14, 0xd3, 0x15, 0xdd, 0x1a, 0x96, 0xcb, 0x1a, 0x65, 0x66, 0xed,
	0x8c, 0x3a, 0x52, 0x4b, 0xbd, 0x3e, 0x29, 0x80, 0x16, 0x1c, 0x92, 0xea, 0x58, 0x10, 0xb3, 0x0c,
	0xff, 0x90, 0xf7, 0x94, 0xa2, 0xe3, 0x88, 0xe0, 0xa7, 0x11, 0x4d, 0x0e, 0x17, 0x77, 0x43, 0x19,
	0x4c, 0x91, 0xc8, 0xd3, 0x1e, 0xe1, 0x42, 0xf2, 0x7a, 0x9c, 0x44, 0x8a, 0x97, 0x3e, 0xc7, 0x53,
	0x18, 0xe5, 0x42, 0x1b, 0x30, 0xdd, 0x53, 0x3c, 0xd5, 0x57, 0x95, 0xd3, 0x65, 0x37, 0xbd, 0x23,
	0x5b, 0xdf, 0xc4, 0xfc, 0xa0, 0x05, 0x87, 0x70, 0xb9, 0x37, 0xdf, 0x03, 0x98, 0x7d, 0x40, 0xb9,
	0x90, 0x7a, 0x7c, 0x14, 0xb4, 0x01, 0x33, 0xc3, 0x34, 0x9c, 0x36, 0xa7, 0x41, 0x87, 0xc4, 0xd1,
	0x7f, 0x9f, 0xfd, 0x9f, 0x31, 0x6d, 0x5b, 0xb2, 0xd0, 0x4d, 0x08, 0x99, 0xe7, 0x90, 0xa8, 0x2d,
	0xba, 0x38, 0x88, 0xef, 0xb5, 0x7c, 0x41, 0xa3, 0x1e, 0x8f, 0x6c, 0x2d, 0xf5, 0x4e, 0x4a, 0xcc,
	0x2b, 0xca, 0x4e, 0x17, 0x07, 0xa5, 0x22, 0x4c, 0xa9, 0xda, 0xcd, 0xc3, 0x99, 0xbb, 0xbb, 0x8d,
	0xed, 0x9d, 0xec, 0x14, 0x82, 0x70, 0xb6, 0xd9, 0x68, 0xd6, 0x1a, 0xad, 0x2c, 0x90, 0xc7, 0xd5,
	0x7a, 0x73, 0x73, 0x2b, 0xab, 0x55, 0xbe, 0x68, 0x30, 0x2d, 0x33, 0xd8, 0x26, 0x51, 0x9f, 0x76,
	0x08, 0xba, 0x07, 0xe7, 0xaa, 0x8e, 0xa3, 0x86, 0x64, 0xb2, 0x64, 0xfa, 0xbf, 0x17, 0xe2, 0x37,
	0xe4, 0x83, 0x33, 0x73, 0x2f, 0x3f, 0x7f, 0x7f, 0xab, 0x21, 0x33, 0xa3, 0x5e, 0x42, 0x7f, 0xcd,
	0x96, 0xa5, 0xe5, 0xeb, 0xa0, 0x84, 0xb6, 0xe0, 0xfc, 0xb8, 0x40, 0x68, 0x79, 0xac, 0x96, 0x2c,
	0x9a, 0x3e, 0x19, 0xc8, 0xfc, 0x47, 0x09, 0x2e, 0xa0, 0x49, 0xc1, 0x55, 0x80, 0x1e, 0xc1, 0x85,
	0x31, 0xb7, 0x76, 0xa4, 0x12, 0x5c, 0x9c, 0x6c, 0x2b, 0xf3, 0x48, 0x52, 0x4d, 0x57, 0x6a, 0x7f,
	0x23, 0x34, 0xa1, 0x66, 0xcb, 0xb9, 0x59, 0x05, 0xe8, 0x09, 0x84, 0x67, 0x93, 0x83, 0xf4, 0x33,
	0x6a, 0x72, 0x9c, 0x92, 0xb2, 0xa6, 0x92, 0x5d, 0xa9, 0x2c, 0x4d, 0xca, 0x3e, 0x97, 0x8b, 0x45,
	0x9d, 0x17, 0xeb, 0xa0, 0x54, 0x7b, 0x05, 0xde, 0x54, 0xef, 0xa3, 0x99, 0xca, 0xf4, 0x9a, 0xb5,
	0x5a, 0x02, 0x5a, 0x74, 0x0b, 0x2e, 0xb9, 0xad, 0x87, 0xb7, 0x8d, 0x58, 0xc8, 0x88, 0x88, 0x7a,
	0x62, 0x2c, 0x3a, 0x42, 0x57, 0xba, 0x42, 0x84, 0x7c, 0xdd, 0xb6, 0x93, 0x1f, 0x81, 0x1b, 0x85,
	0x9d, 0x72, 0x4c, 0x38, 0x1e, 0xe4, 0xc1, 0xa7, 0x41, 0x1e, 0x7c, 0x1b, 0xe4, 0xc1, 0xc7, 0xd3,
	0x3c, 0x38, 0x3e, 0xcd, 0x83, 0xc7, 0xa5, 0x5f, 0x31, 0x86, 0xff, 0xc8, 0x46, 0x6c, 0xed, 0xcd,
	0x2a, 0xf3, 0xda, 0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x59, 0x15, 0x86, 0x88, 0x53, 0x05, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	AddUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*types.Empty, error)
	ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (UserService_ListUsersClient, error)
	ListUsersByRole(ctx context.Context, in *UserRole, opts ...grpc.CallOption) (UserService_ListUsersByRoleClient, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*User, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) AddUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/example.UserService/AddUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (UserService_ListUsersClient, error) {
	stream, err := c.cc.NewStream(ctx, &_UserService_serviceDesc.Streams[0], "/example.UserService/ListUsers", opts...)
	if err != nil {
		return nil, err
	}
	x := &userServiceListUsersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UserService_ListUsersClient interface {
	Recv() (*User, error)
	grpc.ClientStream
}

type userServiceListUsersClient struct {
	grpc.ClientStream
}

func (x *userServiceListUsersClient) Recv() (*User, error) {
	m := new(User)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *userServiceClient) ListUsersByRole(ctx context.Context, in *UserRole, opts ...grpc.CallOption) (UserService_ListUsersByRoleClient, error) {
	stream, err := c.cc.NewStream(ctx, &_UserService_serviceDesc.Streams[1], "/example.UserService/ListUsersByRole", opts...)
	if err != nil {
		return nil, err
	}
	x := &userServiceListUsersByRoleClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UserService_ListUsersByRoleClient interface {
	Recv() (*User, error)
	grpc.ClientStream
}

type userServiceListUsersByRoleClient struct {
	grpc.ClientStream
}

func (x *userServiceListUsersByRoleClient) Recv() (*User, error) {
	m := new(User)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *userServiceClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/example.UserService/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	AddUser(context.Context, *User) (*types.Empty, error)
	ListUsers(*ListUsersRequest, UserService_ListUsersServer) error
	ListUsersByRole(*UserRole, UserService_ListUsersByRoleServer) error
	UpdateUser(context.Context, *UpdateUserRequest) (*User, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) AddUser(ctx context.Context, req *User) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUser not implemented")
}
func (*UnimplementedUserServiceServer) ListUsers(req *ListUsersRequest, srv UserService_ListUsersServer) error {
	return status.Errorf(codes.Unimplemented, "method ListUsers not implemented")
}
func (*UnimplementedUserServiceServer) ListUsersByRole(req *UserRole, srv UserService_ListUsersByRoleServer) error {
	return status.Errorf(codes.Unimplemented, "method ListUsersByRole not implemented")
}
func (*UnimplementedUserServiceServer) UpdateUser(ctx context.Context, req *UpdateUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.UserService/AddUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AddUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ListUsers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListUsersRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UserServiceServer).ListUsers(m, &userServiceListUsersServer{stream})
}

type UserService_ListUsersServer interface {
	Send(*User) error
	grpc.ServerStream
}

type userServiceListUsersServer struct {
	grpc.ServerStream
}

func (x *userServiceListUsersServer) Send(m *User) error {
	return x.ServerStream.SendMsg(m)
}

func _UserService_ListUsersByRole_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(UserRole)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UserServiceServer).ListUsersByRole(m, &userServiceListUsersByRoleServer{stream})
}

type UserService_ListUsersByRoleServer interface {
	Send(*User) error
	grpc.ServerStream
}

type userServiceListUsersByRoleServer struct {
	grpc.ServerStream
}

func (x *userServiceListUsersByRoleServer) Send(m *User) error {
	return x.ServerStream.SendMsg(m)
}

func _UserService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.UserService/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "example.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddUser",
			Handler:    _UserService_AddUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UserService_UpdateUser_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListUsers",
			Handler:       _UserService_ListUsers_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListUsersByRole",
			Handler:       _UserService_ListUsersByRole_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "example.proto",
}

func (m *User) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *User) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *User) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.CreateDate != nil {
		n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.CreateDate, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(*m.CreateDate):])
		if err1 != nil {
			return 0, err1
		}
		i -= n1
		i = encodeVarintExample(dAtA, i, uint64(n1))
		i--
		dAtA[i] = 0x1a
	}
	if m.Role != 0 {
		i = encodeVarintExample(dAtA, i, uint64(m.Role))
		i--
		dAtA[i] = 0x10
	}
	if m.ID != 0 {
		i = encodeVarintExample(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *UserRole) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserRole) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UserRole) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Role != 0 {
		i = encodeVarintExample(dAtA, i, uint64(m.Role))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *UpdateUserRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpdateUserRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UpdateUserRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.UpdateMask != nil {
		{
			size, err := m.UpdateMask.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintExample(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.User != nil {
		{
			size, err := m.User.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintExample(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ListUsersRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListUsersRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListUsersRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.OlderThan != nil {
		n4, err4 := github_com_gogo_protobuf_types.StdDurationMarshalTo(*m.OlderThan, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(*m.OlderThan):])
		if err4 != nil {
			return 0, err4
		}
		i -= n4
		i = encodeVarintExample(dAtA, i, uint64(n4))
		i--
		dAtA[i] = 0x12
	}
	if m.CreatedSince != nil {
		n5, err5 := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.CreatedSince, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(*m.CreatedSince):])
		if err5 != nil {
			return 0, err5
		}
		i -= n5
		i = encodeVarintExample(dAtA, i, uint64(n5))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintExample(dAtA []byte, offset int, v uint64) int {
	offset -= sovExample(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *User) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovExample(uint64(m.ID))
	}
	if m.Role != 0 {
		n += 1 + sovExample(uint64(m.Role))
	}
	if m.CreateDate != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.CreateDate)
		n += 1 + l + sovExample(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *UserRole) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Role != 0 {
		n += 1 + sovExample(uint64(m.Role))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *UpdateUserRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.User != nil {
		l = m.User.Size()
		n += 1 + l + sovExample(uint64(l))
	}
	if m.UpdateMask != nil {
		l = m.UpdateMask.Size()
		n += 1 + l + sovExample(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ListUsersRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CreatedSince != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.CreatedSince)
		n += 1 + l + sovExample(uint64(l))
	}
	if m.OlderThan != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdDuration(*m.OlderThan)
		n += 1 + l + sovExample(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovExample(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozExample(x uint64) (n int) {
	return sovExample(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *User) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExample
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: User: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: User: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExample
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Role", wireType)
			}
			m.Role = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExample
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Role |= Role(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreateDate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExample
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthExample
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthExample
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CreateDate == nil {
				m.CreateDate = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.CreateDate, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExample(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthExample
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthExample
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *UserRole) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExample
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UserRole: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserRole: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Role", wireType)
			}
			m.Role = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExample
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Role |= Role(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipExample(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthExample
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthExample
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *UpdateUserRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExample
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UpdateUserRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpdateUserRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExample
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthExample
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthExample
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.User == nil {
				m.User = &User{}
			}
			if err := m.User.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdateMask", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExample
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthExample
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthExample
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.UpdateMask == nil {
				m.UpdateMask = &types.FieldMask{}
			}
			if err := m.UpdateMask.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExample(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthExample
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthExample
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ListUsersRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExample
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ListUsersRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListUsersRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedSince", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExample
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthExample
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthExample
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CreatedSince == nil {
				m.CreatedSince = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.CreatedSince, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OlderThan", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExample
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthExample
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthExample
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.OlderThan == nil {
				m.OlderThan = new(time.Duration)
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(m.OlderThan, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExample(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthExample
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthExample
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipExample(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowExample
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowExample
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowExample
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthExample
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupExample
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthExample
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthExample        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowExample          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupExample = fmt.Errorf("proto: unexpected end of group")
)
