package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	/*数字字面量*/
	//八进制以0开头，或0o开头
	v1 := 0o57
	fmt.Printf("%o+%d\n", v1, v1)
	//十六进制数
	v2 := 0x1f
	fmt.Printf("%x+%X+%d\n", v2, v2, v2)
	//二进制数
	v3 := 0b01010
	fmt.Printf("%b+%d\n", v3, v3)

	/*无符号整型：uint8,uint16,uint32,uint64,uint
	  有符号整型：int8,int16,int32,int64,int
	  其中int与uint会根据机器的指令长度变化，在涉及跨平台编译时最好不用
	  其它的符号根据变量的实际取值选择，节省空间，整型数据默认类型为int
	  int32与int64为不同的类型，不能相互赋值
	  int8数据范围：-2^7~2^7-1   uint数据范围：0～2^8-1，其他类似
	  uintptr为无符号整型，用于存放一个指针*/
	var age uint8 = 18
	fmt.Println(age)

	/*浮点型：float32,float64
	  可以使用科学计数法来表示，浮点数默认为float64*/
	var f1 float32 = 3.1415
	fmt.Printf("%f\n", f1)
	fmt.Println("max float32 value: ", math.MaxFloat32)
	fmt.Println("max float64 value: ", math.MaxFloat64)

	/*复数：complex64,complex128*/
	var c1 complex64 = 3.2 + 5i
	var c2 = complex(3.4, 6)
	fmt.Println(c1)
	fmt.Printf("real value:", real(c2), "\nimage value:", imag(c2))

	/*字符串：在双引号中，多行字符串放在反引号中
	  其它对字符串的常见操作包含在strings类库中
	  初始化后可以以下标来读取，但不允许修改
	  字符串相加是拼接
	  for index=0;index<len(stringName);index++{}这种方式的遍历是逐字节byte(uint8)遍历
	  for index,val := range stringName{}遍历值的类型是rune(int32)，可以遍历中文*/
	fmt.Println("\nstr := \"usr\\local\\share\\bitcoin\"")
	var s1 = `这
真是
一首
好诗`
	fmt.Println(s1)
	var s2 = strings.Split(s1, "\n")
	fmt.Println(s2)

	/*字符：包括byte类型与rune类型
	  byte类型与uint8等价，表示ASCII码中的一个字符，字符串类型可以通过逐个遍历字节来访问
	  rune类型表示UTF-8编码的一个字符，由一个或多个byte组成，实际上是一个int32类型，不能通过逐个遍历字节来访问*/
	travelString("Hello,朱莉")
	//修改字符串要先转换成[]string或[]rune，完成后再转回string，但都会重新分配内存，复制数组
	changeString()
}

func travelString(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v(%c)\n", s[i], s[i])
	}
	for _, v := range s {
		fmt.Printf("%v(%c)", v, v)
	}
}
func changeString() {
	s1 := "big"
	// 强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'p'
	fmt.Println("\n", string(byteS1))

	s2 := "白萝卜"
	runeS2 := []rune(s2)
	runeS2[0] = '红'
	fmt.Println(string(runeS2))
}
