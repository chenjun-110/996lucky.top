package view
import (
	"fmt"
	// "math/rand"
	"time"
	."model"
	// "strconv"
	"strings"
	"crypto/md5"
	"net/http"
	// "encoding/hex"
	"encoding/json"
	"encoding/base64"
)
var secret string = "123456789"

type UserInfo struct {
	Uid int  `json:"uid"`
	Exp int64  `json:"exp"`
	Username string  `json:"username"`
}

func Login (params Loginrams, outer *Outer) {
	fmt.Printf("username：%v \n", params.Username)
	fmt.Printf("password：%v \n", params.Password)
	//校验用户名
	rows, err := Db.Query("SELECT * FROM user WHERE username=?", params.Username)
	defer rows.Close()
	if err != nil {
		fmt.Println("查询失败", err)
		return
	}
	var (
		username string
		password string
		uid int
	)
	for rows.Next() {
		rows.Scan(&username, &password, &uid)
		fmt.Println("查询登录信息：", username, password, uid)
	}
	// fmt.Println("not: ", rows)
	if params.Username != username {
		outer.Code = "9999"
		outer.Success = false
		outer.Message = "账号不存在"
	} else if password != params.Password {
		//校验密码
		outer.Code = "9999"
		outer.Success = false
		outer.Message = "密码不正确"
	} else {
		outer.Message = "登录成功"
		payload := UserInfo{
			Uid: uid,
			Username: username,
			Exp: time.Now().Unix() + 60,
		}
		payloadStr, _ := json.Marshal(payload)
		payloadBase64 := base64.StdEncoding.EncodeToString(payloadStr)
		original := strings.Join([]string{payloadBase64, secret}, ".")
		md := md5.New()
		md.Write([]byte(original))
		sign := base64.StdEncoding.EncodeToString(md.Sum(nil))
		outer.Value = strings.Join([]string{payloadBase64, sign}, ".")
		fmt.Println("----------------------------------------------------------------- \n")
	}
}

// 检测token 过期 是否有效 是否伪造
func CheckToken(w http.ResponseWriter, req *http.Request) bool  {
	if (req.URL.RequestURI() == "/login") {
		token := req.Header["Token"]
		if (len(token) > 0 && token[0] != "") {
			str := strings.Split(token[0], ".")
			payloadBase64 := str[0]
			original := strings.Join([]string{payloadBase64, secret}, ".")
			md := md5.New()
			md.Write([]byte(original))
			sign := base64.StdEncoding.EncodeToString(md.Sum(nil))
			if sign == str[1] {
				fmt.Println("解析token: 相等")
				return checkExp(payloadBase64, w)
			} else {
				fmt.Println("解析token：不等")
				return false
			}
		}
		return true
	}
	return true
}
//检测过期时间
func checkExp (payloadBase64 string, w http.ResponseWriter) bool {
	var payload UserInfo
	payloadStr,_ := base64.StdEncoding.DecodeString(payloadBase64)
	json.Unmarshal([]byte(payloadStr), &payload)//字符串转struct
	if time.Now().Unix() > payload.Exp {
		fmt.Println("token过期时间 ", payload.Exp)
		outer := Outer{
			Code: "2005",
			Success: false,
			Message: "token过期",
		}
		jsonStr, _ := json.Marshal(outer)
		w.Write(jsonStr)
		return false
	}
	return true
}
/*
map保存session
	rand.Seed(time.Now().UnixNano())
	uidStr := strconv.FormatInt(int64(uid),10)
	sort := strconv.FormatInt(int64(rand.Intn(9999999)),10)
	info := strings.Join([]string{sort, uidStr}, ".")
	md := md5.New()
	md.Write([]byte(info))
	token := hex.EncodeToString(md.Sum(nil))
	outer.Value = token
	UserInfo[token] = User{
		Uid: uid,
		Exp: time.Now().Unix(),
	}
*/