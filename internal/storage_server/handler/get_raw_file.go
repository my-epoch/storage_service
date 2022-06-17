package handler

import (
	"context"
	"github.com/gabriel-vasile/mimetype"
	"github.com/google/uuid"
	gen "github.com/my-epoch/storage_service/gen/go/api/proto/v1"
	"github.com/my-epoch/storage_service/internal/storage_helper"
	"google.golang.org/genproto/googleapis/api/httpbody"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetRawFile(_ context.Context, request *gen.GetRawFileRequest) (*httpbody.HttpBody, error) {
	if !storage_helper.IsFileOrDirectoryExist("storage/" + request.Uuid) {
		return nil, status.Error(codes.NotFound, "file not found on server")
	}

	fileId, err := uuid.Parse(request.Uuid)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "cannot parse file id")
	}

	file, err := storage_helper.GetFileFromStorage(fileId)
	if err != nil {
		return nil, status.Error(codes.Internal, "something went wrong")
	}

	mime := mimetype.Detect(file)

	response := &httpbody.HttpBody{
		ContentType: mime.String(),
		Data:        file,
	}

	return response, nil
}
