package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

//func main() {
//	dsn := "user:password@tcp(127.0.0.1:3306)/class_6"
//	db, err := sql.Open("mysql", dsn)
//	if err != nil {
//		panic(err)
//	}
//	defer db.Close()
//}
// 定义一个全局对象db

// 定义一个初始化数据库的函数
func initDB() (err error) {
	// DSN:Data Source Name
	dsn := "root:root@tcp(127.0.0.1:3306)/class_6?charset=utf8mb4&parseTime=True"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	//校验账号密码
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := initDB() // 调用输出化数据库的函数
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}
}