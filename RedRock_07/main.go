package main

import (
	"RedRock_01/cmd"
	"RedRock_01/dao"
	_ "github.com/go-sql-driver/mysql"
)

func main()  {
	dao.MysqlInit()//链接数据库
	cmd.Entrance()
}
