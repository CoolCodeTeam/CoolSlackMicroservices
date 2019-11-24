// Code generated by protoc-gen-go. DO NOT EDIT.
// source: users.proto

package grpc_utils

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

type UserID struct {
	ID                   uint64   `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserID) Reset()         { *m = UserID{} }
func (m *UserID) String() string { return proto.CompactTextString(m) }
func (*UserID) ProtoMessage()    {}
func (*UserID) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{0}
}

func (m *UserID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserID.Unmarshal(m, b)
}
func (m *UserID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserID.Marshal(b, m, deterministic)
}
func (m *UserID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserID.Merge(m, src)
}
func (m *UserID) XXX_Size() int {
	return xxx_messageInfo_UserID.Size(m)
}
func (m *UserID) XXX_DiscardUnknown() {
	xxx_messageInfo_UserID.DiscardUnknown(m)
}

var xxx_messageInfo_UserID proto.InternalMessageInfo

func (m *UserID) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

type UserEmail struct {
	Email                string   `protobuf:"bytes,1,opt,name=Email,proto3" json:"Email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserEmail) Reset()         { *m = UserEmail{} }
func (m *UserEmail) String() string { return proto.CompactTextString(m) }
func (*UserEmail) ProtoMessage()    {}
func (*UserEmail) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{1}
}

func (m *UserEmail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserEmail.Unmarshal(m, b)
}
func (m *UserEmail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserEmail.Marshal(b, m, deterministic)
}
func (m *UserEmail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserEmail.Merge(m, src)
}
func (m *UserEmail) XXX_Size() int {
	return xxx_messageInfo_UserEmail.Size(m)
}
func (m *UserEmail) XXX_DiscardUnknown() {
	xxx_messageInfo_UserEmail.DiscardUnknown(m)
}

var xxx_messageInfo_UserEmail proto.InternalMessageInfo

func (m *UserEmail) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type UserName struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserName) Reset()         { *m = UserName{} }
func (m *UserName) String() string { return proto.CompactTextString(m) }
func (*UserName) ProtoMessage()    {}
func (*UserName) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{2}
}

func (m *UserName) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserName.Unmarshal(m, b)
}
func (m *UserName) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserName.Marshal(b, m, deterministic)
}
func (m *UserName) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserName.Merge(m, src)
}
func (m *UserName) XXX_Size() int {
	return xxx_messageInfo_UserName.Size(m)
}
func (m *UserName) XXX_DiscardUnknown() {
	xxx_messageInfo_UserName.DiscardUnknown(m)
}

var xxx_messageInfo_UserName proto.InternalMessageInfo

func (m *UserName) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Session struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Session) Reset()         { *m = Session{} }
func (m *Session) String() string { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()    {}
func (*Session) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{3}
}

func (m *Session) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Session.Unmarshal(m, b)
}
func (m *Session) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Session.Marshal(b, m, deterministic)
}
func (m *Session) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Session.Merge(m, src)
}
func (m *Session) XXX_Size() int {
	return xxx_messageInfo_Session.Size(m)
}
func (m *Session) XXX_DiscardUnknown() {
	xxx_messageInfo_Session.DiscardUnknown(m)
}

var xxx_messageInfo_Session proto.InternalMessageInfo

func (m *Session) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type User struct {
	ID                   uint64   `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=Username,proto3" json:"Username,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=Email,proto3" json:"Email,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=Name,proto3" json:"Name,omitempty"`
	Password             string   `protobuf:"bytes,5,opt,name=Password,proto3" json:"Password,omitempty"`
	Status               string   `protobuf:"bytes,6,opt,name=Status,proto3" json:"Status,omitempty"`
	Phone                string   `protobuf:"bytes,7,opt,name=Phone,proto3" json:"Phone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{4}
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

func (m *User) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *User) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

type Users struct {
	Users                []*User  `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Users) Reset()         { *m = Users{} }
func (m *Users) String() string { return proto.CompactTextString(m) }
func (*Users) ProtoMessage()    {}
func (*Users) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{5}
}

func (m *Users) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Users.Unmarshal(m, b)
}
func (m *Users) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Users.Marshal(b, m, deterministic)
}
func (m *Users) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Users.Merge(m, src)
}
func (m *Users) XXX_Size() int {
	return xxx_messageInfo_Users.Size(m)
}
func (m *Users) XXX_DiscardUnknown() {
	xxx_messageInfo_Users.DiscardUnknown(m)
}

var xxx_messageInfo_Users proto.InternalMessageInfo

func (m *Users) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{6}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*UserID)(nil), "grpc_utils.UserID")
	proto.RegisterType((*UserEmail)(nil), "grpc_utils.UserEmail")
	proto.RegisterType((*UserName)(nil), "grpc_utils.UserName")
	proto.RegisterType((*Session)(nil), "grpc_utils.Session")
	proto.RegisterType((*User)(nil), "grpc_utils.User")
	proto.RegisterType((*Users)(nil), "grpc_utils.Users")
	proto.RegisterType((*Empty)(nil), "grpc_utils.Empty")
}

func init() { proto.RegisterFile("users.proto", fileDescriptor_030765f334c86cea) }

var fileDescriptor_030765f334c86cea = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x4f, 0xea, 0x40,
	0x10, 0xc7, 0x53, 0x68, 0xcb, 0x63, 0x78, 0x21, 0xbc, 0x79, 0x68, 0x9a, 0x1e, 0x14, 0xf7, 0x60,
	0x48, 0x8c, 0x18, 0xc1, 0x8b, 0xf1, 0xa6, 0x55, 0xd3, 0xc4, 0x18, 0x42, 0xc3, 0xd9, 0x54, 0xd8,
	0x94, 0x4d, 0x60, 0xdb, 0x74, 0x5b, 0x0c, 0x37, 0x3f, 0x8d, 0x9f, 0xd3, 0xec, 0x2e, 0x14, 0x2c,
	0x1c, 0x3c, 0x75, 0xff, 0x33, 0xff, 0xf9, 0xcd, 0xec, 0x74, 0xa1, 0x91, 0x0b, 0x9a, 0x8a, 0x5e,
	0x92, 0xc6, 0x59, 0x8c, 0x10, 0xa5, 0xc9, 0xe4, 0x2d, 0xcf, 0xd8, 0x5c, 0x10, 0x07, 0xec, 0xb1,
	0xa0, 0xa9, 0xef, 0x61, 0x13, 0x2a, 0xbe, 0xe7, 0x18, 0x1d, 0xa3, 0x6b, 0x8e, 0x2a, 0xbe, 0x47,
	0xce, 0xa0, 0x2e, 0x33, 0x8f, 0x8b, 0x90, 0xcd, 0xb1, 0x0d, 0x96, 0x3a, 0xa8, 0x7c, 0x7d, 0xa4,
	0x05, 0x39, 0x81, 0x3f, 0xd2, 0xf2, 0x1a, 0x2e, 0x28, 0x22, 0x98, 0xf2, 0xbb, 0x36, 0xa8, 0x33,
	0x39, 0x85, 0x5a, 0x40, 0x85, 0x60, 0x31, 0x97, 0x80, 0x65, 0x38, 0xcf, 0x37, 0x79, 0x2d, 0xc8,
	0x97, 0x01, 0xa6, 0x24, 0x94, 0x9b, 0xa3, 0xab, 0xc9, 0x5c, 0x12, 0x2b, 0xaa, 0xa2, 0xd0, 0xdb,
	0x59, 0xaa, 0x3b, 0xb3, 0x14, 0xfd, 0xcd, 0x6d, 0x7f, 0x49, 0x19, 0x86, 0x42, 0x7c, 0xc4, 0xe9,
	0xd4, 0xb1, 0x34, 0x65, 0xa3, 0xf1, 0x18, 0xec, 0x20, 0x0b, 0xb3, 0x5c, 0x38, 0xb6, 0xca, 0xac,
	0x95, 0xa4, 0x0f, 0x67, 0x31, 0xa7, 0x4e, 0x4d, 0xd3, 0x95, 0x20, 0x57, 0x60, 0xc9, 0xfe, 0x02,
	0xcf, 0xc1, 0x52, 0xab, 0x74, 0x8c, 0x4e, 0xb5, 0xdb, 0xe8, 0xb7, 0x7a, 0xdb, 0x5d, 0xf6, 0xa4,
	0x63, 0xa4, 0xd3, 0xa4, 0x26, 0x87, 0x4c, 0xb2, 0x55, 0xff, 0xb3, 0x0a, 0x7f, 0x55, 0x69, 0x40,
	0xd3, 0x25, 0x9b, 0x50, 0x1c, 0x40, 0xe3, 0x99, 0x66, 0x32, 0x74, 0xbf, 0xf2, 0x3d, 0xc4, 0x32,
	0xc1, 0xf7, 0xdc, 0x3d, 0x2a, 0xde, 0x42, 0xb3, 0x28, 0xd2, 0xf7, 0x3d, 0x2a, 0x7b, 0x54, 0xf8,
	0x40, 0xe9, 0x25, 0xd8, 0x01, 0x8b, 0xf8, 0x38, 0xc1, 0xbd, 0x9c, 0xfb, 0x6f, 0x37, 0xa2, 0xe6,
	0xc5, 0x0b, 0xb0, 0x5e, 0xe2, 0x88, 0xf1, 0x03, 0xee, 0x7d, 0xf6, 0x35, 0xc0, 0xc3, 0x2c, 0xe4,
	0x11, 0x55, 0xea, 0x57, 0xfc, 0x1b, 0xa8, 0x3f, 0x31, 0x3e, 0xd5, 0xdb, 0x6c, 0x97, 0x2b, 0xe4,
	0x6f, 0xfb, 0x59, 0xa5, 0x8d, 0x77, 0xd0, 0x2a, 0xee, 0xbf, 0x79, 0x52, 0xff, 0x77, 0x6d, 0xeb,
	0xa0, 0x7b, 0x60, 0x9d, 0xef, 0xb6, 0x7a, 0xf6, 0x83, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa8,
	0xc3, 0xb0, 0x0e, 0x05, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UsersServiceClient is the client API for UsersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UsersServiceClient interface {
	GetUserByID(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*User, error)
	GetUserByEmail(ctx context.Context, in *UserEmail, opts ...grpc.CallOption) (*User, error)
	SignUp(ctx context.Context, in *User, opts ...grpc.CallOption) (*Empty, error)
	Login(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
	ChangeUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*Empty, error)
	FindUsers(ctx context.Context, in *UserName, opts ...grpc.CallOption) (*Users, error)
	GetUserBySession(ctx context.Context, in *Session, opts ...grpc.CallOption) (*UserID, error)
}

type usersServiceClient struct {
	cc *grpc.ClientConn
}

func NewUsersServiceClient(cc *grpc.ClientConn) UsersServiceClient {
	return &usersServiceClient{cc}
}

func (c *usersServiceClient) GetUserByID(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/grpc_utils.UsersService/GetUserByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) GetUserByEmail(ctx context.Context, in *UserEmail, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/grpc_utils.UsersService/GetUserByEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) SignUp(ctx context.Context, in *User, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/grpc_utils.UsersService/SignUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) Login(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/grpc_utils.UsersService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) ChangeUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/grpc_utils.UsersService/ChangeUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) FindUsers(ctx context.Context, in *UserName, opts ...grpc.CallOption) (*Users, error) {
	out := new(Users)
	err := c.cc.Invoke(ctx, "/grpc_utils.UsersService/FindUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) GetUserBySession(ctx context.Context, in *Session, opts ...grpc.CallOption) (*UserID, error) {
	out := new(UserID)
	err := c.cc.Invoke(ctx, "/grpc_utils.UsersService/GetUserBySession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServiceServer is the server API for UsersService service.
type UsersServiceServer interface {
	GetUserByID(context.Context, *UserID) (*User, error)
	GetUserByEmail(context.Context, *UserEmail) (*User, error)
	SignUp(context.Context, *User) (*Empty, error)
	Login(context.Context, *User) (*User, error)
	ChangeUser(context.Context, *User) (*Empty, error)
	FindUsers(context.Context, *UserName) (*Users, error)
	GetUserBySession(context.Context, *Session) (*UserID, error)
}

// UnimplementedUsersServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUsersServiceServer struct {
}

func (*UnimplementedUsersServiceServer) GetUserByID(ctx context.Context, req *UserID) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByID not implemented")
}
func (*UnimplementedUsersServiceServer) GetUserByEmail(ctx context.Context, req *UserEmail) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByEmail not implemented")
}
func (*UnimplementedUsersServiceServer) SignUp(ctx context.Context, req *User) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUp not implemented")
}
func (*UnimplementedUsersServiceServer) Login(ctx context.Context, req *User) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedUsersServiceServer) ChangeUser(ctx context.Context, req *User) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeUser not implemented")
}
func (*UnimplementedUsersServiceServer) FindUsers(ctx context.Context, req *UserName) (*Users, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindUsers not implemented")
}
func (*UnimplementedUsersServiceServer) GetUserBySession(ctx context.Context, req *Session) (*UserID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserBySession not implemented")
}

func RegisterUsersServiceServer(s *grpc.Server, srv UsersServiceServer) {
	s.RegisterService(&_UsersService_serviceDesc, srv)
}

func _UsersService_GetUserByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).GetUserByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_utils.UsersService/GetUserByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).GetUserByID(ctx, req.(*UserID))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_GetUserByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserEmail)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).GetUserByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_utils.UsersService/GetUserByEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).GetUserByEmail(ctx, req.(*UserEmail))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_utils.UsersService/SignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).SignUp(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_utils.UsersService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).Login(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_ChangeUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).ChangeUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_utils.UsersService/ChangeUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).ChangeUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_FindUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).FindUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_utils.UsersService/FindUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).FindUsers(ctx, req.(*UserName))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_GetUserBySession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Session)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).GetUserBySession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_utils.UsersService/GetUserBySession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).GetUserBySession(ctx, req.(*Session))
	}
	return interceptor(ctx, in, info, handler)
}

var _UsersService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc_utils.UsersService",
	HandlerType: (*UsersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserByID",
			Handler:    _UsersService_GetUserByID_Handler,
		},
		{
			MethodName: "GetUserByEmail",
			Handler:    _UsersService_GetUserByEmail_Handler,
		},
		{
			MethodName: "SignUp",
			Handler:    _UsersService_SignUp_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _UsersService_Login_Handler,
		},
		{
			MethodName: "ChangeUser",
			Handler:    _UsersService_ChangeUser_Handler,
		},
		{
			MethodName: "FindUsers",
			Handler:    _UsersService_FindUsers_Handler,
		},
		{
			MethodName: "GetUserBySession",
			Handler:    _UsersService_GetUserBySession_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "users.proto",
}
