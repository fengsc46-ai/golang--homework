package initial

import (
	"blogSystem/bean"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDbConnection() {
	dsn := "root:123456@tcp(localhost:3306)/mydatabase?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	CreateTable(DB)
	setDB(DB)
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

func CreateTable(db *gorm.DB) {
	err := db.AutoMigrate(&bean.User{}, &bean.Post{}, &bean.Comment{})
	if err != nil {
		panic(err)
	}
}
