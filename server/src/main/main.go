package main

import (
	."model"
	"view"
    "fmt"
	"net/http"
	"encoding/json"
)

type MyHandler struct{}

// var Db sql.DB
func (h *MyHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	// w.Header().Set("Access-Control-Expose-Headers", "*")
	outer := Outer{
		Code: "0000",
		Success: true,
	}
	fmt.Printf("请求地址 %v \n", req.URL.RequestURI())
	boolean := view.CheckToken(w, req)
	if (boolean == false) {
		return
	}
	switch req.Method {
		case "GET":
			fmt.Printf("GET")
		case "POST":
			var params Loginrams
			buf := make([]byte, 1024)
			n, _ := req.Body.Read(buf)//json转字符串
			json.Unmarshal(buf[0:n], &params)//字符串转struct
			fmt.Printf("db %T \n", Db)
			if (req.URL.RequestURI() == "/login") {
				view.Login(params, &outer)
				jsonStr, _ := json.Marshal(outer)
				w.Write(jsonStr)
			}
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