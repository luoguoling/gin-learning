package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"newBubble/models"
)

//首页
func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)

}

//新增数据
func CreateTodo(c *gin.Context) {
	var todo models.Todo
	c.BindJSON(&todo)
	err := models.CreateATodo(&todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 2000,
			"msg":  "success",
			"data": todo,
		})
	}
}

//查询数据
func GetTodoList(c *gin.Context) {
	todolist, err := models.GetAllTodo()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todolist)
	}

}

//删除数据
func DeleteATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"err": "无效id"})
		return
	}
	if err := models.DeleteATodo(id); err != nil {
		c.JSON(http.StatusOK, gin.H{"err": err.Error()})

	} else {
		c.JSON(http.StatusOK, gin.H{id: "delete"})
	}

}

//修改数据
func UpdateATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"err": "无效id"})
		return
	}
	todo, err := models.GetATodo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"err": err.Error()})
	}
	c.BindJSON(&todo)
	err = models.UpdateATodo(&models.Todo{})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"err": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			id: "deleted",
		})
	}

}
