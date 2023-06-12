package main

import "fmt"

// new 函数只接受一个参数，这个参数是一个类型，并且返回一个指向该类型内存地址的指针。
// 同时 new 函数会把分配的内存置为零，也就是类型的零值。

// make 也是用于内存分配的，但是和 new 不同，它只用于 chan、map 以及 slice 的内存创建，
// 而且它返回的类型就是这三个类型本身，而不是他们的指针类型，因为这三种类型就是引用类型，所以就没有必要返回他们的指针了。

type Student struct {
	name string
	age  int
}

func main() {
	// 为基本类型的变量分配内存空间
	var sum *int = new(int)
	*sum = 1
	fmt.Println(sum)  // ...
	fmt.Println(*sum) // 1

	// 为自定义类型的变量分配内存空间
	var s1 Student = Student{name: "Tom", age: 18}
	var s2 *Student
	// 未初始化，直接访问会报错，panic: runtime error: invalid memory address or nil pointer dereference
	//fmt.Println(s2.name, s2.age)
	s2 = new(Student)
	s2.name = "Jack"
	s2.age = 22
	fmt.Println(s1.name, s1.age)
	fmt.Println(s2.name, s2.age)
}
