package main

import "fmt"

func main() {
	/*函数定义：func funcName(paramName paraType, ...) ([returnName] returnType){}
	  返回值可只写类型，多个返回值必须用括号扩起来*/
	v1 := 3
	v2 := 4
	sumResult, productResult := sumAndProduct(v1, v2)
	fmt.Println(sumResult, ":", productResult)

	/*参数：
	  1、若相邻参数类型相同，可省略前面的参数类型
	  2、可变参数：要放在所有参数的最后面，通过切片实现，故可以使用对切片的操作*/
	v3 := sumPlus(10)
	v4 := sumPlus(10, 20)
	v5 := sumPlus(10, 20, 30, 40)
	fmt.Println(v3, v4, v5)

	/*返回值：*/
	//返回值含有名称时的用法
	fmt.Println(sumAndSub(5, 2))

	//函数可以作为参数，也可作为返回值

	/*匿名函数：没有函数名 func(paramName paramType) (returnType){}*/
	add := func(x, y int) {
		fmt.Println(x + y)
	}
	add(10, 20) // 通过变量调用匿名函数

	//自执行函数：匿名函数定义完加()直接执行
	func(x, y int) {
		fmt.Println(x + y)
	}(10, 20)

	/*defer语句：对其后的语句延迟处理，先被defer的语句越靠后
	  常用来处理资源释放问题*/
	fmt.Println("Before defer...")
	defer fmt.Println("First defer")
	defer fmt.Println("Second defer")
	defer fmt.Println("Third defer")
	fmt.Println("After defer")

	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())

	//defer面试题
	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y))
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20

	//panic函数与recover函数
	funcA()
	funcB()
	funcC()
}

func funcA() {
	fmt.Println("func A")
}

//recover一定要搭配defer函数使用，且defer需要在引发panic的语句之前定义
func funcB() {
	defer func() {
		err := recover()
		//如果程序出出现了panic错误,可以通过recover恢复过来
		if err != nil {
			fmt.Println("recover in B")
		}
	}()
	panic("panic in B")
}

func funcC() {
	fmt.Println("func C")
}

//**************************************
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

//****************************************
func sumAndSub(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

func sumPlus(x int, y ...int) int {
	fmt.Println(x, y)
	sum := x
	for _, v := range y {
		sum += v
	}
	return sum
}

func sumAndProduct(a, b int) (int, int) {
	return a + b, a * b
}
