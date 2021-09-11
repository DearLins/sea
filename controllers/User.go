package controllers

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"sea_mod/handler"
	"sea_mod/modles"
	"strconv"
	"time"
)

type LoginInfo struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

//用户注册
func Register(c *gin.Context) {
	var user modles.User
	if err := c.BindJSON(&user); err != nil {
		panic(err.Error())
	}
	user.CreateTime = time.Now().Unix()
	user.LoginTime = time.Now().Unix()
	user.Auth = 1

	db := handler.Connect()
	//判断用户是否注册
	que, err := db.Prepare("select id from user where phone = ?")
	if err != nil {
		panic(err)
	}
	defer que.Close()
	res := que.QueryRow(user.Phone)
	var id int
	err = res.Scan(&id) // TODO 这里查询结果为空记录时暂时不知道怎么处理
	//if err != nil {
	//	panic(err)
	//}
	fmt.Println(id)
	if id > 0 {
		panic("手机号码已注册！")
	}

	stmt, err := db.Prepare("insert into `user`(username,phone,password,create_time,login_time,auth,age,sex) values(?,?,?,?,?,?,?,?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(user.Username, user.Phone, user.Password, user.CreateTime, user.LoginTime, user.Auth, user.Age, user.Sex)
	if err != nil {
		panic(err)
	}

	rows, _ := result.RowsAffected()
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": gin.H{
			"affect": rows,
		},
		"msg": "ok",
	})
	db.Close()
}

//用户登录
func Login(c *gin.Context) {
	var info LoginInfo
	//if err := c.ShouldBindBodyWith(&info, binding.JSON); err != nil { //多次绑定避免EOF
	if err := c.BindJSON(&info); err != nil {
		panic(err.Error())
	}
	//获取用户信息
	db := handler.Connect()
	stmt, err := db.Prepare("select id,username,password,phone from user where phone = ? limit 1")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(info.Phone)
	if err != nil {
		panic(err)
	}
	var user modles.User
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Username, &user.Password, &user.Phone)
	}
	if err != nil {
		panic(err)
	}
	if user.Id == 0 {
		panic("账号不存在!")
	}

	if user.Phone != info.Phone || user.Password != info.Password {
		panic("账号或密码错误！")
	}

	var token string
	token = handler.GenerateToken(&handler.UserClaims{
		Id:             user.Id,
		Password:       info.Password,
		Phone:          info.Phone,
		StandardClaims: jwt.StandardClaims{},
	})

	handler.IntoCache(user.Id, token)
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": gin.H{
			"token": token,
		},
		"msg": "ok",
	})
}

//判断用户登录标识是否有效
func IsToken(c *gin.Context) {

}

//用户推出
func Logout(c *gin.Context) {
	handler.JwtVerify(c)
	id,err := c.Get("id")
	if !err {
		panic("获取用户信息失败")
	}
	//fmt.Println("type:", reflect.TypeOf(id))
	rdb := handler.InitClient()
	res := rdb.Del("token" +strconv.Itoa(id.(int)))
	if res != nil{
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "ok",
		})
	}
}

func UserInfo(c *gin.Context) {
	handler.JwtVerify(c)
	user, ok := c.Get("user") //取值 实现了跨中间件取值
	if !ok {
		user = "user is error"
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": user,
	})
}

