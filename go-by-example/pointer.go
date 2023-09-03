package main

import "fmt"

// 值传递
func zeroVal(val int) {
	val = 0
}

// 引用传递
func zeroPtr(val *int) {
	*val = 0
}

func main() {
	i := 1
	fmt.Println(i) // 1

	zeroVal(i)
	fmt.Println(i) // 1

	zeroPtr(&i)
	fmt.Println(i) // 0
	fmt.Println(&i)
}
