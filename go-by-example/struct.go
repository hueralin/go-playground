package main

import "fmt"

type person struct {
	name string
	age  int
}

// 将新结构体的创建封装在构造函数中是惯用的做法
func newPerson(name string) *person {
	p := person{name: name}
	// 没被初始化的属性将会被赋予零值
	return &p
}

type rect struct {
	width  int
	height int
}

// 结构体上的方法 - 值接收器类型
func (r rect) area() int {
	// 方法里的 r 实际上是接收器的一个拷贝
	return r.width * r.height
}

// 结构体上的方法 - 指针接收器类型
func (r *rect) perim() int {
	// 方法里的 r 实际上就是接收器本身
	// 可以修改接收器的值
	r.width *= 2
	return 2*r.width + 2*r.height
}

func main() {
	fmt.Println(*newPerson("tom"))            // {tom 0}
	fmt.Println(person{name: "Bob", age: 20}) // {Bob 20}

	var p = person{name: "jelly"}
	fmt.Println(p.name) // jelly

	var p2 *person = newPerson("jack")
	// 指针也可以通过 . 访问字段，Go 自动进行解引用操作
	fmt.Println(p2.name) // jack

	// 匿名结构类型，一次性使用
	var b = struct {
		name string
	}{
		"haha",
	}
	fmt.Println(b) // {haha}

	// Go 自动处理方法调用的值和指针之间的转换
	// 你可能希望使用指针接收器类型来避免在方法调用上进行复制或允许该方法改变结构本身
	var r = rect{width: 10, height: 10}
	fmt.Println(r.area()) // 100
	rp := &r
	fmt.Println(rp.perim()) // 60
}
