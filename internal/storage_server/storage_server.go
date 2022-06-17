package storage_server

import (
	"context"
	gen "github.com/my-epoch/storage_service/gen/go/api/proto/v1"
	"github.com/my-epoch/storage_service/internal/storage_server/handler"
	"google.golang.org/genproto/googleapis/api/httpbody"
)

type StorageServiceServer struct {
	gen.UnimplementedStorageServiceAPIServer
}

func (StorageServiceServer) UploadBase64EncodedFile(ctx context.Context, request *gen.Base64FileUploadRequest) (*gen.FileUploadResponse, error) {
	return handler.UploadBase64EncodedFile(ctx, request)
}

func (StorageServiceServer) GetRawFile(ctx context.Context, request *gen.GetRawFileRequest) (*httpbody.HttpBody, error) {
	return handler.GetRawFile(ctx, request)
}
