package main

import (
	_ "github.com/go-sql-driver/mysql"
	"test/api"
	"test/dao"
)

func main() {
	dao.InitDB()
	api.InitEngine()
}
