package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

type stu struct {
	name string
	age int8
}
func main() {
	fmt.Printf("Hello world!\n")
	//使用同一个项目中的其他包中的函数（packageName.funcName()）
	logrus.Debug("123!")
	stringFunc()
	s1 := &stu{
		name: "zhangsan",
		age:  9,
	}
	fmt.Println(s1.name)

}

func stringFunc() {
	a := "Hello, 世界！"
	for i, v := range a { //遍历rune数组
		fmt.Println(i, v)
	}
	//for i:=0; i < len(a); i++{遍历字节数组
	//	fmt.Println(a[i])
	//}
}
