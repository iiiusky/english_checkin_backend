package model

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	OpenID      string `gorm:"column:_openid;type:varchar(255);not null;comment:User's OpenID from WeChat." json:"open_id"`
	Name        string `gorm:"column:name;type:varchar(255);comment:User's nickname." json:"name"`
	AvatarURL   string `gorm:"column:avatar_url;type:varchar(1023);comment:URL of the user's avatar." json:"avatar_url"`
	PhoneNumber string `gorm:"column:phone_number;type:varchar(50);comment:User's phone number." json:"phone_number"`
	ExpiryDate  int64  `gorm:"column:expiry_date;type:bigint;comment:Membership expiry date as milliseconds since epoch." json:"expiry_date"`
	Status      string `gorm:"column:status;type:varchar(50);comment:Membership status (e.g., 'active', 'expired')." json:"status"`
}

func (m *Student) TableName() string {
	return "students"
}
