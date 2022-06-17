package handler

import (
	"context"
	"encoding/base64"
	"github.com/google/uuid"
	gen "github.com/my-epoch/storage_service/gen/go/api/proto/v1"
	"github.com/my-epoch/storage_service/internal/storage_helper"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func UploadBase64EncodedFile(_ context.Context, request *gen.Base64FileUploadRequest) (*gen.FileUploadResponse, error) {
	if len(request.Base64FileEncoded) < 1 {
		return nil, status.Error(codes.InvalidArgument, "missed base64 encoded file")
	}

	content, err := base64.StdEncoding.DecodeString(request.Base64FileEncoded)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "cannot decode base64")
	}

	fileId, err := uuid.NewUUID()
	if err != nil {
		return nil, status.Error(codes.Internal, "something went wrong")
	}

	if err = storage_helper.PutFileInStorage(fileId, content); err != nil {
		return nil, status.Error(codes.Internal, "something went wrong")
	}

	response := &gen.FileUploadResponse{
		Uuid: fileId.String(),
		Size: uint32(len(content)),
	}

	return response, nil
}
