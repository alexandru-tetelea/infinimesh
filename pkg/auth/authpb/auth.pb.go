// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/auth/authpb/auth.proto

package authpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Action int32

const (
	Action_NONE  Action = 0
	Action_READ  Action = 1
	Action_WRITE Action = 2
)

var Action_name = map[int32]string{
	0: "NONE",
	1: "READ",
	2: "WRITE",
}

var Action_value = map[string]int32{
	"NONE":  0,
	"READ":  1,
	"WRITE": 2,
}

func (x Action) String() string {
	return proto.EnumName(Action_name, int32(x))
}

func (Action) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_31381d2bd3d06ae7, []int{0}
}

type IsAuthorizedRequest struct {
	Object               string   `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
	Subject              string   `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
	Action               string   `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsAuthorizedRequest) Reset()         { *m = IsAuthorizedRequest{} }
func (m *IsAuthorizedRequest) String() string { return proto.CompactTextString(m) }
func (*IsAuthorizedRequest) ProtoMessage()    {}
func (*IsAuthorizedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_31381d2bd3d06ae7, []int{0}
}

func (m *IsAuthorizedRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsAuthorizedRequest.Unmarshal(m, b)
}
func (m *IsAuthorizedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsAuthorizedRequest.Marshal(b, m, deterministic)
}
func (m *IsAuthorizedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsAuthorizedRequest.Merge(m, src)
}
func (m *IsAuthorizedRequest) XXX_Size() int {
	return xxx_messageInfo_IsAuthorizedRequest.Size(m)
}
func (m *IsAuthorizedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IsAuthorizedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IsAuthorizedRequest proto.InternalMessageInfo

func (m *IsAuthorizedRequest) GetObject() string {
	if m != nil {
		return m.Object
	}
	return ""
}

func (m *IsAuthorizedRequest) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *IsAuthorizedRequest) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

type IsAuthorizedResponse struct {
	Decision             *wrappers.BoolValue `protobuf:"bytes,1,opt,name=decision,proto3" json:"decision,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *IsAuthorizedResponse) Reset()         { *m = IsAuthorizedResponse{} }
func (m *IsAuthorizedResponse) String() string { return proto.CompactTextString(m) }
func (*IsAuthorizedResponse) ProtoMessage()    {}
func (*IsAuthorizedResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_31381d2bd3d06ae7, []int{1}
}

func (m *IsAuthorizedResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsAuthorizedResponse.Unmarshal(m, b)
}
func (m *IsAuthorizedResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsAuthorizedResponse.Marshal(b, m, deterministic)
}
func (m *IsAuthorizedResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsAuthorizedResponse.Merge(m, src)
}
func (m *IsAuthorizedResponse) XXX_Size() int {
	return xxx_messageInfo_IsAuthorizedResponse.Size(m)
}
func (m *IsAuthorizedResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IsAuthorizedResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IsAuthorizedResponse proto.InternalMessageInfo

func (m *IsAuthorizedResponse) GetDecision() *wrappers.BoolValue {
	if m != nil {
		return m.Decision
	}
	return nil
}

type LoginRequest struct {
	ClientID             string   `protobuf:"bytes,1,opt,name=clientID,proto3" json:"clientID,omitempty"`
	ClientSecret         string   `protobuf:"bytes,2,opt,name=clientSecret,proto3" json:"clientSecret,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_31381d2bd3d06ae7, []int{2}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetClientID() string {
	if m != nil {
		return m.ClientID
	}
	return ""
}

func (m *LoginRequest) GetClientSecret() string {
	if m != nil {
		return m.ClientSecret
	}
	return ""
}

type LoginResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_31381d2bd3d06ae7, []int{3}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

type CreateUserRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_31381d2bd3d06ae7, []int{4}
}

func (m *CreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRequest.Unmarshal(m, b)
}
func (m *CreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRequest.Marshal(b, m, deterministic)
}
func (m *CreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRequest.Merge(m, src)
}
func (m *CreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserRequest.Size(m)
}
func (m *CreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserRequest proto.InternalMessageInfo

func (m *CreateUserRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CreateUserResponse struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserResponse) Reset()         { *m = CreateUserResponse{} }
func (m *CreateUserResponse) String() string { return proto.CompactTextString(m) }
func (*CreateUserResponse) ProtoMessage()    {}
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_31381d2bd3d06ae7, []int{5}
}

func (m *CreateUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserResponse.Unmarshal(m, b)
}
func (m *CreateUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserResponse.Marshal(b, m, deterministic)
}
func (m *CreateUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserResponse.Merge(m, src)
}
func (m *CreateUserResponse) XXX_Size() int {
	return xxx_messageInfo_CreateUserResponse.Size(m)
}
func (m *CreateUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserResponse proto.InternalMessageInfo

func (m *CreateUserResponse) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

type AuthorizeRequest struct {
	EntityUid            string   `protobuf:"bytes,1,opt,name=entity_uid,json=entityUid,proto3" json:"entity_uid,omitempty"`
	ResourceUid          string   `protobuf:"bytes,2,opt,name=resource_uid,json=resourceUid,proto3" json:"resource_uid,omitempty"`
	Action               string   `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthorizeRequest) Reset()         { *m = AuthorizeRequest{} }
func (m *AuthorizeRequest) String() string { return proto.CompactTextString(m) }
func (*AuthorizeRequest) ProtoMessage()    {}
func (*AuthorizeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_31381d2bd3d06ae7, []int{6}
}

func (m *AuthorizeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthorizeRequest.Unmarshal(m, b)
}
func (m *AuthorizeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthorizeRequest.Marshal(b, m, deterministic)
}
func (m *AuthorizeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthorizeRequest.Merge(m, src)
}
func (m *AuthorizeRequest) XXX_Size() int {
	return xxx_messageInfo_AuthorizeRequest.Size(m)
}
func (m *AuthorizeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthorizeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthorizeRequest proto.InternalMessageInfo

func (m *AuthorizeRequest) GetEntityUid() string {
	if m != nil {
		return m.EntityUid
	}
	return ""
}

func (m *AuthorizeRequest) GetResourceUid() string {
	if m != nil {
		return m.ResourceUid
	}
	return ""
}

func (m *AuthorizeRequest) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

type AuthorizeResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthorizeResponse) Reset()         { *m = AuthorizeResponse{} }
func (m *AuthorizeResponse) String() string { return proto.CompactTextString(m) }
func (*AuthorizeResponse) ProtoMessage()    {}
func (*AuthorizeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_31381d2bd3d06ae7, []int{7}
}

func (m *AuthorizeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthorizeResponse.Unmarshal(m, b)
}
func (m *AuthorizeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthorizeResponse.Marshal(b, m, deterministic)
}
func (m *AuthorizeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthorizeResponse.Merge(m, src)
}
func (m *AuthorizeResponse) XXX_Size() int {
	return xxx_messageInfo_AuthorizeResponse.Size(m)
}
func (m *AuthorizeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthorizeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthorizeResponse proto.InternalMessageInfo

type SetCredentialsRequest struct {
	Password             string   `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetCredentialsRequest) Reset()         { *m = SetCredentialsRequest{} }
func (m *SetCredentialsRequest) String() string { return proto.CompactTextString(m) }
func (*SetCredentialsRequest) ProtoMessage()    {}
func (*SetCredentialsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_31381d2bd3d06ae7, []int{8}
}

func (m *SetCredentialsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetCredentialsRequest.Unmarshal(m, b)
}
func (m *SetCredentialsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetCredentialsRequest.Marshal(b, m, deterministic)
}
func (m *SetCredentialsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetCredentialsRequest.Merge(m, src)
}
func (m *SetCredentialsRequest) XXX_Size() int {
	return xxx_messageInfo_SetCredentialsRequest.Size(m)
}
func (m *SetCredentialsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetCredentialsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetCredentialsRequest proto.InternalMessageInfo

func (m *SetCredentialsRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type SetCredentialsResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetCredentialsResponse) Reset()         { *m = SetCredentialsResponse{} }
func (m *SetCredentialsResponse) String() string { return proto.CompactTextString(m) }
func (*SetCredentialsResponse) ProtoMessage()    {}
func (*SetCredentialsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_31381d2bd3d06ae7, []int{9}
}

func (m *SetCredentialsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetCredentialsResponse.Unmarshal(m, b)
}
func (m *SetCredentialsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetCredentialsResponse.Marshal(b, m, deterministic)
}
func (m *SetCredentialsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetCredentialsResponse.Merge(m, src)
}
func (m *SetCredentialsResponse) XXX_Size() int {
	return xxx_messageInfo_SetCredentialsResponse.Size(m)
}
func (m *SetCredentialsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetCredentialsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetCredentialsResponse proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("infinimesh.auth.Action", Action_name, Action_value)
	proto.RegisterType((*IsAuthorizedRequest)(nil), "infinimesh.auth.IsAuthorizedRequest")
	proto.RegisterType((*IsAuthorizedResponse)(nil), "infinimesh.auth.IsAuthorizedResponse")
	proto.RegisterType((*LoginRequest)(nil), "infinimesh.auth.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "infinimesh.auth.LoginResponse")
	proto.RegisterType((*CreateUserRequest)(nil), "infinimesh.auth.CreateUserRequest")
	proto.RegisterType((*CreateUserResponse)(nil), "infinimesh.auth.CreateUserResponse")
	proto.RegisterType((*AuthorizeRequest)(nil), "infinimesh.auth.AuthorizeRequest")
	proto.RegisterType((*AuthorizeResponse)(nil), "infinimesh.auth.AuthorizeResponse")
	proto.RegisterType((*SetCredentialsRequest)(nil), "infinimesh.auth.SetCredentialsRequest")
	proto.RegisterType((*SetCredentialsResponse)(nil), "infinimesh.auth.SetCredentialsResponse")
}

func init() { proto.RegisterFile("pkg/auth/authpb/auth.proto", fileDescriptor_31381d2bd3d06ae7) }

var fileDescriptor_31381d2bd3d06ae7 = []byte{
	// 504 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xdf, 0x6f, 0xd3, 0x30,
	0x10, 0xa6, 0x6b, 0x57, 0xda, 0x5b, 0x61, 0x99, 0x07, 0x53, 0x15, 0x69, 0x13, 0x33, 0xb0, 0x22,
	0x1e, 0x52, 0x69, 0x93, 0x78, 0xef, 0xb6, 0x22, 0x55, 0x42, 0x45, 0xca, 0x28, 0x48, 0xf0, 0x30,
	0xa5, 0xc9, 0xad, 0x35, 0x64, 0x71, 0xb0, 0x1d, 0x4d, 0xf0, 0x9f, 0xf1, 0xdf, 0xa1, 0xd8, 0x4e,
	0x96, 0x36, 0xfb, 0xf1, 0xd2, 0xfa, 0x7c, 0xdf, 0x77, 0xdf, 0x77, 0xe7, 0x6b, 0xc1, 0x4d, 0x7f,
	0x2d, 0x86, 0x41, 0xa6, 0x96, 0xfa, 0x23, 0x9d, 0xeb, 0x2f, 0x2f, 0x15, 0x5c, 0x71, 0xb2, 0xcd,
	0x92, 0x2b, 0x96, 0xb0, 0x6b, 0x94, 0x4b, 0x2f, 0xbf, 0x76, 0x0f, 0x16, 0x9c, 0x2f, 0x62, 0x1c,
	0xea, 0xf4, 0x3c, 0xbb, 0x1a, 0xde, 0x88, 0x20, 0x4d, 0x51, 0x48, 0x43, 0xa0, 0x97, 0xb0, 0x3b,
	0x91, 0xa3, 0x4c, 0x2d, 0xb9, 0x60, 0x7f, 0x31, 0xf2, 0xf1, 0x77, 0x86, 0x52, 0x91, 0x3d, 0x68,
	0xf3, 0xf9, 0x4f, 0x0c, 0x55, 0xbf, 0xf1, 0xaa, 0xf1, 0xae, 0xeb, 0xdb, 0x88, 0xf4, 0xe1, 0xa9,
	0xcc, 0x4c, 0x62, 0x43, 0x27, 0x8a, 0x30, 0x67, 0x04, 0xa1, 0x62, 0x3c, 0xe9, 0x37, 0x0d, 0xc3,
	0x44, 0x74, 0x0a, 0x2f, 0x56, 0x05, 0x64, 0xca, 0x13, 0x89, 0xe4, 0x03, 0x74, 0x22, 0x0c, 0x99,
	0xcc, 0x19, 0xb9, 0xc6, 0xd6, 0xb1, 0xeb, 0x19, 0xaf, 0x5e, 0xe1, 0xd5, 0x3b, 0xe5, 0x3c, 0xfe,
	0x1a, 0xc4, 0x19, 0xfa, 0x25, 0x96, 0x4e, 0xa1, 0xf7, 0x89, 0x2f, 0x58, 0x52, 0x38, 0x75, 0xa1,
	0x13, 0xc6, 0x0c, 0x13, 0x35, 0x39, 0xb7, 0x5e, 0xcb, 0x98, 0x50, 0xe8, 0x99, 0xf3, 0x05, 0x86,
	0x02, 0x0b, 0xcb, 0x2b, 0x77, 0x74, 0x1b, 0x9e, 0xd9, 0x7a, 0xc6, 0x18, 0x1d, 0xc0, 0xce, 0x99,
	0xc0, 0x40, 0xe1, 0x4c, 0xa2, 0x28, 0x54, 0x08, 0xb4, 0x92, 0xe0, 0x1a, 0xad, 0x82, 0x3e, 0xd3,
	0x23, 0x20, 0x55, 0xa0, 0xed, 0xcb, 0x81, 0x66, 0xc6, 0x22, 0x0b, 0xcc, 0x8f, 0x34, 0x06, 0xa7,
	0xec, 0xbf, 0xa8, 0xb7, 0x0f, 0x80, 0x89, 0x62, 0xea, 0xcf, 0xe5, 0x2d, 0xb8, 0x6b, 0x6e, 0x66,
	0x2c, 0x22, 0x87, 0xd0, 0x13, 0x28, 0x79, 0x26, 0x42, 0xd4, 0x00, 0x63, 0x7c, 0xab, 0xb8, 0xcb,
	0x21, 0xf7, 0xcd, 0x7b, 0x17, 0x76, 0x2a, 0x6a, 0xb6, 0xa7, 0x13, 0x78, 0x79, 0x81, 0xea, 0x4c,
	0x60, 0x94, 0x6b, 0x04, 0xb1, 0xac, 0x4c, 0x2f, 0x0d, 0xa4, 0xbc, 0xe1, 0xa2, 0x70, 0x51, 0xc6,
	0xb4, 0x0f, 0x7b, 0xeb, 0x24, 0x53, 0xee, 0xfd, 0x00, 0xda, 0x23, 0xad, 0x46, 0x3a, 0xd0, 0x9a,
	0x7e, 0x9e, 0x8e, 0x9d, 0x27, 0xf9, 0xc9, 0x1f, 0x8f, 0xce, 0x9d, 0x06, 0xe9, 0xc2, 0xe6, 0x37,
	0x7f, 0xf2, 0x65, 0xec, 0x6c, 0x1c, 0xff, 0x6b, 0x42, 0x2b, 0x77, 0x43, 0x7e, 0x40, 0xaf, 0xba,
	0x05, 0xe4, 0x8d, 0xb7, 0xb6, 0xa8, 0xde, 0x1d, 0x5b, 0xe8, 0xbe, 0x7d, 0x04, 0x65, 0x47, 0x3e,
	0x03, 0xb8, 0x7d, 0x08, 0x42, 0x6b, 0xa4, 0xda, 0x73, 0xba, 0xaf, 0x1f, 0xc4, 0xd8, 0xb2, 0x01,
	0x3c, 0x5f, 0xed, 0x9f, 0x1c, 0xd5, 0x68, 0x77, 0x4e, 0xd5, 0x1d, 0x3c, 0x8a, 0xb3, 0x12, 0x1f,
	0x61, 0x53, 0x2f, 0x1f, 0xd9, 0xaf, 0x31, 0xaa, 0x4b, 0xee, 0x1e, 0xdc, 0x97, 0xb6, 0x75, 0x7c,
	0xe8, 0x96, 0x73, 0x21, 0x87, 0x35, 0xf0, 0xfa, 0xfa, 0xb9, 0xf4, 0x21, 0x88, 0xa9, 0x79, 0xda,
	0xf9, 0xde, 0x36, 0xff, 0x2f, 0xf3, 0xb6, 0xfe, 0x41, 0x9e, 0xfc, 0x0f, 0x00, 0x00, 0xff, 0xff,
	0x76, 0x78, 0x08, 0x26, 0x79, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthClient interface {
	IsAuthorized(ctx context.Context, in *IsAuthorizedRequest, opts ...grpc.CallOption) (*IsAuthorizedResponse, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	SetCredentials(ctx context.Context, in *SetCredentialsRequest, opts ...grpc.CallOption) (*SetCredentialsResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Authorize(ctx context.Context, in *AuthorizeRequest, opts ...grpc.CallOption) (*AuthorizeResponse, error)
}

type authClient struct {
	cc *grpc.ClientConn
}

func NewAuthClient(cc *grpc.ClientConn) AuthClient {
	return &authClient{cc}
}

func (c *authClient) IsAuthorized(ctx context.Context, in *IsAuthorizedRequest, opts ...grpc.CallOption) (*IsAuthorizedResponse, error) {
	out := new(IsAuthorizedResponse)
	err := c.cc.Invoke(ctx, "/infinimesh.auth.Auth/IsAuthorized", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/infinimesh.auth.Auth/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) SetCredentials(ctx context.Context, in *SetCredentialsRequest, opts ...grpc.CallOption) (*SetCredentialsResponse, error) {
	out := new(SetCredentialsResponse)
	err := c.cc.Invoke(ctx, "/infinimesh.auth.Auth/SetCredentials", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/infinimesh.auth.Auth/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Authorize(ctx context.Context, in *AuthorizeRequest, opts ...grpc.CallOption) (*AuthorizeResponse, error) {
	out := new(AuthorizeResponse)
	err := c.cc.Invoke(ctx, "/infinimesh.auth.Auth/Authorize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
type AuthServer interface {
	IsAuthorized(context.Context, *IsAuthorizedRequest) (*IsAuthorizedResponse, error)
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	SetCredentials(context.Context, *SetCredentialsRequest) (*SetCredentialsResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Authorize(context.Context, *AuthorizeRequest) (*AuthorizeResponse, error)
}

func RegisterAuthServer(s *grpc.Server, srv AuthServer) {
	s.RegisterService(&_Auth_serviceDesc, srv)
}

func _Auth_IsAuthorized_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsAuthorizedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).IsAuthorized(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infinimesh.auth.Auth/IsAuthorized",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).IsAuthorized(ctx, req.(*IsAuthorizedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infinimesh.auth.Auth/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_SetCredentials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetCredentialsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).SetCredentials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infinimesh.auth.Auth/SetCredentials",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).SetCredentials(ctx, req.(*SetCredentialsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infinimesh.auth.Auth/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_Authorize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Authorize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infinimesh.auth.Auth/Authorize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Authorize(ctx, req.(*AuthorizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Auth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "infinimesh.auth.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsAuthorized",
			Handler:    _Auth_IsAuthorized_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _Auth_CreateUser_Handler,
		},
		{
			MethodName: "SetCredentials",
			Handler:    _Auth_SetCredentials_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Auth_Login_Handler,
		},
		{
			MethodName: "Authorize",
			Handler:    _Auth_Authorize_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/auth/authpb/auth.proto",
}
