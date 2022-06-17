package storage_service

import gen "storage_service/gen/go/api/proto/v1"

type StorageServiceServer struct {
	gen.UnimplementedStorageServiceAPIServer
}
