package main

import "fmt"

func main() {
	// 基本使用
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// 条件语句前面还可以加其他语句，其中声明的变量可以在当前分支及其子分支中使用
	if num := 2; num < 0 {
		fmt.Println("is negative")
	} else if num < 10 {
		fmt.Println("has 1 digit")
	} else {
		fmt.Println("have many digit")
	}

	// Go 中没有三元表达式
	//fmt.Println(1 > 0 ? true : false)
}
