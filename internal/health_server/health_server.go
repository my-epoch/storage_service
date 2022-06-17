package health_server

import (
	"context"
	"google.golang.org/grpc/codes"
	gh "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
)

type HealthServer struct {
	gh.UnimplementedHealthServer
}

func (HealthServer) Check(context.Context, *gh.HealthCheckRequest) (*gh.HealthCheckResponse, error) {
	return &gh.HealthCheckResponse{Status: gh.HealthCheckResponse_SERVING}, nil
}

func (HealthServer) Watch(_ *gh.HealthCheckRequest, _ gh.Health_WatchServer) error {
	return status.Error(codes.Unimplemented, "Unimplemented")
}
