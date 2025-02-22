// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: item.proto

package item

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

// Api Endpoints for ItemService service

func NewItemServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for ItemService service

type ItemService interface {
	FindItems(ctx context.Context, in *ItemsRequest, opts ...client.CallOption) (*ItemsResponse, error)
	FindCount(ctx context.Context, in *CountRequest, opts ...client.CallOption) (*CountResponse, error)
	FindKaraCount(ctx context.Context, in *KaraCountRequest, opts ...client.CallOption) (*KaraCountResponse, error)
	FindUnApproveItems(ctx context.Context, in *UnApproveItemsRequest, opts ...client.CallOption) (*UnApproveItemsResponse, error)
	FindItem(ctx context.Context, in *ItemRequest, opts ...client.CallOption) (*ItemResponse, error)
	FindRishiritsu(ctx context.Context, in *RishiritsuRequest, opts ...client.CallOption) (*RishiritsuResponse, error)
	AddItem(ctx context.Context, in *AddRequest, opts ...client.CallOption) (*AddResponse, error)
	ModifyItem(ctx context.Context, in *ModifyRequest, opts ...client.CallOption) (*ModifyResponse, error)
	ConfimItem(ctx context.Context, in *JournalRequest, opts ...client.CallOption) (*JournalResponse, error)
	GenerateItem(ctx context.Context, in *JournalRequest, opts ...client.CallOption) (*JournalResponse, error)
	GenerateShoukyakuItem(ctx context.Context, in *JournalRequest, opts ...client.CallOption) (*JournalResponse, error)
	InventoryItem(ctx context.Context, in *InventoryItemRequest, opts ...client.CallOption) (*InventoryItemResponse, error)
	MutilInventoryItem(ctx context.Context, in *MutilInventoryItemRequest, opts ...client.CallOption) (*MutilInventoryItemResponse, error)
	ResetInventoryItems(ctx context.Context, in *ResetInventoryItemsRequest, opts ...client.CallOption) (*ResetInventoryItemsResponse, error)
	DeleteItem(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
	DeleteDatastoreItems(ctx context.Context, in *DeleteDatastoreItemsRequest, opts ...client.CallOption) (*DeleteResponse, error)
	DeleteItems(ctx context.Context, in *DeleteItemsRequest, opts ...client.CallOption) (*DeleteResponse, error)
	DeleteSelectItems(ctx context.Context, in *SelectedItemsRequest, opts ...client.CallOption) (ItemService_DeleteSelectItemsService, error)
	ChangeOwners(ctx context.Context, in *OwnersRequest, opts ...client.CallOption) (*OwnersResponse, error)
	ChangeSelectOwners(ctx context.Context, in *SelectOwnersRequest, opts ...client.CallOption) (*SelectOwnersResponse, error)
	ChangeItemOwner(ctx context.Context, in *ItemOwnerRequest, opts ...client.CallOption) (*ItemOwnerResponse, error)
	ChangeStatus(ctx context.Context, in *StatusRequest, opts ...client.CallOption) (*StatusResponse, error)
	ChangeLabelTime(ctx context.Context, in *LabelTimeRequest, opts ...client.CallOption) (*LabelTimeResponse, error)
	ChangeDebt(ctx context.Context, in *ChangeDebtRequest, opts ...client.CallOption) (*ChangeDebtResponse, error)
	ContractExpire(ctx context.Context, in *ContractExpireRequest, opts ...client.CallOption) (*ContractExpireResponse, error)
	ModifyContract(ctx context.Context, in *ModifyContractRequest, opts ...client.CallOption) (*ModifyContractResponse, error)
	TerminateContract(ctx context.Context, in *TerminateContractRequest, opts ...client.CallOption) (*TerminateContractResponse, error)
	// double stream
	ImportItem(ctx context.Context, opts ...client.CallOption) (ItemService_ImportItemService, error)
	ImportCheckItem(ctx context.Context, opts ...client.CallOption) (ItemService_ImportCheckItemService, error)
	MappingUpload(ctx context.Context, opts ...client.CallOption) (ItemService_MappingUploadService, error)
	// single stream
	Download(ctx context.Context, in *DownloadRequest, opts ...client.CallOption) (ItemService_DownloadService, error)
	FindAndModifyFile(ctx context.Context, in *FindRequest, opts ...client.CallOption) (ItemService_FindAndModifyFileService, error)
	SwkDownload(ctx context.Context, in *DownloadRequest, opts ...client.CallOption) (ItemService_DownloadService, error)
}

type itemService struct {
	c    client.Client
	name string
}

func NewItemService(name string, c client.Client) ItemService {
	return &itemService{
		c:    c,
		name: name,
	}
}

func (c *itemService) FindItems(ctx context.Context, in *ItemsRequest, opts ...client.CallOption) (*ItemsResponse, error) {
	req := c.c.NewRequest(c.name, "ItemService.FindItems", in)
	out := new(ItemsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemService) FindCount(ctx context.Context, in *CountRequest, opts ...client.CallOption) (*CountResponse, error) {
	req := c.c.NewRequest(c.name, "ItemService.FindCount", in)
	out := new(CountResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemService) FindKaraCount(ctx context.Context, in *KaraCountRequest, opts ...client.CallOption) (*KaraCountResponse, error) {
	req := c.c.NewRequest(c.name, "ItemService.FindKaraCount", in)
	out := new(KaraCountResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemService) FindUnApproveItems(ctx context.Context, in *UnApproveItemsRequest, opts ...client.CallOption) (*UnApproveItemsResponse, error) {
	req := c.c.NewRequest(c.name, "ItemService.FindUnApproveItems", in)
	out := new(UnApproveItemsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemService) FindItem(ctx context.Context, in *ItemRequest, opts ...client.CallOption) (*ItemResponse, error) {
	req := c.c.NewRequest(c.name, "ItemService.FindItem", in)
	out := new(ItemResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemService) FindRishiritsu(ctx context.Context, in *RishiritsuRequest, opts ...client.CallOption) (*RishiritsuResponse, error) {
	req := c.c.NewRequest(c.name, "ItemService.FindRishiritsu", in)
	out := new(RishiritsuResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemService) AddItem(ctx context.Context, in *AddRequest, opts ...client.CallOption) (*AddResponse, error) {
	req := c.c.NewRequest(c.name, "ItemService.AddItem", in)
	out := new(AddResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemService) ModifyItem(ctx context.Context, in *ModifyRequest, opts ...client.CallOption) (*ModifyResponse, error) {
	req := c.c.NewRequest(c.name, "ItemService.ModifyItem", in)
	out := new(ModifyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemService) ConfimItem(ctx context.Context, in *JournalRequest, opts ...client.CallOption) (*JournalResponse, error) {
	req := c.c.NewRequest(c.name, "ItemService.ConfimItem", in)
	out := new(JournalResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemService) GenerateItem(ctx context.Context, in *JournalRequest, opts ...client.CallOption) (*JournalResponse, error) {
	req := c.c.NewRequest(c.name, "ItemService.GenerateItem", in)
	out := new(JournalResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemService) GenerateShoukyakuItem(ctx context.Context, in *JournalRequest, opts ...client.CallOption) (*JournalResponse, error) {
	req := c.c.NewRequest(c.name, "ItemService.GenerateShoukyakuItem", in)
	out := new(JournalResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemService) InventoryItem(ctx context.Context, in *InventoryItemRequest, opts ...client.CallOption) (*InventoryItemResponse, error) {
	req := c.c.NewRequest(c.name, "ItemService.InventoryItem", in)
	out := new(InventoryItemResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemService) MutilInventoryItem(ctx context.Context, in *MutilInventoryItemRequest, opts ...client.CallOption) (*MutilInventoryItemResponse, error) {
	req := c.c.NewRequest(c.name, "ItemService.MutilInventoryItem", in)
	out := new(MutilInventoryItemResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemService) ResetInventoryItems(ctx context.Context, in *ResetInventoryItemsRequest, opts ...client.CallOption) (*ResetInventoryItemsResponse, error) {
	req := c.c.NewRequest(c.name, "ItemService.ResetInventoryItems", in)
	out := new(ResetInventoryItemsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemService) DeleteItem(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.name, "ItemService.DeleteItem", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemService) DeleteDatastoreItems(ctx context.Context, in *DeleteDatastoreItemsRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.name, "ItemService.DeleteDatastoreItems", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemService) DeleteItems(ctx context.Context, in *DeleteItemsRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.name, "ItemService.DeleteItems", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemService) DeleteSelectItems(ctx context.Context, in *SelectedItemsRequest, opts ...client.CallOption) (ItemService_DeleteSelectItemsService, error) {
	req := c.c.NewRequest(c.name, "ItemService.DeleteSelectItems", &SelectedItemsRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &itemServiceDeleteSelectItems{stream}, nil
}

type ItemService_DeleteSelectItemsService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*SelectedItemsResponse, error)
}

type itemServiceDeleteSelectItems struct {
	stream client.Stream
}

func (x *itemServiceDeleteSelectItems) Close() error {
	return x.stream.Close()
}

func (x *itemServiceDeleteSelectItems) Context() context.Context {
	return x.stream.Context()
}

func (x *itemServiceDeleteSelectItems) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *itemServiceDeleteSelectItems) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *itemServiceDeleteSelectItems) Recv() (*SelectedItemsResponse, error) {
	m := new(SelectedItemsResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *itemService) ChangeOwners(ctx context.Context, in *OwnersRequest, opts ...client.CallOption) (*OwnersResponse, error) {
	req := c.c.NewRequest(c.name, "ItemService.ChangeOwners", in)
	out := new(OwnersResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemService) ChangeSelectOwners(ctx context.Context, in *SelectOwnersRequest, opts ...client.CallOption) (*SelectOwnersResponse, error) {
	req := c.c.NewRequest(c.name, "ItemService.ChangeSelectOwners", in)
	out := new(SelectOwnersResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemService) ChangeItemOwner(ctx context.Context, in *ItemOwnerRequest, opts ...client.CallOption) (*ItemOwnerResponse, error) {
	req := c.c.NewRequest(c.name, "ItemService.ChangeItemOwner", in)
	out := new(ItemOwnerResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemService) ChangeStatus(ctx context.Context, in *StatusRequest, opts ...client.CallOption) (*StatusResponse, error) {
	req := c.c.NewRequest(c.name, "ItemService.ChangeStatus", in)
	out := new(StatusResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemService) ChangeLabelTime(ctx context.Context, in *LabelTimeRequest, opts ...client.CallOption) (*LabelTimeResponse, error) {
	req := c.c.NewRequest(c.name, "ItemService.ChangeLabelTime", in)
	out := new(LabelTimeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemService) ChangeDebt(ctx context.Context, in *ChangeDebtRequest, opts ...client.CallOption) (*ChangeDebtResponse, error) {
	req := c.c.NewRequest(c.name, "ItemService.ChangeDebt", in)
	out := new(ChangeDebtResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemService) ContractExpire(ctx context.Context, in *ContractExpireRequest, opts ...client.CallOption) (*ContractExpireResponse, error) {
	req := c.c.NewRequest(c.name, "ItemService.ContractExpire", in)
	out := new(ContractExpireResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemService) ModifyContract(ctx context.Context, in *ModifyContractRequest, opts ...client.CallOption) (*ModifyContractResponse, error) {
	req := c.c.NewRequest(c.name, "ItemService.ModifyContract", in)
	out := new(ModifyContractResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemService) TerminateContract(ctx context.Context, in *TerminateContractRequest, opts ...client.CallOption) (*TerminateContractResponse, error) {
	req := c.c.NewRequest(c.name, "ItemService.TerminateContract", in)
	out := new(TerminateContractResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemService) ImportItem(ctx context.Context, opts ...client.CallOption) (ItemService_ImportItemService, error) {
	req := c.c.NewRequest(c.name, "ItemService.ImportItem", &ImportRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &itemServiceImportItem{stream}, nil
}

type ItemService_ImportItemService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*ImportRequest) error
	Recv() (*ImportResponse, error)
}

type itemServiceImportItem struct {
	stream client.Stream
}

func (x *itemServiceImportItem) Close() error {
	return x.stream.Close()
}

func (x *itemServiceImportItem) Context() context.Context {
	return x.stream.Context()
}

func (x *itemServiceImportItem) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *itemServiceImportItem) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *itemServiceImportItem) Send(m *ImportRequest) error {
	return x.stream.Send(m)
}

func (x *itemServiceImportItem) Recv() (*ImportResponse, error) {
	m := new(ImportResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *itemService) ImportCheckItem(ctx context.Context, opts ...client.CallOption) (ItemService_ImportCheckItemService, error) {
	req := c.c.NewRequest(c.name, "ItemService.ImportCheckItem", &ImportCheckRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &itemServiceImportCheckItem{stream}, nil
}

type ItemService_ImportCheckItemService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*ImportCheckRequest) error
	Recv() (*ImportCheckResponse, error)
}

type itemServiceImportCheckItem struct {
	stream client.Stream
}

func (x *itemServiceImportCheckItem) Close() error {
	return x.stream.Close()
}

func (x *itemServiceImportCheckItem) Context() context.Context {
	return x.stream.Context()
}

func (x *itemServiceImportCheckItem) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *itemServiceImportCheckItem) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *itemServiceImportCheckItem) Send(m *ImportCheckRequest) error {
	return x.stream.Send(m)
}

func (x *itemServiceImportCheckItem) Recv() (*ImportCheckResponse, error) {
	m := new(ImportCheckResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *itemService) MappingUpload(ctx context.Context, opts ...client.CallOption) (ItemService_MappingUploadService, error) {
	req := c.c.NewRequest(c.name, "ItemService.MappingUpload", &MappingUploadRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &itemServiceMappingUpload{stream}, nil
}

type ItemService_MappingUploadService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*MappingUploadRequest) error
	Recv() (*MappingUploadResponse, error)
}

type itemServiceMappingUpload struct {
	stream client.Stream
}

func (x *itemServiceMappingUpload) Close() error {
	return x.stream.Close()
}

func (x *itemServiceMappingUpload) Context() context.Context {
	return x.stream.Context()
}

func (x *itemServiceMappingUpload) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *itemServiceMappingUpload) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *itemServiceMappingUpload) Send(m *MappingUploadRequest) error {
	return x.stream.Send(m)
}

func (x *itemServiceMappingUpload) Recv() (*MappingUploadResponse, error) {
	m := new(MappingUploadResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *itemService) Download(ctx context.Context, in *DownloadRequest, opts ...client.CallOption) (ItemService_DownloadService, error) {
	req := c.c.NewRequest(c.name, "ItemService.Download", &DownloadRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &itemServiceDownload{stream}, nil
}

func (c *itemService) SwkDownload(ctx context.Context, in *DownloadRequest, opts ...client.CallOption) (ItemService_DownloadService, error) {
	req := c.c.NewRequest(c.name, "ItemService.SwkDownload", &DownloadRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &itemServiceDownload{stream}, nil
}


type ItemService_DownloadService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*DownloadResponse, error)
}

type itemServiceDownload struct {
	stream client.Stream
}

func (x *itemServiceDownload) Close() error {
	return x.stream.Close()
}

func (x *itemServiceDownload) Context() context.Context {
	return x.stream.Context()
}

func (x *itemServiceDownload) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *itemServiceDownload) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *itemServiceDownload) Recv() (*DownloadResponse, error) {
	m := new(DownloadResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *itemService) FindAndModifyFile(ctx context.Context, in *FindRequest, opts ...client.CallOption) (ItemService_FindAndModifyFileService, error) {
	req := c.c.NewRequest(c.name, "ItemService.FindAndModifyFile", &FindRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &itemServiceFindAndModifyFile{stream}, nil
}

type ItemService_FindAndModifyFileService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*FindResponse, error)
}

type itemServiceFindAndModifyFile struct {
	stream client.Stream
}

func (x *itemServiceFindAndModifyFile) Close() error {
	return x.stream.Close()
}

func (x *itemServiceFindAndModifyFile) Context() context.Context {
	return x.stream.Context()
}

func (x *itemServiceFindAndModifyFile) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *itemServiceFindAndModifyFile) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *itemServiceFindAndModifyFile) Recv() (*FindResponse, error) {
	m := new(FindResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for ItemService service

type ItemServiceHandler interface {
	FindItems(context.Context, *ItemsRequest, *ItemsResponse) error
	FindCount(context.Context, *CountRequest, *CountResponse) error
	FindKaraCount(context.Context, *KaraCountRequest, *KaraCountResponse) error
	FindUnApproveItems(context.Context, *UnApproveItemsRequest, *UnApproveItemsResponse) error
	FindItem(context.Context, *ItemRequest, *ItemResponse) error
	FindRishiritsu(context.Context, *RishiritsuRequest, *RishiritsuResponse) error
	AddItem(context.Context, *AddRequest, *AddResponse) error
	ModifyItem(context.Context, *ModifyRequest, *ModifyResponse) error
	ConfimItem(context.Context, *JournalRequest, *JournalResponse) error
	GenerateItem(context.Context, *JournalRequest, *JournalResponse) error
	GenerateShoukyakuItem(context.Context, *JournalRequest, *JournalResponse) error
	InventoryItem(context.Context, *InventoryItemRequest, *InventoryItemResponse) error
	MutilInventoryItem(context.Context, *MutilInventoryItemRequest, *MutilInventoryItemResponse) error
	ResetInventoryItems(context.Context, *ResetInventoryItemsRequest, *ResetInventoryItemsResponse) error
	DeleteItem(context.Context, *DeleteRequest, *DeleteResponse) error
	DeleteDatastoreItems(context.Context, *DeleteDatastoreItemsRequest, *DeleteResponse) error
	DeleteItems(context.Context, *DeleteItemsRequest, *DeleteResponse) error
	DeleteSelectItems(context.Context, *SelectedItemsRequest, ItemService_DeleteSelectItemsStream) error
	ChangeOwners(context.Context, *OwnersRequest, *OwnersResponse) error
	ChangeSelectOwners(context.Context, *SelectOwnersRequest, *SelectOwnersResponse) error
	ChangeItemOwner(context.Context, *ItemOwnerRequest, *ItemOwnerResponse) error
	ChangeStatus(context.Context, *StatusRequest, *StatusResponse) error
	ChangeLabelTime(context.Context, *LabelTimeRequest, *LabelTimeResponse) error
	ChangeDebt(context.Context, *ChangeDebtRequest, *ChangeDebtResponse) error
	ContractExpire(context.Context, *ContractExpireRequest, *ContractExpireResponse) error
	ModifyContract(context.Context, *ModifyContractRequest, *ModifyContractResponse) error
	TerminateContract(context.Context, *TerminateContractRequest, *TerminateContractResponse) error
	// double stream
	ImportItem(context.Context, ItemService_ImportItemStream) error
	ImportCheckItem(context.Context, ItemService_ImportCheckItemStream) error
	MappingUpload(context.Context, ItemService_MappingUploadStream) error
	// single stream
	Download(context.Context, *DownloadRequest, ItemService_DownloadStream) error
	FindAndModifyFile(context.Context, *FindRequest, ItemService_FindAndModifyFileStream) error
	SwkDownload(context.Context, *DownloadRequest, ItemService_DownloadStream) error
}

func RegisterItemServiceHandler(s server.Server, hdlr ItemServiceHandler, opts ...server.HandlerOption) error {
	type itemService interface {
		FindItems(ctx context.Context, in *ItemsRequest, out *ItemsResponse) error
		FindCount(ctx context.Context, in *CountRequest, out *CountResponse) error
		FindKaraCount(ctx context.Context, in *KaraCountRequest, out *KaraCountResponse) error
		FindUnApproveItems(ctx context.Context, in *UnApproveItemsRequest, out *UnApproveItemsResponse) error
		FindItem(ctx context.Context, in *ItemRequest, out *ItemResponse) error
		FindRishiritsu(ctx context.Context, in *RishiritsuRequest, out *RishiritsuResponse) error
		AddItem(ctx context.Context, in *AddRequest, out *AddResponse) error
		ModifyItem(ctx context.Context, in *ModifyRequest, out *ModifyResponse) error
		ConfimItem(ctx context.Context, in *JournalRequest, out *JournalResponse) error
		GenerateItem(ctx context.Context, in *JournalRequest, out *JournalResponse) error
		GenerateShoukyakuItem(ctx context.Context, in *JournalRequest, out *JournalResponse) error
		InventoryItem(ctx context.Context, in *InventoryItemRequest, out *InventoryItemResponse) error
		MutilInventoryItem(ctx context.Context, in *MutilInventoryItemRequest, out *MutilInventoryItemResponse) error
		ResetInventoryItems(ctx context.Context, in *ResetInventoryItemsRequest, out *ResetInventoryItemsResponse) error
		DeleteItem(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error
		DeleteDatastoreItems(ctx context.Context, in *DeleteDatastoreItemsRequest, out *DeleteResponse) error
		DeleteItems(ctx context.Context, in *DeleteItemsRequest, out *DeleteResponse) error
		DeleteSelectItems(ctx context.Context, stream server.Stream) error
		ChangeOwners(ctx context.Context, in *OwnersRequest, out *OwnersResponse) error
		ChangeSelectOwners(ctx context.Context, in *SelectOwnersRequest, out *SelectOwnersResponse) error
		ChangeItemOwner(ctx context.Context, in *ItemOwnerRequest, out *ItemOwnerResponse) error
		ChangeStatus(ctx context.Context, in *StatusRequest, out *StatusResponse) error
		ChangeLabelTime(ctx context.Context, in *LabelTimeRequest, out *LabelTimeResponse) error
		ChangeDebt(ctx context.Context, in *ChangeDebtRequest, out *ChangeDebtResponse) error
		ContractExpire(ctx context.Context, in *ContractExpireRequest, out *ContractExpireResponse) error
		ModifyContract(ctx context.Context, in *ModifyContractRequest, out *ModifyContractResponse) error
		TerminateContract(ctx context.Context, in *TerminateContractRequest, out *TerminateContractResponse) error
		ImportItem(ctx context.Context, stream server.Stream) error
		ImportCheckItem(ctx context.Context, stream server.Stream) error
		MappingUpload(ctx context.Context, stream server.Stream) error
		Download(ctx context.Context, stream server.Stream) error
		SwkDownload(ctx context.Context, stream server.Stream) error
		FindAndModifyFile(ctx context.Context, stream server.Stream) error
	}
	type ItemService struct {
		itemService
	}
	h := &itemServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ItemService{h}, opts...))
}

type itemServiceHandler struct {
	ItemServiceHandler
}

func (h *itemServiceHandler) FindItems(ctx context.Context, in *ItemsRequest, out *ItemsResponse) error {
	return h.ItemServiceHandler.FindItems(ctx, in, out)
}

func (h *itemServiceHandler) FindCount(ctx context.Context, in *CountRequest, out *CountResponse) error {
	return h.ItemServiceHandler.FindCount(ctx, in, out)
}

func (h *itemServiceHandler) FindKaraCount(ctx context.Context, in *KaraCountRequest, out *KaraCountResponse) error {
	return h.ItemServiceHandler.FindKaraCount(ctx, in, out)
}

func (h *itemServiceHandler) FindUnApproveItems(ctx context.Context, in *UnApproveItemsRequest, out *UnApproveItemsResponse) error {
	return h.ItemServiceHandler.FindUnApproveItems(ctx, in, out)
}

func (h *itemServiceHandler) FindItem(ctx context.Context, in *ItemRequest, out *ItemResponse) error {
	return h.ItemServiceHandler.FindItem(ctx, in, out)
}

func (h *itemServiceHandler) FindRishiritsu(ctx context.Context, in *RishiritsuRequest, out *RishiritsuResponse) error {
	return h.ItemServiceHandler.FindRishiritsu(ctx, in, out)
}

func (h *itemServiceHandler) AddItem(ctx context.Context, in *AddRequest, out *AddResponse) error {
	return h.ItemServiceHandler.AddItem(ctx, in, out)
}

func (h *itemServiceHandler) ModifyItem(ctx context.Context, in *ModifyRequest, out *ModifyResponse) error {
	return h.ItemServiceHandler.ModifyItem(ctx, in, out)
}

func (h *itemServiceHandler) ConfimItem(ctx context.Context, in *JournalRequest, out *JournalResponse) error {
	return h.ItemServiceHandler.ConfimItem(ctx, in, out)
}

func (h *itemServiceHandler) GenerateItem(ctx context.Context, in *JournalRequest, out *JournalResponse) error {
	return h.ItemServiceHandler.GenerateItem(ctx, in, out)
}

func (h *itemServiceHandler) GenerateShoukyakuItem(ctx context.Context, in *JournalRequest, out *JournalResponse) error {
	return h.ItemServiceHandler.GenerateShoukyakuItem(ctx, in, out)
}

func (h *itemServiceHandler) InventoryItem(ctx context.Context, in *InventoryItemRequest, out *InventoryItemResponse) error {
	return h.ItemServiceHandler.InventoryItem(ctx, in, out)
}

func (h *itemServiceHandler) MutilInventoryItem(ctx context.Context, in *MutilInventoryItemRequest, out *MutilInventoryItemResponse) error {
	return h.ItemServiceHandler.MutilInventoryItem(ctx, in, out)
}

func (h *itemServiceHandler) ResetInventoryItems(ctx context.Context, in *ResetInventoryItemsRequest, out *ResetInventoryItemsResponse) error {
	return h.ItemServiceHandler.ResetInventoryItems(ctx, in, out)
}

func (h *itemServiceHandler) DeleteItem(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error {
	return h.ItemServiceHandler.DeleteItem(ctx, in, out)
}

func (h *itemServiceHandler) DeleteDatastoreItems(ctx context.Context, in *DeleteDatastoreItemsRequest, out *DeleteResponse) error {
	return h.ItemServiceHandler.DeleteDatastoreItems(ctx, in, out)
}

func (h *itemServiceHandler) DeleteItems(ctx context.Context, in *DeleteItemsRequest, out *DeleteResponse) error {
	return h.ItemServiceHandler.DeleteItems(ctx, in, out)
}

func (h *itemServiceHandler) DeleteSelectItems(ctx context.Context, stream server.Stream) error {
	m := new(SelectedItemsRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.ItemServiceHandler.DeleteSelectItems(ctx, m, &itemServiceDeleteSelectItemsStream{stream})
}

type ItemService_DeleteSelectItemsStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*SelectedItemsResponse) error
}

type itemServiceDeleteSelectItemsStream struct {
	stream server.Stream
}

func (x *itemServiceDeleteSelectItemsStream) Close() error {
	return x.stream.Close()
}

func (x *itemServiceDeleteSelectItemsStream) Context() context.Context {
	return x.stream.Context()
}

func (x *itemServiceDeleteSelectItemsStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *itemServiceDeleteSelectItemsStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *itemServiceDeleteSelectItemsStream) Send(m *SelectedItemsResponse) error {
	return x.stream.Send(m)
}

func (h *itemServiceHandler) ChangeOwners(ctx context.Context, in *OwnersRequest, out *OwnersResponse) error {
	return h.ItemServiceHandler.ChangeOwners(ctx, in, out)
}

func (h *itemServiceHandler) ChangeSelectOwners(ctx context.Context, in *SelectOwnersRequest, out *SelectOwnersResponse) error {
	return h.ItemServiceHandler.ChangeSelectOwners(ctx, in, out)
}

func (h *itemServiceHandler) ChangeItemOwner(ctx context.Context, in *ItemOwnerRequest, out *ItemOwnerResponse) error {
	return h.ItemServiceHandler.ChangeItemOwner(ctx, in, out)
}

func (h *itemServiceHandler) ChangeStatus(ctx context.Context, in *StatusRequest, out *StatusResponse) error {
	return h.ItemServiceHandler.ChangeStatus(ctx, in, out)
}

func (h *itemServiceHandler) ChangeLabelTime(ctx context.Context, in *LabelTimeRequest, out *LabelTimeResponse) error {
	return h.ItemServiceHandler.ChangeLabelTime(ctx, in, out)
}

func (h *itemServiceHandler) ChangeDebt(ctx context.Context, in *ChangeDebtRequest, out *ChangeDebtResponse) error {
	return h.ItemServiceHandler.ChangeDebt(ctx, in, out)
}

func (h *itemServiceHandler) ContractExpire(ctx context.Context, in *ContractExpireRequest, out *ContractExpireResponse) error {
	return h.ItemServiceHandler.ContractExpire(ctx, in, out)
}

func (h *itemServiceHandler) ModifyContract(ctx context.Context, in *ModifyContractRequest, out *ModifyContractResponse) error {
	return h.ItemServiceHandler.ModifyContract(ctx, in, out)
}

func (h *itemServiceHandler) TerminateContract(ctx context.Context, in *TerminateContractRequest, out *TerminateContractResponse) error {
	return h.ItemServiceHandler.TerminateContract(ctx, in, out)
}

func (h *itemServiceHandler) ImportItem(ctx context.Context, stream server.Stream) error {
	return h.ItemServiceHandler.ImportItem(ctx, &itemServiceImportItemStream{stream})
}

type ItemService_ImportItemStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*ImportResponse) error
	Recv() (*ImportRequest, error)
}

type itemServiceImportItemStream struct {
	stream server.Stream
}

func (x *itemServiceImportItemStream) Close() error {
	return x.stream.Close()
}

func (x *itemServiceImportItemStream) Context() context.Context {
	return x.stream.Context()
}

func (x *itemServiceImportItemStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *itemServiceImportItemStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *itemServiceImportItemStream) Send(m *ImportResponse) error {
	return x.stream.Send(m)
}

func (x *itemServiceImportItemStream) Recv() (*ImportRequest, error) {
	m := new(ImportRequest)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (h *itemServiceHandler) ImportCheckItem(ctx context.Context, stream server.Stream) error {
	return h.ItemServiceHandler.ImportCheckItem(ctx, &itemServiceImportCheckItemStream{stream})
}

type ItemService_ImportCheckItemStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*ImportCheckResponse) error
	Recv() (*ImportCheckRequest, error)
}

type itemServiceImportCheckItemStream struct {
	stream server.Stream
}

func (x *itemServiceImportCheckItemStream) Close() error {
	return x.stream.Close()
}

func (x *itemServiceImportCheckItemStream) Context() context.Context {
	return x.stream.Context()
}

func (x *itemServiceImportCheckItemStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *itemServiceImportCheckItemStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *itemServiceImportCheckItemStream) Send(m *ImportCheckResponse) error {
	return x.stream.Send(m)
}

func (x *itemServiceImportCheckItemStream) Recv() (*ImportCheckRequest, error) {
	m := new(ImportCheckRequest)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (h *itemServiceHandler) MappingUpload(ctx context.Context, stream server.Stream) error {
	return h.ItemServiceHandler.MappingUpload(ctx, &itemServiceMappingUploadStream{stream})
}

type ItemService_MappingUploadStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*MappingUploadResponse) error
	Recv() (*MappingUploadRequest, error)
}

type itemServiceMappingUploadStream struct {
	stream server.Stream
}

func (x *itemServiceMappingUploadStream) Close() error {
	return x.stream.Close()
}

func (x *itemServiceMappingUploadStream) Context() context.Context {
	return x.stream.Context()
}

func (x *itemServiceMappingUploadStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *itemServiceMappingUploadStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *itemServiceMappingUploadStream) Send(m *MappingUploadResponse) error {
	return x.stream.Send(m)
}

func (x *itemServiceMappingUploadStream) Recv() (*MappingUploadRequest, error) {
	m := new(MappingUploadRequest)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (h *itemServiceHandler) Download(ctx context.Context, stream server.Stream) error {
	m := new(DownloadRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.ItemServiceHandler.Download(ctx, m, &itemServiceDownloadStream{stream})
}

func (h *itemServiceHandler) SwkDownload(ctx context.Context, stream server.Stream) error {
	m := new(DownloadRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.ItemServiceHandler.SwkDownload(ctx, m, &itemServiceDownloadStream{stream})
}

type ItemService_DownloadStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*DownloadResponse) error
}

type itemServiceDownloadStream struct {
	stream server.Stream
}

func (x *itemServiceDownloadStream) Close() error {
	return x.stream.Close()
}

func (x *itemServiceDownloadStream) Context() context.Context {
	return x.stream.Context()
}

func (x *itemServiceDownloadStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *itemServiceDownloadStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *itemServiceDownloadStream) Send(m *DownloadResponse) error {
	return x.stream.Send(m)
}

func (h *itemServiceHandler) FindAndModifyFile(ctx context.Context, stream server.Stream) error {
	m := new(FindRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.ItemServiceHandler.FindAndModifyFile(ctx, m, &itemServiceFindAndModifyFileStream{stream})
}

type ItemService_FindAndModifyFileStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*FindResponse) error
}

type itemServiceFindAndModifyFileStream struct {
	stream server.Stream
}

func (x *itemServiceFindAndModifyFileStream) Close() error {
	return x.stream.Close()
}

func (x *itemServiceFindAndModifyFileStream) Context() context.Context {
	return x.stream.Context()
}

func (x *itemServiceFindAndModifyFileStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *itemServiceFindAndModifyFileStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *itemServiceFindAndModifyFileStream) Send(m *FindResponse) error {
	return x.stream.Send(m)
}
