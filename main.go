package main

import (
	"github.com/gin-gonic/gin"
	"sea_mod/route"
)

func main() {
	//删除注释
	//bug
	gin.SetMode(gin.DebugMode)
	route.Route()
}
