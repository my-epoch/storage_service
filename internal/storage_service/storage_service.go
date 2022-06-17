package storage_service

import gen "github.com/my-epoch/storage_service/gen/go/api/proto/v1"

type StorageServiceServer struct {
	gen.UnimplementedStorageServiceAPIServer
}
