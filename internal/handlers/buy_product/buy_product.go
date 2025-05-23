package buyproduct

import (
	"context"
	"time"

	productmodel "github.com/Alexander-s-Digital-Marketplace/core-service/internal/models/product_model"
	profilemodel "github.com/Alexander-s-Digital-Marketplace/core-service/internal/models/profile_model"
	pb "github.com/Alexander-s-Digital-Marketplace/core-service/internal/services/payment_service/payment_service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func BuyProduct(c *gin.Context) (int, Contract) {
	var code int
	id, exists := c.Get("id")
	if !exists {
		return 400, Contract{}
	}
	buyer := profilemodel.Profile{
		AccountId: int(id.(int)),
	}
	code = buyer.GetFromTableByAccountId()
	if code != 200 {
		return code, Contract{}
	}

	var product productmodel.Product
	code = product.DecodeFromContext(c)
	if code != 200 {
		return code, Contract{}
	}
	code = product.GetFromTable()
	if code != 200 {
		return code, Contract{}
	}
	logrus.Infoln("product.Id", product.Id)

	conn, err := grpc.NewClient(
		"localhost:50054",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		logrus.Errorln("Error connect:", err)
		return 503, Contract{}
	}
	defer conn.Close()

	client := pb.NewPaymentServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.BuyProductRequest{
		WalletIdBuyer:  int32(buyer.WalletId),
		WalletIdSeller: int32(product.Seller.WalletId),
		ProductPrice:   product.Price,
		ProductId:      int32(product.Id),
	}

	logrus.Infoln("req:", req)

	res, err := client.BuyProduct(ctx, req)
	if err != nil {
		logrus.Errorln("Error send message:", err)
		return 503, Contract{}
	}
	contract := Contract{
		OrderId:       int(res.OrderId),
		Address:       res.Address,
		SellerAddress: res.SellerAddress,
		Price:         res.ProductPrice,
	}

	logrus.Info("contract", contract)

	return int(res.Code), contract
}
