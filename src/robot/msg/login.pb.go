// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msg/login.proto

package msg

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

type C_A_Login struct {
	AccountName          string   `protobuf:"bytes,1,opt,name=AccountName,proto3" json:"AccountName,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	Channel              int32    `protobuf:"varint,3,opt,name=Channel,proto3" json:"Channel,omitempty"`
	LoginType            int32    `protobuf:"varint,4,opt,name=LoginType,proto3" json:"LoginType,omitempty"`
	HeadIcon             string   `protobuf:"bytes,5,opt,name=HeadIcon,proto3" json:"HeadIcon,omitempty"`
	Sex                  int32    `protobuf:"varint,6,opt,name=Sex,proto3" json:"Sex,omitempty"`
	NickName             string   `protobuf:"bytes,7,opt,name=NickName,proto3" json:"NickName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C_A_Login) Reset()         { *m = C_A_Login{} }
func (m *C_A_Login) String() string { return proto.CompactTextString(m) }
func (*C_A_Login) ProtoMessage()    {}
func (*C_A_Login) Descriptor() ([]byte, []int) {
	return fileDescriptor_085a91d1568bd742, []int{0}
}

func (m *C_A_Login) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C_A_Login.Unmarshal(m, b)
}
func (m *C_A_Login) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C_A_Login.Marshal(b, m, deterministic)
}
func (m *C_A_Login) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C_A_Login.Merge(m, src)
}
func (m *C_A_Login) XXX_Size() int {
	return xxx_messageInfo_C_A_Login.Size(m)
}
func (m *C_A_Login) XXX_DiscardUnknown() {
	xxx_messageInfo_C_A_Login.DiscardUnknown(m)
}

var xxx_messageInfo_C_A_Login proto.InternalMessageInfo

func (m *C_A_Login) GetAccountName() string {
	if m != nil {
		return m.AccountName
	}
	return ""
}

func (m *C_A_Login) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *C_A_Login) GetChannel() int32 {
	if m != nil {
		return m.Channel
	}
	return 0
}

func (m *C_A_Login) GetLoginType() int32 {
	if m != nil {
		return m.LoginType
	}
	return 0
}

func (m *C_A_Login) GetHeadIcon() string {
	if m != nil {
		return m.HeadIcon
	}
	return ""
}

func (m *C_A_Login) GetSex() int32 {
	if m != nil {
		return m.Sex
	}
	return 0
}

func (m *C_A_Login) GetNickName() string {
	if m != nil {
		return m.NickName
	}
	return ""
}

type A_C_Login struct {
	ErrorStr             string   `protobuf:"bytes,1,opt,name=ErrorStr,proto3" json:"ErrorStr,omitempty"`
	UserID               int32    `protobuf:"varint,12,opt,name=UserID,proto3" json:"UserID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *A_C_Login) Reset()         { *m = A_C_Login{} }
func (m *A_C_Login) String() string { return proto.CompactTextString(m) }
func (*A_C_Login) ProtoMessage()    {}
func (*A_C_Login) Descriptor() ([]byte, []int) {
	return fileDescriptor_085a91d1568bd742, []int{1}
}

func (m *A_C_Login) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_A_C_Login.Unmarshal(m, b)
}
func (m *A_C_Login) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_A_C_Login.Marshal(b, m, deterministic)
}
func (m *A_C_Login) XXX_Merge(src proto.Message) {
	xxx_messageInfo_A_C_Login.Merge(m, src)
}
func (m *A_C_Login) XXX_Size() int {
	return xxx_messageInfo_A_C_Login.Size(m)
}
func (m *A_C_Login) XXX_DiscardUnknown() {
	xxx_messageInfo_A_C_Login.DiscardUnknown(m)
}

var xxx_messageInfo_A_C_Login proto.InternalMessageInfo

func (m *A_C_Login) GetErrorStr() string {
	if m != nil {
		return m.ErrorStr
	}
	return ""
}

func (m *A_C_Login) GetUserID() int32 {
	if m != nil {
		return m.UserID
	}
	return 0
}

type C_S_Login struct {
	UserID               int32    `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C_S_Login) Reset()         { *m = C_S_Login{} }
func (m *C_S_Login) String() string { return proto.CompactTextString(m) }
func (*C_S_Login) ProtoMessage()    {}
func (*C_S_Login) Descriptor() ([]byte, []int) {
	return fileDescriptor_085a91d1568bd742, []int{2}
}

func (m *C_S_Login) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C_S_Login.Unmarshal(m, b)
}
func (m *C_S_Login) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C_S_Login.Marshal(b, m, deterministic)
}
func (m *C_S_Login) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C_S_Login.Merge(m, src)
}
func (m *C_S_Login) XXX_Size() int {
	return xxx_messageInfo_C_S_Login.Size(m)
}
func (m *C_S_Login) XXX_DiscardUnknown() {
	xxx_messageInfo_C_S_Login.DiscardUnknown(m)
}

var xxx_messageInfo_C_S_Login proto.InternalMessageInfo

func (m *C_S_Login) GetUserID() int32 {
	if m != nil {
		return m.UserID
	}
	return 0
}

type S_C_Login struct {
	ErrorStr             string   `protobuf:"bytes,1,opt,name=ErrorStr,proto3" json:"ErrorStr,omitempty"`
	UserID               int32    `protobuf:"varint,2,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Gold                 int64    `protobuf:"varint,3,opt,name=Gold,proto3" json:"Gold,omitempty"`
	Coin                 int64    `protobuf:"varint,4,opt,name=Coin,proto3" json:"Coin,omitempty"`
	Money                int64    `protobuf:"varint,5,opt,name=Money,proto3" json:"Money,omitempty"`
	HeadIconUrl          string   `protobuf:"bytes,6,opt,name=HeadIconUrl,proto3" json:"HeadIconUrl,omitempty"`
	UnderWrite           string   `protobuf:"bytes,7,opt,name=UnderWrite,proto3" json:"UnderWrite,omitempty"`
	Sex                  int32    `protobuf:"varint,8,opt,name=Sex,proto3" json:"Sex,omitempty"`
	ActiveFlag           int64    `protobuf:"varint,9,opt,name=ActiveFlag,proto3" json:"ActiveFlag,omitempty"`
	NickName             string   `protobuf:"bytes,10,opt,name=NickName,proto3" json:"NickName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *S_C_Login) Reset()         { *m = S_C_Login{} }
func (m *S_C_Login) String() string { return proto.CompactTextString(m) }
func (*S_C_Login) ProtoMessage()    {}
func (*S_C_Login) Descriptor() ([]byte, []int) {
	return fileDescriptor_085a91d1568bd742, []int{3}
}

func (m *S_C_Login) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S_C_Login.Unmarshal(m, b)
}
func (m *S_C_Login) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S_C_Login.Marshal(b, m, deterministic)
}
func (m *S_C_Login) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S_C_Login.Merge(m, src)
}
func (m *S_C_Login) XXX_Size() int {
	return xxx_messageInfo_S_C_Login.Size(m)
}
func (m *S_C_Login) XXX_DiscardUnknown() {
	xxx_messageInfo_S_C_Login.DiscardUnknown(m)
}

var xxx_messageInfo_S_C_Login proto.InternalMessageInfo

func (m *S_C_Login) GetErrorStr() string {
	if m != nil {
		return m.ErrorStr
	}
	return ""
}

func (m *S_C_Login) GetUserID() int32 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *S_C_Login) GetGold() int64 {
	if m != nil {
		return m.Gold
	}
	return 0
}

func (m *S_C_Login) GetCoin() int64 {
	if m != nil {
		return m.Coin
	}
	return 0
}

func (m *S_C_Login) GetMoney() int64 {
	if m != nil {
		return m.Money
	}
	return 0
}

func (m *S_C_Login) GetHeadIconUrl() string {
	if m != nil {
		return m.HeadIconUrl
	}
	return ""
}

func (m *S_C_Login) GetUnderWrite() string {
	if m != nil {
		return m.UnderWrite
	}
	return ""
}

func (m *S_C_Login) GetSex() int32 {
	if m != nil {
		return m.Sex
	}
	return 0
}

func (m *S_C_Login) GetActiveFlag() int64 {
	if m != nil {
		return m.ActiveFlag
	}
	return 0
}

func (m *S_C_Login) GetNickName() string {
	if m != nil {
		return m.NickName
	}
	return ""
}

func init() {
	proto.RegisterType((*C_A_Login)(nil), "msg.C_A_Login")
	proto.RegisterType((*A_C_Login)(nil), "msg.A_C_Login")
	proto.RegisterType((*C_S_Login)(nil), "msg.C_S_Login")
	proto.RegisterType((*S_C_Login)(nil), "msg.S_C_Login")
}

func init() { proto.RegisterFile("msg/login.proto", fileDescriptor_085a91d1568bd742) }

var fileDescriptor_085a91d1568bd742 = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xdb, 0x4a, 0xfb, 0x40,
	0x10, 0xc6, 0x49, 0xb7, 0xa7, 0x9d, 0xff, 0x1f, 0x94, 0x45, 0x64, 0x29, 0x52, 0x4a, 0xbd, 0xe9,
	0x95, 0x5e, 0xf8, 0x00, 0x12, 0xe2, 0xa9, 0xa0, 0x45, 0x36, 0x16, 0x2f, 0x4b, 0x4c, 0x96, 0x18,
	0x4c, 0x77, 0xcb, 0x26, 0x1e, 0xfa, 0x10, 0x3e, 0x99, 0x2f, 0x25, 0x3b, 0xcd, 0x36, 0xc9, 0xa5,
	0x77, 0xf3, 0x7d, 0x33, 0xf9, 0x26, 0xf3, 0x4b, 0xe0, 0x60, 0x5d, 0xa4, 0xe7, 0xb9, 0x4e, 0x33,
	0x75, 0xb6, 0x31, 0xba, 0xd4, 0x8c, 0xac, 0x8b, 0x74, 0xfa, 0xe3, 0x01, 0x0d, 0x56, 0xfe, 0xea,
	0xde, 0x36, 0xd8, 0x04, 0xfe, 0xf9, 0x71, 0xac, 0xdf, 0x55, 0xb9, 0x88, 0xd6, 0x92, 0x7b, 0x13,
	0x6f, 0x46, 0x45, 0xd3, 0x62, 0x23, 0x18, 0x3e, 0x46, 0x45, 0xf1, 0xa9, 0x4d, 0xc2, 0x3b, 0xd8,
	0xde, 0x6b, 0xc6, 0x61, 0x10, 0xbc, 0x46, 0x4a, 0xc9, 0x9c, 0x93, 0x89, 0x37, 0xeb, 0x09, 0x27,
	0xd9, 0x09, 0x50, 0x5c, 0xf0, 0xb4, 0xdd, 0x48, 0xde, 0xc5, 0x5e, 0x6d, 0xd8, 0xcc, 0x3b, 0x19,
	0x25, 0xf3, 0x58, 0x2b, 0xde, 0xdb, 0x65, 0x3a, 0xcd, 0x0e, 0x81, 0x84, 0xf2, 0x8b, 0xf7, 0xf1,
	0x19, 0x5b, 0xda, 0xe9, 0x45, 0x16, 0xbf, 0xe1, 0x0b, 0x0e, 0x76, 0xd3, 0x4e, 0x4f, 0x2f, 0x81,
	0xfa, 0xab, 0xa0, 0x3a, 0x66, 0x04, 0xc3, 0x6b, 0x63, 0xb4, 0x09, 0x4b, 0x53, 0x5d, 0xb2, 0xd7,
	0xec, 0x18, 0xfa, 0xcb, 0x42, 0x9a, 0xf9, 0x15, 0xff, 0x8f, 0xc9, 0x95, 0x9a, 0x9e, 0x5a, 0x1a,
	0x61, 0x15, 0x50, 0x0f, 0x79, 0xad, 0xa1, 0xef, 0x0e, 0xd0, 0xf0, 0x8f, 0x6b, 0x3a, 0xcd, 0x04,
	0xc6, 0xa0, 0x7b, 0xab, 0xf3, 0x04, 0x31, 0x11, 0x81, 0xb5, 0xf5, 0x02, 0x9d, 0x29, 0xc4, 0x43,
	0x04, 0xd6, 0xec, 0x08, 0x7a, 0x0f, 0x5a, 0xc9, 0x2d, 0x62, 0x21, 0x62, 0x27, 0xec, 0x57, 0x72,
	0x7c, 0x96, 0x26, 0x47, 0x36, 0x54, 0x34, 0x2d, 0x36, 0x06, 0x58, 0xaa, 0x44, 0x9a, 0x67, 0x93,
	0x95, 0x8e, 0x52, 0xc3, 0x71, 0x54, 0x87, 0x35, 0xd5, 0x31, 0x80, 0x1f, 0x97, 0xd9, 0x87, 0xbc,
	0xc9, 0xa3, 0x94, 0x53, 0x5c, 0xd7, 0x70, 0x5a, 0xd4, 0xa1, 0x4d, 0xfd, 0xa5, 0x8f, 0xff, 0xd3,
	0xc5, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xab, 0xee, 0xc5, 0xa3, 0x62, 0x02, 0x00, 0x00,
}
