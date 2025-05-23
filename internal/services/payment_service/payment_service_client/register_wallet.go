package paymentserviceclient

import (
	"context"
	"time"

	pb "github.com/Alexander-s-Digital-Marketplace/core-service/internal/services/payment_service/payment_service"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RegisterWalet(wallet string) (int, string, int) {
	conn, err := grpc.NewClient(
		"localhost:50054",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		logrus.Errorln("Error connect:", err)
		return 503, "", 0
	}
	defer conn.Close()

	client := pb.NewPaymentServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.RegisterWalletRequest{
		WalletAddress: wallet,
	}

	res, err := client.RegisterWallet(ctx, req)
	if err != nil {
		logrus.Errorln("Error wallet register:", err)
		return 503, "", 0
	}
	logrus.Infoln("res from payment: ", res)

	return int(res.Code), res.Message, int(res.WalletId)
}
