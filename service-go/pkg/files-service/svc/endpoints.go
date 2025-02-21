// Code generated by chaosmojo. DO NOT EDIT.
// Rerunning chaosmojo will overwrite this file.
// Version: 1.0
// Version Date:

package svc

// This file contains methods to make individual endpoints from services,
// request and response types to serve those endpoints, as well as encoders and
// decoders for those types, for all of our supported transport serialization
// formats.

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/endpoint"

	"github.com/fuzzstan/files/go/pkg/files"
	"github.com/mojo-lang/core/go/pkg/mojo/core"

	// this service api
	pb "github.com/fuzzstan/files/go/pkg/files/v1"
)

var (
	_ = core.Timestamp{}
	_ = files.File{}
	_ = core.Null{}
)

// Endpoints collects all of the endpoints that compose an add service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
//
// In a server, it's useful for functions that need to operate on a per-endpoint
// basis. For example, you might pass an Endpoints to a function that produces
// an http.Handler, with each method (endpoint) wired up to a specific path. (It
// is probably a mistake in design to invoke the Service methods on the
// Endpoints struct in a server.)
//
// In a client, it's useful to collect individually constructed endpoints into a
// single type that implements the Service interface. For example, you might
// construct individual endpoints using transport/http.NewClient, combine them into an Endpoints, and return it to the caller as a Service.
type Endpoints struct {
	GetFilesEndpoint   endpoint.Endpoint
	GetFileEndpoint    endpoint.Endpoint
	CreateFileEndpoint endpoint.Endpoint
	UpdateFileEndpoint endpoint.Endpoint
	DeleteFileEndpoint endpoint.Endpoint
}

// Endpoints

func (e Endpoints) GetFiles(ctx context.Context, in *pb.GetFilesRequest) (*pb.GetFilesResponse, error) {
	response, err := e.GetFilesEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.GetFilesResponse), nil
}

func (e Endpoints) GetFile(ctx context.Context, in *pb.GetFileRequest) (*files.File, error) {
	response, err := e.GetFileEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*files.File), nil
}

func (e Endpoints) CreateFile(ctx context.Context, in *pb.CreateFileRequest) (*files.File, error) {
	response, err := e.CreateFileEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*files.File), nil
}

func (e Endpoints) UpdateFile(ctx context.Context, in *pb.UpdateFileRequest) (*files.File, error) {
	response, err := e.UpdateFileEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*files.File), nil
}

func (e Endpoints) DeleteFile(ctx context.Context, in *pb.DeleteFileRequest) (*core.Null, error) {
	response, err := e.DeleteFileEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*core.Null), nil
}

// Make Endpoints

func MakeGetFilesEndpoint(s pb.FilesServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.GetFilesRequest)
		v, err := s.GetFiles(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func MakeGetFileEndpoint(s pb.FilesServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.GetFileRequest)
		v, err := s.GetFile(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func MakeCreateFileEndpoint(s pb.FilesServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.CreateFileRequest)
		v, err := s.CreateFile(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func MakeUpdateFileEndpoint(s pb.FilesServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.UpdateFileRequest)
		v, err := s.UpdateFile(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func MakeDeleteFileEndpoint(s pb.FilesServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.DeleteFileRequest)
		v, err := s.DeleteFile(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

// WrapAllExcept wraps each Endpoint field of struct Endpoints with a
// go-kit/kit/endpoint.Middleware.
// Use this for applying a set of middlewares to every endpoint in the service.
// Optionally, endpoints can be passed in by name to be excluded from being wrapped.
// WrapAllExcept(middleware, "Status", "Ping")
func (e *Endpoints) WrapAllExcept(middleware endpoint.Middleware, excluded ...string) {
	included := map[string]struct{}{
		"get_files":   struct{}{},
		"get_file":    struct{}{},
		"create_file": struct{}{},
		"update_file": struct{}{},
		"delete_file": struct{}{},
	}

	for _, ex := range excluded {
		if _, ok := included[ex]; !ok {
			panic(fmt.Sprintf("Excluded endpoint '%s' does not exist; see middlewares/endpoints.go", ex))
		}
		delete(included, ex)
	}

	for inc, _ := range included {
		if inc == "get_files" {
			e.GetFilesEndpoint = middleware(e.GetFilesEndpoint)
		}
		if inc == "get_file" {
			e.GetFileEndpoint = middleware(e.GetFileEndpoint)
		}
		if inc == "create_file" {
			e.CreateFileEndpoint = middleware(e.CreateFileEndpoint)
		}
		if inc == "update_file" {
			e.UpdateFileEndpoint = middleware(e.UpdateFileEndpoint)
		}
		if inc == "delete_file" {
			e.DeleteFileEndpoint = middleware(e.DeleteFileEndpoint)
		}
	}
}

// LabeledMiddleware will get passed the endpoint name when passed to
// WrapAllLabeledExcept, this can be used to write a generic metrics
// middleware which can send the endpoint name to the metrics collector.
type LabeledMiddleware func(string, endpoint.Endpoint) endpoint.Endpoint

// WrapAllLabeledExcept wraps each Endpoint field of struct Endpoints with a
// LabeledMiddleware, which will receive the name of the endpoint. See
// LabeldMiddleware. See method WrapAllExept for details on excluded
// functionality.
func (e *Endpoints) WrapAllLabeledExcept(middleware func(string, endpoint.Endpoint) endpoint.Endpoint, excluded ...string) {
	included := map[string]struct{}{
		"get_files":   struct{}{},
		"get_file":    struct{}{},
		"create_file": struct{}{},
		"update_file": struct{}{},
		"delete_file": struct{}{},
	}

	for _, ex := range excluded {
		if _, ok := included[ex]; !ok {
			panic(fmt.Sprintf("Excluded endpoint '%s' does not exist; see middlewares/endpoints.go", ex))
		}
		delete(included, ex)
	}

	for inc, _ := range included {
		if inc == "get_files" {
			e.GetFilesEndpoint = middleware("get_files", e.GetFilesEndpoint)
		}
		if inc == "get_file" {
			e.GetFileEndpoint = middleware("get_file", e.GetFileEndpoint)
		}
		if inc == "create_file" {
			e.CreateFileEndpoint = middleware("create_file", e.CreateFileEndpoint)
		}
		if inc == "update_file" {
			e.UpdateFileEndpoint = middleware("update_file", e.UpdateFileEndpoint)
		}
		if inc == "delete_file" {
			e.DeleteFileEndpoint = middleware("delete_file", e.DeleteFileEndpoint)
		}
	}
}
