package main

import (
	"github.com/gin-gonic/gin"
	"sea_mod/route"
)

func main() {
	//注释1122
	//测试bug分支
	gin.SetMode(gin.DebugMode)
	route.Route()
}
