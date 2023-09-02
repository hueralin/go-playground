package main

import "fmt"

func main() {
	// var 用来声明和初始化变量
	var a = "init"
	fmt.Println(a)
	// var 可以同时声明多个变量
	var b, c int = 1, 2
	fmt.Println(b, c)
	// Go 可以推断出初始值的类型
	var d = true
	fmt.Println(d)
	// 没有初始化的变量默认是零值
	var e int
	fmt.Println(e)
	// 短语法，声明并初始化，仅用于函数内
	f := "apple"
	fmt.Println(f)
}
