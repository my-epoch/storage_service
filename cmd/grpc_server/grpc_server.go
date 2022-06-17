package main

import (
	"github.com/my-epoch/object_service/pkg/consul"
	"github.com/my-epoch/object_service/pkg/service_config"
	"github.com/my-epoch/storage_service/internal/grpc_server"
)

func main() {
	consul.RegisterService(service_config.Get())
	defer consul.DeregisterService()

	grpc_server.Serve()
}
