package database

import (
	"task3_gorm/bean"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func CreateDb() *gorm.DB {
	dsn := "root:123456@tcp(localhost:3306)/mydatabase?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	bean.CreateTable(db)
	setDB(db)
	return db
}

func setDB(db *gorm.DB) {
	sqlDB, err := db.DB()

	if err != nil {
		panic(err)
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了可以重新使用连接的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
}
