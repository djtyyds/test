package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"test/model"
	"test/service"
	"test/tool"
	"time"
)

func addMoney(c *gin.Context) {
	IUsername, _ := c.Get("username")
	username := IUsername.(string)
	moneyString := c.PostForm("money")
	money, err := strconv.Atoi(moneyString)
	if err != nil {
		tool.RespErrorWithData(c, "请输入正确的金额")
	}
	err = service.AddMoney(username, money)
	if err != nil {
		fmt.Println("服务器出错")
		tool.RespInternalError(c)
	}
	tool.RespSuccessful(c)
	record := model.Record{
		Name:       username,
		Money:      +money,
		RecordTime: time.Now(),
		Txt:        "充值",
	}
	err = service.AddRecord(record)
	if err != nil {
		fmt.Println("服务器出错")
		tool.RespInternalError(c)
	}
}

func giveMoney(c *gin.Context) {
	IUsername1, _ := c.Get("username")
	username1 := IUsername1.(string)
	IUsername2, _ := c.Get("username")
	username2 := IUsername2.(string)
	moneyString := c.PostForm("money")
	money, err := strconv.Atoi(moneyString)
	if err != nil {
		tool.RespErrorWithData(c, "请输入正确的金额")
	}
	bool, err := service.GiveMoney(username1, username2, money)
	if bool == false {
		tool.RespErrorWithData(c, "输入金额大于本身金额")
	}
	if err != nil {
		fmt.Println("服务器出错")
		tool.RespInternalError(c)
	}
	tool.RespSuccessful(c)
	record1 := model.Record{
		Name:       username1,
		Money:      -money,
		RecordTime: time.Now(),
		Txt:        "转账",
	}
	err = service.AddRecord(record1)
	if err != nil {
		fmt.Println("服务器出错")
		tool.RespInternalError(c)
	}
	tool.RespSuccessful(c)
	record2 := model.Record{
		Name:       username2,
		Money:      +money,
		RecordTime: time.Now(),
		Txt:        "收钱",
	}
	err = service.AddRecord(record2)
	if err != nil {
		fmt.Println("服务器出错")
		tool.RespInternalError(c)
	}
	tool.RespSuccessful(c)
}

func Consume(c *gin.Context) {
	IUsername1, _ := c.Get("username")
	username1 := IUsername1.(string)
	Imoney, _ := c.Get("money")
	money := Imoney.(int)
	bool, err := service.Consume(username1, money)
	if bool == false {
		tool.RespErrorWithData(c, "余额不足")
	}
	if err != nil {
		tool.RespInternalError(c)
	}
	tool.RespSuccessful(c)
	record := model.Record{
		Name:       username1,
		Money:      -money,
		RecordTime: time.Now(),
		Txt:        "消费",
	}
	err = service.AddRecord(record)
	if err != nil {
		fmt.Println("服务器出错")
		tool.RespInternalError(c)
	}
	tool.RespSuccessful(c)
}
