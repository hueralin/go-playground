package main

import "fmt"

// 流程控制：顺序、条件、选择、循环

func hello(name string) {
	fmt.Println("Hello " + name)
}

func hello2(name *string) {
	fmt.Println(*name)
}

var s string = "top"

// defer 在 return 后调用
func after() string {
	defer func() {
		s = "inner"
	}()
	fmt.Println("after " + s)
	return s
}

func main() {
	// 条件语句
	v1 := 80
	if v1 <= 100 {
		fmt.Println("baga!")
	}

	// v2 的作用域在条件判断语句内
	if v2 := 90; v2 <= 80 {
		fmt.Println("baga too!")
	} else {
		fmt.Println(v2)
	}

	// 选择语句
	v3 := "E"
	// 也可以将 v3 的声明和初始化放在 switch 后面，和 if 类似，作用域在 switch 内部
	switch v3 {
	case "A":
		fmt.Println("AAA")
		// Go 的 switch 没有 break，且只执行满足 case 的逻辑
		// 如果想执行下一个逻辑，应使用 fallthrough
		fallthrough
	case "B":
		fmt.Println("BBB")
		// case 后可以接多个条件, 多个条件之间是或的关系, 用逗号 , 相隔
	case "C", "D", "E":
		fmt.Println("CCC DDD EEE")
	default:
		fmt.Println("不及格")
	}

	// switch 后面可以接一个函数, 只要保证 case 后的值类型与函数的返回值一致即可

	// 无表达式的 switch
	// 如果省略该表达式, 则表示这个 switch 语句等同于 switch true, 并且每个 case 表达式都被认定为有效, 相应的代码块也会被执行
	// 下面的代码等同于 if-elseif-else
	score := 60
	switch {
	case score >= 90 && score <= 100:
		fmt.Println("nb!")
	case score >= 80 && score < 90:
		fmt.Println("还行")
	default:
		fmt.Println("小垃圾")
	}

	// 循环语句，和其它语言一样，不过只有 for
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d ", i)
	}

	// defer 延迟调用, defer 语句后面跟着的函数会延迟到 "当前函数" 执行完后再执行
	// Go 会将 defer 后的函数放进栈中，后进先出
	defer hello("A")
	defer hello("B")
	defer hello("C")
	fmt.Println("yeah~")
	// yeah~ C B A

	str := "hello"
	ptr := &str
	defer hello2(ptr) // plane
	*ptr = "plane"
	fmt.Println(*ptr) // plane

	// defer 在 return 后调用
	res := after()            // after top
	fmt.Println("res " + res) // res top
	fmt.Println("s " + s)     // s inner
}
