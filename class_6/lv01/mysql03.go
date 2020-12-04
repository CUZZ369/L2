package main

import (
	"database/sql"
	"fmt"
)

func main ()  {
	DB, _ := sql.Open("mysql", "root:root@tcp(IP:127.0.0.1:3306)/schemas")
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil {
		fmt.Println("open database fail")
		return
	}
	fmt.Println("connnect success")
}
