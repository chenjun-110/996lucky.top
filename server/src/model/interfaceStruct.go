package Db
import (
	"fmt"
)
type Outer struct {
	Code string `json:"code"`
    Success bool `json:"success"`
    Message string `json:"message"`
    Value interface{} `json:"value"`
}
type Loginrams struct {
	Username string `json:"username"` //标记为小写
	Password string `json:"password"`
	Uid int
}
func init()  {
	fmt.Printf("interfaceStruct init \n")
}