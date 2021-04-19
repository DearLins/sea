package controllers

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"github.com/go-redis/redis"
)


type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func Index(c *gin.Context) {
	db, err := sql.Open("mysql", "root:honglin1@tcp(47.94.136.121:3306)/gotest")
	fmt.Println(db)
	defer db.Close()
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	stmt, err1 := db.Prepare("select * from user where id = ?")
	defer stmt.Close()
	if err1 != nil {
		panic(err1)
	}
	//panic(err)
	res := stmt.QueryRow(1)
	var user User
	err1 = res.Scan(&user.Id, &user.Name, &user.Age)

	c.JSON(200, gin.H{
		"user": user,
	})

}


func Redis(c *gin.Context) {
	initClient()
	redisExample()
	// 初始化连接
	var user int
	c.JSON(200, gin.H{
		"user": user,
	})

}

var rdb *redis.Client

func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "47.94.136.121:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func redisExample() {
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

