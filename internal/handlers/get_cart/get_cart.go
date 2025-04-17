package getcart

import (
	cartmodel "github.com/Alexander-s-Digital-Marketplace/core-service/internal/models/cart_model"
	productmodel "github.com/Alexander-s-Digital-Marketplace/core-service/internal/models/product_model"
	"github.com/gin-gonic/gin"
)

func GetCart(c *gin.Context) (int, []productmodel.Product) {
	id, exists := c.Get("id")
	if !exists {
		return 400, []productmodel.Product{}
	}

	var cart cartmodel.Cart
	var carts []cartmodel.Cart
	var code int
	cart.ProfileId = id.(int)
	code, carts = cart.GetAllFromTableByProfileId()
	if code != 200 {
		return code, []productmodel.Product{}
	}
	var products []productmodel.Product
	for _, cartInCarts := range carts {
		products = append(products, cartInCarts.Product)
	}
	return 200, products
}
