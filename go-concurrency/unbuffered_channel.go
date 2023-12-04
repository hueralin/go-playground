package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func ShowTennis() {
	// court: 球场
	court := make(chan int)

	// 计数 +2，表示要等待两个 goroutine
	wg.Add(2)

	// 两名选手入场
	go player("张三", court)
	// 注释掉李四，则会打印 goroutine 1 [sem acquire]，表明 goroutine 需要等待信号量
	go player("李四", court)

	// 将球扔到球场
	court <- 1

	wg.Wait()
}

func player(name string, court chan int) {
	defer wg.Done()

	for {
		ball, ok := <-court
		if !ok {
			// 通道被关闭，赢了！
			fmt.Printf("Player %s Won\n", name)
			return
		}
		// 选随机数，来判断是否丢球
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)
			// 关闭通道，表示认输
			close(court)
			return
		}
		// 显示击球数，并将击球数 +1
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++

		// 将球打向球场，并阻塞在此，直到对方从通道中读取值
		court <- ball
	}
}

// 无缓冲通道要求发送和接受的 goroutine 同时准备好，才能完成发送和接受操作，否则会被阻塞。
// 无缓冲，即该通道不能保存值，写入后需要立刻被读出去，不然会被阻塞住。
// 读操作也同理，如果没有理解写入值，那么读操作也会被阻塞住。

// 无缓冲通道在Go语言中是一种强大的同步工具，可以用于协调多个goroutine之间的通信和同步操作。
// 无缓冲通道的特点是发送和接收操作是同步的，也就是说发送操作会等待接收操作，而接收操作也会等待发送操作。这种同步机制确保了数据的安全传递和有序处理。
// 当一个goroutine尝试向无缓冲通道发送数据时，如果没有其他goroutine正在等待接收数据，发送操作会被阻塞，直到有其他goroutine准备好接收数据为止。
// 同样地，当一个goroutine尝试从无缓冲通道接收数据时，如果没有其他goroutine正在等待发送数据，接收操作也会被阻塞，直到有其他goroutine准备好发送数据为止。
// 这种同步机制使得无缓冲通道成为一个非常有用的工具，可以确保数据在goroutine之间的正确传递和顺序执行。但是，需要注意的是，如果没有足够的goroutine来发送或接收数据，就会发生死锁，所以在使用无缓冲通道时需要小心处理并发的情况。
