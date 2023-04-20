package main

import "fmt"

// 方法, 其实就是一个函数, 在 func 这个关键字和方法名中间加入了一个特殊的接收器类型
// 接收器可以是结构体类型或者是非结构体类型, 接收器是可以在方法的内部访问的
// func (t Type) methodName(parameterList) (returnList) {}

type Person struct {
	name string
	age int
}

type Author struct {
	name string
}

// 在 Go 中, 相同的名字的方法可以定义在不同的类型上, 而相同名字的函数是不被允许的(但同一类型的指针接收器使用同名时编译会报错)

func (p Person) print(name string) {
	fmt.Printf("name: %s, age: %d\n", p.name, p.age)
	p.name = name
}

func (p *Person) print2(name string) {
	fmt.Printf("name: %s, age: %d\n", p.name, p.age)
	p.name = name
}

func (a Author) print() {
	fmt.Printf("name: %s\n", a.name)
}

// 值接收器 & 指针接收器
// 值接收器表示调用者是个原值的拷贝
// 指针接收器表示是调用者本身，在函数内对调用者做的变更会体现在调用者上

func main() {
	zhangsan := Person{ "zhangsan", 18 }
	zhangsan.print("sanzhang")
	fmt.Println(zhangsan.name)

	lisi := Author{ "lisi" }
	lisi.print()

	wangwu := &Person{ "wangwu", 30 }
	wangwu.print2("wuwang")
	fmt.Println(wangwu.name)
}
