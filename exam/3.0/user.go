package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)
var DB *sql.DB

type user struct {
	Username string
	Password string
 }
func main()  {
	MysqlInit()
 	router :=gin.Default()
	router.POST("/register",Register)
	router.GET("/login",Login)
 	router.Run("8080")
}
func Register(ctx *gin.Context)  {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	 stmt , err := DB.Prepare("insert into user(username,password) values (?,?)")
	if err != nil {
		fmt.Printf("mysql prepare failed:%v", err)
		ctx.JSON(http.StatusOK, gin.H{
			"code":    400,
			"message": "failure",})
		return
	}
	_, err = stmt.Exec(username, password)
	if err != nil {
		fmt.Printf("insert failed:%v", err)
		ctx.JSON(http.StatusOK, gin.H{
			"code":    400,
			"message": "failure",})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
	})
}
func Login(ctx *gin.Context)  {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	stmt, err := DB.Query("select * from user where username = ?", username)
	if err!=nil{
		fmt.Printf("query failed:%v", err)
		ctx.JSON(http.StatusOK, gin.H{
			"code":    400,
			"message": "账号密码或密码错误",})
		return
	}
	defer stmt.Close()
	var u user
	for stmt.Next() {
		err = stmt.Scan(&u.Username, &u.Password)
		if err != nil {
			fmt.Printf("scan failed: %v", err)
			return
		}
		if u.Password == password {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "登录成功",
			})
		}


}
}
func MysqlInit() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/users?charset=utf8mb4&parseTime=True")
	if err != nil {
		fmt.Printf("mysql connect failed:%v", err)
		return nil
	}
	db.SetMaxOpenConns(1000)
	err = db.Ping()
	if err != nil {
		fmt.Printf("mysql connect failed:%v", err)
		return nil
	}
	DB = db
	return DB
}