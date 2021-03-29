package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	wg           sync.WaitGroup
	Lock         sync.Mutex
	RLock        sync.RWMutex
	PublicData   int
	PublicData_2 int
)

func hello() {
	fmt.Println("goroutine in hello")
}
func multi_goroutine(i int) {
	wg.Done()
	fmt.Println(i)
}

func main() {
	/*go语言中的并发通过goroutine实现，类似于线程，由go的运行时(runtime)调度管理，区别于操作系统调度OS线程。Go程序会智能地将 goroutine 中的任务合理地分配给每个CPU
	  当需要让某个任务并发执行的时候，只要把这个任务包装成一个函数，开启一个goroutine去执行这个函数就可以了*/

	//需要注意的是，go会为main函数默认创建一个goroutine，当main()函数返回的时候该goroutine就结束了，所有在main函数中启动的goroutine会一同结束
	//所以在此处我们强制让main函数多等待一会，好让hello函数所在的goroutine创建并执行
	go hello()
	fmt.Println("goroutine in main")
	time.Sleep(time.Second)

	//此处启动多个goroutine，每次执行的顺序都不一样，因为goroutine的调度是随机的
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go multi_goroutine(i)
	}
	wg.Wait()

	/*Go语言中的操作系统线程和goroutine的关系：
		一个操作系统线程对应用户态多个goroutine。
		go程序可以同时使用多个操作系统线程。
		goroutine和OS线程是多对多的关系，即m:n。
	  可使用runtime.GOMAXPROCS(num)来设置使用的CPU逻辑核心数，默认是使用所有的
	*/

	/*channel
	  Go语言的并发模型是CSP（Communicating Sequential Processes），提倡通过通信共享内存而不是通过共享内存而实现通信。
	  如果说goroutine是Go程序并发的执行体，channel就是它们之间的连接。channel是可以让一个goroutine发送特定值到另一个goroutine的通信机制。
	  Go 语言中的通道（channel）是一种特殊的类型。通道像一个传送带或者队列，总是遵循先入先出（First In First Out）的规则，保证收发数据的顺序。
	  每一个通道都是一个具体类型的导管，也就是声明channel的时候需要为其指定元素类型。*/

	//通道是引用类型，空值是nil，需要用make函数初始化
	//make(chan eleType, [bufferSize])
	var ch chan int
	fmt.Println(ch)
	ch = make(chan int)
	//channel有send、receive和close操作，发送和接收都使用<-符号
	//关闭后的通道有以下特点：
	//对一个关闭的通道再发送值就会导致panic。
	//对一个关闭的通道进行接收会一直获取值直到通道为空。
	//对一个关闭的并且没有值的通道执行接收操作会得到对应类型的零值。
	//关闭一个已经关闭的通道会导致panic。

	/*无缓冲的通道与有缓冲的通道：
	  我们在使用make函数来为通道初始化时，如果未指定缓冲大小，那该通道为无缓冲的通道，否则为有缓冲的通道；
	  无缓冲的通道在发送数据后，会一直阻塞，直到某一goroutine从该通道中接收数据，因此如果先在main函数所在的goroutine中直接发送数据会死锁
	  两种方法：
		1、先在一个goroutine中接收数据，再在main函数所在goroutine中发送数据
		2、先在一个goroutine中发送数据，然后再在一个goroutine中接收数据*/
	//法1
	println("**********Differences between buffered channel and unbuffered channel**********")
	ch_2 := make(chan string)
	go receiveDataFromChan_new(ch_2)
	ch_2 <- "I'm here."

	//法2
	//把10发送到传送整型的通道ch中
	go sendDataToChan(ch, 10)
	//从通道中获取值
	go receiveDataFromChan(ch)
	//为了防止main函数所在的goroutine结束，导致前面的goroutine还未正常执行就被迫结束了
	time.Sleep(time.Second)

	//对于有缓冲的通道，在使用make函数初始化的bufferSize就是该通道可存放元素的个数
	ch_3 := make(chan int, 3)
	ch_3 <- 8
	receiveDataFromChan(ch_3)

	/*从通道中循环取值：从上面可以知道，当一个通道被关闭后，不能再进行发送和关闭操作
	  而从该通道取值的操作会先取完通道中的值，再然后取到的值一直都是对应类型的零值
	  有两种方式可以判断通道中的元素是否被取完：
	*/
	println("**********Retrieve element from channel*********")
	ch_4 := make(chan int)
	ch_5 := make(chan int)
	// 开启goroutine将0~100的数发送到ch1中
	go func() {
		for i := 0; i < 10; i++ {
			ch_4 <- i
		}
		close(ch_4)
	}()
	// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
	go func() {
		for {
			i, ok := <-ch_4 // 通道关闭后再取值ok=false
			if !ok {
				break
			}
			ch_5 <- i * i
		}
		close(ch_5)
	}()
	// 在主goroutine中从ch2中接收值打印
	for i := range ch_5 { // 通道关闭后会退出for range循环
		fmt.Println(i)
	}
	time.Sleep(time.Second)

	/*单向通道：限制通道在函数中只能发送或只能接收
	  chan <- int是一个只写通道
	  <- chan int是一个只读通道*/
	println("**********Single direction channel**********")
	ch_6 := make(chan int, 10)
	ch_7 := make(chan int, 10)
	go func(in chan<- int) {
		for i := 0; i < 10; i++ {
			in <- i
		}
		close(in)
	}(ch_6)
	go func(in chan<- int, out <-chan int) {
		for i := range out {
			in <- i * i
			fmt.Println("i in ch_6:", i)
		}
		close(in)
	}(ch_7, ch_6)
	time.Sleep(time.Second)
	func(out <-chan int) {
		for i := range out {
			fmt.Println("i in ch_7:", i)
		}
	}(ch_7)

	/*select多路复用：
	  可处理一个或多个channel的发送/接收操作。
	  如果多个case同时满足，select会随机选择一个。
	  对于没有case的select{}会一直等待，可用于阻塞main函数。*/
	println("**********select multi-road**********")
	ch_8 := make(chan int, 1)
	for j := 0; j < 10; j++ {
		select {
		case x := <-ch_8:
			println(x)
		case ch_8 <- j:
		}
	}

	/*并发安全与锁：当多个goroutine同时操作一个临界区资源，会发生竞态问题；此时，就需要对临界区资源加锁
	  常用两种锁：
		1、互斥锁：多个goroutine同时等待一个锁时，唤醒的策略是随机的，使用sync包中的Mutex实现。
		2、读写互斥锁：有很多实际的场景下是读多写少的，当并发的去读取一个资源而不涉及资源修改的时候是没有必要加锁的，这种场景下使用读写锁是更好的一种选择，使用sync包中的RWMutex实现。
		读写锁分为两种：读锁和写锁。当一个goroutine获取读锁之后，其他的goroutine如果是获取读锁会继续获得锁，如果是获取写锁就会等待；
		当一个goroutine获取写锁之后，其他的goroutine无论是获取读锁还是写锁都会等待。*/
	println("**********Locker in Golang**********")
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println("PublicData:", PublicData)

	println("**********Read and write locker in Golang**********")
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()

	/*并发任务的同步：sync.WaitGroup内部维护着一个计数器，计数器的值可以增加和减少。
	  例如当我们启动了N 个并发任务时，就将计数器值增加N。每个任务完成时通过调用Done()方法将计数器减1。
	  通过调用Wait()来等待并发任务执行完，当计数器值为0时，表示所有并发任务已经完成。*/
	println("**********Elegant concurrent task**********")
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("I'm in a concurrent task")
	}()
	wg.Wait()

	/*高并发场景下，go还提供了sync.Once、sync.Map等结构
	  Go语言的sync包中提供了一个开箱即用的并发安全版map——sync.Map；
	  开箱即用表示不用像内置的map一样使用make函数初始化就能直接使用；
	  同时sync.Map内置了诸如Store、Load、LoadOrStore、Delete、Range等操作方法*/
}

func read() {
	//注意不要写成了RLock.RLocker()
	RLock.RLock()
	//假设读操作耗时1毫秒
	time.Sleep(time.Millisecond)
	RLock.RUnlock()
	wg.Done()
}

func write() {
	RLock.Lock()
	PublicData_2 += 1
	//假设写操作耗时1毫秒
	time.Sleep(time.Millisecond)
	RLock.Unlock()
	wg.Done()
}

func add() {
	for i := 0; i < 1000; i++ {
		Lock.Lock()
		PublicData += 1
		Lock.Unlock()
	}
	wg.Done()
}

func receiveDataFromChan_new(ch_2 chan string) {
	fmt.Println(<-ch_2)
}

func receiveDataFromChan(ch chan int) {
	x := <-ch
	fmt.Println(x)
}

func sendDataToChan(ch chan int, i int) {
	ch <- i
}
