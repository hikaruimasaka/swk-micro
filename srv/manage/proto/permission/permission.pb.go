// Code generated by protoc-gen-go. DO NOT EDIT.
// source: permission.proto

package permission

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

// 权限
type Permission struct {
	PermissionId         string    `protobuf:"bytes,1,opt,name=permission_id,json=permissionId,proto3" json:"permission_id"`
	RoleId               string    `protobuf:"bytes,2,opt,name=role_id,json=roleId,proto3" json:"role_id"`
	PermissionType       string    `protobuf:"bytes,3,opt,name=permission_type,json=permissionType,proto3" json:"permission_type"`
	AppId                string    `protobuf:"bytes,4,opt,name=app_id,json=appId,proto3" json:"app_id"`
	ActionType           string    `protobuf:"bytes,5,opt,name=action_type,json=actionType,proto3" json:"action_type"`
	Actions              []*Action `protobuf:"bytes,6,rep,name=actions,proto3" json:"actions"`
	CreatedAt            string    `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	CreatedBy            string    `protobuf:"bytes,8,opt,name=created_by,json=createdBy,proto3" json:"created_by"`
	UpdatedAt            string    `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	UpdatedBy            string    `protobuf:"bytes,10,opt,name=updated_by,json=updatedBy,proto3" json:"updated_by"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Permission) Reset()         { *m = Permission{} }
func (m *Permission) String() string { return proto.CompactTextString(m) }
func (*Permission) ProtoMessage()    {}
func (*Permission) Descriptor() ([]byte, []int) {
	return fileDescriptor_c837ef01cbda0ad8, []int{0}
}

func (m *Permission) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Permission.Unmarshal(m, b)
}
func (m *Permission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Permission.Marshal(b, m, deterministic)
}
func (m *Permission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Permission.Merge(m, src)
}
func (m *Permission) XXX_Size() int {
	return xxx_messageInfo_Permission.Size(m)
}
func (m *Permission) XXX_DiscardUnknown() {
	xxx_messageInfo_Permission.DiscardUnknown(m)
}

var xxx_messageInfo_Permission proto.InternalMessageInfo

func (m *Permission) GetPermissionId() string {
	if m != nil {
		return m.PermissionId
	}
	return ""
}

func (m *Permission) GetRoleId() string {
	if m != nil {
		return m.RoleId
	}
	return ""
}

func (m *Permission) GetPermissionType() string {
	if m != nil {
		return m.PermissionType
	}
	return ""
}

func (m *Permission) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *Permission) GetActionType() string {
	if m != nil {
		return m.ActionType
	}
	return ""
}

func (m *Permission) GetActions() []*Action {
	if m != nil {
		return m.Actions
	}
	return nil
}

func (m *Permission) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Permission) GetCreatedBy() string {
	if m != nil {
		return m.CreatedBy
	}
	return ""
}

func (m *Permission) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *Permission) GetUpdatedBy() string {
	if m != nil {
		return m.UpdatedBy
	}
	return ""
}

type Action struct {
	ObjectId             string          `protobuf:"bytes,1,opt,name=object_id,json=objectId,proto3" json:"object_id"`
	Fields               []string        `protobuf:"bytes,2,rep,name=fields,proto3" json:"fields"`
	ActionMap            map[string]bool `protobuf:"bytes,3,rep,name=action_map,json=actionMap,proto3" json:"action_map" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Action) Reset()         { *m = Action{} }
func (m *Action) String() string { return proto.CompactTextString(m) }
func (*Action) ProtoMessage()    {}
func (*Action) Descriptor() ([]byte, []int) {
	return fileDescriptor_c837ef01cbda0ad8, []int{1}
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

func (m *Action) GetObjectId() string {
	if m != nil {
		return m.ObjectId
	}
	return ""
}

func (m *Action) GetFields() []string {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *Action) GetActionMap() map[string]bool {
	if m != nil {
		return m.ActionMap
	}
	return nil
}

type FindActionsRequest struct {
	RoleId               []string `protobuf:"bytes,1,rep,name=role_id,json=roleId,proto3" json:"role_id"`
	PermissionType       string   `protobuf:"bytes,2,opt,name=permission_type,json=permissionType,proto3" json:"permission_type"`
	AppId                string   `protobuf:"bytes,3,opt,name=app_id,json=appId,proto3" json:"app_id"`
	ActionType           string   `protobuf:"bytes,4,opt,name=action_type,json=actionType,proto3" json:"action_type"`
	ObjectId             string   `protobuf:"bytes,5,opt,name=object_id,json=objectId,proto3" json:"object_id"`
	Database             string   `protobuf:"bytes,6,opt,name=database,proto3" json:"database"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindActionsRequest) Reset()         { *m = FindActionsRequest{} }
func (m *FindActionsRequest) String() string { return proto.CompactTextString(m) }
func (*FindActionsRequest) ProtoMessage()    {}
func (*FindActionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c837ef01cbda0ad8, []int{2}
}

func (m *FindActionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindActionsRequest.Unmarshal(m, b)
}
func (m *FindActionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindActionsRequest.Marshal(b, m, deterministic)
}
func (m *FindActionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindActionsRequest.Merge(m, src)
}
func (m *FindActionsRequest) XXX_Size() int {
	return xxx_messageInfo_FindActionsRequest.Size(m)
}
func (m *FindActionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindActionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindActionsRequest proto.InternalMessageInfo

func (m *FindActionsRequest) GetRoleId() []string {
	if m != nil {
		return m.RoleId
	}
	return nil
}

func (m *FindActionsRequest) GetPermissionType() string {
	if m != nil {
		return m.PermissionType
	}
	return ""
}

func (m *FindActionsRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *FindActionsRequest) GetActionType() string {
	if m != nil {
		return m.ActionType
	}
	return ""
}

func (m *FindActionsRequest) GetObjectId() string {
	if m != nil {
		return m.ObjectId
	}
	return ""
}

func (m *FindActionsRequest) GetDatabase() string {
	if m != nil {
		return m.Database
	}
	return ""
}

type FindActionsResponse struct {
	Actions              []*Action `protobuf:"bytes,1,rep,name=actions,proto3" json:"actions"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *FindActionsResponse) Reset()         { *m = FindActionsResponse{} }
func (m *FindActionsResponse) String() string { return proto.CompactTextString(m) }
func (*FindActionsResponse) ProtoMessage()    {}
func (*FindActionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c837ef01cbda0ad8, []int{3}
}

func (m *FindActionsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindActionsResponse.Unmarshal(m, b)
}
func (m *FindActionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindActionsResponse.Marshal(b, m, deterministic)
}
func (m *FindActionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindActionsResponse.Merge(m, src)
}
func (m *FindActionsResponse) XXX_Size() int {
	return xxx_messageInfo_FindActionsResponse.Size(m)
}
func (m *FindActionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindActionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindActionsResponse proto.InternalMessageInfo

func (m *FindActionsResponse) GetActions() []*Action {
	if m != nil {
		return m.Actions
	}
	return nil
}

type FindPermissionsRequest struct {
	RoleId               string   `protobuf:"bytes,1,opt,name=role_id,json=roleId,proto3" json:"role_id"`
	Database             string   `protobuf:"bytes,2,opt,name=database,proto3" json:"database"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindPermissionsRequest) Reset()         { *m = FindPermissionsRequest{} }
func (m *FindPermissionsRequest) String() string { return proto.CompactTextString(m) }
func (*FindPermissionsRequest) ProtoMessage()    {}
func (*FindPermissionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c837ef01cbda0ad8, []int{4}
}

func (m *FindPermissionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindPermissionsRequest.Unmarshal(m, b)
}
func (m *FindPermissionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindPermissionsRequest.Marshal(b, m, deterministic)
}
func (m *FindPermissionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindPermissionsRequest.Merge(m, src)
}
func (m *FindPermissionsRequest) XXX_Size() int {
	return xxx_messageInfo_FindPermissionsRequest.Size(m)
}
func (m *FindPermissionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindPermissionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindPermissionsRequest proto.InternalMessageInfo

func (m *FindPermissionsRequest) GetRoleId() string {
	if m != nil {
		return m.RoleId
	}
	return ""
}

func (m *FindPermissionsRequest) GetDatabase() string {
	if m != nil {
		return m.Database
	}
	return ""
}

type FindPermissionsResponse struct {
	Permission           []*Permission `protobuf:"bytes,1,rep,name=permission,proto3" json:"permission"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *FindPermissionsResponse) Reset()         { *m = FindPermissionsResponse{} }
func (m *FindPermissionsResponse) String() string { return proto.CompactTextString(m) }
func (*FindPermissionsResponse) ProtoMessage()    {}
func (*FindPermissionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c837ef01cbda0ad8, []int{5}
}

func (m *FindPermissionsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindPermissionsResponse.Unmarshal(m, b)
}
func (m *FindPermissionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindPermissionsResponse.Marshal(b, m, deterministic)
}
func (m *FindPermissionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindPermissionsResponse.Merge(m, src)
}
func (m *FindPermissionsResponse) XXX_Size() int {
	return xxx_messageInfo_FindPermissionsResponse.Size(m)
}
func (m *FindPermissionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindPermissionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindPermissionsResponse proto.InternalMessageInfo

func (m *FindPermissionsResponse) GetPermission() []*Permission {
	if m != nil {
		return m.Permission
	}
	return nil
}

func init() {
	proto.RegisterType((*Permission)(nil), "permission.Permission")
	proto.RegisterType((*Action)(nil), "permission.Action")
	proto.RegisterMapType((map[string]bool)(nil), "permission.Action.ActionMapEntry")
	proto.RegisterType((*FindActionsRequest)(nil), "permission.FindActionsRequest")
	proto.RegisterType((*FindActionsResponse)(nil), "permission.FindActionsResponse")
	proto.RegisterType((*FindPermissionsRequest)(nil), "permission.FindPermissionsRequest")
	proto.RegisterType((*FindPermissionsResponse)(nil), "permission.FindPermissionsResponse")
}

func init() { proto.RegisterFile("permission.proto", fileDescriptor_c837ef01cbda0ad8) }

var fileDescriptor_c837ef01cbda0ad8 = []byte{
	// 489 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x49, 0xb2, 0xa6, 0xcd, 0x2b, 0x6c, 0xe3, 0x01, 0x9d, 0x55, 0x04, 0x2b, 0xd9, 0x81,
	0x1d, 0xd0, 0x0e, 0x43, 0x42, 0x08, 0x71, 0x60, 0x43, 0x20, 0xf5, 0x30, 0x69, 0x04, 0x8e, 0x48,
	0x95, 0x53, 0x1b, 0x29, 0xd0, 0x35, 0x26, 0x76, 0x27, 0xe5, 0xc3, 0x71, 0xe2, 0xc0, 0x27, 0xe0,
	0xfb, 0xa0, 0xd8, 0xc9, 0xec, 0x24, 0x6c, 0x3b, 0xb5, 0xef, 0xfd, 0xfc, 0x9e, 0xfe, 0xef, 0xfd,
	0xed, 0xc0, 0xae, 0xe0, 0xc5, 0x45, 0x26, 0x65, 0x96, 0xaf, 0x8f, 0x44, 0x91, 0xab, 0x1c, 0xc1,
	0x66, 0xe2, 0xbf, 0x3e, 0xc0, 0xf9, 0x55, 0x88, 0x07, 0x70, 0xcf, 0xc2, 0x45, 0xc6, 0x88, 0x37,
	0xf3, 0x0e, 0xa3, 0xe4, 0xae, 0x4d, 0xce, 0x19, 0xee, 0xc1, 0xb0, 0xc8, 0x57, 0xbc, 0xc2, 0xbe,
	0xc6, 0x61, 0x15, 0xce, 0x19, 0x3e, 0x87, 0x1d, 0xa7, 0x5a, 0x95, 0x82, 0x93, 0x40, 0x1f, 0xd8,
	0xb6, 0xe9, 0x2f, 0xa5, 0xe0, 0xf8, 0x08, 0x42, 0x2a, 0x44, 0xd5, 0x60, 0x4b, 0xf3, 0x01, 0x15,
	0x62, 0xce, 0x70, 0x1f, 0xc6, 0x74, 0xa9, 0xae, 0x6a, 0x07, 0x9a, 0x81, 0x49, 0xe9, 0xba, 0x17,
	0x30, 0x34, 0x91, 0x24, 0xe1, 0x2c, 0x38, 0x1c, 0x1f, 0xe3, 0x91, 0x33, 0xdd, 0x89, 0x46, 0x49,
	0x73, 0x04, 0x9f, 0x00, 0x2c, 0x0b, 0x4e, 0x15, 0x67, 0x0b, 0xaa, 0xc8, 0x50, 0x77, 0x8b, 0xea,
	0xcc, 0x89, 0x72, 0x71, 0x5a, 0x92, 0x51, 0x0b, 0x9f, 0x96, 0x15, 0xde, 0x08, 0xd6, 0x54, 0x47,
	0x06, 0xd7, 0x19, 0x53, 0xdd, 0xe0, 0xb4, 0x24, 0xd0, 0xc2, 0xa7, 0x65, 0xfc, 0xcb, 0x83, 0xd0,
	0xe8, 0xc1, 0xc7, 0x10, 0xe5, 0xe9, 0x77, 0xbe, 0x54, 0x76, 0x9f, 0x23, 0x93, 0x98, 0x33, 0x9c,
	0x40, 0xf8, 0x2d, 0xe3, 0x2b, 0x26, 0x89, 0x3f, 0x0b, 0xaa, 0x55, 0x9a, 0x08, 0xdf, 0x41, 0x3d,
	0xf7, 0xe2, 0x82, 0x0a, 0x12, 0xe8, 0x61, 0x9f, 0xf5, 0x87, 0xad, 0x7f, 0xce, 0xa8, 0xf8, 0xb0,
	0x56, 0x45, 0x99, 0x44, 0xb4, 0x89, 0xa7, 0x6f, 0x61, 0xbb, 0x0d, 0x71, 0x17, 0x82, 0x1f, 0xbc,
	0xac, 0x25, 0x54, 0x7f, 0xf1, 0x21, 0x0c, 0x2e, 0xe9, 0x6a, 0xc3, 0xb5, 0x8f, 0xa3, 0xc4, 0x04,
	0x6f, 0xfc, 0xd7, 0x5e, 0xfc, 0xc7, 0x03, 0xfc, 0x98, 0xad, 0x99, 0x69, 0x21, 0x13, 0xfe, 0x73,
	0xc3, 0xa5, 0x72, 0xad, 0xf7, 0x8c, 0xde, 0xeb, 0xad, 0xf7, 0x6f, 0xb1, 0x3e, 0xb8, 0xc1, 0xfa,
	0xad, 0x9e, 0xf5, 0xad, 0x2d, 0x0e, 0x3a, 0x5b, 0x9c, 0xc2, 0x88, 0x51, 0x45, 0x53, 0x2a, 0x39,
	0x09, 0x0d, 0x6b, 0xe2, 0xf8, 0x3d, 0x3c, 0x68, 0x0d, 0x22, 0x45, 0xbe, 0x96, 0xad, 0xab, 0xe4,
	0xdd, 0x7a, 0x95, 0xe2, 0x33, 0x98, 0x54, 0x4d, 0xec, 0x4b, 0xf9, 0xff, 0x46, 0xdc, 0xc7, 0xe0,
	0x6a, 0xf2, 0x3b, 0x9a, 0x3e, 0xc1, 0x5e, 0xaf, 0x5d, 0xad, 0xeb, 0x15, 0x38, 0xcf, 0xb3, 0x96,
	0x36, 0x71, 0xa5, 0xd9, 0xa2, 0xc4, 0x39, 0x79, 0xfc, 0xdb, 0x83, 0xfb, 0x16, 0x7d, 0xe6, 0xc5,
	0x65, 0xb6, 0xe4, 0x78, 0x0e, 0x63, 0x67, 0x78, 0x7c, 0xea, 0x36, 0xea, 0xdb, 0x3b, 0xdd, 0xbf,
	0x96, 0x1b, 0x75, 0xf1, 0x1d, 0xfc, 0x0a, 0x3b, 0x1d, 0xe9, 0x18, 0x77, 0xab, 0xfa, 0x6b, 0x9a,
	0x1e, 0xdc, 0x78, 0xa6, 0xe9, 0x9e, 0x86, 0xfa, 0x0b, 0xf5, 0xf2, 0x5f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x63, 0x6b, 0xb8, 0x7d, 0xb5, 0x04, 0x00, 0x00,
}
