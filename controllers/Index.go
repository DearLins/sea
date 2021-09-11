package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"sea_mod/handler"
	"sea_mod/modles"
)

func Index(c *gin.Context) {
	db := handler.Connect()
	defer db.Close()
	var user modles.User
	stmt, err := db.Prepare("select id,username,phone,password from user where id =?")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(1)
	//fmt.Printf("%v \n%+v \n%#v \n", rows,rows,rows)
	if err != nil{
		panic(err)
	}
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Username,&user.Phone,&user.Password)
	}
	c.JSON(200, gin.H{
		"code": 0,
		"msg": "ok",
		"data": gin.H{
			"id":user.Id,
			"phone":user.Phone,
			"username":user.Username,
			"password":user.Password,
		},
	})
}


func Redis(c *gin.Context) {
	rdb := handler.InitClient()
	redisExample(rdb)
	// 初始化连接
	var user int
	c.JSON(200, gin.H{
		"user": user,
	})

}

func redisExample(rdb *redis.Client) {
	err := rdb.Set("score", 100, 0).Err()
	if err != nil {
		fmt.Printf("set score failed, err:%v\n", err)
		return
	}

	val, err := rdb.Get("score").Result()
	if err != nil {
		fmt.Printf("get score failed, err:%v\n", err)
		return
	}
	fmt.Println("score", val)

	val2, err := rdb.Get("name").Result()
	if err == redis.Nil {
		fmt.Println("name does not exist")
	} else if err != nil {
		fmt.Printf("get name failed, err:%v\n", err)
		return
	} else {
		fmt.Println("name", val2)
	}
}


