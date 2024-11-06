// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: generate.proto

package generate

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

// Api Endpoints for GenerateService service

func NewGenerateServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for GenerateService service

type GenerateService interface {
	FindGenerateConfig(ctx context.Context, in *FindRequest, opts ...client.CallOption) (*FindResponse, error)
	AddGenerateConfig(ctx context.Context, in *AddRequest, opts ...client.CallOption) (*AddResponse, error)
	ModifyGenerateConfig(ctx context.Context, in *ModifyRequest, opts ...client.CallOption) (*ModifyResponse, error)
	UploadData(ctx context.Context, in *UploadRequest, opts ...client.CallOption) (*UploadResponse, error)
	FindRowData(ctx context.Context, in *RowRequest, opts ...client.CallOption) (*RowResponse, error)
	FindColumnData(ctx context.Context, in *ColumnRequest, opts ...client.CallOption) (*ColumnResponse, error)
	DeleteGenerateConfig(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
}

type generateService struct {
	c    client.Client
	name string
}

func NewGenerateService(name string, c client.Client) GenerateService {
	return &generateService{
		c:    c,
		name: name,
	}
}

func (c *generateService) FindGenerateConfig(ctx context.Context, in *FindRequest, opts ...client.CallOption) (*FindResponse, error) {
	req := c.c.NewRequest(c.name, "GenerateService.FindGenerateConfig", in)
	out := new(FindResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *generateService) AddGenerateConfig(ctx context.Context, in *AddRequest, opts ...client.CallOption) (*AddResponse, error) {
	req := c.c.NewRequest(c.name, "GenerateService.AddGenerateConfig", in)
	out := new(AddResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *generateService) ModifyGenerateConfig(ctx context.Context, in *ModifyRequest, opts ...client.CallOption) (*ModifyResponse, error) {
	req := c.c.NewRequest(c.name, "GenerateService.ModifyGenerateConfig", in)
	out := new(ModifyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *generateService) UploadData(ctx context.Context, in *UploadRequest, opts ...client.CallOption) (*UploadResponse, error) {
	req := c.c.NewRequest(c.name, "GenerateService.UploadData", in)
	out := new(UploadResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *generateService) FindRowData(ctx context.Context, in *RowRequest, opts ...client.CallOption) (*RowResponse, error) {
	req := c.c.NewRequest(c.name, "GenerateService.FindRowData", in)
	out := new(RowResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *generateService) FindColumnData(ctx context.Context, in *ColumnRequest, opts ...client.CallOption) (*ColumnResponse, error) {
	req := c.c.NewRequest(c.name, "GenerateService.FindColumnData", in)
	out := new(ColumnResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *generateService) DeleteGenerateConfig(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.name, "GenerateService.DeleteGenerateConfig", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GenerateService service

type GenerateServiceHandler interface {
	FindGenerateConfig(context.Context, *FindRequest, *FindResponse) error
	AddGenerateConfig(context.Context, *AddRequest, *AddResponse) error
	ModifyGenerateConfig(context.Context, *ModifyRequest, *ModifyResponse) error
	UploadData(context.Context, *UploadRequest, *UploadResponse) error
	FindRowData(context.Context, *RowRequest, *RowResponse) error
	FindColumnData(context.Context, *ColumnRequest, *ColumnResponse) error
	DeleteGenerateConfig(context.Context, *DeleteRequest, *DeleteResponse) error
}

func RegisterGenerateServiceHandler(s server.Server, hdlr GenerateServiceHandler, opts ...server.HandlerOption) error {
	type generateService interface {
		FindGenerateConfig(ctx context.Context, in *FindRequest, out *FindResponse) error
		AddGenerateConfig(ctx context.Context, in *AddRequest, out *AddResponse) error
		ModifyGenerateConfig(ctx context.Context, in *ModifyRequest, out *ModifyResponse) error
		UploadData(ctx context.Context, in *UploadRequest, out *UploadResponse) error
		FindRowData(ctx context.Context, in *RowRequest, out *RowResponse) error
		FindColumnData(ctx context.Context, in *ColumnRequest, out *ColumnResponse) error
		DeleteGenerateConfig(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error
	}
	type GenerateService struct {
		generateService
	}
	h := &generateServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&GenerateService{h}, opts...))
}

type generateServiceHandler struct {
	GenerateServiceHandler
}

func (h *generateServiceHandler) FindGenerateConfig(ctx context.Context, in *FindRequest, out *FindResponse) error {
	return h.GenerateServiceHandler.FindGenerateConfig(ctx, in, out)
}

func (h *generateServiceHandler) AddGenerateConfig(ctx context.Context, in *AddRequest, out *AddResponse) error {
	return h.GenerateServiceHandler.AddGenerateConfig(ctx, in, out)
}

func (h *generateServiceHandler) ModifyGenerateConfig(ctx context.Context, in *ModifyRequest, out *ModifyResponse) error {
	return h.GenerateServiceHandler.ModifyGenerateConfig(ctx, in, out)
}

func (h *generateServiceHandler) UploadData(ctx context.Context, in *UploadRequest, out *UploadResponse) error {
	return h.GenerateServiceHandler.UploadData(ctx, in, out)
}

func (h *generateServiceHandler) FindRowData(ctx context.Context, in *RowRequest, out *RowResponse) error {
	return h.GenerateServiceHandler.FindRowData(ctx, in, out)
}

func (h *generateServiceHandler) FindColumnData(ctx context.Context, in *ColumnRequest, out *ColumnResponse) error {
	return h.GenerateServiceHandler.FindColumnData(ctx, in, out)
}

func (h *generateServiceHandler) DeleteGenerateConfig(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error {
	return h.GenerateServiceHandler.DeleteGenerateConfig(ctx, in, out)
}
