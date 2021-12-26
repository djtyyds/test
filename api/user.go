package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"test/model"
	"test/service"
	"test/tool"
)

func register(c *gin.Context) {
	username := c.PostForm("name")
	password := c.PostForm("password")
	user := model.User{
		Name:     username,
		Password: password,
	}
	flag, err := service.IsRepeatName(username)
	if err != nil {
		fmt.Println("judge repeat username err: ", err)
		tool.RespInternalError(c)
		return
	}
	if flag {
		tool.RespErrorWithData(c, "该用户名已存在")
		return
	}
	bool := service.IsPasswordPlausible(user)
	if !bool {
		tool.RespErrorWithData(c, "密码必须大于6位数")
		return
	}
	err = service.Register(user)
	if err != nil {
		fmt.Println("register err")
		tool.RespInternalError(c)
		return
	}

}

func login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	flag, err := service.IsPasswordRight(username, password)
	if err != nil {
		fmt.Println("judge password err:", err)
		tool.RespInternalError(c)
		return
	}
	if !flag {
		tool.RespErrorWithData(c, "密码错误")
		return
	}
	c.SetCookie("username", username, 600, "/", "", false, false)
	tool.RespSuccessful(c)
}
