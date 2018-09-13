// Code generated by protoc-gen-go. DO NOT EDIT.
// source: comments/comments.proto

// COMMENT: package goproto.protoc.proto2

package proto2

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// COMMENT: Message1
type Message1 struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message1) Reset()         { *m = Message1{} }
func (m *Message1) String() string { return proto.CompactTextString(m) }
func (*Message1) ProtoMessage()    {}
func (*Message1) Descriptor() ([]byte, []int) {
	return fileDescriptor_885e8293f1fab554, []int{0}
}
func (m *Message1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message1.Unmarshal(m, b)
}
func (m *Message1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message1.Marshal(b, m, deterministic)
}
func (m *Message1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message1.Merge(m, src)
}
func (m *Message1) XXX_Size() int {
	return xxx_messageInfo_Message1.Size(m)
}
func (m *Message1) XXX_DiscardUnknown() {
	xxx_messageInfo_Message1.DiscardUnknown(m)
}

var xxx_messageInfo_Message1 proto.InternalMessageInfo

// COMMENT: Message2
type Message2 struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message2) Reset()         { *m = Message2{} }
func (m *Message2) String() string { return proto.CompactTextString(m) }
func (*Message2) ProtoMessage()    {}
func (*Message2) Descriptor() ([]byte, []int) {
	return fileDescriptor_885e8293f1fab554, []int{1}
}
func (m *Message2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message2.Unmarshal(m, b)
}
func (m *Message2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message2.Marshal(b, m, deterministic)
}
func (m *Message2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message2.Merge(m, src)
}
func (m *Message2) XXX_Size() int {
	return xxx_messageInfo_Message2.Size(m)
}
func (m *Message2) XXX_DiscardUnknown() {
	xxx_messageInfo_Message2.DiscardUnknown(m)
}

var xxx_messageInfo_Message2 proto.InternalMessageInfo

// COMMENT: Message1A
type Message1_Message1A struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message1_Message1A) Reset()         { *m = Message1_Message1A{} }
func (m *Message1_Message1A) String() string { return proto.CompactTextString(m) }
func (*Message1_Message1A) ProtoMessage()    {}
func (*Message1_Message1A) Descriptor() ([]byte, []int) {
	return fileDescriptor_885e8293f1fab554, []int{0, 0}
}
func (m *Message1_Message1A) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message1_Message1A.Unmarshal(m, b)
}
func (m *Message1_Message1A) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message1_Message1A.Marshal(b, m, deterministic)
}
func (m *Message1_Message1A) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message1_Message1A.Merge(m, src)
}
func (m *Message1_Message1A) XXX_Size() int {
	return xxx_messageInfo_Message1_Message1A.Size(m)
}
func (m *Message1_Message1A) XXX_DiscardUnknown() {
	xxx_messageInfo_Message1_Message1A.DiscardUnknown(m)
}

var xxx_messageInfo_Message1_Message1A proto.InternalMessageInfo

// COMMENT: Message1B
type Message1_Message1B struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message1_Message1B) Reset()         { *m = Message1_Message1B{} }
func (m *Message1_Message1B) String() string { return proto.CompactTextString(m) }
func (*Message1_Message1B) ProtoMessage()    {}
func (*Message1_Message1B) Descriptor() ([]byte, []int) {
	return fileDescriptor_885e8293f1fab554, []int{0, 1}
}
func (m *Message1_Message1B) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message1_Message1B.Unmarshal(m, b)
}
func (m *Message1_Message1B) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message1_Message1B.Marshal(b, m, deterministic)
}
func (m *Message1_Message1B) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message1_Message1B.Merge(m, src)
}
func (m *Message1_Message1B) XXX_Size() int {
	return xxx_messageInfo_Message1_Message1B.Size(m)
}
func (m *Message1_Message1B) XXX_DiscardUnknown() {
	xxx_messageInfo_Message1_Message1B.DiscardUnknown(m)
}

var xxx_messageInfo_Message1_Message1B proto.InternalMessageInfo

// COMMENT: Message2A
type Message2_Message2A struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message2_Message2A) Reset()         { *m = Message2_Message2A{} }
func (m *Message2_Message2A) String() string { return proto.CompactTextString(m) }
func (*Message2_Message2A) ProtoMessage()    {}
func (*Message2_Message2A) Descriptor() ([]byte, []int) {
	return fileDescriptor_885e8293f1fab554, []int{1, 0}
}
func (m *Message2_Message2A) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message2_Message2A.Unmarshal(m, b)
}
func (m *Message2_Message2A) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message2_Message2A.Marshal(b, m, deterministic)
}
func (m *Message2_Message2A) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message2_Message2A.Merge(m, src)
}
func (m *Message2_Message2A) XXX_Size() int {
	return xxx_messageInfo_Message2_Message2A.Size(m)
}
func (m *Message2_Message2A) XXX_DiscardUnknown() {
	xxx_messageInfo_Message2_Message2A.DiscardUnknown(m)
}

var xxx_messageInfo_Message2_Message2A proto.InternalMessageInfo

// COMMENT: Message2B
type Message2_Message2B struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message2_Message2B) Reset()         { *m = Message2_Message2B{} }
func (m *Message2_Message2B) String() string { return proto.CompactTextString(m) }
func (*Message2_Message2B) ProtoMessage()    {}
func (*Message2_Message2B) Descriptor() ([]byte, []int) {
	return fileDescriptor_885e8293f1fab554, []int{1, 1}
}
func (m *Message2_Message2B) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message2_Message2B.Unmarshal(m, b)
}
func (m *Message2_Message2B) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message2_Message2B.Marshal(b, m, deterministic)
}
func (m *Message2_Message2B) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message2_Message2B.Merge(m, src)
}
func (m *Message2_Message2B) XXX_Size() int {
	return xxx_messageInfo_Message2_Message2B.Size(m)
}
func (m *Message2_Message2B) XXX_DiscardUnknown() {
	xxx_messageInfo_Message2_Message2B.DiscardUnknown(m)
}

var xxx_messageInfo_Message2_Message2B proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Message1)(nil), "goproto.protoc.proto2.Message1")
	proto.RegisterType((*Message2)(nil), "goproto.protoc.proto2.Message2")
	proto.RegisterType((*Message1_Message1A)(nil), "goproto.protoc.proto2.Message1.Message1A")
	proto.RegisterType((*Message1_Message1B)(nil), "goproto.protoc.proto2.Message1.Message1B")
	proto.RegisterType((*Message2_Message2A)(nil), "goproto.protoc.proto2.Message2.Message2A")
	proto.RegisterType((*Message2_Message2B)(nil), "goproto.protoc.proto2.Message2.Message2B")
}

func init() { proto.RegisterFile("comments/comments.proto", fileDescriptor_885e8293f1fab554) }

var fileDescriptor_885e8293f1fab554 = []byte{
	// 136 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4f, 0xce, 0xcf, 0xcd,
	0x4d, 0xcd, 0x2b, 0x29, 0xd6, 0x87, 0x31, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x44, 0xd3,
	0xf3, 0xc1, 0x0c, 0x08, 0x37, 0x19, 0x42, 0x19, 0x29, 0xa9, 0x70, 0x71, 0xf8, 0xa6, 0x16, 0x17,
	0x27, 0xa6, 0xa7, 0x1a, 0x4a, 0x71, 0x73, 0x71, 0xc2, 0xd8, 0x8e, 0xc8, 0x1c, 0x27, 0x24, 0x55,
	0x46, 0x48, 0x12, 0x46, 0xc8, 0xaa, 0x8c, 0x9c, 0x9c, 0xac, 0xa3, 0x2c, 0xd3, 0xf3, 0xf3, 0xd3,
	0x73, 0x52, 0xf5, 0xd2, 0xf3, 0x73, 0x12, 0xf3, 0xd2, 0xf5, 0xf2, 0x8b, 0xd2, 0xf5, 0xc1, 0xf6,
	0xe8, 0x27, 0xe7, 0xa6, 0x40, 0x58, 0xc9, 0xba, 0xe9, 0xa9, 0x79, 0xba, 0xe9, 0xf9, 0xfa, 0x25,
	0xa9, 0xc5, 0x25, 0x29, 0x89, 0x25, 0x89, 0x10, 0x61, 0x23, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x32, 0x8f, 0xcd, 0x4f, 0xb9, 0x00, 0x00, 0x00,
}
