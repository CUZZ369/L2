package dao

import (
	"database/sql"
	"fmt"
)

var DB *sql.DB
//链接数据库
func MysqlInit() *sql.DB  {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/users?charset=utf8mb4&parseTime=True")
	if err != nil {
		fmt.Printf("mysql connect failed:%v", err)
		return nil
	}
	db.SetMaxOpenConns(50) //setMaxOpenConns设置与数据库建立连接的最大数目
	err = db.Ping()//验证数据库用户密码
	if err != nil {
		fmt.Printf("mysql connect failed:%v", err)
		return nil
	}

	DB=db
	return DB
}
