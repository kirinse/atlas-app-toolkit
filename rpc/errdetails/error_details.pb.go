// Code generated by protoc-gen-go. DO NOT EDIT.
// source: error_details.proto

package errdetails // import "github.com/kirinse/atlas-app-toolkit/rpc/errdetails"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// TargetInfo is a default representation of error details that conforms
// REST API Syntax Specification
type TargetInfo struct {
	// The status code is an enumerated error code,
	// which should be an enum value of [google.rpc.Code][google.rpc.Code]
	Code int32 `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	// The message is a human-readable non-localized message
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	// The target is a resource name
	Target               string   `protobuf:"bytes,3,opt,name=target" json:"target,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TargetInfo) Reset()         { *m = TargetInfo{} }
func (m *TargetInfo) String() string { return proto.CompactTextString(m) }
func (*TargetInfo) ProtoMessage()    {}
func (*TargetInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_error_details_48e27bb52297ce6d, []int{0}
}
func (m *TargetInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TargetInfo.Unmarshal(m, b)
}
func (m *TargetInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TargetInfo.Marshal(b, m, deterministic)
}
func (dst *TargetInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TargetInfo.Merge(dst, src)
}
func (m *TargetInfo) XXX_Size() int {
	return xxx_messageInfo_TargetInfo.Size(m)
}
func (m *TargetInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TargetInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TargetInfo proto.InternalMessageInfo

func (m *TargetInfo) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *TargetInfo) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *TargetInfo) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func init() {
	proto.RegisterType((*TargetInfo)(nil), "atlas.rpc.TargetInfo")
}

func init() { proto.RegisterFile("error_details.proto", fileDescriptor_error_details_48e27bb52297ce6d) }

var fileDescriptor_error_details_48e27bb52297ce6d = []byte{
	// 175 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x8e, 0xbf, 0xeb, 0xc2, 0x30,
	0x10, 0x47, 0xe9, 0xf7, 0x47, 0xa5, 0x19, 0x23, 0x48, 0xc6, 0xe2, 0xd4, 0xa5, 0xcd, 0xe0, 0xe8,
	0xa6, 0x38, 0xb8, 0x16, 0x27, 0x17, 0x49, 0xd3, 0x6b, 0x0d, 0xa6, 0xbd, 0x70, 0x39, 0xc1, 0x3f,
	0x5f, 0x08, 0x15, 0xb7, 0xf7, 0x38, 0x8e, 0xcf, 0x13, 0x6b, 0x20, 0x42, 0xba, 0xf5, 0xc0, 0xc6,
	0xf9, 0xd8, 0x04, 0x42, 0x46, 0x59, 0x18, 0xf6, 0x26, 0x36, 0x14, 0xec, 0xb6, 0x15, 0xe2, 0x62,
	0x68, 0x04, 0x3e, 0xcf, 0x03, 0x4a, 0x29, 0xfe, 0x2c, 0xf6, 0xa0, 0xb2, 0x32, 0xab, 0xfe, 0xdb,
	0xc4, 0x52, 0x89, 0xd5, 0x04, 0x31, 0x9a, 0x11, 0xd4, 0x4f, 0x99, 0x55, 0x45, 0xfb, 0x51, 0xb9,
	0x11, 0x39, 0xa7, 0x5f, 0xf5, 0x9b, 0x0e, 0x8b, 0x1d, 0x4e, 0xd7, 0xe3, 0xe8, 0xf8, 0xfe, 0xec,
	0x1a, 0x8b, 0x93, 0x76, 0xf3, 0x80, 0x9d, 0xc7, 0x17, 0x06, 0x98, 0x75, 0x1a, 0xae, 0x4d, 0x08,
	0x35, 0x23, 0xfa, 0x87, 0x63, 0x4d, 0xc1, 0x6a, 0x20, 0x5a, 0x02, 0xf7, 0x5f, 0xec, 0xf2, 0x14,
	0xbb, 0x7b, 0x07, 0x00, 0x00, 0xff, 0xff, 0x49, 0x26, 0x7b, 0x22, 0xc3, 0x00, 0x00, 0x00,
}
