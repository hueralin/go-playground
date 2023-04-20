package main

import (
	"errors"
	"fmt"
)

/**
	函数语法：
	func funcName(parameter_list) result_list {
		// logic
	}
 */

func hello() {
	fmt.Println("hello")
}

func sum(a int, b int) int {
	return a + b
}

// 可变参数, 参数类型相同（类似于前端的收集运算符）
func show(args ...int) int {
	num := 0
	for _, val := range args {
		fmt.Printf("%d ", val)
		num++
	}
	return num
}

// 可变参数，参数类型不同（类似于前端的收集运算符）
func show2(args ...interface{}) {
	for _, val := range args {
		fmt.Printf("%T: ", val)
		fmt.Println(val)
	}
}

// 返回值列表可以不命名
func showBookInfo(name, author string) (string, error) {
	if name == "" {
		return "", errors.New("name is empty")
	}
	if author == "" {
		return "", errors.New("author is empty")
	}
	return name + ", " + author, nil
}

// 命名的返回值列表, 直接给返回值赋值即可，不需要将其 return，会自动返回
func showBookInfo2(name string, author string) (info string, error error) {
	if name == "" {
		error = errors.New("name is empty")
		return
	}
	if author == "" {
		error = errors.New("author is empty")
		return
	}
	info = name + ", " + author
	return
}

func main() {
	hello()
	res1 := sum(1, 2)
	fmt.Println(res1)

	res2 := show(1, 2, 3, 4, 5)
	fmt.Println(res2)

	show2(1, 2.3, "hahaha")

	res3, err3 := showBookInfo("Go", "")
	if err3 != nil {
		fmt.Println(err3)
	} else {
		fmt.Println(res3)
	}

	res4, err4 := showBookInfo2("", "Jack")
	if err4 != nil {
		fmt.Println(err4)
	} else {
		fmt.Println(res4)
	}

	// 匿名函数, 只拥有短暂的生命，一般都是定义后立即使用（好像不能定义在外部，类似于 JS 的立即执行函数）
	func (a, b int) {
		fmt.Println(a + b)
	}(1, 2)

	// 当方法的首字母为 大写 时，这个方法对于 "所有包" 都是 Public ，其他包可以随意调用
	// 当方法的首字母为 小写 时，这个方法是 Private ，其他包是无法访问的
}
