package notificationservice

import (
	"context"
	"time"

	productmodel "github.com/Alexander-s-Digital-Marketplace/core-service/internal/models/product_model"
	pb "github.com/Alexander-s-Digital-Marketplace/core-service/internal/services/notification_service/notification_service_gen"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DeliverNotif(product productmodel.Product, email string) (int, string) {
	conn, err := grpc.NewClient(
		"localhost:50053",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		logrus.Errorln("Error connect:", err)
		return 503, ""
	}
	defer conn.Close()

	client := pb.NewNotificationServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.DeliverRequest{
		Email:   email,
		Product: product.Title,
		Item:    product.Item.Content,
	}

	logrus.Infoln("req:", req)

	res, err := client.DeliverNotif(ctx, req)
	if err != nil {
		logrus.Errorln("Error send message:", err)
		return 503, ""
	}

	return int(res.Code), res.Message
}
