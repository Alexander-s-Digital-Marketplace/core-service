package coreserviceserver

import (
	"context"
	"errors"
	"net/http"

	profilemodel "github.com/Alexander-s-Digital-Marketplace/core-service/internal/models/profile_model"
	pb "github.com/Alexander-s-Digital-Marketplace/core-service/internal/services/core_service/core_service_gen"
	"github.com/sirupsen/logrus"
)

func (s *Server) ProfileRegister(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	var profile profilemodel.Profile
	profile.UserName = req.UserName
	profile.Rating = 0.0
	profile.CountRating = 0
	profile.AccountId = int(req.AccountInfoId)

	code := profile.AddToTable()
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
