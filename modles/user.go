package modles

type User struct {
	Id			 int	 `json:"id"`
	Username	 string	 `json:"username"`
	Phone 		 string  `json:"phone"`
	Password	 string  `json:"password"`
	CreateTime	 int64 	 `json:"create_time"`
	LoginTime	 int64   `json:"login_time"`
	Auth 		 int  	 `json:"auth"`
	Age 		 int  	 `json:"age"`
	Sex 		 int  	 `json:"sex"`
}

