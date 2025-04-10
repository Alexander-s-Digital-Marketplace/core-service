package main

import (
	loggerconfig "github.com/Alexander-s-Digital-Marketplace/core-service/internal/config/logger"
	"github.com/Alexander-s-Digital-Marketplace/core-service/internal/database"
	cartmodel "github.com/Alexander-s-Digital-Marketplace/core-service/internal/models/cart_model"
	historymodel "github.com/Alexander-s-Digital-Marketplace/core-service/internal/models/history_model"
	productmodel "github.com/Alexander-s-Digital-Marketplace/core-service/internal/models/product_model"
	profilemodel "github.com/Alexander-s-Digital-Marketplace/core-service/internal/models/profile_model"
)

func main() {
	loggerconfig.Init()

	var db database.DataBase
	db.InitDB()
	defer db.CloseDB()

	var profile profilemodel.Profile
	profile.MigrateToDB(db)
	profile.Seeding(5)

	var product productmodel.Product
	product.MigrateToDB(db)
	product.Seeding(10, 5)

	var cart cartmodel.Cart
	cart.MigrateToDB(db)

	var history historymodel.History
	history.MigrateToDB(db)

}
