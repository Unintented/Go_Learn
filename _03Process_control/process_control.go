package main

import "fmt"

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
}

func gotoFunc() {
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
}
