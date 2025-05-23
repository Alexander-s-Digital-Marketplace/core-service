package switchproduct

import (
	productmodel "github.com/Alexander-s-Digital-Marketplace/core-service/internal/models/product_model"
	profilemodel "github.com/Alexander-s-Digital-Marketplace/core-service/internal/models/profile_model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func SwitchProduct(c *gin.Context) (int, string) {
	id, exists := c.Get("id")
	if !exists {
		return 400, "Error get id"
	}

	var product productmodel.Product
	var code int

	product.DecodeFromContext(c)

	profile := profilemodel.Profile{
		AccountId: int(id.(int)),
	}
	code = profile.GetFromTableByAccountId()
	if code != 200 {
		return code, ""
	}

	product.SellerId = profile.Id
	logrus.Infoln("product", product)
	logrus.Infoln("product.IsSellNow", product.IsSellNow)
	code = product.Switch(!product.IsSellNow)
	if code != 200 {
		return code, "Error publish/remove product"
	}
	logrus.Infoln("product", product)

	if product.IsSellNow {
		return 200, "Success publish product"
	} else {
		return 200, "Success remove publish product"
	}
}
