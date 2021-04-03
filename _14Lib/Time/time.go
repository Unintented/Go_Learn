package main

import (
	"fmt"
	"time"
)

func main() {
	//获取当前时间对象
	println("**********Time object**********")
	now := time.Now()
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Printf("%d-%02d-%02d %02d-%02d-%02d\n", year, month, day, hour, minute, second)

	//获取时间戳
	println("**********Timestamp**********")
	timeStamp_1 := now.Unix()
	timeStamp_2 := now.UnixNano() //纳秒时间戳
	fmt.Printf("%v-%v\n", timeStamp_1, timeStamp_2)
	//	时间戳转换成时间对象
	timeObj := time.Unix(timeStamp_1, 0)
	fmt.Printf("%#v\n %+v\n", timeObj, timeObj)
	fmt.Printf("%d-%02d-%02d %02d-%02d-%02d\n", timeObj.Year(), timeObj.Month(), timeObj.Day(), timeObj.Hour(), timeObj.Minute(), timeObj.Second())

	/*时间操作：
	func (t Time) Add(d Duration) Time
	func (t Time) Sub(u Time) Duration
	func (t Time) Equal(u Time) bool
	func (t Time) Before(u Time) bool
	func (t Time) After(u Time) bool*/

	/*时间格式化:
	  时间类型有一个自带的方法Format进行格式化，需要注意的是Go语言中格式化时间模板不是常见的Y-m-d H:M:S，
	  而是使用Go的诞生时间2006年1月2号15点04分（记忆口诀为2006 1 2 3 4），根据自己的需要进行组合
	*/
	println("**********Time format**********")
	//24小时制
	fmt.Println(now.Format("2006--01--02 03--04--05 Mon Jan"))
	//12小时制
	fmt.Println(now.Format("2006--01--02 03--04--05 PM Mon Jan"))
	fmt.Println(now.Format("Jan Mon 15:04:05 01-02-2006"))

	//解析字符串格式时间
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("Load location error, err:", err)
		return
	}
	timeObj, err = time.ParseInLocation("2006/01/02 15:04:05", "2019/08/04 14:15:20", loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj)

	//定时器：time.Tick(timeGap)，其本质是通道（<-chan Time）,放到最后演示，因为ticker会一直计时
	ticker := time.Tick(time.Second)
	for i := range ticker {
		fmt.Println(i)
	}

}
