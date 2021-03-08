package main

import . "fmt"

//定义全局变量
var v1 int

var v2 string = "Kevin"
var v3 = "曹操"

func main() {
	v1 = 18
	//函数内定义变量的简便形式
	v4 := 5i
	Println("age:", v1)
	Println("English name:", v2)
	Println("Chinese name:", v3)
	Println("Complex num:", v4)
	sum := sumAndProduct(2, 3)
	Print(sum)
}
func sumAndProduct(a int, b int) int {
	c := a + b
	//d := a * b
	return c
}
