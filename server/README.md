版本管理： 
	git目录里go build和工作区里go install的区别是后者编译的执行文件会进入bin目录可被全局调用。
	打包先设置临时变量 export GOPATH=`pwd`

json.Marshal的参数是Map

m1 := make(map[string]interface{})
err = json.Unmarshal(data, &m1)

map[k]v，map中所有的k类型必须相同，所有v的类型也必须相同

结构体字段名首字母必须大写，可能是json是跨包模块必须要大写权限才能解析。

fmt.Printf("请求地址 %v \n", req.URL.RequestURI()) 
	%v 打印值
	%T 打印类型

未定义指针形参的函数都是拷贝副本（值传递）。
同一个目录下的多个文件包名要相同。
请求头map的key全被转换成了首字母大写


hex.EncodeToString(md.Sum(nil)) 用stringf转md5的编码有问题，要用和这个。 json包会对字符串作某种转码。
字符串编码为json字符串。角括号"<"和">"会转义为"\u003c"和"\u003e"以避免某些浏览器吧json输出错误理解为HTML。基于同样的原因，"&"转义为"\u0026"。
数组和切片类型的值编码为json数组，但[]byte编码为base64编码字符串，nil切片编码为null。