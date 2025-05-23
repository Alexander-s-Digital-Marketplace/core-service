package authserviceclient

import (
	"context"
	"time"

	pb "github.com/Alexander-s-Digital-Marketplace/core-service/internal/services/auth_service/auth_service_gen"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GetEmailByAccountId(id int) (int, string) {
	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		logrus.Errorln("Error connect:", err)
		return 503, ""
	}
	defer conn.Close()

	client := pb.NewAuthServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.EmailRequest{Id: int32(id)}

	res, err := client.GetEmailByAccountId(ctx, req)
	if err != nil {
		logrus.Errorln("Error send message:", err)
		return 503, ""
	}

	return int(res.Code), res.Email
}
