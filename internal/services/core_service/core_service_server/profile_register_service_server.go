package coreserviceserver

import (
	"context"
	"errors"
	"net/http"

	profilemodel "github.com/Alexander-s-Digital-Marketplace/core-service/internal/models/profile_model"
	pb "github.com/Alexander-s-Digital-Marketplace/core-service/internal/services/core_service/core_service_gen"
	paymentserviceclient "github.com/Alexander-s-Digital-Marketplace/core-service/internal/services/payment_service/payment_service_client"
	"github.com/sirupsen/logrus"
)

func (s *Server) ProfileRegister(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	code, message, walletId := paymentserviceclient.RegisterWalet(req.Wallet)
	if code != 200 {
		logrus.Errorln(message)
		return &pb.Response{
			Code:    int32(code),
			Message: "Error register new wallet",
		}, errors.New("error register new wallet")
	}

	profile := profilemodel.Profile{
		UserName:    req.UserName,
		Rating:      0.0,
		CountRating: 0,
		AccountId:   int(req.AccountInfoId),
		WalletId:    walletId,
	}
	logrus.Infoln("walletId from payment", walletId)
	logrus.Infoln("New profile register", profile)
	code = profile.AddToTable()
	if code != 200 {
		logrus.Errorln("Error register new profile")
		return &pb.Response{
			Code:    int32(code),
			Message: "Error register new profile",
		}, errors.New("error register new profile")
	}

	return &pb.Response{
		Code:    int32(http.StatusOK),
		Message: "Success register new profile",
	}, nil
}
