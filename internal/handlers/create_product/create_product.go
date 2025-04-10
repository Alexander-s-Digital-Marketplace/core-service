package createproduct

import (
	productmodel "github.com/Alexander-s-Digital-Marketplace/core-service/internal/models/product_model"
	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) (int, string) {
	var product productmodel.Product
	var code int
	code = product.DecodeFromContext(c)
	if code != 200 {
		return code, "Error decode JSON"
	}

	code = product.AddToTable()
	if code != 200 {
		return code, "Error add to table"
	}

	return 200, "Success create product"
}
