package main

import (
	"fmt"
	"github.com/Unintented/Go_learn/basic"
	"github.com/sirupsen/logrus"
)

func main() {
	fmt.Printf("Hello world!\n")
	//使用同一个项目中的其他包中的函数（packageName.funcName()）
	basic.PrintFunc()
	logrus.Debug("123!")
	stringFunc()
}

func stringFunc() {
	a := "Hello, 世界！"
	for i, v := range a{//遍历rune数组
		fmt.Println(i, v)
	}
	//for i:=0; i < len(a); i++{遍历字节数组
	//	fmt.Println(a[i])
	//}
}