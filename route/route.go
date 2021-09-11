package route

import (
	"github.com/gin-gonic/gin"
	"sea_mod/conf"
	"sea_mod/controllers"
	"sea_mod/handler"
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
			tokenOn.GET("/logout", controllers.Logout)
		}
		route.POST("/register", controllers.Register)
		route.POST("/login", controllers.Login)
	}
	route.Run(":" + conf.GetConfiguration().AppPort)
}
