package bean

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name    string
	Age     int
	PostNum int
	Posts   []Post
}

// DB写入初始化数据
func InsertDB(db *gorm.DB) {
	user := User{
		Name:    "John",
		Age:     25,
		PostNum: 2,
		Posts: []Post{
			{Title: "My first post",
				Content:    "This is my first post",
				CommentNum: 2,
				Status:     true,
				Comments:   []Comment{{ContentText: "This is my first comment"}, {ContentText: "This is my second comment"}},
			},
			{Title: "My second post",
				Content:    "This is my second post",
				CommentNum: 1,
				Status:     true,
				Comments:   []Comment{{ContentText: "This is my second comment"}}},
		},
	}
	db.Create(&user)

}
