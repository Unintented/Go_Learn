package main

import "fmt"

func main() {
	/*Go语言中的指针不能进行偏移和运算，是安全指针
	  地址变量还可以取地址；
	  函数可使用指针传值，直接对原始变量进行操作
	  指针变量在32位机器上占用4个字节，在64位机器上占用8个字节
	  指针可以进行多层嵌套*/
	a := 10
	b := &a
	c := *b
	fmt.Printf("value:%p, type:%T\n", b, b)
	fmt.Printf("value:%v, type:%T\n", c, c)

	//指针类型分配内存使用new函数：new(Type)，返回值为*Type
	var v1 *int
	v1 = new(int)
	*v1 = 10
	fmt.Println(*v1)
	v2 := new(string)
	*v2 = "See you around"
	fmt.Println(*v2)

}
