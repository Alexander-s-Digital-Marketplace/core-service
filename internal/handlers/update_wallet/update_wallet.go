package updatewallet

import (
	"net/http"

	profilemodel "github.com/Alexander-s-Digital-Marketplace/core-service/internal/models/profile_model"
	paymentserviceclient "github.com/Alexander-s-Digital-Marketplace/core-service/internal/services/payment_service/payment_service_client"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Wallet struct {
	Wallet string `json:"wallet"`
}

func UpdateWallet(c *gin.Context) (int, string) {
	id, exists := c.Get("id")
	if !exists {
		return 400, "Bad request"
	}
	var wallet Wallet
	if err := c.ShouldBindJSON(&wallet); err != nil {
		logrus.Errorln("Error decode JSON: ", err)
		return 400, "Error decode JSON"
	}

	profile := profilemodel.Profile{
		AccountId: id.(int),
	}
	code := profile.GetFromTableByAccountId()
	if code != 200 {
		return code, "Error get profile"
	}

	var message string
	code, message = paymentserviceclient.UpdateWallet(profile.WalletId, wallet.Wallet)
	if code != 200 {
		return code, "Error update wallet"
	}

	return http.StatusOK, message
}
