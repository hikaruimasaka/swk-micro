// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cache.proto

package cache

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

type SetRequest struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value"`
	Key                  []string `protobuf:"bytes,2,rep,name=key,proto3" json:"key"`
	Ttl                  int64    `protobuf:"varint,3,opt,name=ttl,proto3" json:"ttl"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetRequest) Reset()         { *m = SetRequest{} }
func (m *SetRequest) String() string { return proto.CompactTextString(m) }
func (*SetRequest) ProtoMessage()    {}
func (*SetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fca3b110c9bbf3a, []int{0}
}

func (m *SetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetRequest.Unmarshal(m, b)
}
func (m *SetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetRequest.Marshal(b, m, deterministic)
}
func (m *SetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetRequest.Merge(m, src)
}
func (m *SetRequest) XXX_Size() int {
	return xxx_messageInfo_SetRequest.Size(m)
}
func (m *SetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetRequest proto.InternalMessageInfo

func (m *SetRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *SetRequest) GetKey() []string {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *SetRequest) GetTtl() int64 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

type GetRequest struct {
	Key                  []string `protobuf:"bytes,1,rep,name=key,proto3" json:"key"`
	Ttl                  int64    `protobuf:"varint,2,opt,name=ttl,proto3" json:"ttl"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fca3b110c9bbf3a, []int{1}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetKey() []string {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *GetRequest) GetTtl() int64 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

type DeleteRequest struct {
	Key                  []string `protobuf:"bytes,1,rep,name=key,proto3" json:"key"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fca3b110c9bbf3a, []int{2}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetKey() []string {
	if m != nil {
		return m.Key
	}
	return nil
}

type GetResponse struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fca3b110c9bbf3a, []int{3}
}

func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (m *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(m, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Response struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fca3b110c9bbf3a, []int{4}
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

func init() {
	proto.RegisterType((*SetRequest)(nil), "cache.SetRequest")
	proto.RegisterType((*GetRequest)(nil), "cache.GetRequest")
	proto.RegisterType((*DeleteRequest)(nil), "cache.DeleteRequest")
	proto.RegisterType((*GetResponse)(nil), "cache.GetResponse")
	proto.RegisterType((*Response)(nil), "cache.Response")
}

func init() { proto.RegisterFile("cache.proto", fileDescriptor_5fca3b110c9bbf3a) }

var fileDescriptor_5fca3b110c9bbf3a = []byte{
	// 224 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcd, 0x4e, 0x85, 0x30,
	0x10, 0x85, 0x2d, 0x0d, 0x06, 0x06, 0x8d, 0xda, 0xb0, 0x20, 0xac, 0xb0, 0x6e, 0x58, 0x11, 0x22,
	0x89, 0x2f, 0xa0, 0x91, 0x7d, 0x79, 0x02, 0x24, 0x93, 0x68, 0x24, 0x82, 0x50, 0x48, 0x7c, 0x28,
	0xdf, 0xf1, 0xa6, 0x2d, 0xdc, 0xde, 0x1f, 0xee, 0x6e, 0xe6, 0x64, 0xbe, 0x03, 0x5f, 0x0a, 0x41,
	0x53, 0x37, 0x9f, 0x98, 0xf5, 0x43, 0x27, 0x3b, 0xe6, 0xea, 0x85, 0xbf, 0x03, 0x54, 0x28, 0x05,
	0xfe, 0x4e, 0x38, 0x4a, 0x16, 0x82, 0x3b, 0xd7, 0xed, 0x84, 0x11, 0x49, 0x48, 0xea, 0x0b, 0xb3,
	0xb0, 0x7b, 0xa0, 0xdf, 0xf8, 0x17, 0x39, 0x09, 0x4d, 0x7d, 0xa1, 0x46, 0x95, 0x48, 0xd9, 0x46,
	0x34, 0x21, 0x29, 0x15, 0x6a, 0xe4, 0x39, 0x40, 0x69, 0x7b, 0x16, 0x82, 0x9c, 0x11, 0x8e, 0x25,
	0x1e, 0xe1, 0xf6, 0x0d, 0x5b, 0x94, 0x78, 0x11, 0xe2, 0x4f, 0x10, 0xe8, 0xd2, 0xb1, 0xef, 0x7e,
	0x46, 0xdc, 0xfe, 0x3b, 0x0e, 0xe0, 0xad, 0x17, 0xcf, 0xff, 0x04, 0x6e, 0x5e, 0x95, 0x57, 0x85,
	0xc3, 0xfc, 0xd5, 0x20, 0xcb, 0xc1, 0xab, 0x50, 0xea, 0x88, 0x3d, 0x64, 0xc6, 0xdf, 0xfa, 0xc6,
	0x77, 0x4b, 0xb4, 0x16, 0xf0, 0x2b, 0x56, 0x80, 0x57, 0x9e, 0x12, 0xd6, 0x2c, 0x66, 0x87, 0xd1,
	0x1e, 0x7a, 0x81, 0xc0, 0xb8, 0x18, 0x2e, 0x5c, 0x8e, 0x8e, 0xfc, 0x36, 0x3e, 0xf6, 0x71, 0xad,
	0xdf, 0xa2, 0xd8, 0x05, 0x00, 0x00, 0xff, 0xff, 0x44, 0xa8, 0x51, 0x52, 0x9a, 0x01, 0x00, 0x00,
}
