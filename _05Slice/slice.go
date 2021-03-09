package main

import "fmt"

func main() {
	/*切片（Slice）是一个拥有相同类型元素的可变长度的序列。它是基于数组类型做的一层封装。它非常灵活，支持自动扩容。
	  切片是一个引用类型，它的内部结构包含地址、长度和容量，可以将切片理解为在底层的数组上进行操作，容量为给该切片在底层数组上分配了多少个可使用的存储空间，长度为实际已使用的存储空间。
	  切片的拷贝与从同一个数组上生成的两个切片，其底层指向同一个数组，因此修改具有统一性。
	*/
	//声明一个整型切片
	var a []int
	//声明一个字符串切片并初始化
	var b = []string{}
	fmt.Println(a == nil, b == nil)

	//简单切片表达式：sliceName[lowIndex : highIndex]
	c := [5]int{1, 2, 3, 4, 5}
	s := c[1:3] // s := a[low:high]
	fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))
	//完整切片表达式：sliceName[lowIndex : highIndex : max]，其中额外设置capacity为max-lowIndex，若lowIndex省略，默认为0
	s2 := c[1:3:4]
	fmt.Printf("s2:%v len(s2):%v cap(s2):%v\n", s2, len(s2), cap(s2))
	//make函数构造切片：make([]sliceType, size, capacity)
	s3 := make([]int, 2, 10)
	fmt.Println(s3)
	//判断切片是否为空不能与nil比较
	fmt.Println(len(s3) == 0)

	/*append方法为切片添加容量：每个切片会指向一个底层数组，这个数组的容量够用就添加新增元素。
	  当底层数组不能容纳新增的元素时，切片就会自动按照一定的策略进行“扩容”，此时该切片指向的底层数组就会更换。
	  “扩容”操作往往发生在append()函数调用时，所以我们通常都需要用原变量接收append函数的返回值。*/
	var s4 []int
	s4 = append(s4, 1)
	s4 = append(s4, 2, 3, 4)
	s5 := []int{5, 6, 7}
	//ATTENTION:将一个切片append到另一个切片后，需要加"..."
	s4 = append(s4, s5...)
	fmt.Println(s4)
	//append扩容演示
	var numSlice []int
	for i := 0; i < 10; i++ {
		numSlice = append(numSlice, i)
		fmt.Printf("%v  len:%d  cap:%d  ptr:%p\n", numSlice, len(numSlice), cap(numSlice), numSlice)
	}

	//使用copy(destSlice, srcSlice)函数复制源切片，如此一来二者修改互不影响
	//从切片中删除元素：a = append(a[:index], a[index+1:]...)
	var s6 = []int{1, 2, 3, 4, 5}
	s6 = append(s6[:2], s6[3:]...)
	fmt.Println(s6)

}
