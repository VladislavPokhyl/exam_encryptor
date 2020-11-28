// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/randomStrings.proto

package proto

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

type ListOfStrings struct {
	//list of string to pass to the encryptor
	List                 []string `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListOfStrings) Reset()         { *m = ListOfStrings{} }
func (m *ListOfStrings) String() string { return proto.CompactTextString(m) }
func (*ListOfStrings) ProtoMessage()    {}
func (*ListOfStrings) Descriptor() ([]byte, []int) {
	return fileDescriptor_30bbaf6140154aa6, []int{0}
}

func (m *ListOfStrings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListOfStrings.Unmarshal(m, b)
}
func (m *ListOfStrings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListOfStrings.Marshal(b, m, deterministic)
}
func (m *ListOfStrings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListOfStrings.Merge(m, src)
}
func (m *ListOfStrings) XXX_Size() int {
	return xxx_messageInfo_ListOfStrings.Size(m)
}
func (m *ListOfStrings) XXX_DiscardUnknown() {
	xxx_messageInfo_ListOfStrings.DiscardUnknown(m)
}

var xxx_messageInfo_ListOfStrings proto.InternalMessageInfo

func (m *ListOfStrings) GetList() []string {
	if m != nil {
		return m.List
	}
	return nil
}

//response with encrypted slice of strings
type Response struct {
	List                 []string `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	OneStr               string   `protobuf:"bytes,2,opt,name=oneStr,proto3" json:"oneStr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_30bbaf6140154aa6, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetList() []string {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *Response) GetOneStr() string {
	if m != nil {
		return m.OneStr
	}
	return ""
}

func init() {
	proto.RegisterType((*ListOfStrings)(nil), "proto.listOfStrings")
	proto.RegisterType((*Response)(nil), "proto.Response")
}

func init() {
	proto.RegisterFile("proto/randomStrings.proto", fileDescriptor_30bbaf6140154aa6)
}

var fileDescriptor_30bbaf6140154aa6 = []byte{
	// 207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x4a, 0xcc, 0x4b, 0xc9, 0xcf, 0x0d, 0x2e, 0x29, 0xca, 0xcc, 0x4b, 0x2f, 0xd6,
	0x03, 0x8b, 0x09, 0xb1, 0x82, 0x29, 0x25, 0x65, 0x2e, 0xde, 0x9c, 0xcc, 0xe2, 0x12, 0xff, 0x34,
	0xa8, 0xac, 0x90, 0x10, 0x17, 0x0b, 0x48, 0x40, 0x82, 0x51, 0x81, 0x59, 0x83, 0x33, 0x08, 0xcc,
	0x56, 0x32, 0xe3, 0xe2, 0x08, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0xc5, 0x26, 0x2f, 0x24,
	0xc6, 0xc5, 0x96, 0x9f, 0x97, 0x1a, 0x5c, 0x52, 0x24, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x19, 0x04,
	0xe5, 0x19, 0x79, 0x73, 0xf1, 0x15, 0x43, 0x8c, 0x0d, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x15,
	0xb2, 0xe4, 0x12, 0x70, 0x4f, 0x2d, 0x09, 0xce, 0x48, 0x34, 0x32, 0x35, 0x83, 0xd9, 0x28, 0x02,
	0x71, 0x91, 0x1e, 0x8a, 0x3b, 0xa4, 0xf8, 0xa1, 0xa2, 0x30, 0x8b, 0x9d, 0xec, 0xa2, 0x6c, 0xd2,
	0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0xc3, 0x72, 0x12, 0x53, 0xdc, 0x72,
	0x12, 0x93, 0x52, 0x8b, 0xf4, 0x53, 0x2b, 0x12, 0x73, 0xe3, 0x53, 0xf3, 0x92, 0x8b, 0x2a, 0x0b,
	0x4a, 0xf2, 0x8b, 0xf4, 0xe1, 0x2c, 0xdd, 0x62, 0x88, 0xa5, 0xfa, 0x60, 0x83, 0x92, 0xd8, 0xc0,
	0x94, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xc6, 0xcf, 0x9a, 0xa9, 0x14, 0x01, 0x00, 0x00,
}
