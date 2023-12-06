package main

import "fmt"

type Inner struct {
	Field1 string
	Field2 string
}

type Outer struct {
	// 相当于 Inner Inner，但是需和会实现字段和方法的提升
	Inner
	Field3 string
}

type Outer2 struct {
	*Inner
	Field3 string
}

func StructEmbed() {
	// 初始化时要使用 Inner 作为 key
	a := Outer{
		Inner: Inner{
			Field1: "f1",
			Field2: "f2",
		},
		Field3: "f3",
	}
	fmt.Println(a)        // {{f1 f2} f3}
	fmt.Println(a.Field1) // f1
	fmt.Println(a.Field2) // f2
	fmt.Println(a.Field3) // f3
	fmt.Println(a.Inner)  // {f1 f2}

	b := Outer2{
		Inner: &Inner{
			Field1: "f1",
			Field2: "f2",
		},
		Field3: "f3",
	}
	fmt.Println(b)        // {0xc000062160 f3}
	fmt.Println(b.Field1) // f1
	fmt.Println(b.Field2) // f2
	fmt.Println(b.Field3) // f3
	fmt.Println(b.Inner)  // &{f1 f2}
}
