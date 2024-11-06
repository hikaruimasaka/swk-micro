// Code generated by protoc-gen-go. DO NOT EDIT.
// source: node.proto

package node

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
type Node struct {
	NodeId   string `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id"`
	NodeName string `protobuf:"bytes,4,opt,name=node_name,json=nodeName,proto3" json:"node_name"`
	WfId     string `protobuf:"bytes,2,opt,name=wf_id,json=wfId,proto3" json:"wf_id"`
	NodeType string `protobuf:"bytes,3,opt,name=node_type,json=nodeType,proto3" json:"node_type"`
	// repeated Condition condition =4; // 分支条件
	PrevNode             string   `protobuf:"bytes,5,opt,name=prev_node,json=prevNode,proto3" json:"prev_node"`
	NextNode             string   `protobuf:"bytes,6,opt,name=next_node,json=nextNode,proto3" json:"next_node"`
	Assignees            []string `protobuf:"bytes,7,rep,name=assignees,proto3" json:"assignees"`
	ActType              string   `protobuf:"bytes,8,opt,name=act_type,json=actType,proto3" json:"act_type"`
	NodeGroupId          string   `protobuf:"bytes,11,opt,name=node_group_id,json=nodeGroupId,proto3" json:"node_group_id"`
	CreatedAt            string   `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	CreatedBy            string   `protobuf:"bytes,10,opt,name=created_by,json=createdBy,proto3" json:"created_by"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c843d59d2d938e7, []int{0}
}

func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (m *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(m, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *Node) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *Node) GetWfId() string {
	if m != nil {
		return m.WfId
	}
	return ""
}

func (m *Node) GetNodeType() string {
	if m != nil {
		return m.NodeType
	}
	return ""
}

func (m *Node) GetPrevNode() string {
	if m != nil {
		return m.PrevNode
	}
	return ""
}

func (m *Node) GetNextNode() string {
	if m != nil {
		return m.NextNode
	}
	return ""
}

func (m *Node) GetAssignees() []string {
	if m != nil {
		return m.Assignees
	}
	return nil
}

func (m *Node) GetActType() string {
	if m != nil {
		return m.ActType
	}
	return ""
}

func (m *Node) GetNodeGroupId() string {
	if m != nil {
		return m.NodeGroupId
	}
	return ""
}

func (m *Node) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Node) GetCreatedBy() string {
	if m != nil {
		return m.CreatedBy
	}
	return ""
}

// 查找多条记录
type NodesRequest struct {
	WfId                 string   `protobuf:"bytes,1,opt,name=wf_id,json=wfId,proto3" json:"wf_id"`
	Database             string   `protobuf:"bytes,2,opt,name=database,proto3" json:"database"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodesRequest) Reset()         { *m = NodesRequest{} }
func (m *NodesRequest) String() string { return proto.CompactTextString(m) }
func (*NodesRequest) ProtoMessage()    {}
func (*NodesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c843d59d2d938e7, []int{1}
}

func (m *NodesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodesRequest.Unmarshal(m, b)
}
func (m *NodesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodesRequest.Marshal(b, m, deterministic)
}
func (m *NodesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodesRequest.Merge(m, src)
}
func (m *NodesRequest) XXX_Size() int {
	return xxx_messageInfo_NodesRequest.Size(m)
}
func (m *NodesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NodesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NodesRequest proto.InternalMessageInfo

func (m *NodesRequest) GetWfId() string {
	if m != nil {
		return m.WfId
	}
	return ""
}

func (m *NodesRequest) GetDatabase() string {
	if m != nil {
		return m.Database
	}
	return ""
}

type NodesResponse struct {
	Nodes                []*Node  `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodesResponse) Reset()         { *m = NodesResponse{} }
func (m *NodesResponse) String() string { return proto.CompactTextString(m) }
func (*NodesResponse) ProtoMessage()    {}
func (*NodesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c843d59d2d938e7, []int{2}
}

func (m *NodesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodesResponse.Unmarshal(m, b)
}
func (m *NodesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodesResponse.Marshal(b, m, deterministic)
}
func (m *NodesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodesResponse.Merge(m, src)
}
func (m *NodesResponse) XXX_Size() int {
	return xxx_messageInfo_NodesResponse.Size(m)
}
func (m *NodesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NodesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NodesResponse proto.InternalMessageInfo

func (m *NodesResponse) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

// 查询单条记录
type NodeRequest struct {
	NodeId               string   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id"`
	WfId                 string   `protobuf:"bytes,2,opt,name=wf_id,json=wfId,proto3" json:"wf_id"`
	Database             string   `protobuf:"bytes,3,opt,name=database,proto3" json:"database"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeRequest) Reset()         { *m = NodeRequest{} }
func (m *NodeRequest) String() string { return proto.CompactTextString(m) }
func (*NodeRequest) ProtoMessage()    {}
func (*NodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c843d59d2d938e7, []int{3}
}

func (m *NodeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeRequest.Unmarshal(m, b)
}
func (m *NodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeRequest.Marshal(b, m, deterministic)
}
func (m *NodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeRequest.Merge(m, src)
}
func (m *NodeRequest) XXX_Size() int {
	return xxx_messageInfo_NodeRequest.Size(m)
}
func (m *NodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NodeRequest proto.InternalMessageInfo

func (m *NodeRequest) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *NodeRequest) GetWfId() string {
	if m != nil {
		return m.WfId
	}
	return ""
}

func (m *NodeRequest) GetDatabase() string {
	if m != nil {
		return m.Database
	}
	return ""
}

type NodeResponse struct {
	Node                 *Node    `protobuf:"bytes,1,opt,name=node,proto3" json:"node"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeResponse) Reset()         { *m = NodeResponse{} }
func (m *NodeResponse) String() string { return proto.CompactTextString(m) }
func (*NodeResponse) ProtoMessage()    {}
func (*NodeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c843d59d2d938e7, []int{4}
}

func (m *NodeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeResponse.Unmarshal(m, b)
}
func (m *NodeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeResponse.Marshal(b, m, deterministic)
}
func (m *NodeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeResponse.Merge(m, src)
}
func (m *NodeResponse) XXX_Size() int {
	return xxx_messageInfo_NodeResponse.Size(m)
}
func (m *NodeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NodeResponse proto.InternalMessageInfo

func (m *NodeResponse) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

// 添加数据
type AddRequest struct {
	NodeId   string `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id"`
	NodeName string `protobuf:"bytes,4,opt,name=node_name,json=nodeName,proto3" json:"node_name"`
	WfId     string `protobuf:"bytes,2,opt,name=wf_id,json=wfId,proto3" json:"wf_id"`
	NodeType string `protobuf:"bytes,3,opt,name=node_type,json=nodeType,proto3" json:"node_type"`
	// repeated Condition condition =4; // 分支条件
	PrevNode             string   `protobuf:"bytes,5,opt,name=prev_node,json=prevNode,proto3" json:"prev_node"`
	NextNode             string   `protobuf:"bytes,6,opt,name=next_node,json=nextNode,proto3" json:"next_node"`
	Assignees            []string `protobuf:"bytes,7,rep,name=assignees,proto3" json:"assignees"`
	ActType              string   `protobuf:"bytes,8,opt,name=act_type,json=actType,proto3" json:"act_type"`
	NodeGroupId          string   `protobuf:"bytes,11,opt,name=node_group_id,json=nodeGroupId,proto3" json:"node_group_id"`
	Database             string   `protobuf:"bytes,9,opt,name=database,proto3" json:"database"`
	Writer               string   `protobuf:"bytes,10,opt,name=writer,proto3" json:"writer"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddRequest) Reset()         { *m = AddRequest{} }
func (m *AddRequest) String() string { return proto.CompactTextString(m) }
func (*AddRequest) ProtoMessage()    {}
func (*AddRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c843d59d2d938e7, []int{5}
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

func (m *AddRequest) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *AddRequest) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *AddRequest) GetWfId() string {
	if m != nil {
		return m.WfId
	}
	return ""
}

func (m *AddRequest) GetNodeType() string {
	if m != nil {
		return m.NodeType
	}
	return ""
}

func (m *AddRequest) GetPrevNode() string {
	if m != nil {
		return m.PrevNode
	}
	return ""
}

func (m *AddRequest) GetNextNode() string {
	if m != nil {
		return m.NextNode
	}
	return ""
}

func (m *AddRequest) GetAssignees() []string {
	if m != nil {
		return m.Assignees
	}
	return nil
}

func (m *AddRequest) GetActType() string {
	if m != nil {
		return m.ActType
	}
	return ""
}

func (m *AddRequest) GetNodeGroupId() string {
	if m != nil {
		return m.NodeGroupId
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
	NodeId               string   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddResponse) Reset()         { *m = AddResponse{} }
func (m *AddResponse) String() string { return proto.CompactTextString(m) }
func (*AddResponse) ProtoMessage()    {}
func (*AddResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c843d59d2d938e7, []int{6}
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

func (m *AddResponse) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

// 删除数据记录
type DeleteRequest struct {
	WfId                 string   `protobuf:"bytes,1,opt,name=wf_id,json=wfId,proto3" json:"wf_id"`
	Database             string   `protobuf:"bytes,2,opt,name=database,proto3" json:"database"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c843d59d2d938e7, []int{7}
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

func (m *DeleteRequest) GetWfId() string {
	if m != nil {
		return m.WfId
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
	return fileDescriptor_0c843d59d2d938e7, []int{8}
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
	proto.RegisterType((*Node)(nil), "node.Node")
	proto.RegisterType((*NodesRequest)(nil), "node.NodesRequest")
	proto.RegisterType((*NodesResponse)(nil), "node.NodesResponse")
	proto.RegisterType((*NodeRequest)(nil), "node.NodeRequest")
	proto.RegisterType((*NodeResponse)(nil), "node.NodeResponse")
	proto.RegisterType((*AddRequest)(nil), "node.AddRequest")
	proto.RegisterType((*AddResponse)(nil), "node.AddResponse")
	proto.RegisterType((*DeleteRequest)(nil), "node.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "node.DeleteResponse")
}

func init() { proto.RegisterFile("node.proto", fileDescriptor_0c843d59d2d938e7) }

var fileDescriptor_0c843d59d2d938e7 = []byte{
	// 475 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe4, 0x54, 0x4d, 0x6f, 0x13, 0x31,
	0x14, 0x64, 0xf3, 0xb9, 0x7e, 0x4b, 0x50, 0xeb, 0x22, 0x30, 0xe1, 0x43, 0x91, 0x0f, 0x28, 0xa7,
	0x08, 0x5a, 0x09, 0x89, 0x13, 0x04, 0x21, 0x50, 0x2f, 0x3d, 0x04, 0x24, 0x8e, 0x91, 0x13, 0xbf,
	0x56, 0x2b, 0xd1, 0xdd, 0x65, 0xed, 0x36, 0xcd, 0xaf, 0xe3, 0xc0, 0x8f, 0xe1, 0x6f, 0x20, 0x3f,
	0x3b, 0x59, 0x2f, 0xa2, 0x42, 0xe2, 0xca, 0x2d, 0x9e, 0xf1, 0xf8, 0xcd, 0xce, 0x1b, 0x05, 0xa0,
	0x28, 0x35, 0xce, 0xaa, 0xba, 0xb4, 0x25, 0xef, 0xb9, 0xdf, 0xf2, 0x47, 0x07, 0x7a, 0x67, 0xa5,
	0x46, 0xfe, 0x10, 0x86, 0x0e, 0x58, 0xe6, 0x5a, 0x24, 0x93, 0x64, 0xca, 0x16, 0x03, 0x77, 0x3c,
	0xd5, 0xfc, 0x31, 0x30, 0x22, 0x0a, 0x75, 0x89, 0xa2, 0x47, 0x54, 0xea, 0x80, 0x33, 0x75, 0x89,
	0xfc, 0x08, 0xfa, 0x9b, 0x73, 0xa7, 0xe9, 0x10, 0xd1, 0xdb, 0x9c, 0x47, 0x0a, 0xbb, 0xad, 0x50,
	0x74, 0x1b, 0xc5, 0xe7, 0x6d, 0x85, 0x8e, 0xac, 0x6a, 0xbc, 0x5e, 0x3a, 0x40, 0xf4, 0x3d, 0xe9,
	0x00, 0x32, 0xe1, 0x94, 0x78, 0x63, 0x3d, 0x39, 0x08, 0x4a, 0xbc, 0xb1, 0x44, 0x3e, 0x01, 0xa6,
	0x8c, 0xc9, 0x2f, 0x0a, 0x44, 0x23, 0x86, 0x93, 0xee, 0x94, 0x2d, 0x1a, 0x80, 0x3f, 0x82, 0x54,
	0xad, 0xad, 0x9f, 0x99, 0x92, 0x72, 0xa8, 0xd6, 0x96, 0x46, 0x4a, 0x18, 0x91, 0x9f, 0x8b, 0xba,
	0xbc, 0xaa, 0x9c, 0xd9, 0x8c, 0xf8, 0xcc, 0x81, 0x1f, 0x1d, 0x76, 0xaa, 0xf9, 0x53, 0x80, 0x75,
	0x8d, 0xca, 0xa2, 0x5e, 0x2a, 0x2b, 0x18, 0x5d, 0x60, 0x01, 0x99, 0xdb, 0x98, 0x5e, 0x6d, 0x05,
	0xb4, 0xe8, 0x77, 0x5b, 0xf9, 0x06, 0xee, 0x3a, 0x8b, 0x66, 0x81, 0xdf, 0xae, 0xd0, 0xd8, 0x26,
	0x96, 0x24, 0x8a, 0x65, 0x0c, 0xa9, 0x56, 0x56, 0xad, 0x94, 0xc1, 0x10, 0xd7, 0xfe, 0x2c, 0x5f,
	0xc2, 0x28, 0x3c, 0x60, 0xaa, 0xb2, 0x30, 0xc8, 0x27, 0xd0, 0x77, 0xf6, 0x8c, 0x48, 0x26, 0xdd,
	0x69, 0x76, 0x0c, 0x33, 0xda, 0x9c, 0xbb, 0xb3, 0xf0, 0x84, 0xfc, 0x02, 0x19, 0x1d, 0xc3, 0xc8,
	0x5b, 0xf7, 0xf7, 0xc7, 0x15, 0xc5, 0x5e, 0xba, 0xbf, 0x79, 0x99, 0xf9, 0x8f, 0xd9, 0x5b, 0x79,
	0x06, 0x54, 0x15, 0x7a, 0xb6, 0xed, 0xc4, 0x57, 0xe8, 0x7b, 0x07, 0x60, 0xae, 0xf5, 0x5f, 0x8d,
	0xfc, 0x1f, 0x45, 0x8a, 0x93, 0x65, 0xed, 0x64, 0xf9, 0x03, 0x18, 0x6c, 0xea, 0xdc, 0x62, 0x1d,
	0x1a, 0x14, 0x4e, 0xf2, 0x39, 0x64, 0x14, 0x60, 0x08, 0xfc, 0xb6, 0x04, 0xe5, 0x5b, 0x18, 0xbd,
	0xc7, 0xaf, 0x68, 0xf1, 0x9f, 0x7b, 0x76, 0x00, 0xf7, 0x76, 0x2f, 0xf8, 0x61, 0xc7, 0x3f, 0x13,
	0xdf, 0xa3, 0x4f, 0x58, 0x5f, 0xe7, 0x6b, 0xe4, 0xaf, 0x80, 0x7d, 0xc8, 0x0b, 0x4d, 0x6d, 0xe4,
	0xbc, 0x59, 0xf6, 0xae, 0xdb, 0xe3, 0xa3, 0x16, 0xe6, 0x5f, 0x91, 0x77, 0xf8, 0x09, 0xa4, 0x3b,
	0x1d, 0x3f, 0x8c, 0x3a, 0x12, 0x54, 0x3c, 0x86, 0xf6, 0xa2, 0x17, 0x30, 0x9c, 0x6b, 0xaf, 0x39,
	0xf0, 0x17, 0x9a, 0x22, 0x8d, 0x0f, 0x23, 0x64, 0xaf, 0x78, 0x0d, 0xe0, 0x3f, 0x80, 0x44, 0xc1,
	0x4b, 0x2b, 0x94, 0xf1, 0xfd, 0x36, 0xb8, 0x93, 0xae, 0x06, 0xf4, 0xbf, 0x77, 0xf2, 0x2b, 0x00,
	0x00, 0xff, 0xff, 0xda, 0x21, 0x15, 0xf8, 0x05, 0x05, 0x00, 0x00,
}
