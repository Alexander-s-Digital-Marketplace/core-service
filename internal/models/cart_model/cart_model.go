package cartmodel

import (
	"github.com/Alexander-s-Digital-Marketplace/core-service/internal/database"
	productmodel "github.com/Alexander-s-Digital-Marketplace/core-service/internal/models/product_model"
	profilemodel "github.com/Alexander-s-Digital-Marketplace/core-service/internal/models/profile_model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Cart struct {
	Id   int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Date string `json:"date" gorm:"type:varchar(30)"`

	ProfileId int `json:"profile_id" gorm:"not null"`
	ProductId int `json:"product_id" gorm:"not null"`

	Profile profilemodel.Profile `gorm:"foreignKey:ProfileId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Product productmodel.Product `gorm:"foreignKey:ProductId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (cart *Cart) DecodeFromContext(c *gin.Context) int {

	if err := c.ShouldBindJSON(&cart); err != nil {
		logrus.Errorln("Error decode JSON: ", err)
		return 400
	}
	return 200
}

func (cart *Cart) AddToTable() int {
	var db database.DataBase
	db.InitDB()
	defer db.CloseDB()

	err := db.Connection.Create(&cart).Error
	if err != nil {
		logrus.Errorln("Error add to table: ", err)
		return 503
	}
	return 200
}

func (cart *Cart) GetFromTable() int {
	var db database.DataBase
	db.InitDB()
	defer db.CloseDB()

	err := db.Connection.Preload("Seller").First(&cart).Error
	if err != nil {
		logrus.Errorln("Error get from table: ", err)
		return 503
	}
	return 200
}

func (cart *Cart) GetAllFromTableByProfileId() (int, []Cart) {
	var db database.DataBase
	db.InitDB()
	defer db.CloseDB()

	var carts []Cart
	err := db.Connection.Where("profile_id = ?", cart.ProfileId).Preload("Product").Preload("Product.Seller").Find(&carts).Error
	if err != nil {
		logrus.Errorln("Error get all from table: ", err)
		return 503, []Cart{}
	}
	return 200, carts
}

func (cart *Cart) UpdateInTable() int {
	var db database.DataBase
	db.InitDB()
	defer db.CloseDB()

	err := db.Connection.Save(&cart).Error
	if err != nil {
		logrus.Errorln("Error update in table: ", err)
		return 503
	}
	return 200
}

func (cart *Cart) MigrateToDB(db database.DataBase) error {
	err := db.Connection.AutoMigrate(&Cart{})
	if err != nil {
		logrus.Errorln("Error migrate Cart model :", err)
		return err
	}
	logrus.Infoln("Success migrate Cart model")
	return nil
}
