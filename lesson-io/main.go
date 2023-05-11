package main

import "fmt"

func main() {
	var x int
	var y float64

	// fmt.Println() 打印并换行
	// fmt.Printf() 格式化输出
	// fmt.Print() 打印输出

	fmt.Scanln(&x, &y)
	fmt.Println("x: ", x)
	fmt.Println("y: ", y)
}
