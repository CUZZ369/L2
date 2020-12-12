package models

import (
	"RedRock_01/dao"
	"fmt"
)

type users struct {
	Username string`json:"username"`
	Password string`json:"password"`
}
//写入数据库
func Register(username string, password string) bool{
	stmt, err := dao.DB.Prepare("insert into user(username,password) values (?,?)")
	if err != nil {
		fmt.Printf("mysql prepare failed:%v", err)
		return false
	}
	defer stmt.Close()
	_, err = stmt.Exec(username, password)
	if err != nil {
		fmt.Printf("insert failed:%v", err)
		return false
	}
	return true
}
//与数据库进行验证
func Login(username string, password string) bool {
	stmt, err := dao.DB.Query("select * from user where username = ?", username)
	if err != nil {
		fmt.Printf("query failed:%v", err)
		return false
	}
	defer stmt.Close()
	var u users
	for stmt.Next() {
		err = stmt.Scan(&u.Username, &u.Password)
		if err != nil {
			fmt.Printf("scan failed: %v", err)
			return false
		}
	}
	if u.Password == password {
		return true
	}
	return false
}
