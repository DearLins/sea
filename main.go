package main

import (
	"github.com/gin-gonic/gin"
	"sea_mod/route"
)

func main() {
	//删除注释
	gin.SetMode(gin.DebugMode)
	route.Route()
}
