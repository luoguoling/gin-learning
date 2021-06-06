package models

import "newBubble/dao"

//定义模型
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

//定义对模型的操作
//新增数据
func CreateATodo(todo *Todo) (err error) {
	err = dao.DB.Create(&todo).Error
	return err
}

//查询所有数据
func GetAllTodo() (todolist []*Todo, err error) {
	if err = dao.DB.Find(&todolist).Error; err != nil {
		return nil, err
	}
	return

}

//查询某一个id的数据
func GetATodo(id string) (todo *Todo, err error) {
	if err = dao.DB.Where("id=?", id).Find(&todo).Error; err != nil {
		return nil, err
	}
	return

}

//更新数据
func UpdateATodo(todo *Todo) (err error) {
	err = dao.DB.Save(&todo).Error
	return

}

//删除数据
func DeleteATodo(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&Todo{}).Error
	return

}
