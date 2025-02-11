// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: journal.proto

package journal

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

// Api Endpoints for JournalService service

func NewJournalServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for JournalService service

type JournalService interface {
	FindJournals(ctx context.Context, in *JournalsRequest, opts ...client.CallOption) (*JournalsResponse, error)
	FindJournal(ctx context.Context, in *JournalRequest, opts ...client.CallOption) (*JournalResponse, error)
	ImportJournal(ctx context.Context, in *ImportRequest, opts ...client.CallOption) (*ImportResponse, error)
	ModifyJournal(ctx context.Context, in *ModifyRequest, opts ...client.CallOption) (*ModifyResponse, error)
	AddDownloadSetting(ctx context.Context, in *AddDownloadSettingRequest, opts ...client.CallOption) (*AddDownloadSettingResponse, error)
	FindDownloadSetting(ctx context.Context, in *FindDownloadSettingRequest, opts ...client.CallOption) (*FindDownloadSettingResponse, error)
	FindDownloadSettings(ctx context.Context, in *FindDownloadSettingsRequest, opts ...client.CallOption) (*FindDownloadSettingsResponse, error)
	AddSelectJournals(ctx context.Context, in *AddSelectJournalsRequest, opts ...client.CallOption) (*AddSelectJournalsResponse, error)
	FindSelectJournals(ctx context.Context, in *JournalsRequest, opts ...client.CallOption) (*JournalsResponse, error)
}

type journalService struct {
	c    client.Client
	name string
}

func NewJournalService(name string, c client.Client) JournalService {
	return &journalService{
		c:    c,
		name: name,
	}
}

func (c *journalService) FindJournals(ctx context.Context, in *JournalsRequest, opts ...client.CallOption) (*JournalsResponse, error) {
	req := c.c.NewRequest(c.name, "JournalService.FindJournals", in)
	out := new(JournalsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *journalService) FindJournal(ctx context.Context, in *JournalRequest, opts ...client.CallOption) (*JournalResponse, error) {
	req := c.c.NewRequest(c.name, "JournalService.FindJournal", in)
	out := new(JournalResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *journalService) ImportJournal(ctx context.Context, in *ImportRequest, opts ...client.CallOption) (*ImportResponse, error) {
	req := c.c.NewRequest(c.name, "JournalService.ImportJournal", in)
	out := new(ImportResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *journalService) ModifyJournal(ctx context.Context, in *ModifyRequest, opts ...client.CallOption) (*ModifyResponse, error) {
	req := c.c.NewRequest(c.name, "JournalService.ModifyJournal", in)
	out := new(ModifyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *journalService) AddDownloadSetting(ctx context.Context, in *AddDownloadSettingRequest, opts ...client.CallOption) (*AddDownloadSettingResponse, error) {
	req := c.c.NewRequest(c.name, "JournalService.AddDownloadSetting", in)
	out := new(AddDownloadSettingResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *journalService) FindDownloadSetting(ctx context.Context, in *FindDownloadSettingRequest, opts ...client.CallOption) (*FindDownloadSettingResponse, error) {
	req := c.c.NewRequest(c.name, "JournalService.FindDownloadSetting", in)
	out := new(FindDownloadSettingResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *journalService) FindDownloadSettings(ctx context.Context, in *FindDownloadSettingsRequest, opts ...client.CallOption) (*FindDownloadSettingsResponse, error) {
	req := c.c.NewRequest(c.name, "JournalService.FindDownloadSettings", in)
	out := new(FindDownloadSettingsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *journalService) AddSelectJournals(ctx context.Context, in *AddSelectJournalsRequest, opts ...client.CallOption) (*AddSelectJournalsResponse, error) {
	req := c.c.NewRequest(c.name, "JournalService.AddSelectJournals", in)
	out := new(AddSelectJournalsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *journalService) FindSelectJournals(ctx context.Context, in *JournalsRequest, opts ...client.CallOption) (*JournalsResponse, error) {
	req := c.c.NewRequest(c.name, "JournalService.FindSelectJournals", in)
	out := new(JournalsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for JournalService service

type JournalServiceHandler interface {
	FindJournals(context.Context, *JournalsRequest, *JournalsResponse) error
	FindJournal(context.Context, *JournalRequest, *JournalResponse) error
	ImportJournal(context.Context, *ImportRequest, *ImportResponse) error
	ModifyJournal(context.Context, *ModifyRequest, *ModifyResponse) error
	AddDownloadSetting(context.Context, *AddDownloadSettingRequest, *AddDownloadSettingResponse) error
	FindDownloadSetting(context.Context, *FindDownloadSettingRequest, *FindDownloadSettingResponse) error
	FindDownloadSettings(context.Context, *FindDownloadSettingsRequest, *FindDownloadSettingsResponse) error
	AddSelectJournals(context.Context, *AddSelectJournalsRequest, *AddSelectJournalsResponse) error
	FindSelectJournals(context.Context, *JournalsRequest, *JournalsResponse) error
}

func RegisterJournalServiceHandler(s server.Server, hdlr JournalServiceHandler, opts ...server.HandlerOption) error {
	type journalService interface {
		FindJournals(ctx context.Context, in *JournalsRequest, out *JournalsResponse) error
		FindJournal(ctx context.Context, in *JournalRequest, out *JournalResponse) error
		ImportJournal(ctx context.Context, in *ImportRequest, out *ImportResponse) error
		ModifyJournal(ctx context.Context, in *ModifyRequest, out *ModifyResponse) error
		AddDownloadSetting(ctx context.Context, in *AddDownloadSettingRequest, out *AddDownloadSettingResponse) error
		FindDownloadSetting(ctx context.Context, in *FindDownloadSettingRequest, out *FindDownloadSettingResponse) error
		FindDownloadSettings(ctx context.Context, in *FindDownloadSettingsRequest, out *FindDownloadSettingsResponse) error
		AddSelectJournals(ctx context.Context, in *AddSelectJournalsRequest, out *AddSelectJournalsResponse) error
		FindSelectJournals(ctx context.Context, in *JournalsRequest, out *JournalsResponse) error
	}
	type JournalService struct {
		journalService
	}
	h := &journalServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&JournalService{h}, opts...))
}

type journalServiceHandler struct {
	JournalServiceHandler
}

func (h *journalServiceHandler) FindJournals(ctx context.Context, in *JournalsRequest, out *JournalsResponse) error {
	return h.JournalServiceHandler.FindJournals(ctx, in, out)
}

func (h *journalServiceHandler) FindJournal(ctx context.Context, in *JournalRequest, out *JournalResponse) error {
	return h.JournalServiceHandler.FindJournal(ctx, in, out)
}

func (h *journalServiceHandler) ImportJournal(ctx context.Context, in *ImportRequest, out *ImportResponse) error {
	return h.JournalServiceHandler.ImportJournal(ctx, in, out)
}

func (h *journalServiceHandler) ModifyJournal(ctx context.Context, in *ModifyRequest, out *ModifyResponse) error {
	return h.JournalServiceHandler.ModifyJournal(ctx, in, out)
}

func (h *journalServiceHandler) AddDownloadSetting(ctx context.Context, in *AddDownloadSettingRequest, out *AddDownloadSettingResponse) error {
	return h.JournalServiceHandler.AddDownloadSetting(ctx, in, out)
}

func (h *journalServiceHandler) FindDownloadSetting(ctx context.Context, in *FindDownloadSettingRequest, out *FindDownloadSettingResponse) error {
	return h.JournalServiceHandler.FindDownloadSetting(ctx, in, out)
}

func (h *journalServiceHandler) FindDownloadSettings(ctx context.Context, in *FindDownloadSettingsRequest, out *FindDownloadSettingsResponse) error {
	return h.JournalServiceHandler.FindDownloadSettings(ctx, in, out)
}

func (h *journalServiceHandler) AddSelectJournals(ctx context.Context, in *AddSelectJournalsRequest, out *AddSelectJournalsResponse) error {
	return h.JournalServiceHandler.AddSelectJournals(ctx, in, out)
}

func (h *journalServiceHandler) FindSelectJournals(ctx context.Context, in *JournalsRequest, out *JournalsResponse) error {
	return h.JournalServiceHandler.FindSelectJournals(ctx, in, out)
}
