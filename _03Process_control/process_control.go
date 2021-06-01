package main

import (
	"fmt"
	"os"
)

func main() {
	/*if-else*/
	ifElse()

	/*for-loop*/
	forFunc()

	/*goto*/
	gotoFunc()

	/*switch*/
	switchFunc()
}

func switchFunc() {
	/*Go中的switch不需要break，默认执行完匹配的case后停止执行
	  如果需要继续往下顺序执行，使用fallthrough*/
	//定义新类型Gender
	type Gender bool
	const (
		Male   Gender = true
		Female Gender = false
	)
	Leesin := Male
	switch Leesin {
	case Female:
		fmt.Println("Leesin is a girl.")
	case Male:
		fmt.Println("Leesin is a boy.")
		fallthrough
	default:
		fmt.Println("Leesin is somebody.")
	}

	//	switch variableName := 100; {//可赋初值
	//    	 case variableName > 90:
	// 		 ...
	//	}
	//  switch {//switch后为空可模拟if else
	//		case variableName > 90:
	// }
	//  此外，一个case中可以有多个条件，只要满足其中一个即可
}

func gotoFunc() {
	//使用for range遍历数组时，value是对应索引的值拷贝，不会影响原数据；若想修改，可以设成指针类型
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签
				goto breakTag
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	return
	// 标签
breakTag:
	fmt.Println("结束for循环")
}

func forFunc() {
	for i := 0; i < 3; i++ {
		println("*")
	}
	//	无限循环通过break,goto,return,panic语句强制退出
STOPFOR:
	for {
		println("Infinite for loop...")
		break STOPFOR
	}
	//listdemo := [3]uint8{2, 4, 8}
	//var listdemo = [...]uint8{2,4,8}
	var listdemo = [...]uint8{1: 2, 3: 16}
	for i, v := range listdemo {
		fmt.Println("i:", i, " v:", v)
	}
}

func ifElse() {
	score := 65
	if score >= 90 {
		println("A")
	} else if score >= 70 {
		println("B")
	} else {
		println("C")
	}

	//	打开某一个文件不存在或者由于权限的问题无法操作，通过这种方式判断在go语言中普遍存在，推荐使用
	if fileHandle, err := os.Open("hello.txt"); err != nil {
		//null
		fmt.Println(fileHandle)
		fmt.Println(err.Error())
	} else {
		fmt.Println("获取文件成功")
	}
}
