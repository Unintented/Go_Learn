package main

import "fmt"

type Dog struct {
}
type Barker interface {
	Barking()
}

func main() {
	/*接口：接口区别于我们之前所有的具体类型，接口是一种抽象的类型。
	  接口做的事情就像是定义一个协议（规则），当你看到一个接口类型的值时，你不知道它是什么，唯一知道的是通过它的方法能做什么。
	  定义：一般接口名是要加er；参数列表和返回值列表中的参数变量名可以省略
		type InterfaceNamer interface{
			MethodName (ParameterList) returnParameterList
		}*/

	/*接口的实现：只要实现了接口中的所有方法，就实现了这个接口*/
	dog := Dog{}
	dog.Barking()

	/*实现接口后，就可以使用接口类型的变量来接收各实现该接口的类型的变量，同时可以调用其中的方法，类似于多态*/
	var b Barker
	b = dog
	b.Barking()

	/*值接收者与指针接收者实现接口的区别：
	  	值接收者可以接收值类型或指针类型，二者没有区别，因为指针类型的变量会因为内部语法糖的原因自动求值；
		而指针接收者只能接收指针类型，否则编译不通过*/
	dog_2 := &Dog{}
	b = dog_2
	b.Barking()

	/*一个类型可以实现多个接口，而一个接口可被多个类型实现*/

	/*空接口：指没有定义任何方法的接口，因此任何类型都实现了空接口
	  空接口类型的变量可以存储任何类型的变量；
	  空接口作为函数的参数可以接收任何类型的函数参数；
	  使用空接口可以存储任意值的字典*/
	var v interface{}
	v = "司马·拜登"
	fmt.Println(v)
	v = true
	fmt.Println(v)

	show(8)
	show(dog)
	show(b)

	var StuInfo = make(map[string]interface{})
	StuInfo["name"] = "little bush"
	StuInfo["age"] = 78
	fmt.Println(StuInfo)

	/*判断接口值：使用类型断言
	  x.(T):x为接口类型变量，T为要判断的类型
	  返回两个参数：第一个参数是x转换成T类型后的变量，第二个参数表示断言是否成功*/
	afterConversion, ok := v.(bool)
	if ok {
		fmt.Println("v is a bool type:", afterConversion)
	} else {
		fmt.Println("assertion failed")
	}

	//断言多次
	justifyType(true)
}
func justifyType(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string，value is %v\n", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	default:
		fmt.Println("unsupport type！")
	}
}

func show(p interface{}) {
	fmt.Printf("type:%T value:%v\n", p, p)
}

func (d Dog) Barking() {
	fmt.Println("汪～")
}
