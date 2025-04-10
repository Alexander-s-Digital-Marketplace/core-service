package switchproductcart

import (
	cartmodel "github.com/Alexander-s-Digital-Marketplace/core-service/internal/models/cart_model"
	"github.com/gin-gonic/gin"
)

func SwitchProductCart(c *gin.Context) (int, string) {
	id, exists := c.Get("id")
	if !exists {
		return 400, "Error get id"
	}

	var cart cartmodel.Cart
	var code int
	code = cart.DecodeFromContext(c)
	if code != 200 {
		return code, "Error decode JSON"
	}

	cart.ProfileId = id.(int)

	code = cart.AddToTable()
	if code != 200 {
		return code, "Error add to cart"
	}

	return 200, "Success add to cart"
}
