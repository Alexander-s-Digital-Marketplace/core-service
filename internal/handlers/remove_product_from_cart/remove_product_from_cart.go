package removeproductfromcart

import (
	cartmodel "github.com/Alexander-s-Digital-Marketplace/core-service/internal/models/cart_model"
	profilemodel "github.com/Alexander-s-Digital-Marketplace/core-service/internal/models/profile_model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func RemoveProductFromCart(c *gin.Context) (int, string) {
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
	profile := profilemodel.Profile{
		AccountId: id.(int),
	}

	code = profile.GetFromTableByAccountId()
	if code != 200 {
		return code, "Error get profile"
	}
	cart.ProfileId = profile.Id
	cart.Product.Id = cart.ProductId
	cart.Profile.Id = cart.ProfileId

	logrus.Infoln("cart: ", cart.ProfileId, cart.ProductId)

	if !cart.Product.IsBuy {
		code = cart.DeleteFromTable()
		if code != 200 {
			return code, "Error delete from cart"
		}
	} else {
		return 400, "Product is buyed"
	}

	return 200, "Success delete from cart"
}
