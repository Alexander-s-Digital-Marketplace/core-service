package getwallet

import (
	"net/http"

	profilemodel "github.com/Alexander-s-Digital-Marketplace/core-service/internal/models/profile_model"
	paymentserviceclient "github.com/Alexander-s-Digital-Marketplace/core-service/internal/services/payment_service/payment_service_client"
	"github.com/gin-gonic/gin"
)

func GetWallet(c *gin.Context) (int, string) {
	id, exists := c.Get("id")
	if !exists {
		return 400, "Bad request"
	}

	profile := profilemodel.Profile{
		AccountId: id.(int),
	}
	code := profile.GetFromTableByAccountId()
	if code != 200 {
		return code, "error get profile from table"
	}

	var wallet string
	code, wallet = paymentserviceclient.GetWallet(profile.WalletId)
	if code != 200 {
		return code, "error get wallet"
	}

	return http.StatusOK, wallet
}
