package profilemodel

import (
	"math/rand"
	"time"

	"github.com/Alexander-s-Digital-Marketplace/core-service/internal/database"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Profile struct {
	Id          int     `json:"id" gorm:"primaryKey;autoIncrement"`
	UserName    string  `json:"user_name" gorm:"type:varchar(20)"`
	Rating      float32 `json:"rating" gorm:"type:real"`
	CountRating int     `json:"count_rating" gorm:"type:int"`

	AccountId int `json:"account_id" gorm:"type:bigint"`
	WalletId  int `json:"wallet_id" gorm:"type:bigint"`
}

func (profile *Profile) DecodeFromContext(c *gin.Context) int {

	if err := c.ShouldBindJSON(&profile); err != nil {
		logrus.Errorln("Error decode JSON: ", err)
		return 400
	}
	return 200
}

func (profile *Profile) AddToTable() int {
	var db database.DataBase
	db.InitDB()
	defer db.CloseDB()

	err := db.Connection.Create(&profile).Error
	if err != nil {
		logrus.Errorln("Error add to table: ", err)
		return 503
	}
	return 200
}

func (profile *Profile) GetFromTable() int {
	var db database.DataBase
	db.InitDB()
	defer db.CloseDB()

	err := db.Connection.First(&profile).Error
	if err != nil {
		logrus.Errorln("Error get from table: ", err)
		return 503
	}
	return 200
}

func (profile *Profile) GetFromTableByWalletId() int {
	var db database.DataBase
	db.InitDB()
	defer db.CloseDB()

	err := db.Connection.Where("wallet_id = ?", profile.WalletId).First(&profile).Error
	if err != nil {
		logrus.Errorln("Error get from table: ", err)
		return 503
	}
	return 200
}

func (profile *Profile) UpdateInTable() int {
	var db database.DataBase
	db.InitDB()
	defer db.CloseDB()

	err := db.Connection.Save(&profile).Error
	if err != nil {
		logrus.Errorln("Error update in table: ", err)
		return 503
	}
	return 200
}

func (profile *Profile) MigrateToDB(db database.DataBase) error {
	err := db.Connection.AutoMigrate(&Profile{})
	if err != nil {
		logrus.Errorln("Error migrate Profile model :", err)
		return err
	}
	logrus.Infoln("Success migrate Profile model")
	return nil
}

func (prof *Profile) Seeding(count int) error {
	var db database.DataBase
	db.InitDB()
	defer db.CloseDB()

	for i := 0; i < count; i++ {
		gofakeit.Seed(time.Now().UnixNano())
		source := rand.NewSource(time.Now().UnixNano())
		r := rand.New(source)

		var profile Profile
		profile.UserName = gofakeit.Username()
		profile.Rating = r.Float32() * 10
		profile.CountRating = r.Intn(10) + 5

		profile.AccountId = 1
		profile.WalletId = 1

		err := db.Connection.Create(&profile).Error
		if err != nil {
			logrus.Errorln("Error add to table: ", err)
			logrus.Errorln("Error profile: ", profile)
			return err
		}
	}
	return nil
}
