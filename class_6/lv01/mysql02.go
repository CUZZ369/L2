package main
import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"

)
var db *sql.DB
type uh struct {
	id int
	user string
	password string
	num int
}

func queryRowDemo() {
	sqlStr := "select id, user, password,num from t_2 where id=?"
	var uh uh
	// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	err := db.QueryRow(sqlStr, 2).Scan(&uh.id, &uh.user, &uh.password,&uh.num)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	fmt.Printf("id:%d user:%s password:%s,num:%d\n", uh.id, uh.user, uh.password,uh.num)
}
func main() {
	err := initDB() // 调用输出化数据库的函数
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}
	go queryRowDemo()
	insertRowDemo()
	defer db.Close()
}
func insertRowDemo() {
	sqlStr := "insert into t_2(id,user,password,num,time , go ) values (?,?,?,?,?,?)"
	ret, err := db.Exec(sqlStr, 74,"afdsdgscs","fss",30,1,2)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
}
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