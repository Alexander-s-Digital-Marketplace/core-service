package coreserviceserver

import (
	"context"
	"errors"
	"net/http"
	"time"

	historymodel "github.com/Alexander-s-Digital-Marketplace/core-service/internal/models/history_model"
	productmodel "github.com/Alexander-s-Digital-Marketplace/core-service/internal/models/product_model"
	profilemodel "github.com/Alexander-s-Digital-Marketplace/core-service/internal/models/profile_model"
	pb "github.com/Alexander-s-Digital-Marketplace/core-service/internal/services/core_service/core_service_gen"
	"github.com/sirupsen/logrus"
)

func (s *Server) UpdateSoldProduct(ctx context.Context, req *pb.UpdateSoldProductRequest) (*pb.UpdateSoldProductResponse, error) {
	product := productmodel.Product{
		Id: int(req.ProductId),
	}
	code := product.GetFromTable()
	if code != 200 {
		return &pb.UpdateSoldProductResponse{
			Code:    int32(code),
			Message: "error get product from table",
		}, errors.New("error get product from table")
	}

	product.OrderId = int(req.OrderId)
	product.IsBuy = true
	product.IsSellNow = false
	code = product.UpdateInTable()
	logrus.Infoln("product.IsBuy: ", product.IsBuy)
	if code != 200 {
		return &pb.UpdateSoldProductResponse{
			Code:    int32(code),
			Message: "error update product from table",
		}, errors.New("error update product from table")
	}

	profile := profilemodel.Profile{
		WalletId: int(req.WalletId),
	}

	code = profile.GetFromTableByWalletId()
	if code != 200 {
		return &pb.UpdateSoldProductResponse{
			Code:    int32(code),
			Message: "error get profile from table",
		}, errors.New("error get profile from table")
	}

	history := historymodel.History{
		Date:      time.Now().Format("2006-01-02 15:04"),
		BuyerId:   profile.Id,
		ProductId: int(req.ProductId),
	}

	code = history.AddToTable()
	if code != 200 {
		return &pb.UpdateSoldProductResponse{
			Code:    int32(code),
			Message: "error create history",
		}, errors.New("error create history")
	}
	return &pb.UpdateSoldProductResponse{
		Code:    int32(http.StatusOK),
		Message: "Succes buy product",
	}, nil
}
