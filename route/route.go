package route

import (
	"github.com/gin-gonic/gin"
	"sea/conf"
	"sea/controllers"
	"sea/handler"
)

func Route() {
	route := gin.Default()
	route.Use(handler.Recover)
	{
		tokenOn := route.Group("/")
		tokenOn.Use(handler.TokenOn)
		{
			tokenOn.GET("/index", controllers.Index)
			tokenOn.GET("/redis", controllers.Redis)
			tokenOn.GET("/userinfo", controllers.UserInfo)
		}
		route.POST("/register", controllers.Register)
		route.POST("/login", controllers.Login)
	}
	route.Run(":" + conf.GetConfiguration().AppPort)
}
