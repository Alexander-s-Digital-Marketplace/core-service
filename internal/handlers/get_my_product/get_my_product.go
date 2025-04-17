package getmyproduct

import (
	productmodel "github.com/Alexander-s-Digital-Marketplace/core-service/internal/models/product_model"
	"github.com/gin-gonic/gin"
)

func GetMyProduct(c *gin.Context) (int, []productmodel.Product) {
	id, exists := c.Get("id")
	if !exists {
		return 400, []productmodel.Product{}
	}

	var product productmodel.Product
	var products []productmodel.Product
	var code int

	product.SellerId = id.(int)
	code, products = product.GetAllMyFromTable()
	if code != 200 {
		return code, []productmodel.Product{}
	}

	return 200, products
}
