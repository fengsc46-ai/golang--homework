package bean

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	PostID      uint `gorm:"foreignKey:PostId"`
	UserID      uint `gorm:"foreignKey:UserId"`
	ContentText string
}
