// Code generated by protoc-gen-go. DO NOT EDIT.
// source: task.proto

package task

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

type File struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *File) Reset()         { *m = File{} }
func (m *File) String() string { return proto.CompactTextString(m) }
func (*File) ProtoMessage()    {}
func (*File) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{0}
}

func (m *File) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_File.Unmarshal(m, b)
}
func (m *File) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_File.Marshal(b, m, deterministic)
}
func (m *File) XXX_Merge(src proto.Message) {
	xxx_messageInfo_File.Merge(m, src)
}
func (m *File) XXX_Size() int {
	return xxx_messageInfo_File.Size(m)
}
func (m *File) XXX_DiscardUnknown() {
	xxx_messageInfo_File.DiscardUnknown(m)
}

var xxx_messageInfo_File proto.InternalMessageInfo

func (m *File) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *File) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// 任务数据
type Task struct {
	JobId                string   `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id"`
	JobName              string   `protobuf:"bytes,2,opt,name=job_name,json=jobName,proto3" json:"job_name"`
	Origin               string   `protobuf:"bytes,3,opt,name=origin,proto3" json:"origin"`
	UserId               string   `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id"`
	ShowProgress         bool     `protobuf:"varint,5,opt,name=show_progress,json=showProgress,proto3" json:"show_progress"`
	Progress             int64    `protobuf:"varint,6,opt,name=progress,proto3" json:"progress"`
	StartTime            string   `protobuf:"bytes,7,opt,name=start_time,json=startTime,proto3" json:"start_time"`
	EndTime              string   `protobuf:"bytes,8,opt,name=end_time,json=endTime,proto3" json:"end_time"`
	Message              string   `protobuf:"bytes,9,opt,name=message,proto3" json:"message"`
	File                 *File    `protobuf:"bytes,10,opt,name=file,proto3" json:"file"`
	ErrorFile            *File    `protobuf:"bytes,11,opt,name=error_file,json=errorFile,proto3" json:"error_file"`
	TaskType             string   `protobuf:"bytes,12,opt,name=task_type,json=taskType,proto3" json:"task_type"`
	Steps                []string `protobuf:"bytes,13,rep,name=steps,proto3" json:"steps"`
	CurrentStep          string   `protobuf:"bytes,14,opt,name=current_step,json=currentStep,proto3" json:"current_step"`
	ScheduleId           string   `protobuf:"bytes,15,opt,name=schedule_id,json=scheduleId,proto3" json:"schedule_id"`
	AppId                string   `protobuf:"bytes,17,opt,name=app_id,json=appId,proto3" json:"app_id"`
	Insert               int64    `protobuf:"varint,18,opt,name=insert,proto3" json:"insert"`
	Update               int64    `protobuf:"varint,19,opt,name=update,proto3" json:"update"`
	Total                int64    `protobuf:"varint,20,opt,name=total,proto3" json:"total"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Task) Reset()         { *m = Task{} }
func (m *Task) String() string { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()    {}
func (*Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{1}
}

func (m *Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Task.Unmarshal(m, b)
}
func (m *Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Task.Marshal(b, m, deterministic)
}
func (m *Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Task.Merge(m, src)
}
func (m *Task) XXX_Size() int {
	return xxx_messageInfo_Task.Size(m)
}
func (m *Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Task proto.InternalMessageInfo

func (m *Task) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

func (m *Task) GetJobName() string {
	if m != nil {
		return m.JobName
	}
	return ""
}

func (m *Task) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *Task) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Task) GetShowProgress() bool {
	if m != nil {
		return m.ShowProgress
	}
	return false
}

func (m *Task) GetProgress() int64 {
	if m != nil {
		return m.Progress
	}
	return 0
}

func (m *Task) GetStartTime() string {
	if m != nil {
		return m.StartTime
	}
	return ""
}

func (m *Task) GetEndTime() string {
	if m != nil {
		return m.EndTime
	}
	return ""
}

func (m *Task) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Task) GetFile() *File {
	if m != nil {
		return m.File
	}
	return nil
}

func (m *Task) GetErrorFile() *File {
	if m != nil {
		return m.ErrorFile
	}
	return nil
}

func (m *Task) GetTaskType() string {
	if m != nil {
		return m.TaskType
	}
	return ""
}

func (m *Task) GetSteps() []string {
	if m != nil {
		return m.Steps
	}
	return nil
}

func (m *Task) GetCurrentStep() string {
	if m != nil {
		return m.CurrentStep
	}
	return ""
}

func (m *Task) GetScheduleId() string {
	if m != nil {
		return m.ScheduleId
	}
	return ""
}

func (m *Task) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *Task) GetInsert() int64 {
	if m != nil {
		return m.Insert
	}
	return 0
}

func (m *Task) GetUpdate() int64 {
	if m != nil {
		return m.Update
	}
	return 0
}

func (m *Task) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

// 查找多条记录
type TasksRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id"`
	PageIndex            int64    `protobuf:"varint,2,opt,name=page_index,json=pageIndex,proto3" json:"page_index"`
	PageSize             int64    `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Database             string   `protobuf:"bytes,4,opt,name=database,proto3" json:"database"`
	ScheduleId           string   `protobuf:"bytes,5,opt,name=schedule_id,json=scheduleId,proto3" json:"schedule_id"`
	AppId                string   `protobuf:"bytes,6,opt,name=app_id,json=appId,proto3" json:"app_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TasksRequest) Reset()         { *m = TasksRequest{} }
func (m *TasksRequest) String() string { return proto.CompactTextString(m) }
func (*TasksRequest) ProtoMessage()    {}
func (*TasksRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{2}
}

func (m *TasksRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TasksRequest.Unmarshal(m, b)
}
func (m *TasksRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TasksRequest.Marshal(b, m, deterministic)
}
func (m *TasksRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TasksRequest.Merge(m, src)
}
func (m *TasksRequest) XXX_Size() int {
	return xxx_messageInfo_TasksRequest.Size(m)
}
func (m *TasksRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TasksRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TasksRequest proto.InternalMessageInfo

func (m *TasksRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *TasksRequest) GetPageIndex() int64 {
	if m != nil {
		return m.PageIndex
	}
	return 0
}

func (m *TasksRequest) GetPageSize() int64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *TasksRequest) GetDatabase() string {
	if m != nil {
		return m.Database
	}
	return ""
}

func (m *TasksRequest) GetScheduleId() string {
	if m != nil {
		return m.ScheduleId
	}
	return ""
}

func (m *TasksRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

type TasksResponse struct {
	Tasks                []*Task  `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks"`
	Total                int64    `protobuf:"varint,2,opt,name=total,proto3" json:"total"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TasksResponse) Reset()         { *m = TasksResponse{} }
func (m *TasksResponse) String() string { return proto.CompactTextString(m) }
func (*TasksResponse) ProtoMessage()    {}
func (*TasksResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{3}
}

func (m *TasksResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TasksResponse.Unmarshal(m, b)
}
func (m *TasksResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TasksResponse.Marshal(b, m, deterministic)
}
func (m *TasksResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TasksResponse.Merge(m, src)
}
func (m *TasksResponse) XXX_Size() int {
	return xxx_messageInfo_TasksResponse.Size(m)
}
func (m *TasksResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TasksResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TasksResponse proto.InternalMessageInfo

func (m *TasksResponse) GetTasks() []*Task {
	if m != nil {
		return m.Tasks
	}
	return nil
}

func (m *TasksResponse) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

// 查询单条记录
type TaskRequest struct {
	JobId                string   `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id"`
	Database             string   `protobuf:"bytes,2,opt,name=database,proto3" json:"database"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskRequest) Reset()         { *m = TaskRequest{} }
func (m *TaskRequest) String() string { return proto.CompactTextString(m) }
func (*TaskRequest) ProtoMessage()    {}
func (*TaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{4}
}

func (m *TaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskRequest.Unmarshal(m, b)
}
func (m *TaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskRequest.Marshal(b, m, deterministic)
}
func (m *TaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskRequest.Merge(m, src)
}
func (m *TaskRequest) XXX_Size() int {
	return xxx_messageInfo_TaskRequest.Size(m)
}
func (m *TaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TaskRequest proto.InternalMessageInfo

func (m *TaskRequest) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

func (m *TaskRequest) GetDatabase() string {
	if m != nil {
		return m.Database
	}
	return ""
}

type TaskResponse struct {
	Task                 *Task    `protobuf:"bytes,1,opt,name=task,proto3" json:"task"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskResponse) Reset()         { *m = TaskResponse{} }
func (m *TaskResponse) String() string { return proto.CompactTextString(m) }
func (*TaskResponse) ProtoMessage()    {}
func (*TaskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{5}
}

func (m *TaskResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskResponse.Unmarshal(m, b)
}
func (m *TaskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskResponse.Marshal(b, m, deterministic)
}
func (m *TaskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskResponse.Merge(m, src)
}
func (m *TaskResponse) XXX_Size() int {
	return xxx_messageInfo_TaskResponse.Size(m)
}
func (m *TaskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TaskResponse proto.InternalMessageInfo

func (m *TaskResponse) GetTask() *Task {
	if m != nil {
		return m.Task
	}
	return nil
}

// 添加数据
type AddRequest struct {
	JobId                string   `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id"`
	JobName              string   `protobuf:"bytes,2,opt,name=job_name,json=jobName,proto3" json:"job_name"`
	Origin               string   `protobuf:"bytes,3,opt,name=origin,proto3" json:"origin"`
	UserId               string   `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id"`
	ShowProgress         bool     `protobuf:"varint,5,opt,name=show_progress,json=showProgress,proto3" json:"show_progress"`
	Progress             int64    `protobuf:"varint,6,opt,name=progress,proto3" json:"progress"`
	StartTime            string   `protobuf:"bytes,7,opt,name=start_time,json=startTime,proto3" json:"start_time"`
	Message              string   `protobuf:"bytes,9,opt,name=message,proto3" json:"message"`
	TaskType             string   `protobuf:"bytes,12,opt,name=task_type,json=taskType,proto3" json:"task_type"`
	Steps                []string `protobuf:"bytes,13,rep,name=steps,proto3" json:"steps"`
	CurrentStep          string   `protobuf:"bytes,14,opt,name=current_step,json=currentStep,proto3" json:"current_step"`
	Database             string   `protobuf:"bytes,15,opt,name=database,proto3" json:"database"`
	ScheduleId           string   `protobuf:"bytes,16,opt,name=schedule_id,json=scheduleId,proto3" json:"schedule_id"`
	AppId                string   `protobuf:"bytes,17,opt,name=app_id,json=appId,proto3" json:"app_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddRequest) Reset()         { *m = AddRequest{} }
func (m *AddRequest) String() string { return proto.CompactTextString(m) }
func (*AddRequest) ProtoMessage()    {}
func (*AddRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{6}
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

func (m *AddRequest) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

func (m *AddRequest) GetJobName() string {
	if m != nil {
		return m.JobName
	}
	return ""
}

func (m *AddRequest) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *AddRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *AddRequest) GetShowProgress() bool {
	if m != nil {
		return m.ShowProgress
	}
	return false
}

func (m *AddRequest) GetProgress() int64 {
	if m != nil {
		return m.Progress
	}
	return 0
}

func (m *AddRequest) GetStartTime() string {
	if m != nil {
		return m.StartTime
	}
	return ""
}

func (m *AddRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *AddRequest) GetTaskType() string {
	if m != nil {
		return m.TaskType
	}
	return ""
}

func (m *AddRequest) GetSteps() []string {
	if m != nil {
		return m.Steps
	}
	return nil
}

func (m *AddRequest) GetCurrentStep() string {
	if m != nil {
		return m.CurrentStep
	}
	return ""
}

func (m *AddRequest) GetDatabase() string {
	if m != nil {
		return m.Database
	}
	return ""
}

func (m *AddRequest) GetScheduleId() string {
	if m != nil {
		return m.ScheduleId
	}
	return ""
}

func (m *AddRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

type AddResponse struct {
	JobId                string   `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddResponse) Reset()         { *m = AddResponse{} }
func (m *AddResponse) String() string { return proto.CompactTextString(m) }
func (*AddResponse) ProtoMessage()    {}
func (*AddResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{7}
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

func (m *AddResponse) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

// 更新记录
type ModifyRequest struct {
	JobId                string   `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id"`
	Progress             int64    `protobuf:"varint,2,opt,name=progress,proto3" json:"progress"`
	EndTime              string   `protobuf:"bytes,3,opt,name=end_time,json=endTime,proto3" json:"end_time"`
	Message              string   `protobuf:"bytes,4,opt,name=message,proto3" json:"message"`
	File                 *File    `protobuf:"bytes,5,opt,name=file,proto3" json:"file"`
	ErrorFile            *File    `protobuf:"bytes,6,opt,name=error_file,json=errorFile,proto3" json:"error_file"`
	CurrentStep          string   `protobuf:"bytes,7,opt,name=current_step,json=currentStep,proto3" json:"current_step"`
	Insert               int64    `protobuf:"varint,9,opt,name=insert,proto3" json:"insert"`
	Update               int64    `protobuf:"varint,10,opt,name=update,proto3" json:"update"`
	Total                int64    `protobuf:"varint,11,opt,name=total,proto3" json:"total"`
	Database             string   `protobuf:"bytes,8,opt,name=database,proto3" json:"database"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ModifyRequest) Reset()         { *m = ModifyRequest{} }
func (m *ModifyRequest) String() string { return proto.CompactTextString(m) }
func (*ModifyRequest) ProtoMessage()    {}
func (*ModifyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{8}
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

func (m *ModifyRequest) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

func (m *ModifyRequest) GetProgress() int64 {
	if m != nil {
		return m.Progress
	}
	return 0
}

func (m *ModifyRequest) GetEndTime() string {
	if m != nil {
		return m.EndTime
	}
	return ""
}

func (m *ModifyRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ModifyRequest) GetFile() *File {
	if m != nil {
		return m.File
	}
	return nil
}

func (m *ModifyRequest) GetErrorFile() *File {
	if m != nil {
		return m.ErrorFile
	}
	return nil
}

func (m *ModifyRequest) GetCurrentStep() string {
	if m != nil {
		return m.CurrentStep
	}
	return ""
}

func (m *ModifyRequest) GetInsert() int64 {
	if m != nil {
		return m.Insert
	}
	return 0
}

func (m *ModifyRequest) GetUpdate() int64 {
	if m != nil {
		return m.Update
	}
	return 0
}

func (m *ModifyRequest) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
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
	return fileDescriptor_ce5d8dd45b4a91ff, []int{9}
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
	JobId                string   `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id"`
	AppId                string   `protobuf:"bytes,2,opt,name=app_id,json=appId,proto3" json:"app_id"`
	UserId               string   `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{10}
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

func (m *DeleteRequest) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

func (m *DeleteRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *DeleteRequest) GetUserId() string {
	if m != nil {
		return m.UserId
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
	return fileDescriptor_ce5d8dd45b4a91ff, []int{11}
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
	proto.RegisterType((*File)(nil), "task.File")
	proto.RegisterType((*Task)(nil), "task.Task")
	proto.RegisterType((*TasksRequest)(nil), "task.TasksRequest")
	proto.RegisterType((*TasksResponse)(nil), "task.TasksResponse")
	proto.RegisterType((*TaskRequest)(nil), "task.TaskRequest")
	proto.RegisterType((*TaskResponse)(nil), "task.TaskResponse")
	proto.RegisterType((*AddRequest)(nil), "task.AddRequest")
	proto.RegisterType((*AddResponse)(nil), "task.AddResponse")
	proto.RegisterType((*ModifyRequest)(nil), "task.ModifyRequest")
	proto.RegisterType((*ModifyResponse)(nil), "task.ModifyResponse")
	proto.RegisterType((*DeleteRequest)(nil), "task.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "task.DeleteResponse")
}

func init() { proto.RegisterFile("task.proto", fileDescriptor_ce5d8dd45b4a91ff) }

var fileDescriptor_ce5d8dd45b4a91ff = []byte{
	// 768 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x56, 0x4f, 0x6f, 0x13, 0x3b,
	0x10, 0xef, 0x66, 0x37, 0x7f, 0x76, 0x92, 0xb4, 0xa9, 0xdb, 0xf7, 0x9e, 0x5f, 0x9e, 0xda, 0x17,
	0x16, 0x0e, 0x41, 0x42, 0x15, 0x6a, 0x25, 0x24, 0x6e, 0x54, 0x42, 0x45, 0x39, 0x80, 0x50, 0x5a,
	0x71, 0x8d, 0x9c, 0x7a, 0x9a, 0x6e, 0x9b, 0xec, 0x2e, 0xb6, 0x03, 0xb4, 0x5f, 0x83, 0x33, 0x17,
	0xbe, 0x04, 0x1f, 0x81, 0xaf, 0x85, 0x6c, 0x6f, 0x12, 0x6f, 0xda, 0xb4, 0xbd, 0x70, 0xe0, 0xb6,
	0xf3, 0x1b, 0x8f, 0x3d, 0xf3, 0xfb, 0xcd, 0x8c, 0x16, 0x40, 0x31, 0x79, 0xb9, 0x97, 0x89, 0x54,
	0xa5, 0x24, 0xd0, 0xdf, 0xd1, 0x33, 0x08, 0x8e, 0xe2, 0x31, 0x92, 0x16, 0xf8, 0x53, 0x31, 0xa6,
	0x5e, 0xc7, 0xeb, 0x86, 0x7d, 0xfd, 0x49, 0x08, 0x04, 0x09, 0x9b, 0x20, 0x2d, 0x19, 0xc8, 0x7c,
	0x47, 0xdf, 0x02, 0x08, 0x4e, 0x98, 0xbc, 0x24, 0x7f, 0x41, 0xe5, 0x22, 0x1d, 0x0e, 0x62, 0x9e,
	0x47, 0x94, 0x2f, 0xd2, 0x61, 0x8f, 0x93, 0x7f, 0xa1, 0xa6, 0x61, 0x27, 0xae, 0x7a, 0x91, 0x0e,
	0xdf, 0xb1, 0x09, 0x92, 0xbf, 0xa1, 0x92, 0x8a, 0x78, 0x14, 0x27, 0xd4, 0x37, 0x8e, 0xdc, 0x22,
	0xff, 0x40, 0x75, 0x2a, 0x51, 0xe8, 0xab, 0x02, 0xeb, 0xd0, 0x66, 0x8f, 0x93, 0xc7, 0xd0, 0x94,
	0xe7, 0xe9, 0xe7, 0x41, 0x26, 0xd2, 0x91, 0x40, 0x29, 0x69, 0xb9, 0xe3, 0x75, 0x6b, 0xfd, 0x86,
	0x06, 0xdf, 0xe7, 0x18, 0x69, 0x43, 0x6d, 0xee, 0xaf, 0x74, 0xbc, 0xae, 0xdf, 0x9f, 0xdb, 0x64,
	0x07, 0x40, 0x2a, 0x26, 0xd4, 0x40, 0xc5, 0x13, 0xa4, 0x55, 0x73, 0x79, 0x68, 0x90, 0x93, 0x78,
	0x82, 0x3a, 0x57, 0x4c, 0xb8, 0x75, 0xd6, 0x6c, 0xae, 0x98, 0x70, 0xe3, 0xa2, 0x50, 0x9d, 0xa0,
	0x94, 0x6c, 0x84, 0x34, 0xb4, 0x9e, 0xdc, 0x24, 0xbb, 0x10, 0x9c, 0xc5, 0x63, 0xa4, 0xd0, 0xf1,
	0xba, 0xf5, 0x7d, 0xd8, 0x33, 0x7c, 0x6a, 0x02, 0xfb, 0x06, 0x27, 0x4f, 0x01, 0x50, 0x88, 0x54,
	0x0c, 0xcc, 0xa9, 0xfa, 0x8d, 0x53, 0xa1, 0xf1, 0x1a, 0xc6, 0xff, 0x83, 0x50, 0xe3, 0x03, 0x75,
	0x95, 0x21, 0x6d, 0x98, 0x67, 0x6a, 0x1a, 0x38, 0xb9, 0xca, 0x90, 0x6c, 0x43, 0x59, 0x2a, 0xcc,
	0x24, 0x6d, 0x76, 0x7c, 0x4d, 0xaf, 0x31, 0xc8, 0x23, 0x68, 0x9c, 0x4e, 0x85, 0xc0, 0x44, 0x0d,
	0x34, 0x40, 0xd7, 0x4d, 0x54, 0x3d, 0xc7, 0x8e, 0x15, 0x66, 0xe4, 0x7f, 0xa8, 0xcb, 0xd3, 0x73,
	0xe4, 0xd3, 0x31, 0x6a, 0x4a, 0x37, 0xcc, 0x09, 0x98, 0x41, 0x3d, 0xae, 0x95, 0x63, 0x59, 0xa6,
	0x7d, 0x9b, 0x56, 0x39, 0x96, 0x65, 0x3d, 0xae, 0xe5, 0x89, 0x13, 0x89, 0x42, 0x51, 0x62, 0x68,
	0xcc, 0x2d, 0x8d, 0x4f, 0x33, 0xce, 0x14, 0xd2, 0x2d, 0x8b, 0x5b, 0x4b, 0x27, 0xa8, 0x52, 0xc5,
	0xc6, 0x74, 0xdb, 0xc0, 0xd6, 0x88, 0x7e, 0x78, 0xd0, 0xd0, 0xfd, 0x21, 0xfb, 0xf8, 0x71, 0x8a,
	0x52, 0xb9, 0xea, 0x7a, 0x05, 0x75, 0x77, 0x00, 0x32, 0x36, 0xc2, 0x41, 0x9c, 0x70, 0xfc, 0x62,
	0x7a, 0xc5, 0xef, 0x87, 0x1a, 0xe9, 0x69, 0x40, 0x93, 0x63, 0xdc, 0x32, 0xbe, 0x46, 0xd3, 0x30,
	0x5a, 0x58, 0x36, 0xc2, 0xe3, 0xf8, 0x1a, 0xb5, 0xe8, 0x9c, 0x29, 0x36, 0x64, 0x12, 0xf3, 0x9e,
	0x99, 0xdb, 0xcb, 0xf5, 0x97, 0xef, 0xa8, 0xbf, 0xe2, 0xd4, 0x1f, 0xbd, 0x81, 0x66, 0x9e, 0xb8,
	0xcc, 0xd2, 0x44, 0x22, 0xe9, 0x40, 0x59, 0xab, 0x21, 0xa9, 0xd7, 0xf1, 0x17, 0x22, 0xea, 0x33,
	0x7d, 0xeb, 0x58, 0x50, 0x50, 0x72, 0x29, 0x78, 0x05, 0x75, 0x73, 0x28, 0x27, 0x60, 0xc5, 0xa0,
	0xb8, 0x25, 0x94, 0x8a, 0x25, 0x44, 0x7b, 0x96, 0xc3, 0x79, 0x26, 0xbb, 0x60, 0x46, 0xd5, 0x5c,
	0x50, 0x4c, 0xc4, 0x8e, 0xf0, 0x57, 0x1f, 0xe0, 0x90, 0xf3, 0x7b, 0x5e, 0xfc, 0x93, 0x46, 0x73,
	0xf5, 0xfc, 0xfd, 0x9e, 0xa1, 0x71, 0xd5, 0xd8, 0xb8, 0xbb, 0xa1, 0x5a, 0x0f, 0x1c, 0xa8, 0xe8,
	0x09, 0xd4, 0x8d, 0x28, 0xb9, 0x88, 0xb7, 0xab, 0x12, 0xfd, 0x2c, 0x41, 0xf3, 0x6d, 0xca, 0xe3,
	0xb3, 0xab, 0xfb, 0x1b, 0x66, 0xce, 0x66, 0x69, 0x89, 0x4d, 0x77, 0x93, 0xf9, 0x2b, 0x37, 0x59,
	0x70, 0xfb, 0x26, 0x2b, 0x3f, 0x68, 0x93, 0x55, 0xee, 0xda, 0x64, 0xcb, 0x0c, 0x57, 0x6f, 0x32,
	0xbc, 0x58, 0x2f, 0xe1, 0x8a, 0xf5, 0x02, 0xb7, 0xaf, 0x97, 0xba, 0x33, 0x5b, 0x05, 0x9d, 0x6a,
	0x4b, 0x53, 0xd3, 0x82, 0xf5, 0x19, 0x91, 0x96, 0xf2, 0xe8, 0x03, 0x34, 0x5f, 0xe3, 0x18, 0x15,
	0xde, 0x43, 0xed, 0x42, 0xc0, 0x92, 0xbb, 0x11, 0x9d, 0xee, 0xf7, 0xdd, 0xee, 0xd7, 0x2f, 0xcd,
	0xee, 0xb5, 0x2f, 0xed, 0x7f, 0x2f, 0xd9, 0xa1, 0x3f, 0x46, 0xf1, 0x29, 0x3e, 0x45, 0xf2, 0x02,
	0xc2, 0xa3, 0x38, 0xe1, 0x66, 0xa1, 0x10, 0xb2, 0x18, 0xd8, 0xd9, 0x5a, 0x6c, 0x6f, 0x15, 0xb0,
	0x3c, 0xdf, 0x35, 0x72, 0x00, 0xb5, 0x59, 0x1c, 0xd9, 0x74, 0xe6, 0x3c, 0x8f, 0x22, 0x2e, 0x34,
	0x0f, 0x7a, 0x0e, 0xd5, 0x43, 0x6e, 0x63, 0x5a, 0xf6, 0xc0, 0x62, 0x19, 0xb4, 0x37, 0x1d, 0x64,
	0x1e, 0xf1, 0x12, 0xc0, 0x52, 0x65, 0x82, 0xf2, 0x5c, 0x0a, 0x5d, 0xd8, 0xde, 0x2e, 0x82, 0x6e,
	0xa8, 0xad, 0xdd, 0x0d, 0x2d, 0xb0, 0x3c, 0x0b, 0x2d, 0x52, 0x14, 0xad, 0x0d, 0x2b, 0xe6, 0xb7,
	0xe3, 0xe0, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf9, 0x4b, 0xa4, 0x34, 0x84, 0x08, 0x00, 0x00,
}
