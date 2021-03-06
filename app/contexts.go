// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "hostDog": Application Contexts
//
// Command:
// $ goagen
// --design=github.com/YAWAL/hostDog/design
// --out=$(GOPATH)\src\github.com\YAWAL\hostDog
// --version=v1.3.1

package app

import (
	"context"
	"github.com/goadesign/goa"
	uuid "github.com/gofrs/uuid"
	"net/http"
)

// CreateDogContext provides the dog create action context.
type CreateDogContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *CreateDogPayload
}

// NewCreateDogContext parses the incoming request URL and body, performs validations and creates the
// context used by the dog controller create action.
func NewCreateDogContext(ctx context.Context, r *http.Request, service *goa.Service) (*CreateDogContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := CreateDogContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// Created sends a HTTP response with status code 201.
func (ctx *CreateDogContext) Created() error {
	ctx.ResponseData.WriteHeader(201)
	return nil
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *CreateDogContext) BadRequest() error {
	ctx.ResponseData.WriteHeader(400)
	return nil
}

// DeleteDogContext provides the dog delete action context.
type DeleteDogContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	DogID uuid.UUID
}

// NewDeleteDogContext parses the incoming request URL and body, performs validations and creates the
// context used by the dog controller delete action.
func NewDeleteDogContext(ctx context.Context, r *http.Request, service *goa.Service) (*DeleteDogContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := DeleteDogContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramDogID := req.Params["dogID"]
	if len(paramDogID) > 0 {
		rawDogID := paramDogID[0]
		if dogID, err2 := uuid.FromString(rawDogID); err2 == nil {
			rctx.DogID = dogID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("dogID", rawDogID, "uuid"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *DeleteDogContext) OK(resp []byte) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	}
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *DeleteDogContext) BadRequest() error {
	ctx.ResponseData.WriteHeader(400)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DeleteDogContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// CreateHostContext provides the host create action context.
type CreateHostContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *CreateHostPayload
}

// NewCreateHostContext parses the incoming request URL and body, performs validations and creates the
// context used by the host controller create action.
func NewCreateHostContext(ctx context.Context, r *http.Request, service *goa.Service) (*CreateHostContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := CreateHostContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// Created sends a HTTP response with status code 201.
func (ctx *CreateHostContext) Created() error {
	ctx.ResponseData.WriteHeader(201)
	return nil
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *CreateHostContext) BadRequest() error {
	ctx.ResponseData.WriteHeader(400)
	return nil
}

// DeleteHostContext provides the host delete action context.
type DeleteHostContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	HostID uuid.UUID
}

// NewDeleteHostContext parses the incoming request URL and body, performs validations and creates the
// context used by the host controller delete action.
func NewDeleteHostContext(ctx context.Context, r *http.Request, service *goa.Service) (*DeleteHostContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := DeleteHostContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramHostID := req.Params["hostID"]
	if len(paramHostID) > 0 {
		rawHostID := paramHostID[0]
		if hostID, err2 := uuid.FromString(rawHostID); err2 == nil {
			rctx.HostID = hostID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("hostID", rawHostID, "uuid"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *DeleteHostContext) OK(resp []byte) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	}
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *DeleteHostContext) BadRequest() error {
	ctx.ResponseData.WriteHeader(400)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DeleteHostContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// GetAllHostContext provides the host getAll action context.
type GetAllHostContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewGetAllHostContext parses the incoming request URL and body, performs validations and creates the
// context used by the host controller getAll action.
func NewGetAllHostContext(ctx context.Context, r *http.Request, service *goa.Service) (*GetAllHostContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := GetAllHostContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *GetAllHostContext) OK(r HostsresponseCollection) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "hostsresponse; type=collection")
	}
	if r == nil {
		r = HostsresponseCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *GetAllHostContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// GetByIDHostContext provides the host getByID action context.
type GetByIDHostContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	HostID string
}

// NewGetByIDHostContext parses the incoming request URL and body, performs validations and creates the
// context used by the host controller getByID action.
func NewGetByIDHostContext(ctx context.Context, r *http.Request, service *goa.Service) (*GetByIDHostContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := GetByIDHostContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramHostID := req.Params["hostID"]
	if len(paramHostID) > 0 {
		rawHostID := paramHostID[0]
		rctx.HostID = rawHostID
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *GetByIDHostContext) OK(r *Hostwithdogs) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "hostwithdogs")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *GetByIDHostContext) BadRequest() error {
	ctx.ResponseData.WriteHeader(400)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *GetByIDHostContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// GetDogsByHostIDHostContext provides the host getDogsByHostID action context.
type GetDogsByHostIDHostContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	HostID string
}

// NewGetDogsByHostIDHostContext parses the incoming request URL and body, performs validations and creates the
// context used by the host controller getDogsByHostID action.
func NewGetDogsByHostIDHostContext(ctx context.Context, r *http.Request, service *goa.Service) (*GetDogsByHostIDHostContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := GetDogsByHostIDHostContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramHostID := req.Params["hostID"]
	if len(paramHostID) > 0 {
		rawHostID := paramHostID[0]
		rctx.HostID = rawHostID
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *GetDogsByHostIDHostContext) OK(r *Dogs) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "dogs")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *GetDogsByHostIDHostContext) BadRequest() error {
	ctx.ResponseData.WriteHeader(400)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *GetDogsByHostIDHostContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}
