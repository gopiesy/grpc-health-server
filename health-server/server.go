package server

import (
	"context"

	health "github.com/gopiesy/grpc-health-server/proto"
)

type HealthServer struct {
	health.UnimplementedHealthServer
}

func (h HealthServer) Check(ctx context.Context, req *health.HealthCheckRequest) (*health.HealthCheckResponse, error) {
	if req != nil {
		if req.Service == "health" {
			return &health.HealthCheckResponse{Status: health.HealthCheckResponse_SERVING}, nil
		}
		return &health.HealthCheckResponse{Status: health.HealthCheckResponse_SERVICE_UNKNOWN}, nil
	}
	return &health.HealthCheckResponse{Status: health.HealthCheckResponse_UNKNOWN}, nil
}

func NewHealthServer() HealthServer {
	return HealthServer{}
}
