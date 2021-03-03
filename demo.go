package main

import (
	"fh123.co/fiber/dao"
	"fh123.co/fiber/models"
	"fh123.co/fiber/routers"
)

func main() {
	err := dao.InitMysql()
	if err != nil {
		panic(err)
	}
	dao.DB.AutoMigrate(&models.Todo{})
	//使用fiber加载
	app := routers.SetupRouter()
	app.Listen(":3000")
}
