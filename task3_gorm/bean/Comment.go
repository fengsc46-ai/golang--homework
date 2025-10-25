package bean

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	PostID      uint `gorm:"foreignKey:PostId"`
	ContentText string
}

// BeforeDelete 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
func (comment *Comment) BeforeDelete(db *gorm.DB) (err error) {
	post := &Post{}
	db.Model(&Post{}).Where("id =?", comment.PostID).First(post)

	post.CommentNum -= 1
	if post.CommentNum == 0 {
		post.Status = false
	}

	itx := db.Save(&post)
	if itx.Error != nil {
		return itx.Error
	}
	return nil
}

// BeforeCreate 为 Comment 模型添加一个钩子函数，在评论创建时更新文章的评论数量。
func (comment *Comment) BeforeCreate(db *gorm.DB) (err error) {
	postId := comment.PostID
	post := &Post{}
	itx := db.Model(post).Where("id =?", postId).Update("comment_num", gorm.Expr("comment_num + ?", 1))

	if nil != itx.Error {
		return itx.Error
	}

	return nil
}
