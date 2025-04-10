package api

import (
	createproduct "github.com/Alexander-s-Digital-Marketplace/core-service/internal/handlers/create_product"
	getallfeed "github.com/Alexander-s-Digital-Marketplace/core-service/internal/handlers/get_all_feed"
	getcart "github.com/Alexander-s-Digital-Marketplace/core-service/internal/handlers/get_cart"
	gethistory "github.com/Alexander-s-Digital-Marketplace/core-service/internal/handlers/get_history"
	getmyproduct "github.com/Alexander-s-Digital-Marketplace/core-service/internal/handlers/get_my_product"
	getmyprofile "github.com/Alexander-s-Digital-Marketplace/core-service/internal/handlers/get_my_profile"
	getprofilebyid "github.com/Alexander-s-Digital-Marketplace/core-service/internal/handlers/get_profile_by_id"
	switchproduct "github.com/Alexander-s-Digital-Marketplace/core-service/internal/handlers/switch_product"
	switchproductcart "github.com/Alexander-s-Digital-Marketplace/core-service/internal/handlers/switch_product_cart"
	updateprofile "github.com/Alexander-s-Digital-Marketplace/core-service/internal/handlers/update_profile"
	"github.com/gin-gonic/gin"
)

type DefaultAPI struct {
}

func (api *DefaultAPI) GetAllFeedGet(c *gin.Context) {

	code, allFeed := getallfeed.GetAllFeed(c)

	c.JSON(code, allFeed)
}

func (api *DefaultAPI) GetCartGet(c *gin.Context) {

	code, cart := getcart.GetCart(c)

	c.JSON(code, cart)
}

func (api *DefaultAPI) GetHistoryGet(c *gin.Context) {

	code, history := gethistory.GetHistory(c)

	c.JSON(code, history)
}

func (api *DefaultAPI) GetMyProductGet(c *gin.Context) {

	code, myProduct := getmyproduct.GetMyProduct(c)

	c.JSON(code, myProduct)
}

func (api *DefaultAPI) GetMyProfileGet(c *gin.Context) {

	code, myProfile := getmyprofile.GetMyProfile(c)

	c.JSON(code, myProfile)
}

func (api *DefaultAPI) GetProfileByIdGet(c *gin.Context) {

	code, profileById := getprofilebyid.GetProfileById(c)

	c.JSON(code, profileById)
}

func (api *DefaultAPI) CreateProductPost(c *gin.Context) {

	code, message := createproduct.CreateProduct(c)

	c.JSON(code, message)
}

func (api *DefaultAPI) SwitchProductPost(c *gin.Context) {

	code, message := switchproduct.SwitchProduct(c)

	c.JSON(code, message)
}

func (api *DefaultAPI) SwitchProductCartPost(c *gin.Context) {

	code, message := switchproductcart.SwitchProductCart(c)

	c.JSON(code, message)
}

func (api *DefaultAPI) UpdateProfilePost(c *gin.Context) {

	code, message := updateprofile.UpdateProfile(c)

	c.JSON(code, message)
}
