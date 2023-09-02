package main

import "fmt"

func main() {
	// 元素的类型和个数都是数组类型的一部分
	var a [5]int
	fmt.Println(a)      // [0 0 0 0 0]
	fmt.Println(len(a)) // 5

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)
}
