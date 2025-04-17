package getallfeed

import (
	productmodel "github.com/Alexander-s-Digital-Marketplace/core-service/internal/models/product_model"
	"github.com/gin-gonic/gin"
)

func GetAllFeed(c *gin.Context) (int, []productmodel.Product) {
	var product productmodel.Product
	var products []productmodel.Product
	var code int

	code, products = product.GetAllFromTable()
	if code != 200 {
		return code, []productmodel.Product{}
	}

	return 200, products
}
