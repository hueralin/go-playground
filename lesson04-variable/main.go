package main

import "fmt"

func main() {
	// 变量

	// 方式一：仅声明
	var a int
	fmt.Println("a = ", a)
	fmt.Printf("type of a is %T\n", a)

	// 方式二：声明并初始化
	var b int = 100
	fmt.Printf("b = %d\n", b)
	fmt.Printf("type of b is %T\n", b)

	// 方式三：声明并初始化，但不指定类型，Go 会自动推断
	var c = 200
	fmt.Printf("c = %d\n", c)
	fmt.Printf("type of c is %T\n", c)

	var d string = "haha"
	fmt.Printf("d = %s\n", d)
	fmt.Printf("type of d is %T\n", d)

	// 方式四：短声明
	e := 110
	fmt.Printf("e = %d\n", e)
	fmt.Printf("type of e is %T\n", e)
	f := "who?"
	fmt.Printf("f = %s\n", f)
	fmt.Printf("type of f is %T\n", f)

	// 一次声明多个变量
	var g, h int = 100, 200
	fmt.Println(g, h)
	var i, j = 300, "jj"
	fmt.Println(i, j)
	var (
		k int    = 100
		l string = "lol"
	)
	fmt.Println(k, l)

	// 常量
	const nodeEnv = "development"
	fmt.Println(nodeEnv)

	const (
		BEIJING  = "BEIJING"
		SHANGHAI = "SHANGHAI"
		SHENZHEN = "SHENZHEN"
	)
	fmt.Println(BEIJING, SHANGHAI, SHENZHEN)

	// iota 常量计数器
	const (
		A = iota
		B
		C
	)
	fmt.Println(A, B, C) // 0, 1, 2

	const (
		D = 10 * iota
		E
		F
	)
	fmt.Println(D, E, F) // 0, 10, 20

	// 25 个关键字
	var s3 = []int{1, 2, 3}
	fmt.Printf("len is %v\n", len(s3))
	fmt.Printf("cap is %v\n", cap(s3))
}
