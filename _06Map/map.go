package main

import . "fmt"

func main() {
	/*map定义：map[keyType]valueType，默认初始值为nil，需要
	  使用make函数分配内存：make(map[keyType]valueType, capacity)
	  初始化可以在后面用大括号，赋值键值对
	  简短声明时可直接使用make： mapName := make(map[keyType]valueType)
	  map是无序的，底层通过hash表实现
	  map的值可以是结构体，map也可存进切片中*/
	scoreMap := make(map[string]uint8, 10)
	scoreMap["上官特朗普"] = 99
	scoreMap["司徒拜登"] = 59
	Println(scoreMap)
	Println(scoreMap["司徒拜登"])

	userInfo := map[string]string{
		"username": "GoTrump",
		"password": "make America great again",
	}
	Println(userInfo)

	//判断某个键是否存在
	v, eleExist := scoreMap["司徒拜登"]
	if eleExist {
		Println("司徒拜登的分数：", v)
	} else {
		Println("查无此人")
	}

	//map遍历：使用for range，当只遍历key的时候:k := range mapName
	for k, v := range scoreMap {
		Println(k, "score:", v)
	}

	//使用delete(mapName, key)删除键值对
	delete(scoreMap, "司徒拜登")
	Println(scoreMap)

	//map切片
	var mapSlice = make([]map[string]string, 3)
	for index, value := range mapSlice {
		Printf("index:%d value:%v\n", index, value)
	}
	Println("after init")
	// 对切片中的map元素进行初始化
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "南海骗局"
	mapSlice[0]["password"] = "123456"
	mapSlice[0]["address"] = "中国南海"
	for index, value := range mapSlice {
		Printf("index:%d value:%v\n", index, value)
	}
}
