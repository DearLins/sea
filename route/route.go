package route

import (
	"github.com/gin-gonic/gin"
	"sea/conf"
	"sea/controllers"
)

func Route()  {
	route := gin.Default()
	route.GET("/index", controllers.Index)
	route.GET("/redis", controllers.Redis)
	route.Run(":"+conf.GetConfiguration().AppPort)
}

