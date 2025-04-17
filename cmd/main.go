/*
 * Catering service
 *
 * Auth service.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"log"
	"net"

	loggerconfig "github.com/Alexander-s-Digital-Marketplace/core-service/internal/config/logger"
	routespkg "github.com/Alexander-s-Digital-Marketplace/core-service/internal/routes"
	pb "github.com/Alexander-s-Digital-Marketplace/core-service/internal/services/profile_register_service/profile_register_service_gen"
	profileregisterserviceserver "github.com/Alexander-s-Digital-Marketplace/core-service/internal/services/profile_register_service/profile_register_service_server"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	loggerconfig.Init()

	go func() {
		routes := routespkg.ApiHandleFunctions{}
		logrus.Printf("Server started")
		router := routespkg.NewRouter(routes)
		logrus.Fatal(router.Run(":8081"))
	}()

	go func() {
		listener, err := net.Listen("tcp", ":50052")
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}

		grpcServer := grpc.NewServer()

		pb.RegisterProfileRegisterServiceServer(grpcServer, &profileregisterserviceserver.Server{})

		logrus.Println("gRPC server is running on port :50052")
		if err := grpcServer.Serve(listener); err != nil {
			logrus.Fatalf("Failed to serve gRPC server: %v", err)
		}
	}()

	select {}
}
