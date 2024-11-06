// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: message.proto

package message

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for MessageService service

func NewMessageServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for MessageService service

type MessageService interface {
	// 查找单个通知
	FindMessage(ctx context.Context, in *FindMessageRequest, opts ...client.CallOption) (*FindMessageResponse, error)
	// 查找多个通知
	FindMessages(ctx context.Context, in *FindMessagesRequest, opts ...client.CallOption) (*FindMessagesResponse, error)
	// 查找系统更新通知
	FindUpdateMessage(ctx context.Context, in *FindUpdateMessageRequest, opts ...client.CallOption) (*FindUpdateMessageResponse, error)
	// 添加通知
	AddMessage(ctx context.Context, in *AddMessageRequest, opts ...client.CallOption) (*AddMessageResponse, error)
	// 变更通知状态
	ChangeStatus(ctx context.Context, in *ChangeStatusRequest, opts ...client.CallOption) (*ChangeStatusResponse, error)
	// 硬删除单个通知
	DeleteMessage(ctx context.Context, in *DeleteMessageRequest, opts ...client.CallOption) (*DeleteMessageResponse, error)
	// 硬删除多个通知
	DeleteMessages(ctx context.Context, in *DeleteMessagesRequest, opts ...client.CallOption) (*DeleteMessagesResponse, error)
}

type messageService struct {
	c    client.Client
	name string
}

func NewMessageService(name string, c client.Client) MessageService {
	return &messageService{
		c:    c,
		name: name,
	}
}

func (c *messageService) FindMessage(ctx context.Context, in *FindMessageRequest, opts ...client.CallOption) (*FindMessageResponse, error) {
	req := c.c.NewRequest(c.name, "MessageService.FindMessage", in)
	out := new(FindMessageResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageService) FindMessages(ctx context.Context, in *FindMessagesRequest, opts ...client.CallOption) (*FindMessagesResponse, error) {
	req := c.c.NewRequest(c.name, "MessageService.FindMessages", in)
	out := new(FindMessagesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageService) FindUpdateMessage(ctx context.Context, in *FindUpdateMessageRequest, opts ...client.CallOption) (*FindUpdateMessageResponse, error) {
	req := c.c.NewRequest(c.name, "MessageService.FindUpdateMessage", in)
	out := new(FindUpdateMessageResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageService) AddMessage(ctx context.Context, in *AddMessageRequest, opts ...client.CallOption) (*AddMessageResponse, error) {
	req := c.c.NewRequest(c.name, "MessageService.AddMessage", in)
	out := new(AddMessageResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageService) ChangeStatus(ctx context.Context, in *ChangeStatusRequest, opts ...client.CallOption) (*ChangeStatusResponse, error) {
	req := c.c.NewRequest(c.name, "MessageService.ChangeStatus", in)
	out := new(ChangeStatusResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageService) DeleteMessage(ctx context.Context, in *DeleteMessageRequest, opts ...client.CallOption) (*DeleteMessageResponse, error) {
	req := c.c.NewRequest(c.name, "MessageService.DeleteMessage", in)
	out := new(DeleteMessageResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageService) DeleteMessages(ctx context.Context, in *DeleteMessagesRequest, opts ...client.CallOption) (*DeleteMessagesResponse, error) {
	req := c.c.NewRequest(c.name, "MessageService.DeleteMessages", in)
	out := new(DeleteMessagesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MessageService service

type MessageServiceHandler interface {
	// 查找单个通知
	FindMessage(context.Context, *FindMessageRequest, *FindMessageResponse) error
	// 查找多个通知
	FindMessages(context.Context, *FindMessagesRequest, *FindMessagesResponse) error
	// 查找系统更新通知
	FindUpdateMessage(context.Context, *FindUpdateMessageRequest, *FindUpdateMessageResponse) error
	// 添加通知
	AddMessage(context.Context, *AddMessageRequest, *AddMessageResponse) error
	// 变更通知状态
	ChangeStatus(context.Context, *ChangeStatusRequest, *ChangeStatusResponse) error
	// 硬删除单个通知
	DeleteMessage(context.Context, *DeleteMessageRequest, *DeleteMessageResponse) error
	// 硬删除多个通知
	DeleteMessages(context.Context, *DeleteMessagesRequest, *DeleteMessagesResponse) error
}

func RegisterMessageServiceHandler(s server.Server, hdlr MessageServiceHandler, opts ...server.HandlerOption) error {
	type messageService interface {
		FindMessage(ctx context.Context, in *FindMessageRequest, out *FindMessageResponse) error
		FindMessages(ctx context.Context, in *FindMessagesRequest, out *FindMessagesResponse) error
		FindUpdateMessage(ctx context.Context, in *FindUpdateMessageRequest, out *FindUpdateMessageResponse) error
		AddMessage(ctx context.Context, in *AddMessageRequest, out *AddMessageResponse) error
		ChangeStatus(ctx context.Context, in *ChangeStatusRequest, out *ChangeStatusResponse) error
		DeleteMessage(ctx context.Context, in *DeleteMessageRequest, out *DeleteMessageResponse) error
		DeleteMessages(ctx context.Context, in *DeleteMessagesRequest, out *DeleteMessagesResponse) error
	}
	type MessageService struct {
		messageService
	}
	h := &messageServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&MessageService{h}, opts...))
}

type messageServiceHandler struct {
	MessageServiceHandler
}

func (h *messageServiceHandler) FindMessage(ctx context.Context, in *FindMessageRequest, out *FindMessageResponse) error {
	return h.MessageServiceHandler.FindMessage(ctx, in, out)
}

func (h *messageServiceHandler) FindMessages(ctx context.Context, in *FindMessagesRequest, out *FindMessagesResponse) error {
	return h.MessageServiceHandler.FindMessages(ctx, in, out)
}

func (h *messageServiceHandler) FindUpdateMessage(ctx context.Context, in *FindUpdateMessageRequest, out *FindUpdateMessageResponse) error {
	return h.MessageServiceHandler.FindUpdateMessage(ctx, in, out)
}

func (h *messageServiceHandler) AddMessage(ctx context.Context, in *AddMessageRequest, out *AddMessageResponse) error {
	return h.MessageServiceHandler.AddMessage(ctx, in, out)
}

func (h *messageServiceHandler) ChangeStatus(ctx context.Context, in *ChangeStatusRequest, out *ChangeStatusResponse) error {
	return h.MessageServiceHandler.ChangeStatus(ctx, in, out)
}

func (h *messageServiceHandler) DeleteMessage(ctx context.Context, in *DeleteMessageRequest, out *DeleteMessageResponse) error {
	return h.MessageServiceHandler.DeleteMessage(ctx, in, out)
}

func (h *messageServiceHandler) DeleteMessages(ctx context.Context, in *DeleteMessagesRequest, out *DeleteMessagesResponse) error {
	return h.MessageServiceHandler.DeleteMessages(ctx, in, out)
}
