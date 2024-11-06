// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: datastore.proto

package datastore

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

// Api Endpoints for DataStoreService service

func NewDataStoreServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for DataStoreService service

type DataStoreService interface {
	FindDatastores(ctx context.Context, in *DatastoresRequest, opts ...client.CallOption) (*DatastoresResponse, error)
	FindDatastore(ctx context.Context, in *DatastoreRequest, opts ...client.CallOption) (*DatastoreResponse, error)
	FindDatastoreByKey(ctx context.Context, in *DatastoreKeyRequest, opts ...client.CallOption) (*DatastoreResponse, error)
	FindDatastoreMapping(ctx context.Context, in *MappingRequest, opts ...client.CallOption) (*MappingResponse, error)
	AddDatastore(ctx context.Context, in *AddRequest, opts ...client.CallOption) (*AddResponse, error)
	AddDatastoreMapping(ctx context.Context, in *AddMappingRequest, opts ...client.CallOption) (*AddMappingResponse, error)
	ModifyDatastore(ctx context.Context, in *ModifyRequest, opts ...client.CallOption) (*ModifyResponse, error)
	ModifyDatastoreMapping(ctx context.Context, in *ModifyMappingRequest, opts ...client.CallOption) (*ModifyMappingResponse, error)
	DeleteDatastore(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
	DeleteDatastoreMapping(ctx context.Context, in *DeleteMappingRequest, opts ...client.CallOption) (*DeleteResponse, error)
	DeleteSelectDatastores(ctx context.Context, in *DeleteSelectRequest, opts ...client.CallOption) (*DeleteResponse, error)
	HardDeleteDatastores(ctx context.Context, in *HardDeleteDatastoresRequest, opts ...client.CallOption) (*DeleteResponse, error)
	AddUniqueKey(ctx context.Context, in *AddUniqueRequest, opts ...client.CallOption) (*AddUniqueResponse, error)
	DeleteUniqueKey(ctx context.Context, in *DeleteUniqueRequest, opts ...client.CallOption) (*DeleteUniqueResponse, error)
	AddRelation(ctx context.Context, in *AddRelationRequest, opts ...client.CallOption) (*AddRelationResponse, error)
	DeleteRelation(ctx context.Context, in *DeleteRelationRequest, opts ...client.CallOption) (*DeleteRelationResponse, error)
	ModifyDatastoreMenuSort(ctx context.Context, in *MenuSortRequest, opts ...client.CallOption) (*MenuSortResponse, error)
}

type dataStoreService struct {
	c    client.Client
	name string
}

func NewDataStoreService(name string, c client.Client) DataStoreService {
	return &dataStoreService{
		c:    c,
		name: name,
	}
}

func (c *dataStoreService) FindDatastores(ctx context.Context, in *DatastoresRequest, opts ...client.CallOption) (*DatastoresResponse, error) {
	req := c.c.NewRequest(c.name, "DataStoreService.FindDatastores", in)
	out := new(DatastoresResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataStoreService) FindDatastore(ctx context.Context, in *DatastoreRequest, opts ...client.CallOption) (*DatastoreResponse, error) {
	req := c.c.NewRequest(c.name, "DataStoreService.FindDatastore", in)
	out := new(DatastoreResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataStoreService) FindDatastoreByKey(ctx context.Context, in *DatastoreKeyRequest, opts ...client.CallOption) (*DatastoreResponse, error) {
	req := c.c.NewRequest(c.name, "DataStoreService.FindDatastoreByKey", in)
	out := new(DatastoreResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataStoreService) FindDatastoreMapping(ctx context.Context, in *MappingRequest, opts ...client.CallOption) (*MappingResponse, error) {
	req := c.c.NewRequest(c.name, "DataStoreService.FindDatastoreMapping", in)
	out := new(MappingResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataStoreService) AddDatastore(ctx context.Context, in *AddRequest, opts ...client.CallOption) (*AddResponse, error) {
	req := c.c.NewRequest(c.name, "DataStoreService.AddDatastore", in)
	out := new(AddResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataStoreService) AddDatastoreMapping(ctx context.Context, in *AddMappingRequest, opts ...client.CallOption) (*AddMappingResponse, error) {
	req := c.c.NewRequest(c.name, "DataStoreService.AddDatastoreMapping", in)
	out := new(AddMappingResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataStoreService) ModifyDatastore(ctx context.Context, in *ModifyRequest, opts ...client.CallOption) (*ModifyResponse, error) {
	req := c.c.NewRequest(c.name, "DataStoreService.ModifyDatastore", in)
	out := new(ModifyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataStoreService) ModifyDatastoreMapping(ctx context.Context, in *ModifyMappingRequest, opts ...client.CallOption) (*ModifyMappingResponse, error) {
	req := c.c.NewRequest(c.name, "DataStoreService.ModifyDatastoreMapping", in)
	out := new(ModifyMappingResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataStoreService) DeleteDatastore(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.name, "DataStoreService.DeleteDatastore", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataStoreService) DeleteDatastoreMapping(ctx context.Context, in *DeleteMappingRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.name, "DataStoreService.DeleteDatastoreMapping", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataStoreService) DeleteSelectDatastores(ctx context.Context, in *DeleteSelectRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.name, "DataStoreService.DeleteSelectDatastores", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataStoreService) HardDeleteDatastores(ctx context.Context, in *HardDeleteDatastoresRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.name, "DataStoreService.HardDeleteDatastores", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataStoreService) AddUniqueKey(ctx context.Context, in *AddUniqueRequest, opts ...client.CallOption) (*AddUniqueResponse, error) {
	req := c.c.NewRequest(c.name, "DataStoreService.AddUniqueKey", in)
	out := new(AddUniqueResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataStoreService) DeleteUniqueKey(ctx context.Context, in *DeleteUniqueRequest, opts ...client.CallOption) (*DeleteUniqueResponse, error) {
	req := c.c.NewRequest(c.name, "DataStoreService.DeleteUniqueKey", in)
	out := new(DeleteUniqueResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataStoreService) AddRelation(ctx context.Context, in *AddRelationRequest, opts ...client.CallOption) (*AddRelationResponse, error) {
	req := c.c.NewRequest(c.name, "DataStoreService.AddRelation", in)
	out := new(AddRelationResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataStoreService) DeleteRelation(ctx context.Context, in *DeleteRelationRequest, opts ...client.CallOption) (*DeleteRelationResponse, error) {
	req := c.c.NewRequest(c.name, "DataStoreService.DeleteRelation", in)
	out := new(DeleteRelationResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataStoreService) ModifyDatastoreMenuSort(ctx context.Context, in *MenuSortRequest, opts ...client.CallOption) (*MenuSortResponse, error) {
	req := c.c.NewRequest(c.name, "DataStoreService.ModifyDatastoreMenuSort", in)
	out := new(MenuSortResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DataStoreService service

type DataStoreServiceHandler interface {
	FindDatastores(context.Context, *DatastoresRequest, *DatastoresResponse) error
	FindDatastore(context.Context, *DatastoreRequest, *DatastoreResponse) error
	FindDatastoreByKey(context.Context, *DatastoreKeyRequest, *DatastoreResponse) error
	FindDatastoreMapping(context.Context, *MappingRequest, *MappingResponse) error
	AddDatastore(context.Context, *AddRequest, *AddResponse) error
	AddDatastoreMapping(context.Context, *AddMappingRequest, *AddMappingResponse) error
	ModifyDatastore(context.Context, *ModifyRequest, *ModifyResponse) error
	ModifyDatastoreMapping(context.Context, *ModifyMappingRequest, *ModifyMappingResponse) error
	DeleteDatastore(context.Context, *DeleteRequest, *DeleteResponse) error
	DeleteDatastoreMapping(context.Context, *DeleteMappingRequest, *DeleteResponse) error
	DeleteSelectDatastores(context.Context, *DeleteSelectRequest, *DeleteResponse) error
	HardDeleteDatastores(context.Context, *HardDeleteDatastoresRequest, *DeleteResponse) error
	AddUniqueKey(context.Context, *AddUniqueRequest, *AddUniqueResponse) error
	DeleteUniqueKey(context.Context, *DeleteUniqueRequest, *DeleteUniqueResponse) error
	AddRelation(context.Context, *AddRelationRequest, *AddRelationResponse) error
	DeleteRelation(context.Context, *DeleteRelationRequest, *DeleteRelationResponse) error
	ModifyDatastoreMenuSort(context.Context, *MenuSortRequest, *MenuSortResponse) error
}

func RegisterDataStoreServiceHandler(s server.Server, hdlr DataStoreServiceHandler, opts ...server.HandlerOption) error {
	type dataStoreService interface {
		FindDatastores(ctx context.Context, in *DatastoresRequest, out *DatastoresResponse) error
		FindDatastore(ctx context.Context, in *DatastoreRequest, out *DatastoreResponse) error
		FindDatastoreByKey(ctx context.Context, in *DatastoreKeyRequest, out *DatastoreResponse) error
		FindDatastoreMapping(ctx context.Context, in *MappingRequest, out *MappingResponse) error
		AddDatastore(ctx context.Context, in *AddRequest, out *AddResponse) error
		AddDatastoreMapping(ctx context.Context, in *AddMappingRequest, out *AddMappingResponse) error
		ModifyDatastore(ctx context.Context, in *ModifyRequest, out *ModifyResponse) error
		ModifyDatastoreMapping(ctx context.Context, in *ModifyMappingRequest, out *ModifyMappingResponse) error
		DeleteDatastore(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error
		DeleteDatastoreMapping(ctx context.Context, in *DeleteMappingRequest, out *DeleteResponse) error
		DeleteSelectDatastores(ctx context.Context, in *DeleteSelectRequest, out *DeleteResponse) error
		HardDeleteDatastores(ctx context.Context, in *HardDeleteDatastoresRequest, out *DeleteResponse) error
		AddUniqueKey(ctx context.Context, in *AddUniqueRequest, out *AddUniqueResponse) error
		DeleteUniqueKey(ctx context.Context, in *DeleteUniqueRequest, out *DeleteUniqueResponse) error
		AddRelation(ctx context.Context, in *AddRelationRequest, out *AddRelationResponse) error
		DeleteRelation(ctx context.Context, in *DeleteRelationRequest, out *DeleteRelationResponse) error
		ModifyDatastoreMenuSort(ctx context.Context, in *MenuSortRequest, out *MenuSortResponse) error
	}
	type DataStoreService struct {
		dataStoreService
	}
	h := &dataStoreServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&DataStoreService{h}, opts...))
}

type dataStoreServiceHandler struct {
	DataStoreServiceHandler
}

func (h *dataStoreServiceHandler) FindDatastores(ctx context.Context, in *DatastoresRequest, out *DatastoresResponse) error {
	return h.DataStoreServiceHandler.FindDatastores(ctx, in, out)
}

func (h *dataStoreServiceHandler) FindDatastore(ctx context.Context, in *DatastoreRequest, out *DatastoreResponse) error {
	return h.DataStoreServiceHandler.FindDatastore(ctx, in, out)
}

func (h *dataStoreServiceHandler) FindDatastoreByKey(ctx context.Context, in *DatastoreKeyRequest, out *DatastoreResponse) error {
	return h.DataStoreServiceHandler.FindDatastoreByKey(ctx, in, out)
}

func (h *dataStoreServiceHandler) FindDatastoreMapping(ctx context.Context, in *MappingRequest, out *MappingResponse) error {
	return h.DataStoreServiceHandler.FindDatastoreMapping(ctx, in, out)
}

func (h *dataStoreServiceHandler) AddDatastore(ctx context.Context, in *AddRequest, out *AddResponse) error {
	return h.DataStoreServiceHandler.AddDatastore(ctx, in, out)
}

func (h *dataStoreServiceHandler) AddDatastoreMapping(ctx context.Context, in *AddMappingRequest, out *AddMappingResponse) error {
	return h.DataStoreServiceHandler.AddDatastoreMapping(ctx, in, out)
}

func (h *dataStoreServiceHandler) ModifyDatastore(ctx context.Context, in *ModifyRequest, out *ModifyResponse) error {
	return h.DataStoreServiceHandler.ModifyDatastore(ctx, in, out)
}

func (h *dataStoreServiceHandler) ModifyDatastoreMapping(ctx context.Context, in *ModifyMappingRequest, out *ModifyMappingResponse) error {
	return h.DataStoreServiceHandler.ModifyDatastoreMapping(ctx, in, out)
}

func (h *dataStoreServiceHandler) DeleteDatastore(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error {
	return h.DataStoreServiceHandler.DeleteDatastore(ctx, in, out)
}

func (h *dataStoreServiceHandler) DeleteDatastoreMapping(ctx context.Context, in *DeleteMappingRequest, out *DeleteResponse) error {
	return h.DataStoreServiceHandler.DeleteDatastoreMapping(ctx, in, out)
}

func (h *dataStoreServiceHandler) DeleteSelectDatastores(ctx context.Context, in *DeleteSelectRequest, out *DeleteResponse) error {
	return h.DataStoreServiceHandler.DeleteSelectDatastores(ctx, in, out)
}

func (h *dataStoreServiceHandler) HardDeleteDatastores(ctx context.Context, in *HardDeleteDatastoresRequest, out *DeleteResponse) error {
	return h.DataStoreServiceHandler.HardDeleteDatastores(ctx, in, out)
}

func (h *dataStoreServiceHandler) AddUniqueKey(ctx context.Context, in *AddUniqueRequest, out *AddUniqueResponse) error {
	return h.DataStoreServiceHandler.AddUniqueKey(ctx, in, out)
}

func (h *dataStoreServiceHandler) DeleteUniqueKey(ctx context.Context, in *DeleteUniqueRequest, out *DeleteUniqueResponse) error {
	return h.DataStoreServiceHandler.DeleteUniqueKey(ctx, in, out)
}

func (h *dataStoreServiceHandler) AddRelation(ctx context.Context, in *AddRelationRequest, out *AddRelationResponse) error {
	return h.DataStoreServiceHandler.AddRelation(ctx, in, out)
}

func (h *dataStoreServiceHandler) DeleteRelation(ctx context.Context, in *DeleteRelationRequest, out *DeleteRelationResponse) error {
	return h.DataStoreServiceHandler.DeleteRelation(ctx, in, out)
}

func (h *dataStoreServiceHandler) ModifyDatastoreMenuSort(ctx context.Context, in *MenuSortRequest, out *MenuSortResponse) error {
	return h.DataStoreServiceHandler.ModifyDatastoreMenuSort(ctx, in, out)
}
