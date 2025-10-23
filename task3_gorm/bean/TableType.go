package bean

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string
	Age   int
	Posts []Post
}

type Post struct {
	gorm.Model
	Title    string
	Content  string
	UserID   uint `gorm:"foreignKey:UserId"`
	Comments []Comment
}

type Comment struct {
	gorm.Model
	PostID      uint `gorm:"foreignKey:PostId"`
	ContentText string
}

func CreateTable(db *gorm.DB) {
	err := db.AutoMigrate(&User{}, &Post{}, &Comment{})
	if err != nil {
		panic(err)
	}
}

// DB写入初始化数据
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
