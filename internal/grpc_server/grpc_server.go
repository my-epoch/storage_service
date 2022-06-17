package grpc_server

import (
	"fmt"
	"github.com/my-epoch/api-gateway/pkg/logger"
	"github.com/my-epoch/object_service/pkg/service_config"
	"google.golang.org/grpc"
	gh "google.golang.org/grpc/health/grpc_health_v1"
	"net"
	gen "storage_service/gen/go/api/proto/v1"
	"storage_service/internal/health_server"
	"storage_service/internal/storage_service"
)

func Serve() {
	grpcServer := grpc.NewServer()

	gh.RegisterHealthServer(grpcServer, &health_server.HealthServer{})

	gen.RegisterStorageServiceAPIServer(grpcServer, &storage_service.StorageServiceServer{})

	cfg := service_config.Get()
	listenAddr := fmt.Sprintf("%s:%d", cfg.Addr, cfg.Port)

	lis, err := net.Listen("tcp", listenAddr)
	if err != nil {
		logger.Fatalf("cannot start server: %e", err)
	}

	logger.Infof("gRPC server starts listening on '%s'", cfg.Addr)
	defer grpcServer.Stop()
	if err := grpcServer.Serve(lis); err != nil {
		logger.Fatalf("cannot start gRPC server: %e", err)
	}

}
