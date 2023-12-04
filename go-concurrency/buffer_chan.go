package main

import "fmt"

func ShowBufferCh() {
	// 缓存通道，带缓冲容量的通道，可以通过 cap(ch) 获取通道容量
	// ch := make(chan int, 2)
	// 如果通道满了，那么往通道里写值会被阻塞
	// 如果通道空了，从通道里取值也会被阻塞
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	// 通道已经满了，再往里面写值就会阻塞，代码会报死锁，
	// 因为当前 goroutine 在等待其他 goroutine 取值，但是并没有 goroutine
	// fatal error: all goroutines are asleep - deadlock!
	// ch <- 3
	fmt.Println(<-ch) // 1
	fmt.Println(<-ch) // 2
	// fmt.Println(<-ch) // 同理，fatal error: all goroutines are asleep - deadlock!
}
