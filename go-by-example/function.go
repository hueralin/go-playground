package main

import "fmt"

// Go 要求显式的指明返回值类型
func sum(a int, b int) int {
	return a + b
}

// 返回多个值
func multi() (int, int) {
	return 0, 1
}

// 可变参数函数，参数实际上是切片类型
func sumMulti(nums ...int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

// 匿名函数及闭包
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	a := sum(1, 2)
	fmt.Println(a)

	b, c := multi()
	fmt.Println(b, c)

	// 可以一个一个传
	d := sumMulti(1, 2, 3, 4)
	fmt.Println(d)

	// 也可以使用...直接解构一个现成的切片
	e := []int{1, 2, 3, 4}
	f := sumMulti(e...)
	fmt.Println(f)

	g := intSeq()
	fmt.Println(g()) // 1
	fmt.Println(g()) // 2
	fmt.Println(g()) // 3
	g = intSeq()
	fmt.Println(g()) // 1
}
