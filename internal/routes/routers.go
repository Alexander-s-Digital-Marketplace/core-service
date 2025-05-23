package routespkg

import (
	"net/http"

	api "github.com/Alexander-s-Digital-Marketplace/core-service/internal/api"
	authmiddlewares "github.com/Alexander-s-Digital-Marketplace/core-service/internal/middlewares/auth_middlewares"
	corsmiddleware "github.com/Alexander-s-Digital-Marketplace/core-service/internal/middlewares/cors"
	"github.com/gin-gonic/gin"
)

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method string
	//ex None, Protected
	Type string
	// Pattern is the pattern of the URI.
	Pattern string
	// HandlerFunc is the handler function of this route.
	HandlerFunc gin.HandlerFunc
}

// NewRouter returns a new router.
func NewRouter(handleFunctions ApiHandleFunctions) *gin.Engine {
	return NewRouterWithGinEngine(gin.Default(), handleFunctions)
}

// NewRouter add routes to existing gin engine.
func NewRouterWithGinEngine(router *gin.Engine, handleFunctions ApiHandleFunctions) *gin.Engine {
	router.Use(corsmiddleware.CorsMiddleware())
	protected := router.Group("/Protected")
	protected.Use(authmiddlewares.AuthMiddleware())
	for _, route := range getRoutes(handleFunctions) {
		if route.HandlerFunc == nil {
			route.HandlerFunc = DefaultHandleFunc
		}
		switch route.Type {
		case "Protected":
			switch route.Method {
			case http.MethodGet:
				protected.GET(route.Pattern, route.HandlerFunc)
			case http.MethodPost:
				protected.POST(route.Pattern, route.HandlerFunc)
			case http.MethodPut:
				protected.PUT(route.Pattern, route.HandlerFunc)
			case http.MethodPatch:
				protected.PATCH(route.Pattern, route.HandlerFunc)
			case http.MethodDelete:
				protected.DELETE(route.Pattern, route.HandlerFunc)
			}
		default:
			switch route.Method {
			case http.MethodGet:
				router.GET(route.Pattern, route.HandlerFunc)
			case http.MethodPost:
				router.POST(route.Pattern, route.HandlerFunc)
			case http.MethodPut:
				router.PUT(route.Pattern, route.HandlerFunc)
			case http.MethodPatch:
				router.PATCH(route.Pattern, route.HandlerFunc)
			case http.MethodDelete:
				router.DELETE(route.Pattern, route.HandlerFunc)
			}
		}
	}

	return router
}

// Default handler for not yet implemented routes
func DefaultHandleFunc(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

type ApiHandleFunctions struct {

	// Routes for the DefaultAPI part of the API
	DefaultAPI api.DefaultAPI
}

func getRoutes(handleFunctions ApiHandleFunctions) []Route {
	return []Route{
		{
			"GetAllFeedGet",
			http.MethodGet,
			"Protected",
			"/GetAllFeed",
			handleFunctions.DefaultAPI.GetAllFeedGet,
		},
		{
			"GetCartGet",
			http.MethodGet,
			"Protected",
			"/GetCart",
			handleFunctions.DefaultAPI.GetCartGet,
		},
		{
			"GetHistoryGet",
			http.MethodGet,
			"Protected",
			"/GetHistory",
			handleFunctions.DefaultAPI.GetHistoryGet,
		},
		{
			"GetMyProductGet",
			http.MethodGet,
			"Protected",
			"/GetMyProduct",
			handleFunctions.DefaultAPI.GetMyProductGet,
		},
		{
			"GetMyProfileGet",
			http.MethodGet,
			"Protected",
			"/GetMyProfile",
			handleFunctions.DefaultAPI.GetMyProfileGet,
		},
		{
			"GetProfileByIdGet",
			http.MethodGet,
			"Protected",
			"/GetProfileById",
			handleFunctions.DefaultAPI.GetProfileByIdGet,
		},
		{
			"DeliverProductGet",
			http.MethodGet,
			"Protected",
			"/DeliverProduct",
			handleFunctions.DefaultAPI.DeliverProductGet,
		},
		{
			"CreateProductPost",
			http.MethodPost,
			"Protected",
			"/CreateProduct",
			handleFunctions.DefaultAPI.CreateProductPost,
		},
		{
			"SwitchProductPost",
			http.MethodPost,
			"Protected",
			"/SwitchProduct",
			handleFunctions.DefaultAPI.SwitchProductPost,
		},
		{
			"RemoveProductFromCartPost",
			http.MethodPost,
			"Protected",
			"/RemoveProductFromCart",
			handleFunctions.DefaultAPI.RemoveProductFromCartPost,
		},
		{
			"AddProductToCartPost",
			http.MethodPost,
			"Protected",
			"/AddProductToCart",
			handleFunctions.DefaultAPI.AddProductToCartPost,
		},
		{
			"UpdateProfilePost",
			http.MethodPost,
			"Protected",
			"/UpdateProfile",
			handleFunctions.DefaultAPI.UpdateProfilePost,
		},
		{
			"UploadProductImagePost",
			http.MethodPost,
			"Protected",
			"/UploadProductImage",
			handleFunctions.DefaultAPI.UploadProductImagePost,
		},
		{
			"BuyProductPost",
			http.MethodPost,
			"Protected",
			"/BuyProduct",
			handleFunctions.DefaultAPI.BuyProductPost,
		},
		{
			"RateProductPost",
			http.MethodPost,
			"Protected",
			"/RateProduct",
			handleFunctions.DefaultAPI.RateProductPost,
		},
		{
			"GetWalletGet",
			http.MethodGet,
			"Protected",
			"/GetWallet",
			handleFunctions.DefaultAPI.GetWalletGet,
		},
		{
			"GetBalanceGet",
			http.MethodGet,
			"Protected",
			"/GetBalance",
			handleFunctions.DefaultAPI.GetBalanceGet,
		},
		{
			"UpdateWalletPost",
			http.MethodPost,
			"Protected",
			"/UpdateWallet",
			handleFunctions.DefaultAPI.UpdateWalletPost,
		},
	}
}
