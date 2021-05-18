package handler

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

//用户信息类，作为生成token的参数
type UserClaims struct {
	Id       int    `json:"id"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	//jwt-go提供的标准claim
	jwt.StandardClaims
}

var (
	//自定义的token秘钥
	secret = []byte("123123")
	//该路由下不校验token
	noVerify = []interface{}{"/login", "/ping"}
	//token有效时间（纳秒） 这里因为使用了redis做处理，所以过期时间设置为无限大
	effectTime = 10 * 365 * 24 * 60 * 60 * time.Second
)

// 生成token
func GenerateToken(claims *UserClaims) string {
	//设置token有效期，也可不设置有效期，采用redis的方式
	//   1)将token存储在redis中，设置过期时间，token如没过期，则自动刷新redis过期时间，
	//   2)通过这种方式，可以很方便的为token续期，而且也可以实现长时间不登录的话，强制登录
	//本例只是简单采用 设置token有效期的方式，只是提供了刷新token的方法，并没有做续期处理的逻辑
	claims.ExpiresAt = time.Now().Add(effectTime).Unix()
	//生成token
	sign, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secret)
	if err != nil {
		//这里因为项目接入了统一异常处理，所以使用panic并不会使程序终止，如不接入，可使用原始方式处理错误
		//接入统一异常可参考 https://blog.csdn.net/u014155085/article/details/106733391
		panic(err)
	}
	return sign
}

//验证token
func JwtVerify(c *gin.Context) {
	//过滤是否验证token
	//if utils.IsContainArr(noVerify, c.Request.RequestURI) {
	//	return
	//}
	token := c.GetHeader("token")
	if token == "" {
		panic("token not exist !")
	}
	//验证token，并存储在请求中
	user := parseToken(token)
	c.Set("user", user)
	c.Set("id", user.Id)
	c.Set("phone", user.Phone)
	c.Set("token", token)
}

//设置token的过期时间
func IntoCache(id int, token string) {
	rdb := InitClient()
	err := rdb.Set("token"+strconv.Itoa(id), token, 30*24*60*60*time.Second).Err()
	if err != nil {
		panic(err)
	}
}

//中间件验证用户登录状态
func TokenOn(c *gin.Context) {
	token := c.GetHeader("token")
	if token == "" {
		panic("token not exist !")
	}
	user := parseToken(token)
	rdb := InitClient()
	expire, err := rdb.TTL("token" + strconv.Itoa(user.Id)).Result()
	//fmt.Println(expire)
	if err != nil {
		panic(err)
	}
	if expire < 0 {
		panic("Token Exists")
	}
	IntoCache(user.Id, token)

}

// 解析Token
func parseToken(tokenString string) *UserClaims {
	//解析token
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		panic(err)
	}
	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		panic("token is valid")
	}
	return claims
}

// 更新token
func Refresh(tokenString string) string {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		panic(err)
	}
	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		panic("token is valid")
	}
	jwt.TimeFunc = time.Now
	claims.StandardClaims.ExpiresAt = time.Now().Add(2 * time.Hour).Unix()
	return GenerateToken(claims)
}

