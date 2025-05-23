package getbalance

import (
	"net/http"

	profilemodel "github.com/Alexander-s-Digital-Marketplace/core-service/internal/models/profile_model"
	paymentserviceclient "github.com/Alexander-s-Digital-Marketplace/core-service/internal/services/payment_service/payment_service_client"
	"github.com/gin-gonic/gin"
)

func GetBalance(c *gin.Context) (int, float64) {
	id, exists := c.Get("id")
	if !exists {
		return 400, 0
	}
	profile := profilemodel.Profile{
		AccountId: id.(int),
	}
	code := profile.GetFromTableByAccountId()
	if code != 200 {
		return code, 0
	}

	var balance float64
	code, balance = paymentserviceclient.GetBalance(profile.WalletId)
	if code != 200 {
		return code, 0
	}

	return http.StatusOK, balance
}
