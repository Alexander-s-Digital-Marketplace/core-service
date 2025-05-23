package getallfeed

import (
	productmodel "github.com/Alexander-s-Digital-Marketplace/core-service/internal/models/product_model"
	"github.com/gin-gonic/gin"
)

func GetAllFeed(c *gin.Context) (int, []productmodel.Product) {
	var product productmodel.Product
	var products []productmodel.Product
	var productsToFront []productmodel.Product
	var code int

	code, products = product.GetAllFromTable()
	if code != 200 {
		return code, []productmodel.Product{}
	}

	for _, product := range products {
		if !product.IsBuy && product.IsSellNow {
			productsToFront = append(productsToFront, product)
		}
	}

	return 200, productsToFront
}
