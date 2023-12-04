package main

import (
	"fmt"
	"time"
)

// select 语句让 goroutine 等待多个通信操作。
// select 语句会阻塞，直到某个 case 满足条件，然后去执行那个 case
// 如果有多个 case 满足，那么它会随机选一个（啧）
// 写法上类似于 switch

func fibonacci2(c chan int, quit chan int) {
	x, y := 0, 1
	for {
		select {
		// 无缓冲通道，没值则写入，有值则阻塞
		case c <- x:
			x, y = y, x+y
		// quit 通道没值则阻塞，有值了，就打印 "quit"，结束当前函数
		case <-quit:
			fmt.Println("quit")
			return
		// 也有默认值
		default:
			fmt.Println("...")
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func ShowSelect() {
	c := make(chan int)
	quit := make(chan int)
	// 以下是此 goroutine 和 fibonacci2 通过通道通信，协调完成任务
	go func() {
		for i := 0; i < 10; i++ {
			// 不断从通道 c 取值
			fmt.Println(<-c)
		}
		// 10 次取值完毕，给 quit 通道发消息，并结束当前函数
		quit <- 0
	}()
	fibonacci2(c, quit)
}
