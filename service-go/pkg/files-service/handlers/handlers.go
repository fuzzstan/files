package handlers

import (
	"context"

	"github.com/chaos-io/chaos/logs"

	"github.com/mojo-lang/core/go/pkg/mojo/core"

	"github.com/fuzzstan/files/go/pkg/files"

	// this service api
	pb "github.com/fuzzstan/files/go/pkg/files/v1"
)

var (
	_ = core.Timestamp{}
	_ = files.File{}
	_ = core.Null{}
)

type filesServer struct {
	pb.UnimplementedFilesServer
}

// NewService returns a naive, stateless implementation of Interface.
func NewService() pb.FilesServer {
	return filesServer{}
}

// GetFiles implements Interface.
func (s filesServer) GetFiles(ctx context.Context, in *pb.GetFilesRequest) (*pb.GetFilesResponse, error) {
	logs.Infow("GetFiles", "request", in)

	f := &files.File{
		Id:          "a1b2",
		Name:        "name1",
		Size:        100,
		Permalink:   "https://g.cn",
		DateAdded:   core.Now(),
		LastUpdated: nil,
		IsDeleted:   false,
	}

	resp := &pb.GetFilesResponse{
		Files: []*files.File{f},
		// TotalCount:
		// NextPageToken:
	}
	return resp, nil
}

// GetFile implements Interface.
func (s filesServer) GetFile(ctx context.Context, in *pb.GetFileRequest) (*files.File, error) {
	logs.Infow("GetFile", "request", in)

	resp := &files.File{
		// Id:
		// Name:
		// Size:
		// Permalink:
		// DateAdded:
		// LastUpdated:
		// IsDeleted:
	}
	return resp, nil
}

// CreateFile implements Interface.
func (s filesServer) CreateFile(ctx context.Context, in *pb.CreateFileRequest) (*files.File, error) {
	logs.Infow("CreateFile", "request", in)

	if in.File == nil {
		return nil, core.NewErrorFrom(400, "missing_parameter")
	}

	resp := &files.File{
		// Id:
		// Name:
		// Size:
		// Permalink:
		// DateAdded:
		// LastUpdated:
		// IsDeleted:
	}
	return resp, nil
}

// UpdateFile implements Interface.
func (s filesServer) UpdateFile(ctx context.Context, in *pb.UpdateFileRequest) (*files.File, error) {
	logs.Infow("UpdateFile", "request", in)

	resp := &files.File{
		// Id:
		// Name:
		// Size:
		// Permalink:
		// DateAdded:
		// LastUpdated:
		// IsDeleted:
	}
	return resp, nil
}

// DeleteFile implements Interface.
func (s filesServer) DeleteFile(ctx context.Context, in *pb.DeleteFileRequest) (*core.Null, error) {
	logs.Infow("DeleteFile", "request", in)

	resp := &core.Null{}
	return resp, nil
}
