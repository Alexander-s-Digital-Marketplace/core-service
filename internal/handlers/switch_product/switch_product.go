package switchproduct

import (
	productmodel "github.com/Alexander-s-Digital-Marketplace/core-service/internal/models/product_model"
	"github.com/gin-gonic/gin"
)

func SwitchProduct(c *gin.Context) (int, string) {
	id, exists := c.Get("id")
	if !exists {
		return 400, "Error get id"
	}

	var product productmodel.Product
	var code int

	product.DecodeFromContext(c)
	product.SellerId = id.(int)
	code = product.Switch(product.IsSellNow)
	if code != 200 {
		return code, "Error publish/remove product"
	}

	if product.IsSellNow {
		return 200, "Success publish product"
	} else {
		return 200, "Success remove publish product"
	}
}
