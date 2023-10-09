package po

import "gorm.io/gorm"

type User struct {
	gorm.Model
	AuId     string `gorm:"column:auid"`
	UserName string `gorm:"column:username"`
}

func (r *User) TableName() string {
	return "users"
}
