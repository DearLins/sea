package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"sea/handler"
	"sea/modles"
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

	//写入数据库，暂时不判断是否重复
	db := handler.Connect()
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
	//if err := c.ShouldBindBodyWith(&a, binding.JSON); err != nil { //多次绑定避免EOF
	if err := c.BindJSON(&info); err != nil {
		panic(err.Error())
	}
	//获取用户信息
	db := handler.Connect()
	stmt,err := db.Prepare("select id,username,password,phone from user where phone = ? limit 1")
	if err != nil{
		panic(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(info.Phone)
	if err != nil{
		panic(err)
	}
	var user modles.User
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Username, &user.Password, &user.Phone)
	}
	if err != nil{
		panic(err)
	}
	if user.Id == 0{
		panic(user.Id)
	}

	if user.Phone != info.Phone || user.Password != info.Password{
		panic("账号或密码错误！")
	}


	var token string
	token = handler.GenerateToken(&handler.UserClaims{
		Password:       info.Password,
		Phone:          info.Phone,
		StandardClaims: jwt.StandardClaims{},
	})

	intoCache(user.Id,token)
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": gin.H{
			"token": token,
		},
		"msg": "ok",
	})
}

//设置token的过期时间
func intoCache(id int,token string) {
	rdb := handler.InitClient()
	err := rdb.Set("token"+strconv.Itoa(id), token, 30*24*60*time.Second).Err()
	if err != nil {
		panic(err)
	}
}

//判断用户登录标识是否有效
func IsToken(c *gin.Context)  {

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

