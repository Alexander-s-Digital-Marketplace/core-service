package cartstruct

import (
	productmodel "github.com/Alexander-s-Digital-Marketplace/core-service/internal/models/product_model"
	profilemodel "github.com/Alexander-s-Digital-Marketplace/core-service/internal/models/profile_model"
)

type Cart struct {
	Id   int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Date string `json:"date" gorm:"type:varchar(30)"`

	ProfileId int `json:"profile_id" gorm:"not null"`
	ProductId int `json:"product_id" gorm:"not null"`

	Profile profilemodel.Profile `gorm:"foreignKey:ProfileId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Product productmodel.Product `gorm:"foreignKey:ProductId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
