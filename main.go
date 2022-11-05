package main

import (
	"github.com/gin-gonic/gin"
	"qny/models"
)

func main() {
	r := gin.Default()
	r.POST("/adduser", models.Createuser)
	r.GET("/addroom", models.Addroom)
	r.POST("/updateusername", models.Updateusername)
	r.POST("/createroom", models.Createroom)
	r.GET("/getroomlist", models.Getroomlist)
	r.Run(":8080")
}
