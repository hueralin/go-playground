package main

import (
	"fmt"
	"slices"
)

func main() {
	// 切片的类型只取决于元素的类型
	// 未初始化的切片的值是 nil，长度为 0
	var a []int
	fmt.Println(a)        // []
	fmt.Println(len(a))   // 0
	fmt.Println(a == nil) // true

	// 要想创建一个长度不为 0 的空切片，可以使用 make 函数
	// 第二个参数用来指定切片的初始长度
	// 第三个可选参数用来指定切片的容量，默认等于长度
	var b = make([]int, 3)
	fmt.Println(b)              // [0 0 0] 默认零值
	fmt.Println(len(b), cap(b)) // 3 3

	// 向切片中追加值
	// append 返回切片
	b = append(b, 1)
	b = append(b, 2, 3)
	fmt.Println(b) // [0 0 0 1 2 3]

	// 切片拷贝
	c := make([]int, len(b))
	copy(c, b)
	fmt.Println(c)                  // [0 0 0 1 2 3]
	fmt.Println(&c == &b)           // false
	fmt.Println(slices.Equal(c, b)) // true，内容一样

	// 切片操作，左闭右开
	c1 := c[2:5]
	fmt.Println(c1) // [0 1 2]
	c2 := c[:5]
	fmt.Println(c2) // [0 0 0 1 2]
	c3 := c[2:]
	fmt.Println(c3) // [0 1 2 3]

	// 声明并初始化
	d := []string{"a", "b", "c"}
	fmt.Println(d) // [a b c]

	// slices package 中有一些实用的方法
	// slices.Equal 用来判断切片的内容是否一样
	e := []string{"a", "b", "c"}
	fmt.Println(slices.Equal(d, e)) // true，内容一样
	e = append(e, "d")
	fmt.Println(slices.Equal(d, e)) // false，内容不一样

	// 请注意，虽然切片与数组的类型不同，但它们由 fmt.Println 呈现类似。
}
