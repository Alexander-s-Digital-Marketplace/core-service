package getcart

import (
	cartmodel "github.com/Alexander-s-Digital-Marketplace/core-service/internal/models/cart_model"
	productmodel "github.com/Alexander-s-Digital-Marketplace/core-service/internal/models/product_model"
	profilemodel "github.com/Alexander-s-Digital-Marketplace/core-service/internal/models/profile_model"
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
	profile := profilemodel.Profile{
		AccountId: id.(int),
	}
	code = profile.GetFromTableByAccountId()
	if code != 200 {
		return code, []productmodel.Product{}
	}
	cart.ProfileId = profile.Id
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
