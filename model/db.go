package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	var err error
	// 连接数据库
	DB, err = gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database:\n" + err.Error())
	}
	// 迁移模型
	err = DB.AutoMigrate(&Char{}, &Page{}, &Artwork{})
	if err != nil {
		panic("failed to migrate schema:\n" + err.Error())
	}
}
