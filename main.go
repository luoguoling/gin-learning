package main

import (
	//"github.com/gin-gonic/gin"
	"newBubble/dao"
	"newBubble/models"
	"newBubble/routers"
)

func main() {
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer dao.Close()
	//模型绑定
	dao.DB.AutoMigrate(&models.Todo{})
	r := routers.SetupRouter()
	r.Run(":9000")
}
