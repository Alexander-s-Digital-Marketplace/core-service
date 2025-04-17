package switchproductcart

import (
	"time"

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
	cart.Product.Id = cart.ProductId
	cart.Profile.Id = cart.ProfileId
	cart.Date = time.Now().Format("2006-01-02 15:04:05")

	code = cart.AddToTable()
	if code != 200 {
		return code, "Error add to cart"
	}

	return 200, "Success add to cart"
}
