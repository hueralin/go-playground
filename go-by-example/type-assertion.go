package main

import "fmt"

// 类型断言：检查一个类型是否为 nil; 检查一个值是否为某个类型;
// 方式 1：t := i.(T)，断言成功就会将值返回给 t，断言失败就会触发 panic
// 方式 2：t, ok := i.(T)，断言成功就会将值返回给 t，ok 会被置为 true，断言失败不会触发 panic，t 会被置为 T 的零值，ok 会被置为 false

func main() {
	var a interface{} = 10
	t1 := a.(int)
	fmt.Println(t1) // 10
	//t2 := a.(string)
	//fmt.Println(t2) // panic: interface conversion: interface {} is int, not string

	t3, ok := a.(int)
	fmt.Println(t3, ok) // 10 true
	t4, ok2 := a.(string)
	fmt.Println(t4, ok2) // '' false
}
