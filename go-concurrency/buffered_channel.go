package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 四个 goroutine
var numberGoroutines = 4

// 十个任务
var taskLoad = 10
var wg3 sync.WaitGroup

func init() {
	rand.Seed(time.Now().Unix())
}

func ShowWorker() {
	taskChan := make(chan string, 10)
	wg3.Add(numberGoroutines)

	// worker 就绪
	for i := 1; i <= numberGoroutines; i++ {
		go worker(taskChan, i)
	}

	// 发任务咯
	for i := 1; i <= taskLoad; i++ {
		taskChan <- fmt.Sprintf("Task %d", i)
	}

	close(taskChan)
	wg3.Wait()
}

func worker(taskChan chan string, worker int) {
	defer wg3.Done()

	for {
		task, ok := <-taskChan
		if !ok {
			// 说明通道被关闭，且没有值了
			fmt.Printf("Worker: %d : Shutting Down\n", worker)
			return
		}

		fmt.Printf("Worker: %d : Started %s\n", worker, task)

		// 随机模拟工作时长
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		fmt.Printf("Worker: %d : Completed %s\n", worker, task)
	}
}

/**
对于有缓冲的通道，关闭后，不能再向其中写入值，但是可以从里面读取值（如果有的话）。
从一个已经关闭，且没有数据的通道中取值，总是能立刻返回，并返回一个通道类型的零值。
*/
