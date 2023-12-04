package main

import (
	"fmt"
	"sync"
	"time"
)

var wg2 sync.WaitGroup

func ShowRace() {
	baton := make(chan int)

	wg2.Add(1)

	go runner(baton)

	baton <- 1

	wg2.Wait()
}

func runner(baton chan int) {
	var newRunner int

	// 等待接力棒
	currRunner := <-baton
	fmt.Printf("Runner %d Running With Baton\n", currRunner)

	// 创建下一个跑者，让其准备好
	if currRunner < 4 {
		newRunner = currRunner + 1
		go runner(baton)
	}

	// 围绕赛道跑
	time.Sleep(100 * time.Millisecond)
	// 跑到位置了（可能是交接位置，也可能是终点）
	fmt.Printf("Runner %d To The Line\n", currRunner)

	// 比赛结束了吗？
	if currRunner == 4 {
		fmt.Printf("Runner %d Finished, Race Over\n", currRunner)
		wg2.Done()
		return
	}

	// 将接力棒交给下一个跑者
	fmt.Printf("Runner %d Exchange With Runner %d\n", currRunner, newRunner)
	baton <- newRunner
}
