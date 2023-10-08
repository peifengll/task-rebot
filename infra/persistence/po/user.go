package po

import "gorm.io/gorm"

type User struct {
	gorm.Model
	AuId     string
	UserName string
}

func (r *User) TableName() string {
	return "users"
}
