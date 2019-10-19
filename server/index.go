package main

import (
    "fmt"
	"net/http"
	"encoding/json"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Outer struct {
	Code string
    Message string
    Describe string
    Value Loginrams
}
type Loginrams struct {
	Username string `json:"username"` //标记为小写
	Password string `json:"password"`
}
type MyHandler struct{}

var Db sql.DB
func (h *MyHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	outer := Outer{
		Code: "0000",
		Message: "success",
		Describe: "",
	}
	fmt.Printf(req.URL.RequestURI())
	if req.URL.RequestURI() == "/" {
		fmt.Printf("这是/")
	}
	switch req.Method {
		case "GET":
			fmt.Printf("GET")
		case "POST":
			
			buf := make([]byte, 1024)
			n, _ := req.Body.Read(buf)//json转字符串
			var loginrams Loginrams
			json.Unmarshal(buf[0:n], &loginrams)//字符串转struct
			fmt.Println(loginrams.Username)
			fmt.Printf("type:%T\n", loginrams.Username)
			fmt.Println("ssss ", Db)
			fmt.Println("还是空不：： ", Db.Query)
			
			//校验用户名
			rows, err := Db.Query("SELECT * FROM user WHERE username=?", loginrams.Username)
			if err != nil {
				fmt.Println("err", err)
				return
			}
			var (
				username string
				password string
			)
			for rows.Next() {
				rows.Scan(&username, &password)
				fmt.Println("查询登录信息：", username, password)
			}
			rows.Close()
			// fmt.Println("not: ", rows)
			if loginrams.Username != username {
				outer.Code = "9999"
				outer.Message = "fail"
				outer.Describe = "账号不存在"
			} else if password != loginrams.Password {
				//校验密码
				outer.Code = "9999"
				outer.Message = "fail"
				outer.Describe = "密码不正确"
			} else {
				outer.Describe = "登录成功"
				outer.Value.Username = "222"
				outer.Value.Password = "4444"
			}
			jsonStr, _ := json.Marshal(outer)
			w.Write(jsonStr)
		case "OPTIONS":
			fmt.Printf("OPTIONS")
	}
}

func main() {
	db,_:=sql.Open("mysql","root:@(47.106.93.213:3306)/996") // 设置连接数据库的参数
	Db = *db
	// Db = db
    defer db.Close()    //关闭数据库
	err:=db.Ping()      //连接数据库
	// rows, err := db.Query("SELECT * FROM user WHERE username=?", "abc")
    if err!=nil{
        fmt.Println("数据库连接失败", err)
        return
	}

    handler := MyHandler{}
    server := http.Server{
        Addr:    "127.0.0.1:8080",
        Handler: &handler,  //处理器没有路由功能，所有请求都是同样的回应
    }
    server.ListenAndServe()  //开启web服务
}