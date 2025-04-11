package deliverproduct

import (
	"strconv"

	cryproconfig "github.com/Alexander-s-Digital-Marketplace/core-service/internal/config/crypto"
	itemmodel "github.com/Alexander-s-Digital-Marketplace/core-service/internal/models/item_model"
	productmodel "github.com/Alexander-s-Digital-Marketplace/core-service/internal/models/product_model"
	notificationservice "github.com/Alexander-s-Digital-Marketplace/core-service/internal/services/notification_service/notification_service_client"
	"github.com/gin-gonic/gin"
)

func DeliverProduct(c *gin.Context) (int, string) {
	id, exists := c.Get("id")
	if !exists {
		return 400, "Bad request"
	}
	productId := c.Query("product_id")
	if id == "" {
		return 400, "Bad request"
	}

	var product productmodel.Product
	var item itemmodel.Item
	var code int
	var err error
	product.Id, err = strconv.Atoi(productId)
	if err != nil {
		return 400, "Bad request"
	}
	code = product.GetFromTable()
	if code != 200 {
		return code, "Error get from table"
	}
	code = item.GetFromTable()
	if code != 200 {
		return code, "Error get from table"
	}
	item.Decode(cryproconfig.KEY)
	product.Item = item

	var message string
	code, message = notificationservice.DeliverNotif(product, "pm12.kusakin@gmail.com")
	if code != 200 {
		return code, message
	}

	return 200, message
}
