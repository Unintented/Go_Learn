package main

import . "fmt"

//定义全局变量
var v1 int

var v2 string = "Kevin"
var v3 = "曹操"

//变量就是内存中的一块存储空间

//var v int32  自动赋予类型初值
//var v int64 = 10
//var v1, v2 = "lotus",10
/*
var(
v1 string
v2 bool
v3 interface{}
v4 []int
)
*/
func main() {
	v1 = 18
	//函数内定义变量的简便形式，也可以一次赋值多个变量
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
