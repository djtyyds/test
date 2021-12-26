package api

import "github.com/gin-gonic/gin"

func InitEngine() {
	en := gin.Default()
	en.POST("/register", register) //注册
	en.POST("login", login)        //登录
	userGroup := en.Group("/user")
	{
		userGroup.Use(auth)
		userGroup.POST("/addMoney", addMoney)   //充值
		userGroup.POST("/giveMoney", giveMoney) //转账
		userGroup.POST("/consume", Consume)     // 花费
	}
	queryGroup := en.Group("/query")
	{
		queryGroup.Use(auth)
		queryGroup.POST("/record", QueryRecord) //查询记录
	}
}
