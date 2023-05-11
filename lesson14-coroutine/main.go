package main

import (
	"fmt"
	"time"
)

/**
	协程
	// 定义函数
	func funcName(params) (results) {}
	// 调用函数
    funcName(params)
	// 通过协程调用函数
	go funcName(params)
*/

func main() {
	// hello 和 main 函数并发执行，下面的代码不会等待 hello 执行完毕
	// main 运行在主协程上，主协程结束，其他协程也会结束
	//go hello()
	//time.Sleep(1 * time.Second)
	//fmt.Println("main")

	go PrintNum(1)
	go PrintNum(2)
	time.Sleep(1 * time.Second)
}

func hello() {
	fmt.Println("hello, world")
}

func PrintNum(num int) {
	for i := 0; i < 3; i++ {
		fmt.Println(num)
		time.Sleep(100 * time.Millisecond)
	}
}
