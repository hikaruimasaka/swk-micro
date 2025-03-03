// Code generated by protoc-gen-go. DO NOT EDIT.
// source: allow.proto

package allow

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

// 许可类型
type Allow struct {
	AllowId              string    `protobuf:"bytes,2,opt,name=allow_id,json=allowId,proto3" json:"allow_id"`
	AllowName            string    `protobuf:"bytes,3,opt,name=allow_name,json=allowName,proto3" json:"allow_name"`
	AllowType            string    `protobuf:"bytes,4,opt,name=allow_type,json=allowType,proto3" json:"allow_type"`
	ObjectType           string    `protobuf:"bytes,5,opt,name=object_type,json=objectType,proto3" json:"object_type"`
	Actions              []*Action `protobuf:"bytes,6,rep,name=actions,proto3" json:"actions"`
	CreatedAt            string    `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	CreatedBy            string    `protobuf:"bytes,8,opt,name=created_by,json=createdBy,proto3" json:"created_by"`
	UpdatedAt            string    `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	UpdatedBy            string    `protobuf:"bytes,10,opt,name=updated_by,json=updatedBy,proto3" json:"updated_by"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Allow) Reset()         { *m = Allow{} }
func (m *Allow) String() string { return proto.CompactTextString(m) }
func (*Allow) ProtoMessage()    {}
func (*Allow) Descriptor() ([]byte, []int) {
	return fileDescriptor_88d161ea6832a49b, []int{0}
}

func (m *Allow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Allow.Unmarshal(m, b)
}
func (m *Allow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Allow.Marshal(b, m, deterministic)
}
func (m *Allow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Allow.Merge(m, src)
}
func (m *Allow) XXX_Size() int {
	return xxx_messageInfo_Allow.Size(m)
}
func (m *Allow) XXX_DiscardUnknown() {
	xxx_messageInfo_Allow.DiscardUnknown(m)
}

var xxx_messageInfo_Allow proto.InternalMessageInfo

func (m *Allow) GetAllowId() string {
	if m != nil {
		return m.AllowId
	}
	return ""
}

func (m *Allow) GetAllowName() string {
	if m != nil {
		return m.AllowName
	}
	return ""
}

func (m *Allow) GetAllowType() string {
	if m != nil {
		return m.AllowType
	}
	return ""
}

func (m *Allow) GetObjectType() string {
	if m != nil {
		return m.ObjectType
	}
	return ""
}

func (m *Allow) GetActions() []*Action {
	if m != nil {
		return m.Actions
	}
	return nil
}

func (m *Allow) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Allow) GetCreatedBy() string {
	if m != nil {
		return m.CreatedBy
	}
	return ""
}

func (m *Allow) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *Allow) GetUpdatedBy() string {
	if m != nil {
		return m.UpdatedBy
	}
	return ""
}

type Action struct {
	ApiKey               string   `protobuf:"bytes,1,opt,name=api_key,json=apiKey,proto3" json:"api_key"`
	GroupKey             string   `protobuf:"bytes,2,opt,name=group_key,json=groupKey,proto3" json:"group_key"`
	ActionName           string   `protobuf:"bytes,3,opt,name=action_name,json=actionName,proto3" json:"action_name"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Action) Reset()         { *m = Action{} }
func (m *Action) String() string { return proto.CompactTextString(m) }
func (*Action) ProtoMessage()    {}
func (*Action) Descriptor() ([]byte, []int) {
	return fileDescriptor_88d161ea6832a49b, []int{1}
}

func (m *Action) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Action.Unmarshal(m, b)
}
func (m *Action) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Action.Marshal(b, m, deterministic)
}
func (m *Action) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Action.Merge(m, src)
}
func (m *Action) XXX_Size() int {
	return xxx_messageInfo_Action.Size(m)
}
func (m *Action) XXX_DiscardUnknown() {
	xxx_messageInfo_Action.DiscardUnknown(m)
}

var xxx_messageInfo_Action proto.InternalMessageInfo

func (m *Action) GetApiKey() string {
	if m != nil {
		return m.ApiKey
	}
	return ""
}

func (m *Action) GetGroupKey() string {
	if m != nil {
		return m.GroupKey
	}
	return ""
}

func (m *Action) GetActionName() string {
	if m != nil {
		return m.ActionName
	}
	return ""
}

type FindAllowsRequest struct {
	AllowType            string   `protobuf:"bytes,2,opt,name=allow_type,json=allowType,proto3" json:"allow_type"`
	ObjectType           string   `protobuf:"bytes,3,opt,name=object_type,json=objectType,proto3" json:"object_type"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindAllowsRequest) Reset()         { *m = FindAllowsRequest{} }
func (m *FindAllowsRequest) String() string { return proto.CompactTextString(m) }
func (*FindAllowsRequest) ProtoMessage()    {}
func (*FindAllowsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_88d161ea6832a49b, []int{2}
}

func (m *FindAllowsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindAllowsRequest.Unmarshal(m, b)
}
func (m *FindAllowsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindAllowsRequest.Marshal(b, m, deterministic)
}
func (m *FindAllowsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindAllowsRequest.Merge(m, src)
}
func (m *FindAllowsRequest) XXX_Size() int {
	return xxx_messageInfo_FindAllowsRequest.Size(m)
}
func (m *FindAllowsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindAllowsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindAllowsRequest proto.InternalMessageInfo

func (m *FindAllowsRequest) GetAllowType() string {
	if m != nil {
		return m.AllowType
	}
	return ""
}

func (m *FindAllowsRequest) GetObjectType() string {
	if m != nil {
		return m.ObjectType
	}
	return ""
}

type FindAllowsResponse struct {
	Allows               []*Allow `protobuf:"bytes,1,rep,name=allows,proto3" json:"allows"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindAllowsResponse) Reset()         { *m = FindAllowsResponse{} }
func (m *FindAllowsResponse) String() string { return proto.CompactTextString(m) }
func (*FindAllowsResponse) ProtoMessage()    {}
func (*FindAllowsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_88d161ea6832a49b, []int{3}
}

func (m *FindAllowsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindAllowsResponse.Unmarshal(m, b)
}
func (m *FindAllowsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindAllowsResponse.Marshal(b, m, deterministic)
}
func (m *FindAllowsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindAllowsResponse.Merge(m, src)
}
func (m *FindAllowsResponse) XXX_Size() int {
	return xxx_messageInfo_FindAllowsResponse.Size(m)
}
func (m *FindAllowsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindAllowsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindAllowsResponse proto.InternalMessageInfo

func (m *FindAllowsResponse) GetAllows() []*Allow {
	if m != nil {
		return m.Allows
	}
	return nil
}

type FindLevelAllowsRequest struct {
	AllowList            []string `protobuf:"bytes,2,rep,name=allow_list,json=allowList,proto3" json:"allow_list"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindLevelAllowsRequest) Reset()         { *m = FindLevelAllowsRequest{} }
func (m *FindLevelAllowsRequest) String() string { return proto.CompactTextString(m) }
func (*FindLevelAllowsRequest) ProtoMessage()    {}
func (*FindLevelAllowsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_88d161ea6832a49b, []int{4}
}

func (m *FindLevelAllowsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindLevelAllowsRequest.Unmarshal(m, b)
}
func (m *FindLevelAllowsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindLevelAllowsRequest.Marshal(b, m, deterministic)
}
func (m *FindLevelAllowsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindLevelAllowsRequest.Merge(m, src)
}
func (m *FindLevelAllowsRequest) XXX_Size() int {
	return xxx_messageInfo_FindLevelAllowsRequest.Size(m)
}
func (m *FindLevelAllowsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindLevelAllowsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindLevelAllowsRequest proto.InternalMessageInfo

func (m *FindLevelAllowsRequest) GetAllowList() []string {
	if m != nil {
		return m.AllowList
	}
	return nil
}

type FindLevelAllowsResponse struct {
	Allows               []*Allow `protobuf:"bytes,1,rep,name=allows,proto3" json:"allows"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindLevelAllowsResponse) Reset()         { *m = FindLevelAllowsResponse{} }
func (m *FindLevelAllowsResponse) String() string { return proto.CompactTextString(m) }
func (*FindLevelAllowsResponse) ProtoMessage()    {}
func (*FindLevelAllowsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_88d161ea6832a49b, []int{5}
}

func (m *FindLevelAllowsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindLevelAllowsResponse.Unmarshal(m, b)
}
func (m *FindLevelAllowsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindLevelAllowsResponse.Marshal(b, m, deterministic)
}
func (m *FindLevelAllowsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindLevelAllowsResponse.Merge(m, src)
}
func (m *FindLevelAllowsResponse) XXX_Size() int {
	return xxx_messageInfo_FindLevelAllowsResponse.Size(m)
}
func (m *FindLevelAllowsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindLevelAllowsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindLevelAllowsResponse proto.InternalMessageInfo

func (m *FindLevelAllowsResponse) GetAllows() []*Allow {
	if m != nil {
		return m.Allows
	}
	return nil
}

type FindAllowRequest struct {
	AllowId              string   `protobuf:"bytes,1,opt,name=allow_id,json=allowId,proto3" json:"allow_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindAllowRequest) Reset()         { *m = FindAllowRequest{} }
func (m *FindAllowRequest) String() string { return proto.CompactTextString(m) }
func (*FindAllowRequest) ProtoMessage()    {}
func (*FindAllowRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_88d161ea6832a49b, []int{6}
}

func (m *FindAllowRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindAllowRequest.Unmarshal(m, b)
}
func (m *FindAllowRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindAllowRequest.Marshal(b, m, deterministic)
}
func (m *FindAllowRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindAllowRequest.Merge(m, src)
}
func (m *FindAllowRequest) XXX_Size() int {
	return xxx_messageInfo_FindAllowRequest.Size(m)
}
func (m *FindAllowRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindAllowRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindAllowRequest proto.InternalMessageInfo

func (m *FindAllowRequest) GetAllowId() string {
	if m != nil {
		return m.AllowId
	}
	return ""
}

type FindAllowResponse struct {
	Allow                *Allow   `protobuf:"bytes,1,opt,name=allow,proto3" json:"allow"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindAllowResponse) Reset()         { *m = FindAllowResponse{} }
func (m *FindAllowResponse) String() string { return proto.CompactTextString(m) }
func (*FindAllowResponse) ProtoMessage()    {}
func (*FindAllowResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_88d161ea6832a49b, []int{7}
}

func (m *FindAllowResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindAllowResponse.Unmarshal(m, b)
}
func (m *FindAllowResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindAllowResponse.Marshal(b, m, deterministic)
}
func (m *FindAllowResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindAllowResponse.Merge(m, src)
}
func (m *FindAllowResponse) XXX_Size() int {
	return xxx_messageInfo_FindAllowResponse.Size(m)
}
func (m *FindAllowResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindAllowResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindAllowResponse proto.InternalMessageInfo

func (m *FindAllowResponse) GetAllow() *Allow {
	if m != nil {
		return m.Allow
	}
	return nil
}

type AddAllowRequest struct {
	AllowName            string    `protobuf:"bytes,2,opt,name=allow_name,json=allowName,proto3" json:"allow_name"`
	AllowType            string    `protobuf:"bytes,3,opt,name=allow_type,json=allowType,proto3" json:"allow_type"`
	ObjectType           string    `protobuf:"bytes,4,opt,name=object_type,json=objectType,proto3" json:"object_type"`
	Actions              []*Action `protobuf:"bytes,5,rep,name=actions,proto3" json:"actions"`
	Writer               string    `protobuf:"bytes,6,opt,name=writer,proto3" json:"writer"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *AddAllowRequest) Reset()         { *m = AddAllowRequest{} }
func (m *AddAllowRequest) String() string { return proto.CompactTextString(m) }
func (*AddAllowRequest) ProtoMessage()    {}
func (*AddAllowRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_88d161ea6832a49b, []int{8}
}

func (m *AddAllowRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddAllowRequest.Unmarshal(m, b)
}
func (m *AddAllowRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddAllowRequest.Marshal(b, m, deterministic)
}
func (m *AddAllowRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddAllowRequest.Merge(m, src)
}
func (m *AddAllowRequest) XXX_Size() int {
	return xxx_messageInfo_AddAllowRequest.Size(m)
}
func (m *AddAllowRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddAllowRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddAllowRequest proto.InternalMessageInfo

func (m *AddAllowRequest) GetAllowName() string {
	if m != nil {
		return m.AllowName
	}
	return ""
}

func (m *AddAllowRequest) GetAllowType() string {
	if m != nil {
		return m.AllowType
	}
	return ""
}

func (m *AddAllowRequest) GetObjectType() string {
	if m != nil {
		return m.ObjectType
	}
	return ""
}

func (m *AddAllowRequest) GetActions() []*Action {
	if m != nil {
		return m.Actions
	}
	return nil
}

func (m *AddAllowRequest) GetWriter() string {
	if m != nil {
		return m.Writer
	}
	return ""
}

type AddAllowResponse struct {
	AllowId              string   `protobuf:"bytes,1,opt,name=allow_id,json=allowId,proto3" json:"allow_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddAllowResponse) Reset()         { *m = AddAllowResponse{} }
func (m *AddAllowResponse) String() string { return proto.CompactTextString(m) }
func (*AddAllowResponse) ProtoMessage()    {}
func (*AddAllowResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_88d161ea6832a49b, []int{9}
}

func (m *AddAllowResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddAllowResponse.Unmarshal(m, b)
}
func (m *AddAllowResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddAllowResponse.Marshal(b, m, deterministic)
}
func (m *AddAllowResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddAllowResponse.Merge(m, src)
}
func (m *AddAllowResponse) XXX_Size() int {
	return xxx_messageInfo_AddAllowResponse.Size(m)
}
func (m *AddAllowResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddAllowResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddAllowResponse proto.InternalMessageInfo

func (m *AddAllowResponse) GetAllowId() string {
	if m != nil {
		return m.AllowId
	}
	return ""
}

type ModifyAllowRequest struct {
	AllowId              string    `protobuf:"bytes,2,opt,name=allow_id,json=allowId,proto3" json:"allow_id"`
	AllowName            string    `protobuf:"bytes,3,opt,name=allow_name,json=allowName,proto3" json:"allow_name"`
	AllowType            string    `protobuf:"bytes,4,opt,name=allow_type,json=allowType,proto3" json:"allow_type"`
	ObjectType           string    `protobuf:"bytes,5,opt,name=object_type,json=objectType,proto3" json:"object_type"`
	Actions              []*Action `protobuf:"bytes,6,rep,name=actions,proto3" json:"actions"`
	Writer               string    `protobuf:"bytes,7,opt,name=writer,proto3" json:"writer"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ModifyAllowRequest) Reset()         { *m = ModifyAllowRequest{} }
func (m *ModifyAllowRequest) String() string { return proto.CompactTextString(m) }
func (*ModifyAllowRequest) ProtoMessage()    {}
func (*ModifyAllowRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_88d161ea6832a49b, []int{10}
}

func (m *ModifyAllowRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModifyAllowRequest.Unmarshal(m, b)
}
func (m *ModifyAllowRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModifyAllowRequest.Marshal(b, m, deterministic)
}
func (m *ModifyAllowRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModifyAllowRequest.Merge(m, src)
}
func (m *ModifyAllowRequest) XXX_Size() int {
	return xxx_messageInfo_ModifyAllowRequest.Size(m)
}
func (m *ModifyAllowRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ModifyAllowRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ModifyAllowRequest proto.InternalMessageInfo

func (m *ModifyAllowRequest) GetAllowId() string {
	if m != nil {
		return m.AllowId
	}
	return ""
}

func (m *ModifyAllowRequest) GetAllowName() string {
	if m != nil {
		return m.AllowName
	}
	return ""
}

func (m *ModifyAllowRequest) GetAllowType() string {
	if m != nil {
		return m.AllowType
	}
	return ""
}

func (m *ModifyAllowRequest) GetObjectType() string {
	if m != nil {
		return m.ObjectType
	}
	return ""
}

func (m *ModifyAllowRequest) GetActions() []*Action {
	if m != nil {
		return m.Actions
	}
	return nil
}

func (m *ModifyAllowRequest) GetWriter() string {
	if m != nil {
		return m.Writer
	}
	return ""
}

type ModifyAllowResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ModifyAllowResponse) Reset()         { *m = ModifyAllowResponse{} }
func (m *ModifyAllowResponse) String() string { return proto.CompactTextString(m) }
func (*ModifyAllowResponse) ProtoMessage()    {}
func (*ModifyAllowResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_88d161ea6832a49b, []int{11}
}

func (m *ModifyAllowResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModifyAllowResponse.Unmarshal(m, b)
}
func (m *ModifyAllowResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModifyAllowResponse.Marshal(b, m, deterministic)
}
func (m *ModifyAllowResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModifyAllowResponse.Merge(m, src)
}
func (m *ModifyAllowResponse) XXX_Size() int {
	return xxx_messageInfo_ModifyAllowResponse.Size(m)
}
func (m *ModifyAllowResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ModifyAllowResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ModifyAllowResponse proto.InternalMessageInfo

type DeleteAllowRequest struct {
	AllowId              string   `protobuf:"bytes,1,opt,name=allow_id,json=allowId,proto3" json:"allow_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteAllowRequest) Reset()         { *m = DeleteAllowRequest{} }
func (m *DeleteAllowRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteAllowRequest) ProtoMessage()    {}
func (*DeleteAllowRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_88d161ea6832a49b, []int{12}
}

func (m *DeleteAllowRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteAllowRequest.Unmarshal(m, b)
}
func (m *DeleteAllowRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteAllowRequest.Marshal(b, m, deterministic)
}
func (m *DeleteAllowRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteAllowRequest.Merge(m, src)
}
func (m *DeleteAllowRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteAllowRequest.Size(m)
}
func (m *DeleteAllowRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteAllowRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteAllowRequest proto.InternalMessageInfo

func (m *DeleteAllowRequest) GetAllowId() string {
	if m != nil {
		return m.AllowId
	}
	return ""
}

type DeleteAllowResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteAllowResponse) Reset()         { *m = DeleteAllowResponse{} }
func (m *DeleteAllowResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteAllowResponse) ProtoMessage()    {}
func (*DeleteAllowResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_88d161ea6832a49b, []int{13}
}

func (m *DeleteAllowResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteAllowResponse.Unmarshal(m, b)
}
func (m *DeleteAllowResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteAllowResponse.Marshal(b, m, deterministic)
}
func (m *DeleteAllowResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteAllowResponse.Merge(m, src)
}
func (m *DeleteAllowResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteAllowResponse.Size(m)
}
func (m *DeleteAllowResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteAllowResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteAllowResponse proto.InternalMessageInfo

type DeleteAllowsRequest struct {
	AllowIds             []string `protobuf:"bytes,1,rep,name=allow_ids,json=allowIds,proto3" json:"allow_ids"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteAllowsRequest) Reset()         { *m = DeleteAllowsRequest{} }
func (m *DeleteAllowsRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteAllowsRequest) ProtoMessage()    {}
func (*DeleteAllowsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_88d161ea6832a49b, []int{14}
}

func (m *DeleteAllowsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteAllowsRequest.Unmarshal(m, b)
}
func (m *DeleteAllowsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteAllowsRequest.Marshal(b, m, deterministic)
}
func (m *DeleteAllowsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteAllowsRequest.Merge(m, src)
}
func (m *DeleteAllowsRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteAllowsRequest.Size(m)
}
func (m *DeleteAllowsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteAllowsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteAllowsRequest proto.InternalMessageInfo

func (m *DeleteAllowsRequest) GetAllowIds() []string {
	if m != nil {
		return m.AllowIds
	}
	return nil
}

type DeleteAllowsResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteAllowsResponse) Reset()         { *m = DeleteAllowsResponse{} }
func (m *DeleteAllowsResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteAllowsResponse) ProtoMessage()    {}
func (*DeleteAllowsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_88d161ea6832a49b, []int{15}
}

func (m *DeleteAllowsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteAllowsResponse.Unmarshal(m, b)
}
func (m *DeleteAllowsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteAllowsResponse.Marshal(b, m, deterministic)
}
func (m *DeleteAllowsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteAllowsResponse.Merge(m, src)
}
func (m *DeleteAllowsResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteAllowsResponse.Size(m)
}
func (m *DeleteAllowsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteAllowsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteAllowsResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Allow)(nil), "allow.Allow")
	proto.RegisterType((*Action)(nil), "allow.Action")
	proto.RegisterType((*FindAllowsRequest)(nil), "allow.FindAllowsRequest")
	proto.RegisterType((*FindAllowsResponse)(nil), "allow.FindAllowsResponse")
	proto.RegisterType((*FindLevelAllowsRequest)(nil), "allow.FindLevelAllowsRequest")
	proto.RegisterType((*FindLevelAllowsResponse)(nil), "allow.FindLevelAllowsResponse")
	proto.RegisterType((*FindAllowRequest)(nil), "allow.FindAllowRequest")
	proto.RegisterType((*FindAllowResponse)(nil), "allow.FindAllowResponse")
	proto.RegisterType((*AddAllowRequest)(nil), "allow.AddAllowRequest")
	proto.RegisterType((*AddAllowResponse)(nil), "allow.AddAllowResponse")
	proto.RegisterType((*ModifyAllowRequest)(nil), "allow.ModifyAllowRequest")
	proto.RegisterType((*ModifyAllowResponse)(nil), "allow.ModifyAllowResponse")
	proto.RegisterType((*DeleteAllowRequest)(nil), "allow.DeleteAllowRequest")
	proto.RegisterType((*DeleteAllowResponse)(nil), "allow.DeleteAllowResponse")
	proto.RegisterType((*DeleteAllowsRequest)(nil), "allow.DeleteAllowsRequest")
	proto.RegisterType((*DeleteAllowsResponse)(nil), "allow.DeleteAllowsResponse")
}

func init() { proto.RegisterFile("allow.proto", fileDescriptor_88d161ea6832a49b) }

var fileDescriptor_88d161ea6832a49b = []byte{
	// 613 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x55, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xad, 0xf3, 0xe1, 0xc4, 0x93, 0xa0, 0xc2, 0x96, 0x26, 0xdb, 0x44, 0xa5, 0x91, 0x85, 0x44,
	0x2f, 0x14, 0x29, 0x1c, 0x2a, 0x21, 0x21, 0x48, 0x41, 0x95, 0x2a, 0x0a, 0x07, 0x97, 0x7b, 0xe4,
	0xc4, 0x03, 0x32, 0xa4, 0xb1, 0xb1, 0x37, 0xad, 0x7c, 0xe6, 0xbf, 0x70, 0xe7, 0x97, 0xf0, 0x97,
	0x90, 0x77, 0xbd, 0x9b, 0xdd, 0x38, 0x49, 0xe1, 0xd6, 0x4b, 0x94, 0x9d, 0xb7, 0xf3, 0x66, 0xf6,
	0xcd, 0x1b, 0x19, 0x5a, 0xfe, 0x6c, 0x16, 0xdd, 0x9e, 0xc4, 0x49, 0xc4, 0x22, 0x52, 0xe7, 0x07,
	0xf7, 0x57, 0x05, 0xea, 0xa3, 0xfc, 0x1f, 0x39, 0x80, 0x26, 0x0f, 0x8d, 0xc3, 0x80, 0x56, 0x06,
	0xd6, 0xb1, 0xe3, 0x35, 0xf8, 0xf9, 0x22, 0x20, 0x87, 0x00, 0x02, 0x9a, 0xfb, 0xd7, 0x48, 0xab,
	0x1c, 0x74, 0x78, 0xe4, 0x93, 0x7f, 0x8d, 0x4b, 0x98, 0x65, 0x31, 0xd2, 0x9a, 0x06, 0x7f, 0xce,
	0x62, 0x24, 0x47, 0xd0, 0x8a, 0x26, 0xdf, 0x70, 0xca, 0x04, 0x5e, 0xe7, 0x38, 0x88, 0x10, 0xbf,
	0xf0, 0x0c, 0x1a, 0xfe, 0x94, 0x85, 0xd1, 0x3c, 0xa5, 0xf6, 0xa0, 0x7a, 0xdc, 0x1a, 0x3e, 0x38,
	0x11, 0x9d, 0x8e, 0x78, 0xd4, 0x93, 0x68, 0x5e, 0x68, 0x9a, 0xa0, 0xcf, 0x30, 0x18, 0xfb, 0x8c,
	0x36, 0x44, 0xa1, 0x22, 0x32, 0x62, 0x3a, 0x3c, 0xc9, 0x68, 0xd3, 0x80, 0xcf, 0xb2, 0x1c, 0x5e,
	0xc4, 0x81, 0xcc, 0x76, 0x04, 0x5c, 0x44, 0x44, 0xb6, 0x84, 0x27, 0x19, 0x05, 0x03, 0x3e, 0xcb,
	0xdc, 0x31, 0xd8, 0xa2, 0x1d, 0xd2, 0x85, 0x86, 0x1f, 0x87, 0xe3, 0xef, 0x98, 0x51, 0x8b, 0xdf,
	0xb2, 0xfd, 0x38, 0xfc, 0x80, 0x19, 0xe9, 0x83, 0xf3, 0x35, 0x89, 0x16, 0x31, 0x87, 0x84, 0x84,
	0x4d, 0x1e, 0xc8, 0xc1, 0x23, 0x68, 0x89, 0x67, 0xe8, 0x22, 0x82, 0x08, 0xe5, 0x2a, 0xba, 0x57,
	0xf0, 0xe8, 0x3c, 0x9c, 0x07, 0x7c, 0x18, 0xa9, 0x87, 0x3f, 0x16, 0x98, 0xb2, 0x15, 0x69, 0x2b,
	0x77, 0x48, 0x5b, 0x5d, 0x95, 0xd6, 0x7d, 0x05, 0x44, 0x27, 0x4d, 0xe3, 0x68, 0x9e, 0x22, 0x79,
	0x0a, 0x36, 0xe7, 0x48, 0xa9, 0xc5, 0xf5, 0x6e, 0x4b, 0xbd, 0xf3, 0x5f, 0xaf, 0xc0, 0xdc, 0x53,
	0xe8, 0xe4, 0xb9, 0x97, 0x78, 0x83, 0xb3, 0x0d, 0x5d, 0xcd, 0xc2, 0x94, 0xd1, 0xca, 0xa0, 0xaa,
	0xba, 0xba, 0x0c, 0x53, 0xe6, 0xbe, 0x81, 0x6e, 0x29, 0xf1, 0xbf, 0x2a, 0x3f, 0x87, 0x87, 0xaa,
	0x6b, 0x59, 0x53, 0xb7, 0xa7, 0x65, 0xd8, 0xd3, 0x3d, 0xd5, 0x94, 0x53, 0x95, 0x5c, 0x10, 0x0e,
	0xe7, 0x97, 0x57, 0x0b, 0x15, 0xe6, 0xff, 0x6d, 0xc1, 0xee, 0x28, 0x30, 0xeb, 0x98, 0x5e, 0xaf,
	0x6c, 0xf7, 0x7a, 0xf5, 0x8e, 0x81, 0xd4, 0xb6, 0x79, 0xbd, 0xbe, 0xd5, 0xeb, 0x1d, 0xb0, 0x6f,
	0x93, 0x90, 0x61, 0x42, 0x6d, 0x61, 0x32, 0x71, 0xca, 0xb5, 0x59, 0xb6, 0x5c, 0xbc, 0x75, 0x8b,
	0x36, 0x7f, 0x2c, 0x20, 0x1f, 0xa3, 0x20, 0xfc, 0x92, 0x6d, 0x54, 0xf3, 0x9e, 0x2e, 0xfb, 0x52,
	0x80, 0x86, 0x21, 0xc0, 0x3e, 0xec, 0x19, 0x0f, 0x12, 0x1a, 0xb8, 0x2f, 0x80, 0xbc, 0xc7, 0x19,
	0x32, 0xfc, 0x57, 0xd7, 0xec, 0xc3, 0x9e, 0x91, 0x50, 0xf0, 0x0c, 0x8d, 0xb0, 0xb2, 0x7c, 0x1f,
	0x1c, 0x49, 0x24, 0xbc, 0xeb, 0x78, 0xcd, 0x82, 0x29, 0x75, 0x3b, 0xf0, 0xd8, 0xcc, 0x11, 0x5c,
	0xc3, 0x9f, 0x35, 0x68, 0xf3, 0xd0, 0x15, 0x26, 0x37, 0xe1, 0x14, 0xc9, 0x3b, 0x80, 0xe5, 0x3a,
	0x12, 0x5a, 0xbc, 0xbc, 0xb4, 0xf6, 0xbd, 0x83, 0x35, 0x48, 0xd1, 0xdf, 0x0e, 0xf1, 0x60, 0x77,
	0x65, 0xbd, 0xc8, 0xa1, 0x76, 0xbf, 0xbc, 0xaf, 0xbd, 0x27, 0x9b, 0x60, 0xc5, 0xf9, 0x16, 0x1c,
	0x55, 0x8b, 0x74, 0x57, 0xab, 0x4b, 0x1e, 0x5a, 0x06, 0x14, 0xc3, 0x6b, 0x68, 0x4a, 0x5f, 0x92,
	0x8e, 0x1c, 0xa9, 0xb9, 0x5b, 0xbd, 0x6e, 0x29, 0xae, 0xd2, 0xcf, 0xa1, 0xa5, 0x4d, 0x95, 0x48,
	0x01, 0xca, 0xd6, 0xed, 0xf5, 0xd6, 0x41, 0x3a, 0x8f, 0x36, 0x0a, 0xc5, 0x53, 0xb6, 0x86, 0xe2,
	0x59, 0x67, 0x82, 0x1d, 0x72, 0x01, 0x6d, 0x7d, 0xa4, 0x64, 0xcd, 0x6d, 0x25, 0x6f, 0x7f, 0x2d,
	0x26, 0xa9, 0x26, 0x36, 0xff, 0xe0, 0xbe, 0xfc, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xba, 0x0b, 0xbe,
	0x67, 0x7f, 0x07, 0x00, 0x00,
}
