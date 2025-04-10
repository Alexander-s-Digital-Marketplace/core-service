package profilemodel

type Profile struct {
	Id          int     `json:"id" gorm:"primaryKey;autoIncrement"`
	UserName    string  `json:"user_name" gorm:"type:varchar(20)"`
	Rating      float32 `json:"rating" gorm:"type:real"`
	CountRating int     `json:"count_rating" gorm:"type:int"`

	AccountId int `json:"account_id" gorm:"type:bigint"`
	WalletId  int `json:"wallet_id" gorm:"type:bigint"`
}
