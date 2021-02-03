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
}
