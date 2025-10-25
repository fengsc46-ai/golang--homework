package bean

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title      string
	Content    string
	UserID     uint `gorm:"foreignKey:UserId"`
	CommentNum int
	Status     bool
	Comments   []Comment
}

// BeforeCreate 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。/**
func (post *Post) BeforeCreate(db *gorm.DB) (err error) {
	userId := post.UserID
	user := &User{}
	itx := db.Model(user).Where("id =?", userId).Update("post_num", gorm.Expr("post_num + ?", 1))

	if nil != itx.Error {
		return itx.Error
	}

	return nil
}
