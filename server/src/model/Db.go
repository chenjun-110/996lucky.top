package Db

import (
    "fmt"
	// "net/http"
	// "encoding/json"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)
var Db sql.DB
func init () {
	db, _ := sql.Open("mysql","root:@(47.106.93.213:3306)/996") // 设置连接数据库的参数
	Db = *db
	// Db = db
    // defer db.Close()    //关闭数据库
	err := db.Ping()      //连接数据库
	// rows, err := db.Query("SELECT * FROM user WHERE username=?", "abc")
    if err != nil {
        fmt.Println("数据库连接失败", err)
        return
	}
	fmt.Println("db already line")
}