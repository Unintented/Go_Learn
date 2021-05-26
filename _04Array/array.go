package main

import "fmt"

func main() {
	/*数组长度固定且元素类型相同，其长度也属于数组类型的一部分，故长度不同的两数组也是不同类型
	 数组定义及初始化：
	 var listname [elenum]eletype
	 var listname = [elenum / ...]eletype{initial values}
	 listname := [elenum / ...]eletype{initial values}
	 vat listname = [elenum / ...]eletype{index: vaule, ...}
	注意：数组一般在创建时通过字面量初始化，否则初始化的都是元素的默认值，没有意义
	*/
	//listdemo := [3]uint8{2, 4, 8}
	//var listdemo = [...]uint8{2,4,8}
	var listdemo = [...]uint8{1: 2, 3: 16}
	fmt.Println(listdemo)

	/*数组遍历*/
	for i := 0; i < len(listdemo); i++ {
		println(listdemo[i])
	}
	for i, v := range listdemo {
		fmt.Printf("index:%d, value:%d\n", i, v)
	}
	//二维数组
	a := [...][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s\t", v2)
		}
		fmt.Println()
	}
	/*数组作为参数传递时以引用形式传递，可以提高执行效率
	  而以值传递时，是值拷贝，浪费空间不说，还无法修改原始数组内容*/
	arrDemo := &[...]uint8{1: 2}
	changeArray(arrDemo)
	println(arrDemo)
}
func changeArray(arr *[2]uint8) {
	(*arr)[1] = 3
}
