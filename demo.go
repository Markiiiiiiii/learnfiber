package main

import (
	"fmt"

	"github.com/gofiber/template/html"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

type Todo struct {
	ID     int    `json:"id"`
	Tiltle string `json:"title"`
	Status bool   `json:"status"`
}

func initMysql() (err error) {

	dns := "root:qaz@78963@(34.85.112.152:3306)/todo?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		fmt.Println("mysql content faild, err:", err)
		return
	}
	return
}

func main() {
	err := initMysql()
	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(&Todo{})
	//使用fiber加载
	engine := html.New("./templattes", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/static", "./static")
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})
	// 注册一个路由组
	v1 := app.Group("/v1")
	// 添加
	v1.Post("/todo", func(c *fiber.Ctx) error {
		var todo Todo
		c.BodyParser(&todo)
		if err = DB.Create(&todo).Error; err != nil {
			return c.JSON(fiber.Map{
				"code": 2001,
				"msg":  "add a todo message faild",
			})
		} else {
			return c.JSON(todo)
		}

	})
	// 查看
	v1.Get("/todo", func(c *fiber.Ctx) error {
		var todolist []Todo
		if err = DB.Find(&todolist).Error; err != nil {
			return c.JSON(fiber.Map{
				"code": 2002,
				"msg":  "don't get todo list",
			})
		} else {
			return c.JSON(todolist)
		}

	})
	// 修改
	v1.Put("/todo/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var todo Todo
		if err = DB.Where("id=?", id).First(&todo).Error; err != nil {
			return c.JSON(fiber.Map{
				"code": 2003,
				"msg":  "don't search todo message by id ",
			})
		}
		c.BodyParser(&todo)
		if err = DB.Save(&todo).Error; err != nil {
			return c.JSON(fiber.Map{
				"code": 2004,
				"msg":  "don't update todo message by id",
			})
		} else {
			return c.JSON(todo)
		}
	})

	//删除
	v1.Delete("/todo/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		if err = DB.Where("id=?", id).Delete(Todo{}).Error; err != nil {
			return c.JSON(fiber.Map{
				"code": 2003,
				"msg":  "don't search todo message by id",
			})
		} else {
			return c.JSON(fiber.Map{
				"code": 2000,
				"msg":  "delete todo massage success ",
			})
		}
	})
	app.Listen(":3000")
}
