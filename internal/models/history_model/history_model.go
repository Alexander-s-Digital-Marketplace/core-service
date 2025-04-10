package historymodel

import (
	"github.com/Alexander-s-Digital-Marketplace/core-service/internal/database"
	productmodel "github.com/Alexander-s-Digital-Marketplace/core-service/internal/models/product_model"
	profilemodel "github.com/Alexander-s-Digital-Marketplace/core-service/internal/models/profile_model"
	"github.com/sirupsen/logrus"
)

type History struct {
	Id   int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Date string `json:"date" gorm:"type:varchar(30)"`

	BuyerId   int `json:"buyer_id" gorm:"type:bigint"`
	ProductId int `json:"product_id" gorm:"type:bigint;not null"`

	Buyer   profilemodel.Profile `gorm:"foreignKey:BuyerId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Product productmodel.Product `gorm:"foreignKey:ProductId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (history *History) AddToTable() int {
	var db database.DataBase
	db.InitDB()
	defer db.CloseDB()

	err := db.Connection.Create(&history).Error
	if err != nil {
		logrus.Errorln("Error add to table: ", err)
		return 503
	}
	return 200
}

func (history *History) GetFromTable() int {
	var db database.DataBase
	db.InitDB()
	defer db.CloseDB()

	err := db.Connection.
		Preload("Buyer").
		Preload("Product").
		First(&history).Error
	if err != nil {
		logrus.Errorln("Error get from table: ", err)
		return 503
	}
	return 200
}

func (history *History) GetAllFromTable() (int, []History) {
	var db database.DataBase
	db.InitDB()
	defer db.CloseDB()

	var histories []History
	err := db.Connection.
		Preload("Buyer").
		Preload("Product").
		Find(&history).Error
	if err != nil {
		logrus.Errorln("Error get all from table: ", err)
		return 503, []History{}
	}
	return 200, histories
}

func (history *History) MigrateToDB(db database.DataBase) error {
	err := db.Connection.AutoMigrate(&History{})
	if err != nil {
		logrus.Errorln("Error migrate History model :", err)
		return err
	}
	logrus.Infoln("Success migrate History model")
	return nil
}
