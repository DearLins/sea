package main

import (
	"github.com/gin-gonic/gin"
	"sea/controllers"
)

func main()  {
	route := gin.Default()
	route.GET("/index", controllers.Index)
	route.GET("/redis", controllers.Redis)
	route.Run(":8090")

}
