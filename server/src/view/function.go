package view
// import (
// 	"fmt"
// 	"net/http"
// 	// "encoding/json"
// 	// ."model"
// 	"reflect"
// )
//检测token 过期 是否有效 是否伪造
// func CheckToken(w http.ResponseWriter, req *http.Request) bool  {
// 	if (req.URL.RequestURI() == "/login") {
// 		token := req.Header["Token"]
// 		if (len(token) > 0 && token[0] != "") {
// 			if _, ok := UserInfo[token[0]]; ok {
// 				fmt.Printf("服务端有指定token \n ")
// 			} else {
// 				fmt.Printf("服务端木有 \n ")
// 				for k, v := range UserInfo {
// 					fmt.Println(k, v)
// 					fmt.Printf("类型 %v \n ", reflect.TypeOf(k))
// 					if k == token[0] {
// 						fmt.Println("找到了！！！")
// 					} else {
// 						fmt.Println("找到个屁了1！！！", token[0])
// 						fmt.Println("找到个屁了2！！！", k)
// 					}
// 				}
// 			}
// 			fmt.Printf("token请求类型 %v \n ", reflect.TypeOf(token[0]))
// 		}
// 		// for k, v := range req.Header {  //Header 类型为map，通过range循环打印出来
// 		// 	fmt.Println(k+":", v) //+为输出顺序
// 		// }
// 	}
// 	return true
// }