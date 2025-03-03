// Code generated by protoc-gen-go. DO NOT EDIT.
// source: copy.proto

package copy

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

type Value struct {
	DataType             string   `protobuf:"bytes,1,opt,name=data_type,json=dataType,proto3" json:"data_type"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Value) Reset()         { *m = Value{} }
func (m *Value) String() string { return proto.CompactTextString(m) }
func (*Value) ProtoMessage()    {}
func (*Value) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b8564a7da00bcc7, []int{0}
}

func (m *Value) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Value.Unmarshal(m, b)
}
func (m *Value) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Value.Marshal(b, m, deterministic)
}
func (m *Value) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Value.Merge(m, src)
}
func (m *Value) XXX_Size() int {
	return xxx_messageInfo_Value.Size(m)
}
func (m *Value) XXX_DiscardUnknown() {
	xxx_messageInfo_Value.DiscardUnknown(m)
}

var xxx_messageInfo_Value proto.InternalMessageInfo

func (m *Value) GetDataType() string {
	if m != nil {
		return m.DataType
	}
	return ""
}

func (m *Value) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// 恢复台账数据
type BulkItem struct {
	Items                map[string]*Value `protobuf:"bytes,3,rep,name=items,proto3" json:"items" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *BulkItem) Reset()         { *m = BulkItem{} }
func (m *BulkItem) String() string { return proto.CompactTextString(m) }
func (*BulkItem) ProtoMessage()    {}
func (*BulkItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b8564a7da00bcc7, []int{1}
}

func (m *BulkItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BulkItem.Unmarshal(m, b)
}
func (m *BulkItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BulkItem.Marshal(b, m, deterministic)
}
func (m *BulkItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BulkItem.Merge(m, src)
}
func (m *BulkItem) XXX_Size() int {
	return xxx_messageInfo_BulkItem.Size(m)
}
func (m *BulkItem) XXX_DiscardUnknown() {
	xxx_messageInfo_BulkItem.DiscardUnknown(m)
}

var xxx_messageInfo_BulkItem proto.InternalMessageInfo

func (m *BulkItem) GetItems() map[string]*Value {
	if m != nil {
		return m.Items
	}
	return nil
}

// 复制数据
type CopyItemsRequest struct {
	AppId                string   `protobuf:"bytes,1,opt,name=app_id,json=appId,proto3" json:"app_id"`
	DatastoreId          string   `protobuf:"bytes,2,opt,name=datastore_id,json=datastoreId,proto3" json:"datastore_id"`
	CopyAppId            string   `protobuf:"bytes,3,opt,name=copy_app_id,json=copyAppId,proto3" json:"copy_app_id"`
	CopyDatastoreId      string   `protobuf:"bytes,4,opt,name=copy_datastore_id,json=copyDatastoreId,proto3" json:"copy_datastore_id"`
	Database             string   `protobuf:"bytes,5,opt,name=database,proto3" json:"database"`
	WithFile             bool     `protobuf:"varint,6,opt,name=with_file,json=withFile,proto3" json:"with_file"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CopyItemsRequest) Reset()         { *m = CopyItemsRequest{} }
func (m *CopyItemsRequest) String() string { return proto.CompactTextString(m) }
func (*CopyItemsRequest) ProtoMessage()    {}
func (*CopyItemsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b8564a7da00bcc7, []int{2}
}

func (m *CopyItemsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CopyItemsRequest.Unmarshal(m, b)
}
func (m *CopyItemsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CopyItemsRequest.Marshal(b, m, deterministic)
}
func (m *CopyItemsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CopyItemsRequest.Merge(m, src)
}
func (m *CopyItemsRequest) XXX_Size() int {
	return xxx_messageInfo_CopyItemsRequest.Size(m)
}
func (m *CopyItemsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CopyItemsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CopyItemsRequest proto.InternalMessageInfo

func (m *CopyItemsRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *CopyItemsRequest) GetDatastoreId() string {
	if m != nil {
		return m.DatastoreId
	}
	return ""
}

func (m *CopyItemsRequest) GetCopyAppId() string {
	if m != nil {
		return m.CopyAppId
	}
	return ""
}

func (m *CopyItemsRequest) GetCopyDatastoreId() string {
	if m != nil {
		return m.CopyDatastoreId
	}
	return ""
}

func (m *CopyItemsRequest) GetDatabase() string {
	if m != nil {
		return m.Database
	}
	return ""
}

func (m *CopyItemsRequest) GetWithFile() bool {
	if m != nil {
		return m.WithFile
	}
	return false
}

type CopyItemsResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CopyItemsResponse) Reset()         { *m = CopyItemsResponse{} }
func (m *CopyItemsResponse) String() string { return proto.CompactTextString(m) }
func (*CopyItemsResponse) ProtoMessage()    {}
func (*CopyItemsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b8564a7da00bcc7, []int{3}
}

func (m *CopyItemsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CopyItemsResponse.Unmarshal(m, b)
}
func (m *CopyItemsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CopyItemsResponse.Marshal(b, m, deterministic)
}
func (m *CopyItemsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CopyItemsResponse.Merge(m, src)
}
func (m *CopyItemsResponse) XXX_Size() int {
	return xxx_messageInfo_CopyItemsResponse.Size(m)
}
func (m *CopyItemsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CopyItemsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CopyItemsResponse proto.InternalMessageInfo

// 恢复数据
type BulkAddRequest struct {
	AppId                string      `protobuf:"bytes,1,opt,name=app_id,json=appId,proto3" json:"app_id"`
	DatastoreId          string      `protobuf:"bytes,2,opt,name=datastore_id,json=datastoreId,proto3" json:"datastore_id"`
	Items                []*BulkItem `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Writer               string      `protobuf:"bytes,4,opt,name=writer,proto3" json:"writer"`
	Owners               []string    `protobuf:"bytes,5,rep,name=owners,proto3" json:"owners"`
	Database             string      `protobuf:"bytes,6,opt,name=database,proto3" json:"database"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *BulkAddRequest) Reset()         { *m = BulkAddRequest{} }
func (m *BulkAddRequest) String() string { return proto.CompactTextString(m) }
func (*BulkAddRequest) ProtoMessage()    {}
func (*BulkAddRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b8564a7da00bcc7, []int{4}
}

func (m *BulkAddRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BulkAddRequest.Unmarshal(m, b)
}
func (m *BulkAddRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BulkAddRequest.Marshal(b, m, deterministic)
}
func (m *BulkAddRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BulkAddRequest.Merge(m, src)
}
func (m *BulkAddRequest) XXX_Size() int {
	return xxx_messageInfo_BulkAddRequest.Size(m)
}
func (m *BulkAddRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BulkAddRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BulkAddRequest proto.InternalMessageInfo

func (m *BulkAddRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *BulkAddRequest) GetDatastoreId() string {
	if m != nil {
		return m.DatastoreId
	}
	return ""
}

func (m *BulkAddRequest) GetItems() []*BulkItem {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *BulkAddRequest) GetWriter() string {
	if m != nil {
		return m.Writer
	}
	return ""
}

func (m *BulkAddRequest) GetOwners() []string {
	if m != nil {
		return m.Owners
	}
	return nil
}

func (m *BulkAddRequest) GetDatabase() string {
	if m != nil {
		return m.Database
	}
	return ""
}

type BulkAddResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BulkAddResponse) Reset()         { *m = BulkAddResponse{} }
func (m *BulkAddResponse) String() string { return proto.CompactTextString(m) }
func (*BulkAddResponse) ProtoMessage()    {}
func (*BulkAddResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b8564a7da00bcc7, []int{5}
}

func (m *BulkAddResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BulkAddResponse.Unmarshal(m, b)
}
func (m *BulkAddResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BulkAddResponse.Marshal(b, m, deterministic)
}
func (m *BulkAddResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BulkAddResponse.Merge(m, src)
}
func (m *BulkAddResponse) XXX_Size() int {
	return xxx_messageInfo_BulkAddResponse.Size(m)
}
func (m *BulkAddResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BulkAddResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BulkAddResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Value)(nil), "copy.Value")
	proto.RegisterType((*BulkItem)(nil), "copy.BulkItem")
	proto.RegisterMapType((map[string]*Value)(nil), "copy.BulkItem.ItemsEntry")
	proto.RegisterType((*CopyItemsRequest)(nil), "copy.CopyItemsRequest")
	proto.RegisterType((*CopyItemsResponse)(nil), "copy.CopyItemsResponse")
	proto.RegisterType((*BulkAddRequest)(nil), "copy.BulkAddRequest")
	proto.RegisterType((*BulkAddResponse)(nil), "copy.BulkAddResponse")
}

func init() { proto.RegisterFile("copy.proto", fileDescriptor_8b8564a7da00bcc7) }

var fileDescriptor_8b8564a7da00bcc7 = []byte{
	// 417 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x65, 0xeb, 0xda, 0xb2, 0xc7, 0x55, 0xdb, 0x2c, 0x6d, 0x31, 0x46, 0x42, 0xae, 0xc5, 0x21,
	0xe2, 0x10, 0xa4, 0x70, 0x41, 0x95, 0x40, 0x2a, 0x50, 0xa4, 0x5c, 0x0d, 0xe2, 0x6a, 0xb9, 0xf5,
	0x20, 0xac, 0x38, 0xf1, 0xe2, 0xdd, 0x24, 0xf2, 0x95, 0x33, 0x9f, 0xc4, 0x4f, 0xf0, 0x47, 0x68,
	0x76, 0x37, 0x71, 0x1c, 0x8e, 0x5c, 0xac, 0x9d, 0x37, 0x33, 0x6f, 0xe6, 0xbd, 0x91, 0x01, 0x1e,
	0x1a, 0xd1, 0x4d, 0x44, 0xdb, 0xa8, 0x86, 0x1f, 0xd3, 0x3b, 0xbd, 0x01, 0xf7, 0x6b, 0x51, 0xaf,
	0x90, 0x3f, 0x83, 0xa0, 0x2c, 0x54, 0x91, 0xab, 0x4e, 0x60, 0xc4, 0x12, 0x36, 0x0e, 0x32, 0x9f,
	0x80, 0x2f, 0x9d, 0x40, 0x7e, 0x01, 0xee, 0x9a, 0xaa, 0xa2, 0x23, 0x9d, 0x30, 0x41, 0xfa, 0x93,
	0x81, 0xff, 0x7e, 0x55, 0xcf, 0x67, 0x0a, 0x17, 0xfc, 0x15, 0xb8, 0x95, 0xc2, 0x85, 0x8c, 0x9c,
	0xc4, 0x19, 0x87, 0xd3, 0xa7, 0x13, 0x3d, 0x6a, 0x9b, 0x9e, 0xd0, 0x47, 0xde, 0x2d, 0x55, 0xdb,
	0x65, 0xa6, 0x2e, 0xbe, 0x03, 0xe8, 0x41, 0x7e, 0x0e, 0xce, 0x1c, 0x3b, 0x3b, 0x98, 0x9e, 0xfc,
	0x7a, 0x7f, 0x66, 0x38, 0x0d, 0x0d, 0xa1, 0x5e, 0xd6, 0x2e, 0x70, 0x73, 0xf4, 0x86, 0xa5, 0x7f,
	0x18, 0x9c, 0x7f, 0x68, 0x44, 0xa7, 0xb9, 0x32, 0xfc, 0xb1, 0x42, 0xa9, 0xf8, 0x25, 0x78, 0x85,
	0x10, 0x79, 0x55, 0x5a, 0x42, 0xb7, 0x10, 0x62, 0x56, 0xf2, 0x6b, 0x38, 0x21, 0x49, 0x52, 0x35,
	0x2d, 0x52, 0xd2, 0xa8, 0x09, 0x77, 0xd8, 0xac, 0xe4, 0xcf, 0x21, 0xa4, 0x39, 0xb9, 0x6d, 0x77,
	0x74, 0x45, 0x40, 0xd0, 0xad, 0xa6, 0x78, 0x09, 0x23, 0x9d, 0x1f, 0xf0, 0x1c, 0xeb, 0xaa, 0x33,
	0x4a, 0x7c, 0xdc, 0xe3, 0x8a, 0x41, 0x3b, 0x78, 0x5f, 0x48, 0x8c, 0xdc, 0xde, 0x51, 0x8a, 0xc9,
	0xee, 0x4d, 0xa5, 0xbe, 0xe7, 0xdf, 0xaa, 0x1a, 0x23, 0x2f, 0x61, 0x63, 0x3f, 0xf3, 0x09, 0xf8,
	0x54, 0xd5, 0x98, 0x3e, 0x86, 0xd1, 0x9e, 0x24, 0x29, 0x9a, 0xa5, 0xc4, 0xf4, 0x37, 0x83, 0x53,
	0xb2, 0xf3, 0xb6, 0x2c, 0xff, 0x5f, 0xe6, 0x8b, 0xe1, 0xb5, 0x4e, 0x87, 0xd7, 0xb2, 0x27, 0xe2,
	0x57, 0xe0, 0x6d, 0xda, 0x4a, 0x61, 0x6b, 0x15, 0xda, 0x88, 0xf0, 0x66, 0xb3, 0xc4, 0x56, 0x46,
	0x6e, 0xe2, 0x10, 0x6e, 0xa2, 0x81, 0x60, 0x6f, 0x28, 0x38, 0x1d, 0xc1, 0xd9, 0x6e, 0x7b, 0xa3,
	0x68, 0xfa, 0x8b, 0x41, 0x48, 0x3a, 0x3f, 0x63, 0xbb, 0xae, 0x1e, 0x90, 0xbf, 0x83, 0x60, 0x27,
	0x9b, 0x5f, 0x99, 0x95, 0x0e, 0x4f, 0x1b, 0x3f, 0xf9, 0x07, 0xb7, 0xfe, 0x3c, 0xe2, 0x6f, 0xe1,
	0xc4, 0x8e, 0x30, 0x14, 0x17, 0xbd, 0xaa, 0xde, 0xb4, 0xf8, 0xf2, 0x00, 0xdd, 0xb6, 0xdf, 0x7b,
	0xfa, 0xbf, 0x78, 0xfd, 0x37, 0x00, 0x00, 0xff, 0xff, 0xca, 0x83, 0x20, 0x36, 0x25, 0x03, 0x00,
	0x00,
}
