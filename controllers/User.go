package controllers

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"sea/handler"
)

//type User struct {
//	Id			 int
//	UserName	 string
//	Password	 string
//	CreateTime	 int
//	LoginTime	 int
//	Auth 		 int
//	Age 		 int
//	Sex 		 int
//}

func Login(c *gin.Context)  {
	var token 		string
	var phone 		string
	var password	string
	phone,err = c.Get("user")
	if err != nil{

	}
	password,err = c.Get("password")
	token = handler.GenerateToken(&handler.UserClaims{
		Password:       password,
		Phone:          phone,
		StandardClaims: jwt.StandardClaims{},
	})
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
func UserInfo(c *gin.Context) {
	handler.JwtVerify(c)
	fmt.Println(c.Get("user"))
}

