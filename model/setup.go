package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	//setup ke db local mu cuy
	dsn := "root:pass@tcp(127.0.0.1:3306)/go_jwt_mux?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())

	}

	db.AutoMigrate(&User{})

	DB = db
}
