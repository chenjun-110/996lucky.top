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