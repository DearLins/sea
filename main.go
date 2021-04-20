package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"sea/conf"
	"sea/controllers"
)
var config *conf.Config
func main()  {
	config, err := conf.ParseConfig("conf/config.json")
	fmt.Println(config.AppHost)
	fmt.Println(err)
	route := gin.Default()
	route.GET("/index", controllers.Index)
	route.GET("/redis", controllers.Redis)
	route.Run(":8090")

}
