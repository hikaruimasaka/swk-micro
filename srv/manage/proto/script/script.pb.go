// Code generated by protoc-gen-go. DO NOT EDIT.
// source: script.proto

package script

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

// 用户
type ScriptJob struct {
	ScriptId             string   `protobuf:"bytes,1,opt,name=script_id,json=scriptId,proto3" json:"script_id"`
	ScriptName           string   `protobuf:"bytes,2,opt,name=script_name,json=scriptName,proto3" json:"script_name"`
	ScriptDesc           string   `protobuf:"bytes,3,opt,name=script_desc,json=scriptDesc,proto3" json:"script_desc"`
	ScriptType           string   `protobuf:"bytes,4,opt,name=script_type,json=scriptType,proto3" json:"script_type"`
	ScriptData           string   `protobuf:"bytes,5,opt,name=script_data,json=scriptData,proto3" json:"script_data"`
	ScriptFunc           string   `protobuf:"bytes,6,opt,name=script_func,json=scriptFunc,proto3" json:"script_func"`
	ScriptVersion        string   `protobuf:"bytes,7,opt,name=script_version,json=scriptVersion,proto3" json:"script_version"`
	RunLogs              []string `protobuf:"bytes,8,rep,name=run_logs,json=runLogs,proto3" json:"run_logs"`
	CreatedAt            string   `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	CreatedBy            string   `protobuf:"bytes,10,opt,name=created_by,json=createdBy,proto3" json:"created_by"`
	RanAt                string   `protobuf:"bytes,11,opt,name=ran_at,json=ranAt,proto3" json:"ran_at"`
	RanBy                string   `protobuf:"bytes,12,opt,name=ran_by,json=ranBy,proto3" json:"ran_by"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScriptJob) Reset()         { *m = ScriptJob{} }
func (m *ScriptJob) String() string { return proto.CompactTextString(m) }
func (*ScriptJob) ProtoMessage()    {}
func (*ScriptJob) Descriptor() ([]byte, []int) {
	return fileDescriptor_01e8311d1fb507af, []int{0}
}

func (m *ScriptJob) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScriptJob.Unmarshal(m, b)
}
func (m *ScriptJob) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScriptJob.Marshal(b, m, deterministic)
}
func (m *ScriptJob) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScriptJob.Merge(m, src)
}
func (m *ScriptJob) XXX_Size() int {
	return xxx_messageInfo_ScriptJob.Size(m)
}
func (m *ScriptJob) XXX_DiscardUnknown() {
	xxx_messageInfo_ScriptJob.DiscardUnknown(m)
}

var xxx_messageInfo_ScriptJob proto.InternalMessageInfo

func (m *ScriptJob) GetScriptId() string {
	if m != nil {
		return m.ScriptId
	}
	return ""
}

func (m *ScriptJob) GetScriptName() string {
	if m != nil {
		return m.ScriptName
	}
	return ""
}

func (m *ScriptJob) GetScriptDesc() string {
	if m != nil {
		return m.ScriptDesc
	}
	return ""
}

func (m *ScriptJob) GetScriptType() string {
	if m != nil {
		return m.ScriptType
	}
	return ""
}

func (m *ScriptJob) GetScriptData() string {
	if m != nil {
		return m.ScriptData
	}
	return ""
}

func (m *ScriptJob) GetScriptFunc() string {
	if m != nil {
		return m.ScriptFunc
	}
	return ""
}

func (m *ScriptJob) GetScriptVersion() string {
	if m != nil {
		return m.ScriptVersion
	}
	return ""
}

func (m *ScriptJob) GetRunLogs() []string {
	if m != nil {
		return m.RunLogs
	}
	return nil
}

func (m *ScriptJob) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *ScriptJob) GetCreatedBy() string {
	if m != nil {
		return m.CreatedBy
	}
	return ""
}

func (m *ScriptJob) GetRanAt() string {
	if m != nil {
		return m.RanAt
	}
	return ""
}

func (m *ScriptJob) GetRanBy() string {
	if m != nil {
		return m.RanBy
	}
	return ""
}

type FindScriptJobsRequest struct {
	ScriptType           string   `protobuf:"bytes,1,opt,name=script_type,json=scriptType,proto3" json:"script_type"`
	ScriptVersion        string   `protobuf:"bytes,2,opt,name=script_version,json=scriptVersion,proto3" json:"script_version"`
	RanBy                string   `protobuf:"bytes,3,opt,name=ran_by,json=ranBy,proto3" json:"ran_by"`
	Database             string   `protobuf:"bytes,4,opt,name=database,proto3" json:"database"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindScriptJobsRequest) Reset()         { *m = FindScriptJobsRequest{} }
func (m *FindScriptJobsRequest) String() string { return proto.CompactTextString(m) }
func (*FindScriptJobsRequest) ProtoMessage()    {}
func (*FindScriptJobsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_01e8311d1fb507af, []int{1}
}

func (m *FindScriptJobsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindScriptJobsRequest.Unmarshal(m, b)
}
func (m *FindScriptJobsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindScriptJobsRequest.Marshal(b, m, deterministic)
}
func (m *FindScriptJobsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindScriptJobsRequest.Merge(m, src)
}
func (m *FindScriptJobsRequest) XXX_Size() int {
	return xxx_messageInfo_FindScriptJobsRequest.Size(m)
}
func (m *FindScriptJobsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindScriptJobsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindScriptJobsRequest proto.InternalMessageInfo

func (m *FindScriptJobsRequest) GetScriptType() string {
	if m != nil {
		return m.ScriptType
	}
	return ""
}

func (m *FindScriptJobsRequest) GetScriptVersion() string {
	if m != nil {
		return m.ScriptVersion
	}
	return ""
}

func (m *FindScriptJobsRequest) GetRanBy() string {
	if m != nil {
		return m.RanBy
	}
	return ""
}

func (m *FindScriptJobsRequest) GetDatabase() string {
	if m != nil {
		return m.Database
	}
	return ""
}

type FindScriptJobsResponse struct {
	ScriptJobs           []*ScriptJob `protobuf:"bytes,1,rep,name=script_jobs,json=scriptJobs,proto3" json:"script_jobs"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *FindScriptJobsResponse) Reset()         { *m = FindScriptJobsResponse{} }
func (m *FindScriptJobsResponse) String() string { return proto.CompactTextString(m) }
func (*FindScriptJobsResponse) ProtoMessage()    {}
func (*FindScriptJobsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_01e8311d1fb507af, []int{2}
}

func (m *FindScriptJobsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindScriptJobsResponse.Unmarshal(m, b)
}
func (m *FindScriptJobsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindScriptJobsResponse.Marshal(b, m, deterministic)
}
func (m *FindScriptJobsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindScriptJobsResponse.Merge(m, src)
}
func (m *FindScriptJobsResponse) XXX_Size() int {
	return xxx_messageInfo_FindScriptJobsResponse.Size(m)
}
func (m *FindScriptJobsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindScriptJobsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindScriptJobsResponse proto.InternalMessageInfo

func (m *FindScriptJobsResponse) GetScriptJobs() []*ScriptJob {
	if m != nil {
		return m.ScriptJobs
	}
	return nil
}

type DeleteScriptsRequest struct {
	Database             string   `protobuf:"bytes,1,opt,name=database,proto3" json:"database"`
	ScriptIds            []string `protobuf:"bytes,2,rep,name=script_ids,json=scriptIds,proto3" json:"script_ids"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteScriptsRequest) Reset()         { *m = DeleteScriptsRequest{} }
func (m *DeleteScriptsRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteScriptsRequest) ProtoMessage()    {}
func (*DeleteScriptsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_01e8311d1fb507af, []int{3}
}

func (m *DeleteScriptsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteScriptsRequest.Unmarshal(m, b)
}
func (m *DeleteScriptsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteScriptsRequest.Marshal(b, m, deterministic)
}
func (m *DeleteScriptsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteScriptsRequest.Merge(m, src)
}
func (m *DeleteScriptsRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteScriptsRequest.Size(m)
}
func (m *DeleteScriptsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteScriptsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteScriptsRequest proto.InternalMessageInfo

func (m *DeleteScriptsRequest) GetDatabase() string {
	if m != nil {
		return m.Database
	}
	return ""
}

func (m *DeleteScriptsRequest) GetScriptIds() []string {
	if m != nil {
		return m.ScriptIds
	}
	return nil
}

type DeleteScriptsResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteScriptsResponse) Reset()         { *m = DeleteScriptsResponse{} }
func (m *DeleteScriptsResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteScriptsResponse) ProtoMessage()    {}
func (*DeleteScriptsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_01e8311d1fb507af, []int{4}
}

func (m *DeleteScriptsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteScriptsResponse.Unmarshal(m, b)
}
func (m *DeleteScriptsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteScriptsResponse.Marshal(b, m, deterministic)
}
func (m *DeleteScriptsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteScriptsResponse.Merge(m, src)
}
func (m *DeleteScriptsResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteScriptsResponse.Size(m)
}
func (m *DeleteScriptsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteScriptsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteScriptsResponse proto.InternalMessageInfo

type FindScriptJobRequest struct {
	ScriptId             string   `protobuf:"bytes,1,opt,name=script_id,json=scriptId,proto3" json:"script_id"`
	Database             string   `protobuf:"bytes,2,opt,name=database,proto3" json:"database"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindScriptJobRequest) Reset()         { *m = FindScriptJobRequest{} }
func (m *FindScriptJobRequest) String() string { return proto.CompactTextString(m) }
func (*FindScriptJobRequest) ProtoMessage()    {}
func (*FindScriptJobRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_01e8311d1fb507af, []int{5}
}

func (m *FindScriptJobRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindScriptJobRequest.Unmarshal(m, b)
}
func (m *FindScriptJobRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindScriptJobRequest.Marshal(b, m, deterministic)
}
func (m *FindScriptJobRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindScriptJobRequest.Merge(m, src)
}
func (m *FindScriptJobRequest) XXX_Size() int {
	return xxx_messageInfo_FindScriptJobRequest.Size(m)
}
func (m *FindScriptJobRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindScriptJobRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindScriptJobRequest proto.InternalMessageInfo

func (m *FindScriptJobRequest) GetScriptId() string {
	if m != nil {
		return m.ScriptId
	}
	return ""
}

func (m *FindScriptJobRequest) GetDatabase() string {
	if m != nil {
		return m.Database
	}
	return ""
}

type FindScriptJobResponse struct {
	ScriptJob            *ScriptJob `protobuf:"bytes,1,opt,name=script_job,json=scriptJob,proto3" json:"script_job"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *FindScriptJobResponse) Reset()         { *m = FindScriptJobResponse{} }
func (m *FindScriptJobResponse) String() string { return proto.CompactTextString(m) }
func (*FindScriptJobResponse) ProtoMessage()    {}
func (*FindScriptJobResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_01e8311d1fb507af, []int{6}
}

func (m *FindScriptJobResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindScriptJobResponse.Unmarshal(m, b)
}
func (m *FindScriptJobResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindScriptJobResponse.Marshal(b, m, deterministic)
}
func (m *FindScriptJobResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindScriptJobResponse.Merge(m, src)
}
func (m *FindScriptJobResponse) XXX_Size() int {
	return xxx_messageInfo_FindScriptJobResponse.Size(m)
}
func (m *FindScriptJobResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindScriptJobResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindScriptJobResponse proto.InternalMessageInfo

func (m *FindScriptJobResponse) GetScriptJob() *ScriptJob {
	if m != nil {
		return m.ScriptJob
	}
	return nil
}

type AddRequest struct {
	ScriptId             string   `protobuf:"bytes,1,opt,name=script_id,json=scriptId,proto3" json:"script_id"`
	ScriptName           string   `protobuf:"bytes,2,opt,name=script_name,json=scriptName,proto3" json:"script_name"`
	ScriptDesc           string   `protobuf:"bytes,3,opt,name=script_desc,json=scriptDesc,proto3" json:"script_desc"`
	ScriptType           string   `protobuf:"bytes,4,opt,name=script_type,json=scriptType,proto3" json:"script_type"`
	ScriptData           string   `protobuf:"bytes,5,opt,name=script_data,json=scriptData,proto3" json:"script_data"`
	ScriptFunc           string   `protobuf:"bytes,6,opt,name=script_func,json=scriptFunc,proto3" json:"script_func"`
	ScriptVersion        string   `protobuf:"bytes,7,opt,name=script_version,json=scriptVersion,proto3" json:"script_version"`
	CreatedAt            string   `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	Writer               string   `protobuf:"bytes,9,opt,name=writer,proto3" json:"writer"`
	Database             string   `protobuf:"bytes,10,opt,name=database,proto3" json:"database"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddRequest) Reset()         { *m = AddRequest{} }
func (m *AddRequest) String() string { return proto.CompactTextString(m) }
func (*AddRequest) ProtoMessage()    {}
func (*AddRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_01e8311d1fb507af, []int{7}
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

func (m *AddRequest) GetScriptId() string {
	if m != nil {
		return m.ScriptId
	}
	return ""
}

func (m *AddRequest) GetScriptName() string {
	if m != nil {
		return m.ScriptName
	}
	return ""
}

func (m *AddRequest) GetScriptDesc() string {
	if m != nil {
		return m.ScriptDesc
	}
	return ""
}

func (m *AddRequest) GetScriptType() string {
	if m != nil {
		return m.ScriptType
	}
	return ""
}

func (m *AddRequest) GetScriptData() string {
	if m != nil {
		return m.ScriptData
	}
	return ""
}

func (m *AddRequest) GetScriptFunc() string {
	if m != nil {
		return m.ScriptFunc
	}
	return ""
}

func (m *AddRequest) GetScriptVersion() string {
	if m != nil {
		return m.ScriptVersion
	}
	return ""
}

func (m *AddRequest) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *AddRequest) GetWriter() string {
	if m != nil {
		return m.Writer
	}
	return ""
}

func (m *AddRequest) GetDatabase() string {
	if m != nil {
		return m.Database
	}
	return ""
}

type AddResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddResponse) Reset()         { *m = AddResponse{} }
func (m *AddResponse) String() string { return proto.CompactTextString(m) }
func (*AddResponse) ProtoMessage()    {}
func (*AddResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_01e8311d1fb507af, []int{8}
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

type ModifyRequest struct {
	ScriptId             string   `protobuf:"bytes,1,opt,name=script_id,json=scriptId,proto3" json:"script_id"`
	ScriptName           string   `protobuf:"bytes,2,opt,name=script_name,json=scriptName,proto3" json:"script_name"`
	ScriptDesc           string   `protobuf:"bytes,3,opt,name=script_desc,json=scriptDesc,proto3" json:"script_desc"`
	ScriptData           string   `protobuf:"bytes,4,opt,name=script_data,json=scriptData,proto3" json:"script_data"`
	ScriptFunc           string   `protobuf:"bytes,5,opt,name=script_func,json=scriptFunc,proto3" json:"script_func"`
	ScriptVersion        string   `protobuf:"bytes,6,opt,name=script_version,json=scriptVersion,proto3" json:"script_version"`
	Writer               string   `protobuf:"bytes,7,opt,name=writer,proto3" json:"writer"`
	Database             string   `protobuf:"bytes,8,opt,name=database,proto3" json:"database"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ModifyRequest) Reset()         { *m = ModifyRequest{} }
func (m *ModifyRequest) String() string { return proto.CompactTextString(m) }
func (*ModifyRequest) ProtoMessage()    {}
func (*ModifyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_01e8311d1fb507af, []int{9}
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

func (m *ModifyRequest) GetScriptId() string {
	if m != nil {
		return m.ScriptId
	}
	return ""
}

func (m *ModifyRequest) GetScriptName() string {
	if m != nil {
		return m.ScriptName
	}
	return ""
}

func (m *ModifyRequest) GetScriptDesc() string {
	if m != nil {
		return m.ScriptDesc
	}
	return ""
}

func (m *ModifyRequest) GetScriptData() string {
	if m != nil {
		return m.ScriptData
	}
	return ""
}

func (m *ModifyRequest) GetScriptFunc() string {
	if m != nil {
		return m.ScriptFunc
	}
	return ""
}

func (m *ModifyRequest) GetScriptVersion() string {
	if m != nil {
		return m.ScriptVersion
	}
	return ""
}

func (m *ModifyRequest) GetWriter() string {
	if m != nil {
		return m.Writer
	}
	return ""
}

func (m *ModifyRequest) GetDatabase() string {
	if m != nil {
		return m.Database
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
	return fileDescriptor_01e8311d1fb507af, []int{10}
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

type StartRequest struct {
	ScriptId             string   `protobuf:"bytes,1,opt,name=script_id,json=scriptId,proto3" json:"script_id"`
	Writer               string   `protobuf:"bytes,2,opt,name=writer,proto3" json:"writer"`
	Database             string   `protobuf:"bytes,3,opt,name=database,proto3" json:"database"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartRequest) Reset()         { *m = StartRequest{} }
func (m *StartRequest) String() string { return proto.CompactTextString(m) }
func (*StartRequest) ProtoMessage()    {}
func (*StartRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_01e8311d1fb507af, []int{11}
}

func (m *StartRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartRequest.Unmarshal(m, b)
}
func (m *StartRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartRequest.Marshal(b, m, deterministic)
}
func (m *StartRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartRequest.Merge(m, src)
}
func (m *StartRequest) XXX_Size() int {
	return xxx_messageInfo_StartRequest.Size(m)
}
func (m *StartRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StartRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StartRequest proto.InternalMessageInfo

func (m *StartRequest) GetScriptId() string {
	if m != nil {
		return m.ScriptId
	}
	return ""
}

func (m *StartRequest) GetWriter() string {
	if m != nil {
		return m.Writer
	}
	return ""
}

func (m *StartRequest) GetDatabase() string {
	if m != nil {
		return m.Database
	}
	return ""
}

type StartResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartResponse) Reset()         { *m = StartResponse{} }
func (m *StartResponse) String() string { return proto.CompactTextString(m) }
func (*StartResponse) ProtoMessage()    {}
func (*StartResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_01e8311d1fb507af, []int{12}
}

func (m *StartResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartResponse.Unmarshal(m, b)
}
func (m *StartResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartResponse.Marshal(b, m, deterministic)
}
func (m *StartResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartResponse.Merge(m, src)
}
func (m *StartResponse) XXX_Size() int {
	return xxx_messageInfo_StartResponse.Size(m)
}
func (m *StartResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StartResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StartResponse proto.InternalMessageInfo

type AddScriptLogRequest struct {
	ScriptId             string   `protobuf:"bytes,1,opt,name=script_id,json=scriptId,proto3" json:"script_id"`
	RunLog               string   `protobuf:"bytes,2,opt,name=run_log,json=runLog,proto3" json:"run_log"`
	Writer               string   `protobuf:"bytes,3,opt,name=writer,proto3" json:"writer"`
	Database             string   `protobuf:"bytes,4,opt,name=database,proto3" json:"database"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddScriptLogRequest) Reset()         { *m = AddScriptLogRequest{} }
func (m *AddScriptLogRequest) String() string { return proto.CompactTextString(m) }
func (*AddScriptLogRequest) ProtoMessage()    {}
func (*AddScriptLogRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_01e8311d1fb507af, []int{13}
}

func (m *AddScriptLogRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddScriptLogRequest.Unmarshal(m, b)
}
func (m *AddScriptLogRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddScriptLogRequest.Marshal(b, m, deterministic)
}
func (m *AddScriptLogRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddScriptLogRequest.Merge(m, src)
}
func (m *AddScriptLogRequest) XXX_Size() int {
	return xxx_messageInfo_AddScriptLogRequest.Size(m)
}
func (m *AddScriptLogRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddScriptLogRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddScriptLogRequest proto.InternalMessageInfo

func (m *AddScriptLogRequest) GetScriptId() string {
	if m != nil {
		return m.ScriptId
	}
	return ""
}

func (m *AddScriptLogRequest) GetRunLog() string {
	if m != nil {
		return m.RunLog
	}
	return ""
}

func (m *AddScriptLogRequest) GetWriter() string {
	if m != nil {
		return m.Writer
	}
	return ""
}

func (m *AddScriptLogRequest) GetDatabase() string {
	if m != nil {
		return m.Database
	}
	return ""
}

type AddScriptLogResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddScriptLogResponse) Reset()         { *m = AddScriptLogResponse{} }
func (m *AddScriptLogResponse) String() string { return proto.CompactTextString(m) }
func (*AddScriptLogResponse) ProtoMessage()    {}
func (*AddScriptLogResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_01e8311d1fb507af, []int{14}
}

func (m *AddScriptLogResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddScriptLogResponse.Unmarshal(m, b)
}
func (m *AddScriptLogResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddScriptLogResponse.Marshal(b, m, deterministic)
}
func (m *AddScriptLogResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddScriptLogResponse.Merge(m, src)
}
func (m *AddScriptLogResponse) XXX_Size() int {
	return xxx_messageInfo_AddScriptLogResponse.Size(m)
}
func (m *AddScriptLogResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddScriptLogResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddScriptLogResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ScriptJob)(nil), "script.ScriptJob")
	proto.RegisterType((*FindScriptJobsRequest)(nil), "script.FindScriptJobsRequest")
	proto.RegisterType((*FindScriptJobsResponse)(nil), "script.FindScriptJobsResponse")
	proto.RegisterType((*DeleteScriptsRequest)(nil), "script.DeleteScriptsRequest")
	proto.RegisterType((*DeleteScriptsResponse)(nil), "script.DeleteScriptsResponse")
	proto.RegisterType((*FindScriptJobRequest)(nil), "script.FindScriptJobRequest")
	proto.RegisterType((*FindScriptJobResponse)(nil), "script.FindScriptJobResponse")
	proto.RegisterType((*AddRequest)(nil), "script.AddRequest")
	proto.RegisterType((*AddResponse)(nil), "script.AddResponse")
	proto.RegisterType((*ModifyRequest)(nil), "script.ModifyRequest")
	proto.RegisterType((*ModifyResponse)(nil), "script.ModifyResponse")
	proto.RegisterType((*StartRequest)(nil), "script.StartRequest")
	proto.RegisterType((*StartResponse)(nil), "script.StartResponse")
	proto.RegisterType((*AddScriptLogRequest)(nil), "script.AddScriptLogRequest")
	proto.RegisterType((*AddScriptLogResponse)(nil), "script.AddScriptLogResponse")
}

func init() { proto.RegisterFile("script.proto", fileDescriptor_01e8311d1fb507af) }

var fileDescriptor_01e8311d1fb507af = []byte{
	// 703 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x56, 0xcb, 0x4e, 0xdb, 0x4c,
	0x14, 0x26, 0x0e, 0x24, 0xf1, 0xc9, 0x85, 0xff, 0x1f, 0x92, 0x30, 0x0d, 0xd0, 0x22, 0x4b, 0x95,
	0x58, 0xa1, 0x8a, 0xae, 0xba, 0xaa, 0x82, 0x10, 0x12, 0x2d, 0xa5, 0x02, 0xaa, 0x6e, 0xba, 0x88,
	0xc6, 0x9e, 0x01, 0x19, 0x81, 0x9d, 0x7a, 0x26, 0xb4, 0x5e, 0xf5, 0x01, 0xba, 0xeb, 0xba, 0x0f,
	0xd4, 0x77, 0xea, 0xa6, 0x8a, 0x67, 0x3c, 0x1e, 0x0f, 0x4e, 0x9b, 0x55, 0x17, 0xdd, 0xe1, 0xef,
	0xdc, 0xbe, 0xf9, 0xce, 0x25, 0x40, 0x87, 0x07, 0x49, 0x38, 0x15, 0xfb, 0xd3, 0x24, 0x16, 0x31,
	0x6a, 0xc8, 0x2f, 0xef, 0xa7, 0x03, 0xee, 0x65, 0xf6, 0xe7, 0xab, 0xd8, 0x47, 0x5b, 0xe0, 0x4a,
	0x7c, 0x12, 0x52, 0x5c, 0xdb, 0xad, 0xed, 0xb9, 0x17, 0x2d, 0x09, 0x9c, 0x50, 0xf4, 0x04, 0xda,
	0xca, 0x18, 0x91, 0x3b, 0x86, 0x9d, 0xcc, 0x0c, 0x12, 0x3a, 0x23, 0x77, 0xcc, 0x70, 0xa0, 0x8c,
	0x07, 0xb8, 0x6e, 0x3a, 0x1c, 0x31, 0x1e, 0x18, 0x0e, 0x22, 0x9d, 0x32, 0xbc, 0x6a, 0x3a, 0xbc,
	0x4b, 0xa7, 0xa5, 0x0c, 0x44, 0x10, 0xbc, 0x56, 0xca, 0x40, 0x04, 0x31, 0x1c, 0xae, 0x66, 0x51,
	0x80, 0x1b, 0xa6, 0xc3, 0xf1, 0x2c, 0x0a, 0xd0, 0x53, 0xe8, 0x29, 0x87, 0x7b, 0x96, 0xf0, 0x30,
	0x8e, 0x70, 0x33, 0xf3, 0xe9, 0x4a, 0xf4, 0xbd, 0x04, 0xd1, 0x23, 0x68, 0x25, 0xb3, 0x68, 0x72,
	0x1b, 0x5f, 0x73, 0xdc, 0xda, 0xad, 0xef, 0xb9, 0x17, 0xcd, 0x64, 0x16, 0x9d, 0xc6, 0xd7, 0x1c,
	0xed, 0x00, 0x04, 0x09, 0x23, 0x82, 0xd1, 0x09, 0x11, 0xd8, 0xcd, 0xa2, 0x5d, 0x85, 0x8c, 0x85,
	0x69, 0xf6, 0x53, 0x0c, 0x25, 0xf3, 0x61, 0x8a, 0x06, 0xd0, 0x48, 0x48, 0x34, 0x8f, 0x6c, 0x67,
	0xa6, 0xb5, 0x84, 0x44, 0x63, 0x91, 0xc3, 0x7e, 0x8a, 0x3b, 0x1a, 0x3e, 0x4c, 0xbd, 0x6f, 0x35,
	0x18, 0x1c, 0x87, 0x11, 0xd5, 0x1d, 0xe0, 0x17, 0xec, 0xe3, 0x8c, 0x71, 0x61, 0x4b, 0x55, 0x7b,
	0x20, 0xd5, 0xc3, 0x87, 0x3a, 0x55, 0x0f, 0x2d, 0x0a, 0xd7, 0x8d, 0xc2, 0x68, 0x04, 0xad, 0xb9,
	0xc2, 0x3e, 0xe1, 0x79, 0x1b, 0xf4, 0xb7, 0x77, 0x0a, 0x43, 0x9b, 0x13, 0x9f, 0xc6, 0x11, 0x67,
	0xe8, 0x40, 0x93, 0xba, 0x89, 0x7d, 0x8e, 0x6b, 0xbb, 0xf5, 0xbd, 0xf6, 0xc1, 0xff, 0xfb, 0x6a,
	0xb0, 0x74, 0x40, 0xce, 0x73, 0x1e, 0xeb, 0x9d, 0x43, 0xff, 0x88, 0xdd, 0x32, 0xc1, 0xa4, 0x59,
	0x3f, 0xd0, 0x64, 0x50, 0x2b, 0x33, 0x98, 0x6b, 0xac, 0xc7, 0x90, 0x63, 0x27, 0xeb, 0x8f, 0x9b,
	0xcf, 0x21, 0xf7, 0x36, 0x61, 0x60, 0xa5, 0x94, 0xfc, 0xbc, 0xb7, 0xd0, 0x2f, 0x31, 0xcf, 0x6b,
	0xfd, 0x76, 0xac, 0x4d, 0x22, 0x8e, 0x25, 0xc5, 0x89, 0xd5, 0x1e, 0xad, 0xc4, 0x33, 0xcd, 0xf0,
	0x26, 0xf6, 0xb3, 0x94, 0x95, 0x42, 0xb8, 0x5a, 0x08, 0xef, 0x87, 0x03, 0x30, 0xa6, 0x74, 0x29,
	0x4a, 0xff, 0xd2, 0xa6, 0x95, 0xd7, 0xa9, 0x65, 0xaf, 0xd3, 0x10, 0x1a, 0x9f, 0x92, 0x50, 0xb0,
	0x44, 0x6d, 0x9a, 0xfa, 0x2a, 0x75, 0x05, 0xac, 0xae, 0x74, 0xa1, 0x9d, 0x29, 0xa9, 0xba, 0xfe,
	0xd5, 0x81, 0xee, 0x9b, 0x98, 0x86, 0x57, 0xe9, 0x5f, 0x17, 0x37, 0xd3, 0x6e, 0xf5, 0x4f, 0xda,
	0xad, 0x2d, 0xa1, 0x5d, 0xa3, 0x4a, 0xbb, 0x42, 0x9c, 0xe6, 0x42, 0x71, 0x5a, 0x96, 0x38, 0xff,
	0x41, 0x2f, 0x17, 0x43, 0xe9, 0x33, 0x81, 0xce, 0xa5, 0x20, 0x89, 0x58, 0x4a, 0x9d, 0xa2, 0xa4,
	0xb3, 0xb0, 0x64, 0xdd, 0x2a, 0xb9, 0x0e, 0x5d, 0x55, 0x40, 0x55, 0xfc, 0x02, 0x1b, 0x63, 0xaa,
	0xb6, 0xe6, 0x34, 0xbe, 0x5e, 0xaa, 0xf0, 0x26, 0x34, 0xd5, 0x45, 0xce, 0x2b, 0xcb, 0x83, 0x6c,
	0x30, 0xaa, 0x2f, 0x64, 0x64, 0x9f, 0xb0, 0x21, 0xf4, 0xcb, 0x04, 0x24, 0xb1, 0x83, 0xef, 0xab,
	0xd0, 0x95, 0xe8, 0x25, 0x4b, 0xee, 0xc3, 0x80, 0xa1, 0x73, 0xe8, 0x95, 0x8f, 0x1d, 0xda, 0xc9,
	0xd7, 0xb8, 0xf2, 0x30, 0x8f, 0x1e, 0x2f, 0x32, 0xab, 0xb7, 0xaf, 0xa0, 0x33, 0xe8, 0x96, 0x6c,
	0x68, 0xbb, 0x32, 0x24, 0x4f, 0xb8, 0xb3, 0xc0, 0xaa, 0xf3, 0xbd, 0x80, 0x8e, 0x7e, 0xcc, 0x3c,
	0x1d, 0xca, 0x03, 0x8a, 0x73, 0x32, 0xda, 0x28, 0x61, 0x3a, 0xf4, 0x10, 0xd6, 0xe5, 0x30, 0x14,
	0xd1, 0x83, 0xdc, 0xb3, 0xb4, 0x32, 0xa3, 0xa1, 0x0d, 0xeb, 0x1c, 0x2f, 0xa1, 0x97, 0x75, 0xb7,
	0x48, 0xd1, 0xd7, 0x87, 0xce, 0x18, 0xab, 0xd1, 0xc0, 0x42, 0x75, 0x82, 0xd7, 0x06, 0xff, 0x79,
	0x43, 0xb7, 0x0c, 0xae, 0xf6, 0x8c, 0x8c, 0xb6, 0xab, 0x8d, 0x3a, 0xd9, 0x07, 0x18, 0xc9, 0xdb,
	0x7f, 0x34, 0x9b, 0xde, 0x86, 0x01, 0x11, 0x6c, 0x1c, 0xd1, 0x31, 0xa5, 0x27, 0x11, 0x65, 0x9f,
	0x0b, 0xa5, 0xab, 0x7e, 0x72, 0x0a, 0xa5, 0xab, 0x7f, 0x3d, 0x56, 0xfc, 0x46, 0xf6, 0xbf, 0xd1,
	0xf3, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd3, 0x9a, 0x8b, 0x25, 0x2b, 0x09, 0x00, 0x00,
}
