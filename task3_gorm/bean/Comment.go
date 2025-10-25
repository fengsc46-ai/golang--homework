package bean

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	PostID      uint `gorm:"foreignKey:PostId"`
	ContentText string
}

// AfterDelete 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。/*
func (comment *Comment) AfterDelete(db *gorm.DB) {
	post := &Post{}
	db.Model(&Post{}).Where("id =?", comment.PostID).First(post)

	if post.CommentNum == 1 {
		post.Status = false
	}
	post.CommentNum--
	db.Save(&post)
}
