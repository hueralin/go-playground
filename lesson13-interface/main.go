package main

import "fmt"

// 接口：方法签名(Method Signature)的集合

// 创建类型或者结构体，并为其绑定接口定义的方法，接收者为该类型或结构体，方法名为接口中定义的方法名，这样就说该类型或者结构体实现了该接口

// Go 语言提供了另外一种数据类型即接口，它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口。
// 接口可以让我们将不同的类型绑定到一组公共的方法上，从而实现多态和灵活的设计。
// Go 语言中的接口是隐式实现的，也就是说，如果一个类型实现了一个接口定义的所有方法，那么它就自动地实现了该接口。
// 因此，我们可以通过将接口作为参数来实现对不同类型的调用，从而实现多态。

// Go官方文档中表示: interface 本身是引用类型，即接口类型本身是指针类型。

type Study interface {
	learn()
}

type Student struct {
	name string
	book string
}

type Worker struct {
	name string
	book string
	by   string
}

// Student 结构体隐式实现 Study 接口
// 表明了对于 Student 来说，既可以使用值调用，也可以使用指针调用
func (s Student) learn() {
	fmt.Println(s.name + " study " + s.book)
}

// Worker 结构体隐式实现 Study 接口
// 表明了对于 Worker 来说，只能使用指针调用
func (w *Worker) learn() {
	fmt.Println(w.name + " study " + w.book + " by " + w.by)
}

// 接口可以看作是 type 和 value 的组合，type 是底层数据的类型，value 是底层数据的值
func showInterface(s Study) {
	fmt.Printf("接口类型: %T, 接口值: %v\n", s, s)
}

// 空接口参数
func showEmptyInterface(i interface{}) {
	fmt.Printf("type: %T, value: %v\n", i, i)
}

// 新例子

type Phone interface {
	call()
}

type NokiaPhone struct{}

func (p *NokiaPhone) call() {
	fmt.Println("I am Nokia")
}

type IPhone struct{}

func (p *IPhone) call() {
	fmt.Println("I am IPhone")
}

func main() {
	s := Student{
		name: "zhangsan",
		book: "《时间简史》",
	}
	s.learn()

	w := Worker{
		name: "lisi",
		book: "《团结就是力量》",
		by:   "听说",
	}
	w.learn()

	// 声明两个接口变量
	var s1 Study
	var s2 Study

	s1 = s
	s1.learn()

	s1 = &s
	s1.learn()

	//s2 = w
	s2 = &w
	s2.learn()

	// 接口类型: main.Student, 接口值: {zhangsan 《时间简史》}
	showInterface(s)

	// 空接口参数，可以说任何类型都至少实现了空接口
	str := "hello"
	pi := 3.14
	showEmptyInterface(str)
	showEmptyInterface(pi)

	// 定义一个空接口，有两个字段：类型和值
	var i interface{}
	// type: <nil>, value: <nil>
	fmt.Printf("type: %T, value: %v\n", i, i)

	// 空接口变量可以接受任何类型的值
	i = str
	fmt.Println(i) // hello
	i = pi
	fmt.Println(i) // 3.14

	// 但是反过来不行，即空接口变量被赋予一个确定类型的值后，空接口变量就不能赋值给其他类型的变量
	// var xxx = "wow"
	// xxx= i // Cannot use 'i' (type interface{}) as the type string

	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()
}
