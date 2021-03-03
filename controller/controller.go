package controller

import (
	"fh123.co/fiber/models"
	"github.com/gofiber/fiber/v2"
)

func IndexHandler(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{})
}

func CreateATodo(c *fiber.Ctx) error {
	var todo models.Todo
	c.BodyParser(&todo)
	if err := models.CreateATodoInfo(&todo); err != nil {
		return c.JSON(fiber.Map{
			"code": 2001,
			"msg":  "add a todo message faild",
		})
	} else {
		return c.JSON(todo)
	}
}

func GetAllTodo(c *fiber.Ctx) error {
	if todolist, err := models.GetAllTodoInfos(); err != nil {
		return c.JSON(fiber.Map{
			"code": 2002,
			"msg":  "don't get todo list",
		})
	} else {
		return c.JSON(todolist)
	}

}

func EditATodo(c *fiber.Ctx) error {
	id := c.Params("id")
	// var todo Todo
	todo, err := models.SearchATodoByID(id)
	if err != nil {
		return c.JSON(fiber.Map{
			"code": 2003,
			"msg":  "don't search todo message by id ",
		})
	}
	c.BodyParser(&todo)
	if err = models.UpdateATodo(todo); err != nil {
		return c.JSON(fiber.Map{
			"code": 2004,
			"msg":  "don't update todo message by id",
		})
	} else {
		return c.JSON(todo)
	}
}

func DeleteATodo(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := models.DeleteATodoInfo(id); err == nil {
		return c.JSON(fiber.Map{
			"code": 2000,
			"msg":  "delete todo massage success ",
		})
	} else {
		return err
	}
}
