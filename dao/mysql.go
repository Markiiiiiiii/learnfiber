package dao

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitMysql() (err error) {

	dns := "root:qaz@78963@(34.85.112.152:3306)/todo?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		fmt.Println("mysql content faild, err:", err)
		return
	}

	return
}
