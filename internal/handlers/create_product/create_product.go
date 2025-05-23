package createproduct

import (
	"time"

	productmodel "github.com/Alexander-s-Digital-Marketplace/core-service/internal/models/product_model"
	profilemodel "github.com/Alexander-s-Digital-Marketplace/core-service/internal/models/profile_model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func CreateProduct(c *gin.Context) (int, string) {
	id, exists := c.Get("id")
	if !exists {
		return 400, "Bad request"
	}

	var product productmodel.Product
	var code int
	code = product.DecodeFromContext(c)
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
	product.SellerId = profile.Id
	product.PubDate = time.Now().Format("2006-01-02 15:04")
	product.IsBuy = false
	product.IsSellNow = false
	product.IsRated = false
	logrus.Infoln("product", product)
	code = product.AddToTable()
	if code != 200 {
		return code, "Error add to table"
	}

	return 200, "Success create product"
}
