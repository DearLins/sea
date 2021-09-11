package main

import (
	"github.com/gin-gonic/gin"
	"sea_mod/route"
)

func main() {
	//注释1122
	//测试bug
	//main冲突
	gin.SetMode(gin.DebugMode)
	route.Route()
}

