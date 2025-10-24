package main

import (
	"fmt"
	"task3_gorm/bean"
	"task3_gorm/database"
)

func main() {
	// create a database connection
	db := database.CreateDb()
	//insert data into the database
	//bean.InsertDB(db)
	var user bean.User
	//使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
	tx := db.Preload("Posts.Comments").Find(&user, 1)
	if tx.Error != nil {
		panic(tx.Error)
	}
	fmt.Printf("%v\n", user)

	var comment bean.Comment

	//查询评论数量最多的文章信息。
	subQuery := db.Model(&comment).Debug().
		Select("post_id").
		Group("post_id").
		Order("COUNT(post_id) DESC").
		Limit(1)

	var post bean.Post
	tx1 := db.Model(&post).First(&post, subQuery)
	if tx1.Error != nil {
		panic(tx1.Error)
	}
	fmt.Printf("%v\n", post)
	//print the result
	//fmt.Println(user)
	// close the database connection
	//defer func(sqlDB *gorm.DB) {
	//	err := sqlDB.()
	//	if err != nil {
	//		panic(err)
	//	}
	//}(db)
}
