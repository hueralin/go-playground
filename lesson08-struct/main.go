package main

import "fmt"

// 结构体

// Person 命名结构体
type Person struct {
	name   string
	age    int
	gender int
}

// City 结构体的匿名字段
// 第二个字段没有指定字段名，默认就是类型名
type City struct {
	name string
	int
}

// 将方法和结构体绑定，就可以使用 结构体.方法 了
func (p Person) print() {
	fmt.Println(p.name, p.age, p.gender)
}

// 将方法和结构体指针绑定，就能在方法中改变结构体的字段了
func (p *Person) changeName(name string) {
	p.name = name
}

// 方法不报错，能调用，但是不会改变原结构体
func (p Person) changeName2(name string) {
	p.name = name
}

func main() {
	// 指定字段的初始化
	zhangsan := Person{
		name:   "zhangsan",
		age:    18,
		gender: 0,
	}
	// 不指定字段，按照结构体中的字段声明顺序初始化
	lisi := Person{"lisi", 20, 1}
	fmt.Println(zhangsan)
	fmt.Println(lisi)

	// 匿名结构体
	wangwu := struct {
		name   string
		age    int
		gender int
	}{
		name:   "wangwu",
		age:    24,
		gender: 0,
	}
	fmt.Println(wangwu)

	// 结构体的零值，各个字段的零值
	var p = Person{}
	fmt.Println(p) // "" 0 0

	// 部分字段初始化，必须指定字段名
	var p2 = Person{
		name:   "p2",
		gender: 1,
	}
	fmt.Println(p2)
	fmt.Println(p2.name)
	p2.name = "baga"
	fmt.Println(p2.name)

	// 指向结构体的指针
	maliu := &Person{"malin", 24, 1}
	fmt.Println(maliu)
	fmt.Println(*maliu)
	(*maliu).name = "ahhhh"
	fmt.Println((*maliu).name)
	// 结构体指针.xxx 能代替 (*结构体指针).xxx 的解引用写法
	fmt.Println(maliu.age)

	// 匿名字段
	city := City{"beijing", 1}
	fmt.Println(city.int)

	// 结构体嵌套(就是普通的嵌套)

	// 字段提升，如果结构体 A 中有 "匿名的结构体 B" 字段，
	// 那么该 "匿名结构体 B" 里面的字段就被称为提升字段，可以在 A 中直接访问
	// 假设 B { wx }，则应该使用 A.wx，而不是 A.B.wx

	// 结构体比较，其实就是结构体内各个字段的比较
	// 如果结构体的全部成员都是可以比较的，那么结构体也是可以比较的，
	// 那样的话两个结构体将可以使用 == 或 != 运算符进行比较
	// 可以通过 == 运算符或 DeeplyEqual() 函数比较两个结构相同的类型并包含相同的字段值

	// Go 结构体内不能声明方法，但可以通过接收器的方式给结构体绑定一个方法，类似于 JS 的 call, apply
	zhangsan.print()
	// 如果绑定结构体的方法中要改变结构体的属性时，必须使用指针作为方法的接收者
	zhangsan.changeName("lisi")
	zhangsan.changeName2("wangwu")
	fmt.Println(zhangsan)
}
