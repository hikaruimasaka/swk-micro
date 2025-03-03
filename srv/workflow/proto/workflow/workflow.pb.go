// Code generated by protoc-gen-go. DO NOT EDIT.
// source: workflow.proto

package workflow

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

// 流程定义
type Workflow struct {
	WfId                 string            `protobuf:"bytes,1,opt,name=wf_id,json=wfId,proto3" json:"wf_id"`
	WfName               string            `protobuf:"bytes,2,opt,name=wf_name,json=wfName,proto3" json:"wf_name"`
	MenuName             string            `protobuf:"bytes,3,opt,name=menu_name,json=menuName,proto3" json:"menu_name"`
	IsValid              bool              `protobuf:"varint,4,opt,name=is_valid,json=isValid,proto3" json:"is_valid"`
	GroupId              string            `protobuf:"bytes,5,opt,name=group_id,json=groupId,proto3" json:"group_id"`
	AppId                string            `protobuf:"bytes,6,opt,name=app_id,json=appId,proto3" json:"app_id"`
	AcceptOrDismiss      bool              `protobuf:"varint,7,opt,name=accept_or_dismiss,json=acceptOrDismiss,proto3" json:"accept_or_dismiss"`
	WorkflowType         string            `protobuf:"bytes,8,opt,name=workflow_type,json=workflowType,proto3" json:"workflow_type"`
	Params               map[string]string `protobuf:"bytes,9,rep,name=params,proto3" json:"params" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	CreatedAt            string            `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	CreatedBy            string            `protobuf:"bytes,11,opt,name=created_by,json=createdBy,proto3" json:"created_by"`
	UpdatedAt            string            `protobuf:"bytes,12,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	UpdatedBy            string            `protobuf:"bytes,13,opt,name=updated_by,json=updatedBy,proto3" json:"updated_by"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Workflow) Reset()         { *m = Workflow{} }
func (m *Workflow) String() string { return proto.CompactTextString(m) }
func (*Workflow) ProtoMessage()    {}
func (*Workflow) Descriptor() ([]byte, []int) {
	return fileDescriptor_892c7f566756b0be, []int{0}
}

func (m *Workflow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Workflow.Unmarshal(m, b)
}
func (m *Workflow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Workflow.Marshal(b, m, deterministic)
}
func (m *Workflow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Workflow.Merge(m, src)
}
func (m *Workflow) XXX_Size() int {
	return xxx_messageInfo_Workflow.Size(m)
}
func (m *Workflow) XXX_DiscardUnknown() {
	xxx_messageInfo_Workflow.DiscardUnknown(m)
}

var xxx_messageInfo_Workflow proto.InternalMessageInfo

func (m *Workflow) GetWfId() string {
	if m != nil {
		return m.WfId
	}
	return ""
}

func (m *Workflow) GetWfName() string {
	if m != nil {
		return m.WfName
	}
	return ""
}

func (m *Workflow) GetMenuName() string {
	if m != nil {
		return m.MenuName
	}
	return ""
}

func (m *Workflow) GetIsValid() bool {
	if m != nil {
		return m.IsValid
	}
	return false
}

func (m *Workflow) GetGroupId() string {
	if m != nil {
		return m.GroupId
	}
	return ""
}

func (m *Workflow) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *Workflow) GetAcceptOrDismiss() bool {
	if m != nil {
		return m.AcceptOrDismiss
	}
	return false
}

func (m *Workflow) GetWorkflowType() string {
	if m != nil {
		return m.WorkflowType
	}
	return ""
}

func (m *Workflow) GetParams() map[string]string {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *Workflow) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Workflow) GetCreatedBy() string {
	if m != nil {
		return m.CreatedBy
	}
	return ""
}

func (m *Workflow) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *Workflow) GetUpdatedBy() string {
	if m != nil {
		return m.UpdatedBy
	}
	return ""
}

// 查找多条记录
type WorkflowsRequest struct {
	IsValid              string   `protobuf:"bytes,1,opt,name=is_valid,json=isValid,proto3" json:"is_valid"`
	AppId                string   `protobuf:"bytes,2,opt,name=app_id,json=appId,proto3" json:"app_id"`
	ObjectId             string   `protobuf:"bytes,4,opt,name=object_id,json=objectId,proto3" json:"object_id"`
	GroupId              string   `protobuf:"bytes,5,opt,name=group_id,json=groupId,proto3" json:"group_id"`
	Action               string   `protobuf:"bytes,6,opt,name=action,proto3" json:"action"`
	Database             string   `protobuf:"bytes,7,opt,name=database,proto3" json:"database"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorkflowsRequest) Reset()         { *m = WorkflowsRequest{} }
func (m *WorkflowsRequest) String() string { return proto.CompactTextString(m) }
func (*WorkflowsRequest) ProtoMessage()    {}
func (*WorkflowsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_892c7f566756b0be, []int{1}
}

func (m *WorkflowsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowsRequest.Unmarshal(m, b)
}
func (m *WorkflowsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowsRequest.Marshal(b, m, deterministic)
}
func (m *WorkflowsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowsRequest.Merge(m, src)
}
func (m *WorkflowsRequest) XXX_Size() int {
	return xxx_messageInfo_WorkflowsRequest.Size(m)
}
func (m *WorkflowsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowsRequest proto.InternalMessageInfo

func (m *WorkflowsRequest) GetIsValid() string {
	if m != nil {
		return m.IsValid
	}
	return ""
}

func (m *WorkflowsRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *WorkflowsRequest) GetObjectId() string {
	if m != nil {
		return m.ObjectId
	}
	return ""
}

func (m *WorkflowsRequest) GetGroupId() string {
	if m != nil {
		return m.GroupId
	}
	return ""
}

func (m *WorkflowsRequest) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *WorkflowsRequest) GetDatabase() string {
	if m != nil {
		return m.Database
	}
	return ""
}

type WorkflowsResponse struct {
	Workflows            []*Workflow `protobuf:"bytes,1,rep,name=workflows,proto3" json:"workflows"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *WorkflowsResponse) Reset()         { *m = WorkflowsResponse{} }
func (m *WorkflowsResponse) String() string { return proto.CompactTextString(m) }
func (*WorkflowsResponse) ProtoMessage()    {}
func (*WorkflowsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_892c7f566756b0be, []int{2}
}

func (m *WorkflowsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowsResponse.Unmarshal(m, b)
}
func (m *WorkflowsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowsResponse.Marshal(b, m, deterministic)
}
func (m *WorkflowsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowsResponse.Merge(m, src)
}
func (m *WorkflowsResponse) XXX_Size() int {
	return xxx_messageInfo_WorkflowsResponse.Size(m)
}
func (m *WorkflowsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowsResponse proto.InternalMessageInfo

func (m *WorkflowsResponse) GetWorkflows() []*Workflow {
	if m != nil {
		return m.Workflows
	}
	return nil
}

// 查找用户的流程记录
type UserWorkflowsRequest struct {
	AppId                string   `protobuf:"bytes,1,opt,name=app_id,json=appId,proto3" json:"app_id"`
	ObjectId             string   `protobuf:"bytes,2,opt,name=object_id,json=objectId,proto3" json:"object_id"`
	GroupId              string   `protobuf:"bytes,3,opt,name=group_id,json=groupId,proto3" json:"group_id"`
	Action               string   `protobuf:"bytes,4,opt,name=action,proto3" json:"action"`
	Database             string   `protobuf:"bytes,5,opt,name=database,proto3" json:"database"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserWorkflowsRequest) Reset()         { *m = UserWorkflowsRequest{} }
func (m *UserWorkflowsRequest) String() string { return proto.CompactTextString(m) }
func (*UserWorkflowsRequest) ProtoMessage()    {}
func (*UserWorkflowsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_892c7f566756b0be, []int{3}
}

func (m *UserWorkflowsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserWorkflowsRequest.Unmarshal(m, b)
}
func (m *UserWorkflowsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserWorkflowsRequest.Marshal(b, m, deterministic)
}
func (m *UserWorkflowsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserWorkflowsRequest.Merge(m, src)
}
func (m *UserWorkflowsRequest) XXX_Size() int {
	return xxx_messageInfo_UserWorkflowsRequest.Size(m)
}
func (m *UserWorkflowsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserWorkflowsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserWorkflowsRequest proto.InternalMessageInfo

func (m *UserWorkflowsRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *UserWorkflowsRequest) GetObjectId() string {
	if m != nil {
		return m.ObjectId
	}
	return ""
}

func (m *UserWorkflowsRequest) GetGroupId() string {
	if m != nil {
		return m.GroupId
	}
	return ""
}

func (m *UserWorkflowsRequest) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *UserWorkflowsRequest) GetDatabase() string {
	if m != nil {
		return m.Database
	}
	return ""
}

type UserWorkflowsResponse struct {
	Workflows            []*Workflow `protobuf:"bytes,1,rep,name=workflows,proto3" json:"workflows"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *UserWorkflowsResponse) Reset()         { *m = UserWorkflowsResponse{} }
func (m *UserWorkflowsResponse) String() string { return proto.CompactTextString(m) }
func (*UserWorkflowsResponse) ProtoMessage()    {}
func (*UserWorkflowsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_892c7f566756b0be, []int{4}
}

func (m *UserWorkflowsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserWorkflowsResponse.Unmarshal(m, b)
}
func (m *UserWorkflowsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserWorkflowsResponse.Marshal(b, m, deterministic)
}
func (m *UserWorkflowsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserWorkflowsResponse.Merge(m, src)
}
func (m *UserWorkflowsResponse) XXX_Size() int {
	return xxx_messageInfo_UserWorkflowsResponse.Size(m)
}
func (m *UserWorkflowsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserWorkflowsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserWorkflowsResponse proto.InternalMessageInfo

func (m *UserWorkflowsResponse) GetWorkflows() []*Workflow {
	if m != nil {
		return m.Workflows
	}
	return nil
}

// 查询单条记录
type WorkflowRequest struct {
	WfId                 string   `protobuf:"bytes,1,opt,name=wf_id,json=wfId,proto3" json:"wf_id"`
	Database             string   `protobuf:"bytes,2,opt,name=database,proto3" json:"database"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorkflowRequest) Reset()         { *m = WorkflowRequest{} }
func (m *WorkflowRequest) String() string { return proto.CompactTextString(m) }
func (*WorkflowRequest) ProtoMessage()    {}
func (*WorkflowRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_892c7f566756b0be, []int{5}
}

func (m *WorkflowRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowRequest.Unmarshal(m, b)
}
func (m *WorkflowRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowRequest.Marshal(b, m, deterministic)
}
func (m *WorkflowRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowRequest.Merge(m, src)
}
func (m *WorkflowRequest) XXX_Size() int {
	return xxx_messageInfo_WorkflowRequest.Size(m)
}
func (m *WorkflowRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowRequest proto.InternalMessageInfo

func (m *WorkflowRequest) GetWfId() string {
	if m != nil {
		return m.WfId
	}
	return ""
}

func (m *WorkflowRequest) GetDatabase() string {
	if m != nil {
		return m.Database
	}
	return ""
}

type WorkflowResponse struct {
	Workflow             *Workflow `protobuf:"bytes,1,opt,name=workflow,proto3" json:"workflow"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *WorkflowResponse) Reset()         { *m = WorkflowResponse{} }
func (m *WorkflowResponse) String() string { return proto.CompactTextString(m) }
func (*WorkflowResponse) ProtoMessage()    {}
func (*WorkflowResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_892c7f566756b0be, []int{6}
}

func (m *WorkflowResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowResponse.Unmarshal(m, b)
}
func (m *WorkflowResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowResponse.Marshal(b, m, deterministic)
}
func (m *WorkflowResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowResponse.Merge(m, src)
}
func (m *WorkflowResponse) XXX_Size() int {
	return xxx_messageInfo_WorkflowResponse.Size(m)
}
func (m *WorkflowResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowResponse proto.InternalMessageInfo

func (m *WorkflowResponse) GetWorkflow() *Workflow {
	if m != nil {
		return m.Workflow
	}
	return nil
}

// 添加数据
type AddRequest struct {
	WfName               string            `protobuf:"bytes,1,opt,name=wf_name,json=wfName,proto3" json:"wf_name"`
	MenuName             string            `protobuf:"bytes,2,opt,name=menu_name,json=menuName,proto3" json:"menu_name"`
	IsValid              bool              `protobuf:"varint,3,opt,name=is_valid,json=isValid,proto3" json:"is_valid"`
	GroupId              string            `protobuf:"bytes,4,opt,name=group_id,json=groupId,proto3" json:"group_id"`
	AppId                string            `protobuf:"bytes,5,opt,name=app_id,json=appId,proto3" json:"app_id"`
	AcceptOrDismiss      bool              `protobuf:"varint,6,opt,name=accept_or_dismiss,json=acceptOrDismiss,proto3" json:"accept_or_dismiss"`
	WorkflowType         string            `protobuf:"bytes,7,opt,name=workflow_type,json=workflowType,proto3" json:"workflow_type"`
	Params               map[string]string `protobuf:"bytes,8,rep,name=params,proto3" json:"params" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Database             string            `protobuf:"bytes,9,opt,name=database,proto3" json:"database"`
	Writer               string            `protobuf:"bytes,10,opt,name=writer,proto3" json:"writer"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *AddRequest) Reset()         { *m = AddRequest{} }
func (m *AddRequest) String() string { return proto.CompactTextString(m) }
func (*AddRequest) ProtoMessage()    {}
func (*AddRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_892c7f566756b0be, []int{7}
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

func (m *AddRequest) GetWfName() string {
	if m != nil {
		return m.WfName
	}
	return ""
}

func (m *AddRequest) GetMenuName() string {
	if m != nil {
		return m.MenuName
	}
	return ""
}

func (m *AddRequest) GetIsValid() bool {
	if m != nil {
		return m.IsValid
	}
	return false
}

func (m *AddRequest) GetGroupId() string {
	if m != nil {
		return m.GroupId
	}
	return ""
}

func (m *AddRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *AddRequest) GetAcceptOrDismiss() bool {
	if m != nil {
		return m.AcceptOrDismiss
	}
	return false
}

func (m *AddRequest) GetWorkflowType() string {
	if m != nil {
		return m.WorkflowType
	}
	return ""
}

func (m *AddRequest) GetParams() map[string]string {
	if m != nil {
		return m.Params
	}
	return nil
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
	WfId                 string   `protobuf:"bytes,1,opt,name=wf_id,json=wfId,proto3" json:"wf_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddResponse) Reset()         { *m = AddResponse{} }
func (m *AddResponse) String() string { return proto.CompactTextString(m) }
func (*AddResponse) ProtoMessage()    {}
func (*AddResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_892c7f566756b0be, []int{8}
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

func (m *AddResponse) GetWfId() string {
	if m != nil {
		return m.WfId
	}
	return ""
}

// 更新记录
type ModifyRequest struct {
	WfId                 string            `protobuf:"bytes,1,opt,name=wf_id,json=wfId,proto3" json:"wf_id"`
	WfName               string            `protobuf:"bytes,2,opt,name=wf_name,json=wfName,proto3" json:"wf_name"`
	MenuName             string            `protobuf:"bytes,3,opt,name=menu_name,json=menuName,proto3" json:"menu_name"`
	IsValid              string            `protobuf:"bytes,4,opt,name=is_valid,json=isValid,proto3" json:"is_valid"`
	AcceptOrDismiss      string            `protobuf:"bytes,5,opt,name=accept_or_dismiss,json=acceptOrDismiss,proto3" json:"accept_or_dismiss"`
	Params               map[string]string `protobuf:"bytes,6,rep,name=params,proto3" json:"params" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Database             string            `protobuf:"bytes,7,opt,name=database,proto3" json:"database"`
	Writer               string            `protobuf:"bytes,8,opt,name=writer,proto3" json:"writer"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ModifyRequest) Reset()         { *m = ModifyRequest{} }
func (m *ModifyRequest) String() string { return proto.CompactTextString(m) }
func (*ModifyRequest) ProtoMessage()    {}
func (*ModifyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_892c7f566756b0be, []int{9}
}

func (m *ModifyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModifyRequest.Unmarshal(m, b)
}
func (m *ModifyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModifyRequest.Marshal(b, m, deterministic)
}
func (m *ModifyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModifyRequest.Merge(m, src)
}
func (m *ModifyRequest) XXX_Size() int {
	return xxx_messageInfo_ModifyRequest.Size(m)
}
func (m *ModifyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ModifyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ModifyRequest proto.InternalMessageInfo

func (m *ModifyRequest) GetWfId() string {
	if m != nil {
		return m.WfId
	}
	return ""
}

func (m *ModifyRequest) GetWfName() string {
	if m != nil {
		return m.WfName
	}
	return ""
}

func (m *ModifyRequest) GetMenuName() string {
	if m != nil {
		return m.MenuName
	}
	return ""
}

func (m *ModifyRequest) GetIsValid() string {
	if m != nil {
		return m.IsValid
	}
	return ""
}

func (m *ModifyRequest) GetAcceptOrDismiss() string {
	if m != nil {
		return m.AcceptOrDismiss
	}
	return ""
}

func (m *ModifyRequest) GetParams() map[string]string {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *ModifyRequest) GetDatabase() string {
	if m != nil {
		return m.Database
	}
	return ""
}

func (m *ModifyRequest) GetWriter() string {
	if m != nil {
		return m.Writer
	}
	return ""
}

type ModifyResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ModifyResponse) Reset()         { *m = ModifyResponse{} }
func (m *ModifyResponse) String() string { return proto.CompactTextString(m) }
func (*ModifyResponse) ProtoMessage()    {}
func (*ModifyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_892c7f566756b0be, []int{10}
}

func (m *ModifyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModifyResponse.Unmarshal(m, b)
}
func (m *ModifyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModifyResponse.Marshal(b, m, deterministic)
}
func (m *ModifyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModifyResponse.Merge(m, src)
}
func (m *ModifyResponse) XXX_Size() int {
	return xxx_messageInfo_ModifyResponse.Size(m)
}
func (m *ModifyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ModifyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ModifyResponse proto.InternalMessageInfo

// 删除数据记录
type DeleteRequest struct {
	Workflows            []string `protobuf:"bytes,1,rep,name=workflows,proto3" json:"workflows"`
	Database             string   `protobuf:"bytes,2,opt,name=database,proto3" json:"database"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_892c7f566756b0be, []int{11}
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

func (m *DeleteRequest) GetWorkflows() []string {
	if m != nil {
		return m.Workflows
	}
	return nil
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
	return fileDescriptor_892c7f566756b0be, []int{12}
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
	proto.RegisterType((*Workflow)(nil), "workflow.Workflow")
	proto.RegisterMapType((map[string]string)(nil), "workflow.Workflow.ParamsEntry")
	proto.RegisterType((*WorkflowsRequest)(nil), "workflow.WorkflowsRequest")
	proto.RegisterType((*WorkflowsResponse)(nil), "workflow.WorkflowsResponse")
	proto.RegisterType((*UserWorkflowsRequest)(nil), "workflow.UserWorkflowsRequest")
	proto.RegisterType((*UserWorkflowsResponse)(nil), "workflow.UserWorkflowsResponse")
	proto.RegisterType((*WorkflowRequest)(nil), "workflow.WorkflowRequest")
	proto.RegisterType((*WorkflowResponse)(nil), "workflow.WorkflowResponse")
	proto.RegisterType((*AddRequest)(nil), "workflow.AddRequest")
	proto.RegisterMapType((map[string]string)(nil), "workflow.AddRequest.ParamsEntry")
	proto.RegisterType((*AddResponse)(nil), "workflow.AddResponse")
	proto.RegisterType((*ModifyRequest)(nil), "workflow.ModifyRequest")
	proto.RegisterMapType((map[string]string)(nil), "workflow.ModifyRequest.ParamsEntry")
	proto.RegisterType((*ModifyResponse)(nil), "workflow.ModifyResponse")
	proto.RegisterType((*DeleteRequest)(nil), "workflow.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "workflow.DeleteResponse")
}

func init() { proto.RegisterFile("workflow.proto", fileDescriptor_892c7f566756b0be) }

var fileDescriptor_892c7f566756b0be = []byte{
	// 773 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xdb, 0x6e, 0x9b, 0x4c,
	0x10, 0x0e, 0xc6, 0xc6, 0x30, 0x8e, 0x13, 0x67, 0xff, 0xe4, 0x0f, 0x21, 0xff, 0xc1, 0x22, 0x37,
	0x51, 0x2f, 0xac, 0x2a, 0x95, 0xaa, 0xf4, 0x70, 0x13, 0x2b, 0x69, 0xe5, 0x4a, 0x3d, 0x88, 0xb6,
	0xc9, 0x25, 0x5a, 0x9b, 0xa5, 0xa2, 0xb1, 0x0d, 0x05, 0x1c, 0xc4, 0x23, 0xf4, 0x11, 0xfa, 0x16,
	0x7d, 0x88, 0xde, 0xf6, 0x25, 0xfa, 0x24, 0x15, 0xcb, 0x2e, 0x07, 0x1b, 0x9c, 0xaa, 0xcd, 0x5d,
	0x66, 0xbe, 0xdd, 0x2f, 0x33, 0xf3, 0xed, 0x7c, 0x18, 0xb6, 0x22, 0xd7, 0xbf, 0xb6, 0xa7, 0x6e,
	0x34, 0xf0, 0x7c, 0x37, 0x74, 0x91, 0xcc, 0x63, 0xfd, 0x87, 0x08, 0xf2, 0x15, 0x0b, 0xd0, 0x5f,
	0xd0, 0x8a, 0x6c, 0xd3, 0xb1, 0x54, 0xa1, 0x2f, 0x1c, 0x2b, 0x46, 0x33, 0xb2, 0x47, 0x16, 0xda,
	0x87, 0x76, 0x64, 0x9b, 0x73, 0x3c, 0x23, 0x6a, 0x83, 0xa6, 0xa5, 0xc8, 0x7e, 0x85, 0x67, 0x04,
	0x1d, 0x82, 0x32, 0x23, 0xf3, 0x45, 0x0a, 0x89, 0x14, 0x92, 0x93, 0x04, 0x05, 0x0f, 0x40, 0x76,
	0x02, 0xf3, 0x06, 0x4f, 0x1d, 0x4b, 0x6d, 0xf6, 0x85, 0x63, 0xd9, 0x68, 0x3b, 0xc1, 0x65, 0x12,
	0x26, 0xd0, 0x07, 0xdf, 0x5d, 0x78, 0xc9, 0x3f, 0x6a, 0xd1, 0x6b, 0x6d, 0x1a, 0x8f, 0x2c, 0xb4,
	0x07, 0x12, 0xf6, 0x28, 0x20, 0x51, 0xa0, 0x85, 0xbd, 0x24, 0x7d, 0x0f, 0x76, 0xf0, 0x64, 0x42,
	0xbc, 0xd0, 0x74, 0x7d, 0xd3, 0x72, 0x82, 0x99, 0x13, 0x04, 0x6a, 0x9b, 0xb2, 0x6e, 0xa7, 0xc0,
	0x6b, 0xff, 0x3c, 0x4d, 0xa3, 0x23, 0xe8, 0xf2, 0xe6, 0xcc, 0x30, 0xf6, 0x88, 0x2a, 0x53, 0xa6,
	0x4d, 0x9e, 0x7c, 0x17, 0x7b, 0x04, 0x3d, 0x04, 0xc9, 0xc3, 0x3e, 0x9e, 0x05, 0xaa, 0xd2, 0x17,
	0x8f, 0x3b, 0x27, 0xff, 0x0d, 0xb2, 0x01, 0xf1, 0x61, 0x0c, 0xde, 0xd0, 0x03, 0x17, 0xf3, 0xd0,
	0x8f, 0x0d, 0x76, 0x1a, 0xfd, 0x0b, 0x30, 0xf1, 0x09, 0x0e, 0x89, 0x65, 0xe2, 0x50, 0x05, 0xca,
	0xac, 0xb0, 0xcc, 0x59, 0x58, 0x84, 0xc7, 0xb1, 0xda, 0x29, 0xc1, 0xc3, 0x38, 0x81, 0x17, 0x9e,
	0xc5, 0x6f, 0x6f, 0xa6, 0x30, 0xcb, 0xa4, 0xb7, 0x39, 0x3c, 0x8e, 0xd5, 0x6e, 0x09, 0x1e, 0xc6,
	0xda, 0x23, 0xe8, 0x14, 0x4a, 0x42, 0x3d, 0x10, 0xaf, 0x49, 0xcc, 0x94, 0x4a, 0xfe, 0x44, 0xbb,
	0xd0, 0xba, 0xc1, 0xd3, 0x05, 0x97, 0x29, 0x0d, 0x1e, 0x37, 0x4e, 0x05, 0xfd, 0xab, 0x00, 0x3d,
	0xde, 0x57, 0x60, 0x90, 0x4f, 0x0b, 0x12, 0x84, 0x25, 0x85, 0x52, 0x96, 0x4c, 0xa1, 0x5c, 0x86,
	0x46, 0x51, 0x86, 0x43, 0x50, 0xdc, 0xf1, 0x47, 0x32, 0x09, 0x4d, 0x26, 0xaa, 0x62, 0xc8, 0x69,
	0x62, 0xb4, 0x56, 0xd5, 0xbf, 0x41, 0xc2, 0x93, 0xd0, 0x71, 0xe7, 0x4c, 0x55, 0x16, 0x21, 0x0d,
	0x64, 0x0b, 0x87, 0x78, 0x8c, 0x03, 0x42, 0xd5, 0x54, 0x8c, 0x2c, 0xd6, 0x2f, 0x60, 0xa7, 0x50,
	0x71, 0xe0, 0xb9, 0xf3, 0x80, 0xa0, 0xfb, 0xa0, 0x70, 0x9d, 0x02, 0x55, 0xa0, 0xca, 0xa1, 0x55,
	0xe5, 0x8c, 0xfc, 0x90, 0xfe, 0x45, 0x80, 0xdd, 0xf7, 0x01, 0xf1, 0x57, 0xba, 0xcf, 0x5b, 0x14,
	0x6a, 0x5b, 0x6c, 0xac, 0x69, 0x51, 0xac, 0x6b, 0xb1, 0x59, 0xdb, 0x62, 0x6b, 0xa9, 0xc5, 0x11,
	0xec, 0x2d, 0x95, 0xf6, 0xdb, 0x6d, 0x0e, 0x61, 0x3b, 0x4b, 0xb3, 0x06, 0x2b, 0x77, 0xb9, 0x58,
	0x4e, 0x63, 0xa9, 0x9c, 0x61, 0xfe, 0x46, 0xb2, 0x4a, 0x06, 0x90, 0x39, 0x05, 0xe5, 0xa9, 0x2e,
	0x24, 0x77, 0x93, 0xcf, 0x22, 0xc0, 0x99, 0x65, 0xf1, 0x1a, 0x0a, 0xd6, 0x21, 0xd4, 0x5b, 0x47,
	0x63, 0x8d, 0x75, 0x88, 0xf5, 0xd6, 0xd1, 0xac, 0xb3, 0x8e, 0xd6, 0xad, 0xd6, 0x21, 0xfd, 0xa2,
	0x75, 0xb4, 0x2b, 0xac, 0xe3, 0x34, 0xb3, 0x0e, 0x99, 0x2a, 0xd3, 0xcf, 0x07, 0x92, 0x77, 0x5e,
	0x69, 0x1e, 0xc5, 0xe1, 0x2b, 0xe5, 0xe1, 0x27, 0xef, 0x27, 0xf2, 0x9d, 0x90, 0xf8, 0xcc, 0x54,
	0x58, 0xf4, 0x27, 0x4b, 0xaf, 0x43, 0x87, 0x16, 0xc4, 0xa4, 0xac, 0x7a, 0x0f, 0xfa, 0xf7, 0x06,
	0x74, 0x5f, 0xba, 0x96, 0x63, 0xc7, 0x6b, 0x9f, 0xcd, 0xdd, 0x7c, 0x02, 0x0a, 0x06, 0x53, 0xa9,
	0x4a, 0xaa, 0xdb, 0x8a, 0x2a, 0x4f, 0xb2, 0x81, 0x4b, 0x74, 0xe0, 0x47, 0xf9, 0xc0, 0x4b, 0xa5,
	0xdf, 0x3a, 0xf3, 0x76, 0xed, 0xcc, 0xe5, 0xbb, 0x9a, 0x79, 0x0f, 0xb6, 0x78, 0x4d, 0xe9, 0xd8,
	0xf5, 0x11, 0x74, 0xcf, 0xc9, 0x94, 0x84, 0x84, 0x0f, 0xf8, 0x9f, 0xe5, 0xe5, 0x56, 0x0a, 0x8b,
	0xbc, 0x76, 0x41, 0x7b, 0xb0, 0xc5, 0xa9, 0x52, 0xf2, 0x93, 0x6f, 0x22, 0x28, 0x57, 0xf6, 0x5b,
	0xe2, 0xdf, 0x38, 0x13, 0x82, 0x5e, 0x40, 0xf7, 0x99, 0x33, 0xb7, 0xae, 0x72, 0xb2, 0xd5, 0x5d,
	0xe5, 0xfe, 0xa7, 0x1d, 0x56, 0x62, 0xac, 0xe8, 0x0d, 0x74, 0x09, 0x3b, 0x09, 0x57, 0xc9, 0x9f,
	0x50, 0xe1, 0x2b, 0x59, 0xe5, 0xa9, 0xda, 0xff, 0xb5, 0x78, 0xc6, 0xfb, 0x1c, 0x36, 0x8b, 0x35,
	0xa2, 0x83, 0x0a, 0x3b, 0x61, 0x6c, 0x5a, 0x15, 0x94, 0x11, 0x3d, 0xa5, 0xaf, 0x3b, 0xe3, 0xd9,
	0xad, 0xda, 0x42, 0x6d, 0x6f, 0x29, 0x9b, 0xdd, 0xbe, 0xe0, 0x3a, 0x65, 0x04, 0xfb, 0x35, 0xaf,
	0x4a, 0x53, 0x57, 0x81, 0x22, 0x4d, 0xaa, 0x48, 0x15, 0x4d, 0x49, 0xf6, 0x22, 0x4d, 0x59, 0x44,
	0x7d, 0x63, 0x2c, 0xd1, 0x1f, 0x65, 0x0f, 0x7e, 0x06, 0x00, 0x00, 0xff, 0xff, 0x94, 0x21, 0x29,
	0xb2, 0xa6, 0x09, 0x00, 0x00,
}
