// Code generated by protoc-gen-go. DO NOT EDIT.
// source: relation.proto

package relation

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

// 流程实例定义
type Relation struct {
	AppId                string   `protobuf:"bytes,1,opt,name=app_id,json=appId,proto3" json:"app_id"`
	ObjectId             string   `protobuf:"bytes,2,opt,name=object_id,json=objectId,proto3" json:"object_id"`
	GroupId              string   `protobuf:"bytes,3,opt,name=group_id,json=groupId,proto3" json:"group_id"`
	WorkflowId           string   `protobuf:"bytes,4,opt,name=workflow_id,json=workflowId,proto3" json:"workflow_id"`
	Action               string   `protobuf:"bytes,5,opt,name=action,proto3" json:"action"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Relation) Reset()         { *m = Relation{} }
func (m *Relation) String() string { return proto.CompactTextString(m) }
func (*Relation) ProtoMessage()    {}
func (*Relation) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc8ceb4d3910fa25, []int{0}
}

func (m *Relation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Relation.Unmarshal(m, b)
}
func (m *Relation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Relation.Marshal(b, m, deterministic)
}
func (m *Relation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Relation.Merge(m, src)
}
func (m *Relation) XXX_Size() int {
	return xxx_messageInfo_Relation.Size(m)
}
func (m *Relation) XXX_DiscardUnknown() {
	xxx_messageInfo_Relation.DiscardUnknown(m)
}

var xxx_messageInfo_Relation proto.InternalMessageInfo

func (m *Relation) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *Relation) GetObjectId() string {
	if m != nil {
		return m.ObjectId
	}
	return ""
}

func (m *Relation) GetGroupId() string {
	if m != nil {
		return m.GroupId
	}
	return ""
}

func (m *Relation) GetWorkflowId() string {
	if m != nil {
		return m.WorkflowId
	}
	return ""
}

func (m *Relation) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

// 查找多条记录
type RelationsRequest struct {
	AppId                string   `protobuf:"bytes,1,opt,name=app_id,json=appId,proto3" json:"app_id"`
	ObjectId             string   `protobuf:"bytes,2,opt,name=object_id,json=objectId,proto3" json:"object_id"`
	GroupId              string   `protobuf:"bytes,3,opt,name=group_id,json=groupId,proto3" json:"group_id"`
	WorkflowId           string   `protobuf:"bytes,4,opt,name=workflow_id,json=workflowId,proto3" json:"workflow_id"`
	Action               string   `protobuf:"bytes,5,opt,name=action,proto3" json:"action"`
	Database             string   `protobuf:"bytes,6,opt,name=database,proto3" json:"database"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RelationsRequest) Reset()         { *m = RelationsRequest{} }
func (m *RelationsRequest) String() string { return proto.CompactTextString(m) }
func (*RelationsRequest) ProtoMessage()    {}
func (*RelationsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc8ceb4d3910fa25, []int{1}
}

func (m *RelationsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RelationsRequest.Unmarshal(m, b)
}
func (m *RelationsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RelationsRequest.Marshal(b, m, deterministic)
}
func (m *RelationsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RelationsRequest.Merge(m, src)
}
func (m *RelationsRequest) XXX_Size() int {
	return xxx_messageInfo_RelationsRequest.Size(m)
}
func (m *RelationsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RelationsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RelationsRequest proto.InternalMessageInfo

func (m *RelationsRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *RelationsRequest) GetObjectId() string {
	if m != nil {
		return m.ObjectId
	}
	return ""
}

func (m *RelationsRequest) GetGroupId() string {
	if m != nil {
		return m.GroupId
	}
	return ""
}

func (m *RelationsRequest) GetWorkflowId() string {
	if m != nil {
		return m.WorkflowId
	}
	return ""
}

func (m *RelationsRequest) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *RelationsRequest) GetDatabase() string {
	if m != nil {
		return m.Database
	}
	return ""
}

type RelationsResponse struct {
	Relations            []*Relation `protobuf:"bytes,1,rep,name=relations,proto3" json:"relations"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *RelationsResponse) Reset()         { *m = RelationsResponse{} }
func (m *RelationsResponse) String() string { return proto.CompactTextString(m) }
func (*RelationsResponse) ProtoMessage()    {}
func (*RelationsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc8ceb4d3910fa25, []int{2}
}

func (m *RelationsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RelationsResponse.Unmarshal(m, b)
}
func (m *RelationsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RelationsResponse.Marshal(b, m, deterministic)
}
func (m *RelationsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RelationsResponse.Merge(m, src)
}
func (m *RelationsResponse) XXX_Size() int {
	return xxx_messageInfo_RelationsResponse.Size(m)
}
func (m *RelationsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RelationsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RelationsResponse proto.InternalMessageInfo

func (m *RelationsResponse) GetRelations() []*Relation {
	if m != nil {
		return m.Relations
	}
	return nil
}

// 添加数据
type AddRequest struct {
	AppId                string   `protobuf:"bytes,1,opt,name=app_id,json=appId,proto3" json:"app_id"`
	ObjectId             string   `protobuf:"bytes,2,opt,name=object_id,json=objectId,proto3" json:"object_id"`
	GroupId              string   `protobuf:"bytes,3,opt,name=group_id,json=groupId,proto3" json:"group_id"`
	WorkflowId           string   `protobuf:"bytes,4,opt,name=workflow_id,json=workflowId,proto3" json:"workflow_id"`
	Action               string   `protobuf:"bytes,5,opt,name=action,proto3" json:"action"`
	Database             string   `protobuf:"bytes,6,opt,name=database,proto3" json:"database"`
	Writer               string   `protobuf:"bytes,7,opt,name=writer,proto3" json:"writer"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddRequest) Reset()         { *m = AddRequest{} }
func (m *AddRequest) String() string { return proto.CompactTextString(m) }
func (*AddRequest) ProtoMessage()    {}
func (*AddRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc8ceb4d3910fa25, []int{3}
}

func (m *AddRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddRequest.Unmarshal(m, b)
}
func (m *AddRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddRequest.Marshal(b, m, deterministic)
}
func (m *AddRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddRequest.Merge(m, src)
}
func (m *AddRequest) XXX_Size() int {
	return xxx_messageInfo_AddRequest.Size(m)
}
func (m *AddRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddRequest proto.InternalMessageInfo

func (m *AddRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *AddRequest) GetObjectId() string {
	if m != nil {
		return m.ObjectId
	}
	return ""
}

func (m *AddRequest) GetGroupId() string {
	if m != nil {
		return m.GroupId
	}
	return ""
}

func (m *AddRequest) GetWorkflowId() string {
	if m != nil {
		return m.WorkflowId
	}
	return ""
}

func (m *AddRequest) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *AddRequest) GetDatabase() string {
	if m != nil {
		return m.Database
	}
	return ""
}

func (m *AddRequest) GetWriter() string {
	if m != nil {
		return m.Writer
	}
	return ""
}

type AddResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddResponse) Reset()         { *m = AddResponse{} }
func (m *AddResponse) String() string { return proto.CompactTextString(m) }
func (*AddResponse) ProtoMessage()    {}
func (*AddResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc8ceb4d3910fa25, []int{4}
}

func (m *AddResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddResponse.Unmarshal(m, b)
}
func (m *AddResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddResponse.Marshal(b, m, deterministic)
}
func (m *AddResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddResponse.Merge(m, src)
}
func (m *AddResponse) XXX_Size() int {
	return xxx_messageInfo_AddResponse.Size(m)
}
func (m *AddResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddResponse proto.InternalMessageInfo

func (m *AddResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// 删除数据记录
type DeleteRequest struct {
	AppId                string   `protobuf:"bytes,1,opt,name=app_id,json=appId,proto3" json:"app_id"`
	ObjectId             string   `protobuf:"bytes,2,opt,name=object_id,json=objectId,proto3" json:"object_id"`
	GroupId              string   `protobuf:"bytes,3,opt,name=group_id,json=groupId,proto3" json:"group_id"`
	WorkflowId           string   `protobuf:"bytes,4,opt,name=workflow_id,json=workflowId,proto3" json:"workflow_id"`
	Database             string   `protobuf:"bytes,5,opt,name=database,proto3" json:"database"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc8ceb4d3910fa25, []int{5}
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

func (m *DeleteRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *DeleteRequest) GetObjectId() string {
	if m != nil {
		return m.ObjectId
	}
	return ""
}

func (m *DeleteRequest) GetGroupId() string {
	if m != nil {
		return m.GroupId
	}
	return ""
}

func (m *DeleteRequest) GetWorkflowId() string {
	if m != nil {
		return m.WorkflowId
	}
	return ""
}

func (m *DeleteRequest) GetDatabase() string {
	if m != nil {
		return m.Database
	}
	return ""
}

type DeleteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc8ceb4d3910fa25, []int{6}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Relation)(nil), "relation.Relation")
	proto.RegisterType((*RelationsRequest)(nil), "relation.RelationsRequest")
	proto.RegisterType((*RelationsResponse)(nil), "relation.RelationsResponse")
	proto.RegisterType((*AddRequest)(nil), "relation.AddRequest")
	proto.RegisterType((*AddResponse)(nil), "relation.AddResponse")
	proto.RegisterType((*DeleteRequest)(nil), "relation.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "relation.DeleteResponse")
}

func init() { proto.RegisterFile("relation.proto", fileDescriptor_fc8ceb4d3910fa25) }

var fileDescriptor_fc8ceb4d3910fa25 = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x54, 0x41, 0x4e, 0xc2, 0x40,
	0x14, 0x75, 0x40, 0x4a, 0xf9, 0x04, 0xc4, 0x89, 0x68, 0x2d, 0x31, 0x92, 0x59, 0xb1, 0x22, 0x06,
	0xb7, 0x6e, 0x4c, 0xc4, 0xa4, 0x2e, 0xeb, 0x01, 0xcc, 0xc0, 0x7c, 0x4d, 0x95, 0x30, 0x75, 0x3a,
	0xd8, 0x7b, 0xb8, 0xf5, 0x1e, 0x9e, 0xc1, 0x73, 0x78, 0x12, 0xd3, 0x69, 0xa7, 0x45, 0x70, 0xaf,
	0x2e, 0xdf, 0xbc, 0xff, 0xff, 0xbc, 0xf7, 0xdf, 0xb4, 0xd0, 0x55, 0xb8, 0xe0, 0x3a, 0x92, 0xcb,
	0x71, 0xac, 0xa4, 0x96, 0xd4, 0xb5, 0x98, 0xbd, 0x12, 0x70, 0xc3, 0x02, 0xd0, 0x3e, 0x38, 0x3c,
	0x8e, 0xef, 0x22, 0xe1, 0x91, 0x21, 0x19, 0xb5, 0xc2, 0x06, 0x8f, 0xe3, 0x40, 0xd0, 0x01, 0xb4,
	0xe4, 0xec, 0x11, 0xe7, 0x3a, 0x63, 0x6a, 0x86, 0x71, 0xf3, 0x83, 0x40, 0xd0, 0x63, 0x70, 0x1f,
	0x94, 0x5c, 0x99, 0xae, 0xba, 0xe1, 0x9a, 0x06, 0x07, 0x82, 0x9e, 0x42, 0x3b, 0x95, 0xea, 0xe9,
	0x7e, 0x21, 0xd3, 0x8c, 0xdd, 0x35, 0x2c, 0xd8, 0xa3, 0x40, 0xd0, 0x43, 0x70, 0xf8, 0x3c, 0xbb,
	0xd9, 0x6b, 0x18, 0xae, 0x40, 0xec, 0x9d, 0x40, 0xcf, 0x8a, 0x4a, 0x42, 0x7c, 0x5e, 0x61, 0xa2,
	0xff, 0x8a, 0x38, 0xea, 0x83, 0x2b, 0xb8, 0xe6, 0x33, 0x9e, 0xa0, 0xe7, 0xe4, 0xf7, 0x59, 0xcc,
	0xa6, 0xb0, 0xbf, 0xa6, 0x3b, 0x89, 0xe5, 0x32, 0x41, 0x7a, 0x06, 0x2d, 0xbb, 0xee, 0xc4, 0x23,
	0xc3, 0xfa, 0xa8, 0x3d, 0xa1, 0xe3, 0x32, 0x10, 0x5b, 0x1f, 0x56, 0x45, 0xec, 0x83, 0x00, 0x5c,
	0x0a, 0xf1, 0x8f, 0x9c, 0x67, 0x3d, 0xa9, 0x8a, 0x34, 0x2a, 0xaf, 0x99, 0xf7, 0xe4, 0x88, 0x9d,
	0x40, 0xdb, 0x38, 0x29, 0x76, 0xd1, 0x85, 0x5a, 0x69, 0xa3, 0x16, 0x09, 0xf6, 0x46, 0xa0, 0x73,
	0x85, 0x0b, 0xd4, 0xf8, 0x5b, 0x66, 0xd7, 0x4d, 0x35, 0x36, 0xe2, 0xec, 0x41, 0xd7, 0x8a, 0xcb,
	0xf5, 0x4f, 0x3e, 0x09, 0xec, 0xd9, 0xc4, 0x6e, 0x51, 0xbd, 0x44, 0x73, 0xa4, 0x37, 0xd0, 0xb9,
	0x8e, 0x96, 0xa2, 0x0c, 0x9e, 0xfa, 0xdb, 0xe9, 0xda, 0x57, 0xec, 0x0f, 0x7e, 0xe4, 0xf2, 0xe9,
	0x6c, 0x87, 0x5e, 0x14, 0xeb, 0x2a, 0x3e, 0xc8, 0x83, 0xaa, 0xba, 0x7a, 0x0f, 0x7e, 0x7f, 0xe3,
	0xb4, 0xec, 0x9e, 0x56, 0x7a, 0x8b, 0x01, 0x47, 0x55, 0xe9, 0xb7, 0x35, 0xfb, 0xde, 0x36, 0x61,
	0xc7, 0xcc, 0x1c, 0xf3, 0x93, 0x38, 0xff, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x70, 0xee, 0x23, 0x84,
	0x36, 0x04, 0x00, 0x00,
}
