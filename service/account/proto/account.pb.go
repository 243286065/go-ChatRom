// Code generated by protoc-gen-go. DO NOT EDIT.
// source: account.proto

package go_micro_service_account

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type ReqSignup struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqSignup) Reset()         { *m = ReqSignup{} }
func (m *ReqSignup) String() string { return proto.CompactTextString(m) }
func (*ReqSignup) ProtoMessage()    {}
func (*ReqSignup) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{0}
}

func (m *ReqSignup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqSignup.Unmarshal(m, b)
}
func (m *ReqSignup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqSignup.Marshal(b, m, deterministic)
}
func (m *ReqSignup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqSignup.Merge(m, src)
}
func (m *ReqSignup) XXX_Size() int {
	return xxx_messageInfo_ReqSignup.Size(m)
}
func (m *ReqSignup) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqSignup.DiscardUnknown(m)
}

var xxx_messageInfo_ReqSignup proto.InternalMessageInfo

func (m *ReqSignup) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *ReqSignup) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type RespSignup struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RespSignup) Reset()         { *m = RespSignup{} }
func (m *RespSignup) String() string { return proto.CompactTextString(m) }
func (*RespSignup) ProtoMessage()    {}
func (*RespSignup) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{1}
}

func (m *RespSignup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RespSignup.Unmarshal(m, b)
}
func (m *RespSignup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RespSignup.Marshal(b, m, deterministic)
}
func (m *RespSignup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RespSignup.Merge(m, src)
}
func (m *RespSignup) XXX_Size() int {
	return xxx_messageInfo_RespSignup.Size(m)
}
func (m *RespSignup) XXX_DiscardUnknown() {
	xxx_messageInfo_RespSignup.DiscardUnknown(m)
}

var xxx_messageInfo_RespSignup proto.InternalMessageInfo

func (m *RespSignup) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RespSignup) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ReqSignin struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqSignin) Reset()         { *m = ReqSignin{} }
func (m *ReqSignin) String() string { return proto.CompactTextString(m) }
func (*ReqSignin) ProtoMessage()    {}
func (*ReqSignin) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{2}
}

func (m *ReqSignin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqSignin.Unmarshal(m, b)
}
func (m *ReqSignin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqSignin.Marshal(b, m, deterministic)
}
func (m *ReqSignin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqSignin.Merge(m, src)
}
func (m *ReqSignin) XXX_Size() int {
	return xxx_messageInfo_ReqSignin.Size(m)
}
func (m *ReqSignin) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqSignin.DiscardUnknown(m)
}

var xxx_messageInfo_ReqSignin proto.InternalMessageInfo

func (m *ReqSignin) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *ReqSignin) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type RespSignin struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	Message              string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RespSignin) Reset()         { *m = RespSignin{} }
func (m *RespSignin) String() string { return proto.CompactTextString(m) }
func (*RespSignin) ProtoMessage()    {}
func (*RespSignin) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{3}
}

func (m *RespSignin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RespSignin.Unmarshal(m, b)
}
func (m *RespSignin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RespSignin.Marshal(b, m, deterministic)
}
func (m *RespSignin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RespSignin.Merge(m, src)
}
func (m *RespSignin) XXX_Size() int {
	return xxx_messageInfo_RespSignin.Size(m)
}
func (m *RespSignin) XXX_DiscardUnknown() {
	xxx_messageInfo_RespSignin.DiscardUnknown(m)
}

var xxx_messageInfo_RespSignin proto.InternalMessageInfo

func (m *RespSignin) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RespSignin) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *RespSignin) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ReqUpdateToken struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqUpdateToken) Reset()         { *m = ReqUpdateToken{} }
func (m *ReqUpdateToken) String() string { return proto.CompactTextString(m) }
func (*ReqUpdateToken) ProtoMessage()    {}
func (*ReqUpdateToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{4}
}

func (m *ReqUpdateToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqUpdateToken.Unmarshal(m, b)
}
func (m *ReqUpdateToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqUpdateToken.Marshal(b, m, deterministic)
}
func (m *ReqUpdateToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqUpdateToken.Merge(m, src)
}
func (m *ReqUpdateToken) XXX_Size() int {
	return xxx_messageInfo_ReqUpdateToken.Size(m)
}
func (m *ReqUpdateToken) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqUpdateToken.DiscardUnknown(m)
}

var xxx_messageInfo_ReqUpdateToken proto.InternalMessageInfo

func (m *ReqUpdateToken) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *ReqUpdateToken) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type RespUpdateToken struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RespUpdateToken) Reset()         { *m = RespUpdateToken{} }
func (m *RespUpdateToken) String() string { return proto.CompactTextString(m) }
func (*RespUpdateToken) ProtoMessage()    {}
func (*RespUpdateToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{5}
}

func (m *RespUpdateToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RespUpdateToken.Unmarshal(m, b)
}
func (m *RespUpdateToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RespUpdateToken.Marshal(b, m, deterministic)
}
func (m *RespUpdateToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RespUpdateToken.Merge(m, src)
}
func (m *RespUpdateToken) XXX_Size() int {
	return xxx_messageInfo_RespUpdateToken.Size(m)
}
func (m *RespUpdateToken) XXX_DiscardUnknown() {
	xxx_messageInfo_RespUpdateToken.DiscardUnknown(m)
}

var xxx_messageInfo_RespUpdateToken proto.InternalMessageInfo

func (m *RespUpdateToken) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

type ReqGetToken struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqGetToken) Reset()         { *m = ReqGetToken{} }
func (m *ReqGetToken) String() string { return proto.CompactTextString(m) }
func (*ReqGetToken) ProtoMessage()    {}
func (*ReqGetToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{6}
}

func (m *ReqGetToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqGetToken.Unmarshal(m, b)
}
func (m *ReqGetToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqGetToken.Marshal(b, m, deterministic)
}
func (m *ReqGetToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqGetToken.Merge(m, src)
}
func (m *ReqGetToken) XXX_Size() int {
	return xxx_messageInfo_ReqGetToken.Size(m)
}
func (m *ReqGetToken) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqGetToken.DiscardUnknown(m)
}

var xxx_messageInfo_ReqGetToken proto.InternalMessageInfo

func (m *ReqGetToken) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type RespGetToken struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RespGetToken) Reset()         { *m = RespGetToken{} }
func (m *RespGetToken) String() string { return proto.CompactTextString(m) }
func (*RespGetToken) ProtoMessage()    {}
func (*RespGetToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{7}
}

func (m *RespGetToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RespGetToken.Unmarshal(m, b)
}
func (m *RespGetToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RespGetToken.Marshal(b, m, deterministic)
}
func (m *RespGetToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RespGetToken.Merge(m, src)
}
func (m *RespGetToken) XXX_Size() int {
	return xxx_messageInfo_RespGetToken.Size(m)
}
func (m *RespGetToken) XXX_DiscardUnknown() {
	xxx_messageInfo_RespGetToken.DiscardUnknown(m)
}

var xxx_messageInfo_RespGetToken proto.InternalMessageInfo

func (m *RespGetToken) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RespGetToken) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*ReqSignup)(nil), "go.micro.service.account.ReqSignup")
	proto.RegisterType((*RespSignup)(nil), "go.micro.service.account.RespSignup")
	proto.RegisterType((*ReqSignin)(nil), "go.micro.service.account.ReqSignin")
	proto.RegisterType((*RespSignin)(nil), "go.micro.service.account.RespSignin")
	proto.RegisterType((*ReqUpdateToken)(nil), "go.micro.service.account.ReqUpdateToken")
	proto.RegisterType((*RespUpdateToken)(nil), "go.micro.service.account.RespUpdateToken")
	proto.RegisterType((*ReqGetToken)(nil), "go.micro.service.account.ReqGetToken")
	proto.RegisterType((*RespGetToken)(nil), "go.micro.service.account.RespGetToken")
}

func init() { proto.RegisterFile("account.proto", fileDescriptor_8e28828dcb8d24f0) }

var fileDescriptor_8e28828dcb8d24f0 = []byte{
	// 317 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xbf, 0x4f, 0xfb, 0x30,
	0x10, 0xc5, 0xd3, 0x6f, 0xbf, 0x2d, 0xed, 0x15, 0x8a, 0x64, 0x31, 0x44, 0x99, 0x90, 0xa1, 0xa8,
	0x5d, 0x3c, 0xc0, 0x82, 0xd8, 0x80, 0x81, 0x15, 0xb9, 0x74, 0x62, 0x0a, 0xc9, 0xa9, 0xb2, 0x50,
	0xec, 0x24, 0x4e, 0x60, 0xe6, 0x3f, 0x47, 0xf9, 0xe1, 0x60, 0xa4, 0x60, 0x22, 0xb6, 0x3c, 0xdd,
	0xf3, 0xe7, 0x5e, 0x7c, 0x67, 0x38, 0x0a, 0xa3, 0x48, 0x95, 0xb2, 0x60, 0x69, 0xae, 0x0a, 0x45,
	0xfc, 0xbd, 0x62, 0x89, 0x88, 0x72, 0xc5, 0x34, 0xe6, 0x6f, 0x22, 0x42, 0xd6, 0xd6, 0xe9, 0x3d,
	0xcc, 0x39, 0x66, 0x5b, 0xb1, 0x97, 0x65, 0x4a, 0x02, 0x98, 0x95, 0x1a, 0x73, 0x19, 0x26, 0xe8,
	0x8f, 0x4e, 0x47, 0xeb, 0x39, 0xef, 0x74, 0x55, 0x4b, 0x43, 0xad, 0xdf, 0x55, 0x1e, 0xfb, 0xff,
	0x9a, 0x9a, 0xd1, 0xf4, 0x06, 0x80, 0xa3, 0x4e, 0x5b, 0x0a, 0x81, 0xff, 0x91, 0x8a, 0x1b, 0xc2,
	0x84, 0xd7, 0xdf, 0xc4, 0x87, 0x83, 0x04, 0xb5, 0x0e, 0xf7, 0xd8, 0x1e, 0x36, 0xd2, 0x0a, 0x20,
	0xe4, 0x9f, 0x03, 0x3c, 0x7e, 0x05, 0x10, 0xb2, 0x37, 0xc0, 0x09, 0x4c, 0x0a, 0xf5, 0x8a, 0xb2,
	0x3d, 0xda, 0x08, 0x3b, 0xd6, 0xf8, 0x7b, 0xac, 0x3b, 0x58, 0x72, 0xcc, 0x76, 0x69, 0x1c, 0x16,
	0xf8, 0x54, 0x7b, 0x5d, 0xd9, 0x7a, 0xe9, 0x74, 0x05, 0xc7, 0x55, 0x2a, 0x1b, 0xd2, 0x13, 0x8d,
	0x6e, 0x60, 0xc1, 0x31, 0x7b, 0xc0, 0xe2, 0xd7, 0x3e, 0xf4, 0x1a, 0x0e, 0x2b, 0x62, 0xe7, 0x1d,
	0xfc, 0xa7, 0x97, 0x1f, 0x63, 0x58, 0xde, 0x36, 0x33, 0xdf, 0x36, 0x2b, 0x40, 0x76, 0x30, 0x6d,
	0x27, 0x76, 0xc6, 0x7e, 0xda, 0x0f, 0xd6, 0x2d, 0x47, 0x70, 0xee, 0x32, 0x99, 0xe1, 0x53, 0xcf,
	0x60, 0x85, 0x1c, 0x80, 0x15, 0x72, 0x08, 0x56, 0x48, 0xea, 0x91, 0x18, 0x16, 0xf6, 0x45, 0xae,
	0x9d, 0x6c, 0xcb, 0x19, 0x6c, 0xdc, 0x0d, 0x2c, 0x2b, 0xf5, 0xc8, 0x33, 0xcc, 0xba, 0xcb, 0x5d,
	0x39, 0x5b, 0x18, 0x5b, 0x70, 0xe1, 0xe6, 0x1b, 0x1f, 0xf5, 0x5e, 0xa6, 0xf5, 0x63, 0xbc, 0xfa,
	0x0c, 0x00, 0x00, 0xff, 0xff, 0x9f, 0x60, 0xf7, 0xfa, 0x9d, 0x03, 0x00, 0x00,
}
