package coreserviceserver

import (
	pb "github.com/Alexander-s-Digital-Marketplace/core-service/internal/services/core_service/core_service_gen"
)

type Server struct {
	pb.UnimplementedCoreServiceServer
}
