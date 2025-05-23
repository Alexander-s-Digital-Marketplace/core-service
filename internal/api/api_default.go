package api

import (
	addproducttocart "github.com/Alexander-s-Digital-Marketplace/core-service/internal/handlers/add_product_to_cart"
	buyproduct "github.com/Alexander-s-Digital-Marketplace/core-service/internal/handlers/buy_product"
	createproduct "github.com/Alexander-s-Digital-Marketplace/core-service/internal/handlers/create_product"
	deliverproduct "github.com/Alexander-s-Digital-Marketplace/core-service/internal/handlers/deliver_product"
	getallfeed "github.com/Alexander-s-Digital-Marketplace/core-service/internal/handlers/get_all_feed"
	getbalance "github.com/Alexander-s-Digital-Marketplace/core-service/internal/handlers/get_balance"
	getcart "github.com/Alexander-s-Digital-Marketplace/core-service/internal/handlers/get_cart"
	gethistory "github.com/Alexander-s-Digital-Marketplace/core-service/internal/handlers/get_history"
	getmyproduct "github.com/Alexander-s-Digital-Marketplace/core-service/internal/handlers/get_my_product"
	getmyprofile "github.com/Alexander-s-Digital-Marketplace/core-service/internal/handlers/get_my_profile"
	getprofilebyid "github.com/Alexander-s-Digital-Marketplace/core-service/internal/handlers/get_profile_by_id"
	getwallet "github.com/Alexander-s-Digital-Marketplace/core-service/internal/handlers/get_wallet"
	rateproduct "github.com/Alexander-s-Digital-Marketplace/core-service/internal/handlers/rate_product"
	removeproductfromcart "github.com/Alexander-s-Digital-Marketplace/core-service/internal/handlers/remove_product_from_cart"
	switchproduct "github.com/Alexander-s-Digital-Marketplace/core-service/internal/handlers/switch_product"
	updateprofile "github.com/Alexander-s-Digital-Marketplace/core-service/internal/handlers/update_profile"
	updatewallet "github.com/Alexander-s-Digital-Marketplace/core-service/internal/handlers/update_wallet"
	uploadproductimage "github.com/Alexander-s-Digital-Marketplace/core-service/internal/handlers/upload_product_image"
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

func (api *DefaultAPI) AddProductToCartPost(c *gin.Context) {

	code, message := addproducttocart.AddProductToCart(c)

	c.JSON(code, message)
}

func (api *DefaultAPI) RemoveProductFromCartPost(c *gin.Context) {

	code, message := removeproductfromcart.RemoveProductFromCart(c)

	c.JSON(code, message)
}

func (api *DefaultAPI) UpdateProfilePost(c *gin.Context) {

	code, message := updateprofile.UpdateProfile(c)

	c.JSON(code, message)
}

func (api *DefaultAPI) DeliverProductGet(c *gin.Context) {

	code, message := deliverproduct.DeliverProduct(c)

	c.JSON(code, message)
}

func (api *DefaultAPI) UploadProductImagePost(c *gin.Context) {

	code, message := uploadproductimage.UploadProductImage(c)

	c.JSON(code, gin.H{"url": message})
}

func (api *DefaultAPI) BuyProductPost(c *gin.Context) {

	code, contract := buyproduct.BuyProduct(c)

	c.JSON(code, contract)
}

func (api *DefaultAPI) RateProductPost(c *gin.Context) {

	code, message := rateproduct.RateProduct(c)

	c.JSON(code, message)
}

func (api *DefaultAPI) GetWalletGet(c *gin.Context) {

	code, wallet := getwallet.GetWallet(c)

	c.JSON(code, wallet)
}

func (api *DefaultAPI) GetBalanceGet(c *gin.Context) {

	code, wallet := getbalance.GetBalance(c)

	c.JSON(code, wallet)
}

func (api *DefaultAPI) UpdateWalletPost(c *gin.Context) {

	code, message := updatewallet.UpdateWallet(c)

	c.JSON(code, message)
}
