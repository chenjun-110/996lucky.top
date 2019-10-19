json.Marshal的参数是Map

m1 := make(map[string]interface{})
err = json.Unmarshal(data, &m1)

map[k]v，map中所有的k类型必须相同，所有v的类型也必须相同

结构体字段名首字母必须大写，可能是json是跨包模块必须要大写权限才能解析。