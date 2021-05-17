package route

import (
	"github.com/gin-gonic/gin"
	"sea/conf"
	"sea/controllers"
	"sea/handler"
)

func Route()  {
	route := gin.Default()
	route.Use(handler.Recover)

	route.GET("/index", controllers.Index)
	route.GET("/redis", controllers.Redis)
	route.POST("/login", controllers.Login)
	route.GET("/userinfo", controllers.UserInfo)
	route.POST("/register", controllers.Register)
	route.Run(":"+conf.GetConfiguration().AppPort)
}


