package routers

import (
	"fh123.co/fiber/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func SetupRouter() *fiber.App {
	engine := html.New("./templattes", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/static", "./static")
	app.Get("/", controller.IndexHandler)
	// 注册一个路由组
	v1 := app.Group("/v1")
	// 添加
	v1.Post("/todo", controller.CreateATodo)
	// 查看
	v1.Get("/todo", controller.GetAllTodo)
	// 修改
	v1.Put("/todo/:id", controller.EditATodo)

	//删除
	v1.Delete("/todo/:id", controller.DeleteATodo)

	return app
}
