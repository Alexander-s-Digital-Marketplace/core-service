package historystruct

import (
	itemmodel "github.com/Alexander-s-Digital-Marketplace/core-service/internal/models/item_model"
	productmodel "github.com/Alexander-s-Digital-Marketplace/core-service/internal/models/product_model"
	profilemodel "github.com/Alexander-s-Digital-Marketplace/core-service/internal/models/profile_model"
)

type History struct {
	Id   int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Date string `json:"date" gorm:"type:varchar(30)"`

	BuyerId   int `json:"buyer_id" gorm:"type:bigint"`
	ProductId int `json:"product_id" gorm:"type:bigint;not null"`
	ItemId    int `json:"item_id" gorm:"type:bigint;not null"`

	Buyer   profilemodel.Profile `gorm:"foreignKey:BuyerId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Product productmodel.Product `gorm:"foreignKey:ProductId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Item    itemmodel.Item       `gorm:"foreignKey:ItemId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
