package rateproduct

import (
	productmodel "github.com/Alexander-s-Digital-Marketplace/core-service/internal/models/product_model"
	profilemodel "github.com/Alexander-s-Digital-Marketplace/core-service/internal/models/profile_model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Rate struct {
	ProductId int `json:"product_id"`
	Rate      int `json:"rate"`
}

func RateProduct(c *gin.Context) (int, string) {
	_, exists := c.Get("id")
	if !exists {
		return 400, "Bad request"
	}
	var rate Rate
	if err := c.ShouldBindJSON(&rate); err != nil {
		logrus.Errorln("Error decode JSON: ", err)
		return 400, "Error decode JSON"
	}

	product := productmodel.Product{
		Id: rate.ProductId,
	}
	code := product.GetFromTable()
	if code != 200 {
		return 404, "Error find product"
	}

	profile := profilemodel.Profile{
		Id: product.SellerId,
	}
	code = profile.GetFromTable()
	if code != 200 {
		return 404, "Error find profile"
	}

	logrus.Infoln("profile", profile)
	logrus.Infoln("product", product.Title)
	logrus.Infoln("OldRating", profile.Rating)
	profile.Rating = (profile.Rating*float32(profile.CountRating) + float32(rate.Rate)) / (float32(profile.CountRating) + 1)
	profile.CountRating += 1
	logrus.Infoln("NewRating", profile.Rating)

	code = profile.UpdateInTable()
	if code != 200 {
		return 404, "Error set rating"
	}
	product.IsRated = true
	code = product.UpdateInTable()
	if code != 200 {
		return 404, "Error update product"
	}
	return 200, "Error set rating"
}
