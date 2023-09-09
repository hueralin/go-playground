package main

import "fmt"

func main() {
	// for 是 Go 唯一的循环语句
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	var j = 1
	for j <= 10 {
		fmt.Println(j * 2)
		j++
	}

	for {
		fmt.Println("loop break")
		break
	}
}
