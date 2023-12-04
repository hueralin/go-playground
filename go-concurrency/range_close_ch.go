package main

import "fmt"

// 发送者可以使用 close 关闭通道，以表明没有值发送了
// 接收者可以通过从通道中取值操作的第二个参数 ok 来判断还有没有值，以及通道是否被关闭
// 如果为 false，说明没值了，或者通道被关闭了
// for i := range ch，用来不断地从通道取值，直到通道被关闭，不关闭的话容易阻塞报错
// 只有发送者才可以去关闭通道，往一个已关闭的通道里写值会 panic
// 通道不像文件，一般不需要关闭，除了要明确关闭通道的场景，如 for range
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		// 把值写入通道
		c <- x
		// 更新 x，y 的值
		t := x
		x = y
		y = t + x
	}
	// 关闭通道
	// 如果注释掉，则程序会正常输出，然后报错死锁
	close(c)
}

func ShowRangeAndCloseCh() {
	ch := make(chan int, 10)
	go fibonacci(cap(ch), ch)
	for i := range ch {
		fmt.Println(i)
	}
}
