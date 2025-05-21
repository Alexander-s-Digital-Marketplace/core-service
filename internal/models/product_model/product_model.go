package productmodel

import (
	"math/rand"
	"time"

	cryproconfig "github.com/Alexander-s-Digital-Marketplace/core-service/internal/config/crypto"
	"github.com/Alexander-s-Digital-Marketplace/core-service/internal/database"
	itemmodel "github.com/Alexander-s-Digital-Marketplace/core-service/internal/models/item_model"
	profilemodel "github.com/Alexander-s-Digital-Marketplace/core-service/internal/models/profile_model"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Product struct {
	Id          int     `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string  `json:"title" gorm:"type:varchar(100)"`
	Description string  `json:"description" gorm:"type:text"`
	Price       float64 `json:"price" gorm:"type:numeric"`
	PubDate     string  `json:"pub_date" gorm:"type:varchar(30)"`
	IsBuy       bool    `json:"is_buy" gorm:"type:boolean"`
	IsSellNow   bool    `json:"is_sell_now" gorm:"type:boolean"`
	Image       string  `json:"image" gorm:"type:varchar(255)"`
	OrderId     int     `json:"order_id" gorm:"type:integer"`

	SellerId int `json:"seller_id" gorm:"not null"`
	ItemId   int `json:"item_id" gorm:"type:bigint"`

	Seller profilemodel.Profile `gorm:"foreignKey:SellerId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Item   itemmodel.Item       `gorm:"foreignKey:ItemId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (product *Product) DecodeFromContext(c *gin.Context) int {

	if err := c.ShouldBindJSON(&product); err != nil {
		logrus.Errorln("Error decode JSON: ", err)
		return 400
	}
	return 200
}

func (product *Product) AddToTable() int {
	var db database.DataBase
	db.InitDB()
	defer db.CloseDB()

	err := db.Connection.Create(&product).Error
	if err != nil {
		logrus.Errorln("Error add to table: ", err)
		return 503
	}
	return 200
}

func (product *Product) GetFromTable() int {
	var db database.DataBase
	db.InitDB()
	defer db.CloseDB()

	err := db.Connection.Preload("Seller").First(&product).Error
	if err != nil {
		logrus.Errorln("Error get from table: ", err)
		return 503
	}
	return 200
}

func (product *Product) GetAllFromTable() (int, []Product) {
	var db database.DataBase
	db.InitDB()
	defer db.CloseDB()

	var products []Product
	err := db.Connection.Preload("Seller").Find(&products).Error
	if err != nil {
		logrus.Errorln("Error get all from table: ", err)
		return 503, []Product{}
	}
	return 200, products
}

func (product *Product) GetAllMyFromTable() (int, []Product) {
	var db database.DataBase
	db.InitDB()
	defer db.CloseDB()

	var products []Product
	err := db.Connection.Where("seller_id = ?", product.SellerId).Find(&products).Error
	if err != nil {
		logrus.Errorln("Error get all my from table: ", err)
		return 503, []Product{}
	}
	return 200, products
}

func (product *Product) UpdateInTable() int {
	var db database.DataBase
	db.InitDB()
	defer db.CloseDB()

	err := db.Connection.Save(&product).Error
	if err != nil {
		logrus.Errorln("Error update in table: ", err)
		return 503
	}
	return 200
}

func (product *Product) Switch(action bool) int {
	var db database.DataBase
	db.InitDB()
	defer db.CloseDB()

	err := db.Connection.Model(&product).Where("id = ? AND seller_id = ?", product.Id, product.SellerId).Update("is_sell_now", action).Error
	if err != nil {
		logrus.Errorln("Error switch: ", err)
		return 503
	}
	return 200
}

func (product *Product) MigrateToDB(db database.DataBase) error {
	err := db.Connection.AutoMigrate(&Product{})
	if err != nil {
		logrus.Errorln("Error migrate Product model :", err)
		return err
	}
	logrus.Infoln("Success migrate Product model")
	return nil
}

func (prod *Product) Seeding(count int, countProf int) error {
	var db database.DataBase
	db.InitDB()
	defer db.CloseDB()

	for i := 0; i < count; i++ {
		gofakeit.Seed(time.Now().UnixNano())
		source := rand.NewSource(time.Now().UnixNano())
		r := rand.New(source)

		var item itemmodel.Item
		item.Content = gofakeit.SSN()
		item.Encode(cryproconfig.KEY)

		err := db.Connection.Create(&item).Error
		if err != nil {
			logrus.Errorln("Error add to table: ", err)
			logrus.Errorln("Error item: ", item)
			return err
		}

		var product Product
		product.Title = gofakeit.ProductName()
		product.Description = gofakeit.Sentence(100)
		product.Price = gofakeit.Price(0.00000001, 2.0)
		product.PubDate = gofakeit.Date().Format("2006-01-01 15:00")
		product.IsBuy = false
		product.IsSellNow = true
		product.Image = "image.com"
		product.SellerId = r.Intn(countProf) + 1
		product.ItemId = item.Id

		err = db.Connection.Create(&product).Error
		if err != nil {
			logrus.Errorln("Error add to table: ", err)
			logrus.Errorln("Error product: ", product)
			return err
		}
	}
	return nil
}
