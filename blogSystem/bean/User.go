package bean

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	username string
	password string
	email    string
	Posts    []Post
}

//DB写入初始化数据
//func InsertDB(db *gorm.DB) {
//	user := User{
//		Name: "John",
//		Age:  25,
//		Posts: []Post{
//			{Title: "My first post",
//				Content:  "This is my first post",
//				Comments: []Comment{{ContentText: "This is my first comment"}, {ContentText: "This is my second comment"}},
//			},
//			{Title: "My second post",
//				Content:  "This is my second post",
//				Comments: []Comment{{ContentText: "This is my second comment"}}},
//		},
//	}
//	db.Create(&user)
//
//}
