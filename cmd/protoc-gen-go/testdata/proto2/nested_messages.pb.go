// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto2/nested_messages.proto

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

type Layer1 struct {
	L2                   *Layer1_Layer2        `protobuf:"bytes,1,opt,name=l2" json:"l2,omitempty"`
	L3                   *Layer1_Layer2_Layer3 `protobuf:"bytes,2,opt,name=l3" json:"l3,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Layer1) Reset()         { *m = Layer1{} }
func (m *Layer1) String() string { return proto.CompactTextString(m) }
func (*Layer1) ProtoMessage()    {}
func (*Layer1) Descriptor() ([]byte, []int) {
	return fileDescriptor_7417ee157699d191, []int{0}
}
func (m *Layer1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Layer1.Unmarshal(m, b)
}
func (m *Layer1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Layer1.Marshal(b, m, deterministic)
}
func (m *Layer1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Layer1.Merge(m, src)
}
func (m *Layer1) XXX_Size() int {
	return xxx_messageInfo_Layer1.Size(m)
}
func (m *Layer1) XXX_DiscardUnknown() {
	xxx_messageInfo_Layer1.DiscardUnknown(m)
}

var xxx_messageInfo_Layer1 proto.InternalMessageInfo

func (m *Layer1) GetL2() *Layer1_Layer2 {
	if m != nil {
		return m.L2
	}
	return nil
}

func (m *Layer1) GetL3() *Layer1_Layer2_Layer3 {
	if m != nil {
		return m.L3
	}
	return nil
}

type Layer1_Layer2 struct {
	L3                   *Layer1_Layer2_Layer3 `protobuf:"bytes,1,opt,name=l3" json:"l3,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Layer1_Layer2) Reset()         { *m = Layer1_Layer2{} }
func (m *Layer1_Layer2) String() string { return proto.CompactTextString(m) }
func (*Layer1_Layer2) ProtoMessage()    {}
func (*Layer1_Layer2) Descriptor() ([]byte, []int) {
	return fileDescriptor_7417ee157699d191, []int{0, 0}
}
func (m *Layer1_Layer2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Layer1_Layer2.Unmarshal(m, b)
}
func (m *Layer1_Layer2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Layer1_Layer2.Marshal(b, m, deterministic)
}
func (m *Layer1_Layer2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Layer1_Layer2.Merge(m, src)
}
func (m *Layer1_Layer2) XXX_Size() int {
	return xxx_messageInfo_Layer1_Layer2.Size(m)
}
func (m *Layer1_Layer2) XXX_DiscardUnknown() {
	xxx_messageInfo_Layer1_Layer2.DiscardUnknown(m)
}

var xxx_messageInfo_Layer1_Layer2 proto.InternalMessageInfo

func (m *Layer1_Layer2) GetL3() *Layer1_Layer2_Layer3 {
	if m != nil {
		return m.L3
	}
	return nil
}

type Layer1_Layer2_Layer3 struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Layer1_Layer2_Layer3) Reset()         { *m = Layer1_Layer2_Layer3{} }
func (m *Layer1_Layer2_Layer3) String() string { return proto.CompactTextString(m) }
func (*Layer1_Layer2_Layer3) ProtoMessage()    {}
func (*Layer1_Layer2_Layer3) Descriptor() ([]byte, []int) {
	return fileDescriptor_7417ee157699d191, []int{0, 0, 0}
}
func (m *Layer1_Layer2_Layer3) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Layer1_Layer2_Layer3.Unmarshal(m, b)
}
func (m *Layer1_Layer2_Layer3) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Layer1_Layer2_Layer3.Marshal(b, m, deterministic)
}
func (m *Layer1_Layer2_Layer3) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Layer1_Layer2_Layer3.Merge(m, src)
}
func (m *Layer1_Layer2_Layer3) XXX_Size() int {
	return xxx_messageInfo_Layer1_Layer2_Layer3.Size(m)
}
func (m *Layer1_Layer2_Layer3) XXX_DiscardUnknown() {
	xxx_messageInfo_Layer1_Layer2_Layer3.DiscardUnknown(m)
}

var xxx_messageInfo_Layer1_Layer2_Layer3 proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Layer1)(nil), "goproto.protoc.proto2.Layer1")
	proto.RegisterType((*Layer1_Layer2)(nil), "goproto.protoc.proto2.Layer1.Layer2")
	proto.RegisterType((*Layer1_Layer2_Layer3)(nil), "goproto.protoc.proto2.Layer1.Layer2.Layer3")
}

func init() { proto.RegisterFile("proto2/nested_messages.proto", fileDescriptor_7417ee157699d191) }

var fileDescriptor_7417ee157699d191 = []byte{
	// 184 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x29, 0x28, 0xca, 0x2f,
	0xc9, 0x37, 0xd2, 0xcf, 0x4b, 0x2d, 0x2e, 0x49, 0x4d, 0x89, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c,
	0x4f, 0x2d, 0xd6, 0x03, 0x0b, 0x0b, 0x89, 0xa6, 0xe7, 0x83, 0x19, 0x10, 0x6e, 0x32, 0x84, 0x32,
	0x52, 0x3a, 0xc3, 0xc8, 0xc5, 0xe6, 0x93, 0x58, 0x99, 0x5a, 0x64, 0x28, 0x64, 0xc2, 0xc5, 0x94,
	0x63, 0x24, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x6d, 0xa4, 0xa2, 0x87, 0x55, 0xb9, 0x1e, 0x44, 0x29,
	0x84, 0x32, 0x0a, 0x62, 0xca, 0x31, 0x12, 0xb2, 0xe6, 0x62, 0xca, 0x31, 0x96, 0x60, 0x02, 0xeb,
	0xd2, 0x26, 0x46, 0x17, 0x84, 0x32, 0x0e, 0x62, 0xca, 0x31, 0x96, 0xf2, 0x87, 0x5a, 0x0e, 0x33,
	0x86, 0x91, 0x3c, 0x63, 0x38, 0xa0, 0xc6, 0x18, 0x3b, 0x59, 0x47, 0x59, 0xa6, 0xe7, 0xe7, 0xa7,
	0xe7, 0xa4, 0xea, 0xa5, 0xe7, 0xe7, 0x24, 0xe6, 0xa5, 0xeb, 0xe5, 0x17, 0xa5, 0xeb, 0x83, 0xb5,
	0xeb, 0x27, 0xe7, 0xa6, 0x40, 0x58, 0xc9, 0xba, 0xe9, 0xa9, 0x79, 0xba, 0xe9, 0xf9, 0xfa, 0x25,
	0xa9, 0xc5, 0x25, 0x29, 0x89, 0x25, 0x89, 0x10, 0x61, 0x23, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x47, 0x0d, 0xa3, 0x19, 0x41, 0x01, 0x00, 0x00,
}
