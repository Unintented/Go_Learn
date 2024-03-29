package main

import (
	"fmt"
	bean "github.com/Unintented/Go_learn/_10.1Bean"
)

func main() {
	/*包：包名与文件夹的名称可以不一致，但同一个文件夹下面的包名必须相同，不允许有不同包名，会报错，且同一个包下的文件不能放在不同的文件夹中*/

	/*可见性：只有首字母大写的标识符才能被其他包所引用，包括变量、常量、函数、结构体类型等*/
	const MODE = 1

	type person struct { // 首字母小写，外部包不可见，只能在当前包内使用
		name string
	}
	//	只有当结构体名与内部的字段名首字母都大写时，其他包才能访问结构体内部字段
	type Student struct {
		Name  string //可在包外访问的方法
		class string //仅限包内访问的字段
	}

	type Payer interface {
		init() //仅限包内访问的方法
		Pay()  //可在包外访问的方法
	}

	//测试可见性
	fmt.Printf("%f\n", bean.TestPackageVisit())

	// 包的引用方式：绝对路径（必须在GOPATH/src或GOROOT/scr路径下）、相对路径（只能引用GOPATH/src下的包）

	// 在同一个包中，go函数不支持重载

	// 函数可以作为参数类型

}

// 首字母大写，外部包可见，可在其他包中使用
func Add(x, y int) int {
	return x + y
}

func age() { // 首字母小写，外部包不可见，只能在当前包内使用
	var Age = 18 // 函数局部变量，外部包不可见，只能在当前函数内使用
	fmt.Println(Age)
}
