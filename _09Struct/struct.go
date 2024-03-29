package main

import (
	"encoding/json"
	"fmt"
)

/*结构体：Go通过结构体实现面向对象
  结构体也是一种类型，类似于其他语言中的class，需要实例化才会分配内存并使用，可以用var来声明
  ！要注意结构体声明的位置，要保证接收者等能够访问到*/

/*
自定义类型
结构体是值类型，直接使用等号赋值后是两个不同的变量，地址不相同，一般放在main函数外或其他包中
两大种使用方法（指针、非指针），四小种使用方法
两种结构体在字段完全相同的情况下可以进行强制类型转换
结构体可以作为结构体字段的类型
内存分布：结构体中内存地址是连续的
*/
/*结构体类型中的方法与接收者
  个人理解就是结构体中的函数无需在结构体中进行定义，当然任何类型都可以使用接收者来添加方法
  如果需要通过方法来修改结构体中的值，可以设定指针类型的接收者；此后，为了保证一致性，需要其他的方法也要使用指针类型接收者*/

//接受者可以是值类型（原数据拷贝），可以是指针类型；调用者可以是值类型，也可以是指针类型。
//此外，接受者的名称一般设定为this或self

/*方法与函数的区别：
1、定义方式不同，有无接收者
2、调用方式不同，方法需要通过调用者调用，函数直接写函数名*/
type person struct {
	name, work string
	age        int8
}

type Person_2 struct {
	name   string
	age    int8
	dreams []string
}

func main() {
	/*自定义类型与类型别名*/
	//自定义类型，是一种新的类型，具有int特性，二者之间不能直接赋值，可以通过强制类型转换
	type NewInt int
	//类型别名与原类型是一种
	type MyInt = int

	var a NewInt
	var b MyInt
	fmt.Printf("type of a: %T\n", a)
	fmt.Printf("type of b: %T\n", b)
	//var p1 = person{...} 直接赋值
	var p1 person
	fmt.Printf("%+v\n", p1)
	p1.name = "上官·特朗普"
	p1.work = "mess the world up"
	p1.age = 73
	fmt.Printf("%+v\n", p1)

	/*指针类型结构体：可以使用new进行初始化得到的就是指向该结构体类型的指针；同样的，直接使用&来对结构体进行初始化得到的也是指针
	  默认初始化得到的都是对应类型的零值
	  需要注意的是，Go语言中支持对结构体指针直接使用"."来访问结构体成员，因此不论是结构体变量，还是指向结构体的指针变量，访问结构体成员的方式是一样的*/
	// var p2 *person，可以先声明类型，再分配内存；也可直接分配内存，让编译器自动识别类型
	var p2 = new(person)
	fmt.Printf("%+v\n", p2)
	(*p2).name = "司马·拜登"
	(*p2).work = "mess China up"
	p2.age = 78
	fmt.Printf("%+v\n", *p2)

	// var p3 *person = &person{}
	var p3 = &person{}
	fmt.Printf("%+v\n", p3)
	p3.name = "诸葛·彭斯"
	p3.work = "assist 上官·特朗普"
	p3.age = 63
	fmt.Printf("%+v\n", p3)

	//可使用键值对进行初始化
	var p4 = person{
		name: "独孤·奥巴马",
		work: "doing speech",
		age:  62,
	}
	fmt.Printf("%+v\n", p4)

	/*interview question*/
	type student struct {
		name string
		age  int
	}

	m := make(map[string]*student)
	stus := []student{
		{name: "诸葛亮", age: 18},
		{name: "司马懿", age: 23},
		{name: "周瑜", age: 9000},
	}

	for _, stu := range stus {
		//可不可以理解为&stu都是stus数组的头指针，其在遍历过程中还使用额外的一个偏移量
		fmt.Printf("%T\n", stu)
		fmt.Printf("%p\n", &stu)
		fmt.Println((&stu).name)
		fmt.Println(stu)
		m[stu.name] = &stu
	}
	for k, v := range m {
		fmt.Printf("%p\n", v)
		fmt.Println(k, "=>", v.name)
	}

	/*构造函数：go语言结构体没有构造函数，需要自己实现*/
	var p5 = newPerson("罗斯福", "beat Japan", 66)
	fmt.Printf("%+v\n", p5)
	p5.Dream()
	p5.SetAge(18)
	fmt.Printf("%+v\n", p5)

	/*结构体中的匿名字段：在声明时只有类型，没有字段名，默认采用类型名作为字段名，因此同一类型的匿名字段只能有一个*/
	type Worker struct {
		string
		int8
	}
	var s1 = Worker{
		"Lily",
		22,
	}
	fmt.Printf("%+v\n", s1)
	fmt.Printf("name:%s, age:%d\n", s1.string, s1.int8)

	/*嵌套结构体*/
	type Address struct {
		Province string
		City     string
	}
	type User struct {
		Username string
		Sex      bool
		Address  Address
	}
	var u1 = User{
		"BlockCheater",
		true,
		Address{
			"Hainan",
			"Sanya",
		},
	}
	fmt.Printf("%+v\n", u1)

	/*JSON序列化*/
	//Student 学生
	type Student struct {
		ID     int
		Gender string
		Name   string
	}

	//Class 班级
	type Class struct {
		Title    string
		Students []*Student
	}
	c := &Class{
		Title:    "101",
		Students: make([]*Student, 0, 200),
	}
	for i := 0; i < 10; i++ {
		stu := &Student{
			Name:   fmt.Sprintf("stu%02d", i),
			Gender: "男",
			ID:     i,
		}
		c.Students = append(c.Students, stu)
	}
	//JSON序列化：结构体-->JSON格式的字符串
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}
	fmt.Printf("json:%s\n", data)
	//JSON反序列化：JSON格式的字符串-->结构体
	str := `{"Title":"101","Students":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"},{"ID":4,"Gender":"男","Name":"stu04"},{"ID":5,"Gender":"男","Name":"stu05"},{"ID":6,"Gender":"男","Name":"stu06"},{"ID":7,"Gender":"男","Name":"stu07"},{"ID":8,"Gender":"男","Name":"stu08"},{"ID":9,"Gender":"男","Name":"stu09"}]}`
	c1 := &Class{}
	err = json.Unmarshal([]byte(str), c1)
	if err != nil {
		fmt.Println("json unmarshal failed!")
		return
	}
	fmt.Printf("%#v\n", c1)

	/*结构体标签：是结构体的元信息，在运行时通过反射的机制读出来，定义在结构体字段的后方*/
	//例如通过指定tag实现json序列化该字段时的key
	type Philosopher struct {
		Name string `json:"Mynameis"`
		Age  int8
		//首字母小写为私有字段，只能在当前结构体中访问，JSON序列化访问不到
		hobby string
	}
	var p6 = Philosopher{
		"Niche",
		33,
		"thinking",
	}
	data_2, err_2 := json.Marshal(p6)
	if err_2 != nil {
		fmt.Println("json marshal failed!")
		return
	}
	fmt.Printf("json str: %s\n", data_2)

	/*需要注意的是，在使用数组、切片、map等对结构体进行赋值时，最好先使用make函数声明内存，然后用copy函数拷贝，
	  否则因为这几个类型包含了指针，当对外界数据进行修改时，结构体内部数据也会被修改
	func (p *Person) SetDreams(dreams []string) {
		p.dreams = make([]string, len(dreams))
		copy(p.dreams, dreams)
	}
	*/
	p7 := Person_2{name: "小王子", age: 18}
	data_3 := []string{"吃饭", "睡觉", "打豆豆"}
	p7.SetDreams(data_3)

	// 你真的想要修改 p1.dreams 吗？
	data_3[1] = "不睡觉"
	fmt.Println(p7.dreams)
}

/*结构体类型中的方法与接收者
  个人理解就是结构体中的函数无需在结构体中进行定义，当然任何类型都可以使用接收者来添加方法
  如果需要通过方法来修改结构体中的值，可以设定指针类型的接收者；此后，为了保证一致性，需要其他的方法也要使用指针类型接收者*/

//接受者可以是值类型，可以是指针类型；调用者可以是值类型，也可以是指针类型。
//此外，接受者的名称一般设定为this或self
func (p person) Dream() {
	fmt.Printf("%s is daydreaming.\n", p.name)
}

func (this *person) SetAge(newAge int8) {
	this.age = newAge
}

func (p *Person_2) SetDreams(dreams []string) {
	p.dreams = dreams
}

func newPerson(name, work string, age int8) *person {
	return &person{
		name: name,
		work: work,
		age:  age,
	}
}
