package models

import "fh123.co/fiber/dao"

type Todo struct {
	ID     int    `json:"id"`
	Tiltle string `json:"title"`
	Status bool   `json:"status"`
}

func CreateATodoInfo(todo *Todo) (err error) {
	err = dao.DB.Create(&todo).Error
	return
}

func GetAllTodoInfos() (todo []*Todo, err error) {
	if err := dao.DB.Find(&todo).Error; err != nil {
		return nil, err
	}
	return
}

func SearchATodoByID(id string) (todo *Todo, err error) {
	todo = new(Todo)
	if err := dao.DB.Where("id=?", id).First(&todo).Error; err != nil {
		return nil, err
	}
	return
}

func UpdateATodo(todo *Todo) (err error) {
	err = dao.DB.Save(todo).Error
	return
}

func DeleteATodoInfo(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(Todo{}).Error
	return
}
