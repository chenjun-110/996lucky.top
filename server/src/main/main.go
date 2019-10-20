package main

import (
	."model/Db"
    "fmt"
	"net/http"
	"encoding/json"
	// "database/sql"
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

// var Db sql.DB
func (h *MyHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	outer := Outer{
		Code: "0000",
		Message: "success",
		Describe: "",
	}
	fmt.Printf("请求地址 %v \n", req.URL.RequestURI())
	switch req.Method {
		case "GET":
			fmt.Printf("GET")
		case "POST":
			
			buf := make([]byte, 1024)
			n, _ := req.Body.Read(buf)//json转字符串
			var loginrams Loginrams
			json.Unmarshal(buf[0:n], &loginrams)//字符串转struct
			fmt.Printf("用户名：%v", loginrams.Username)
			fmt.Printf("密码：%v", loginrams.Password)
			fmt.Printf("ssss %T", Db)
			
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
	

    handler := MyHandler{}
    server := http.Server{
        Addr:    "127.0.0.1:8080",
        Handler: &handler,  //处理器没有路由功能，所有请求都是同样的回应
    }
    server.ListenAndServe()  //开启web服务
}