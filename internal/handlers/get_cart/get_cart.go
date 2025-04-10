package getcart

import (
	cartmodel "github.com/Alexander-s-Digital-Marketplace/core-service/internal/models/cart_model"
	"github.com/gin-gonic/gin"
)

func GetCart(c *gin.Context) (int, []cartmodel.Cart) {
	var cart cartmodel.Cart
	var carts []cartmodel.Cart
	var code int
	//	add get user id from gin.context
	code, carts = cart.GetAllFromTableByProfileId()
	if code != 200 {
		return code, []cartmodel.Cart{}
	}
	return 200, carts
}
