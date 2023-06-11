package main

// 指针

import (
	"fmt"
	"unsafe"
)

// 声明指针变量：var var_name *var-type
// 在指针类型前面加上 * 号（前缀）来获取指针所指向的内容

func test01(ptr *int) {
	// 修改了指向的变量的值
	*ptr = 999
	// 更改了 ptr 的指向，对之前指向的变量没有影响
	num := 0
	ptr = &num
}

// 参数是一个整型数组的指针
func changeArrByPtr(arr *[3]int) {
	(*arr)[0] = 1
	(*arr)[1] = 2
	(*arr)[2] = 3
}

func changeArrBySlice(slice []int) {
	slice[0] = 111
	slice[1] = 222
	slice[2] = 333
}

func main() {
	var num int = 100

	// 仅声明指针变量
	var ptrNum *int
	fmt.Printf("ptrNum's type is %T, size is %d, value is %p\n", ptrNum, unsafe.Sizeof(ptrNum), ptrNum)

	// 声明并初始化
	var ptrNum2 *int = &num
	fmt.Println(ptrNum2)
	fmt.Println(*ptrNum2)

	// 使用 new 分配内存并返回地址，然后赋值给指针变量
	ptrNum3 := new(int)
	ptrNum3 = &num
	fmt.Println(*ptrNum3)
	*ptrNum3 = 200
	fmt.Println(*ptrNum3)

	// & 操作符, 获取一个变量的地址
	// * 操作符, 如果出现在赋值操作符的左边(LHS)，则表示该指针指向的变量
	//		    如果出现在赋值操作符的右边(RHS)，则表示该指针指向的变量的值（即指针的解引用）

	str := "hello"
	num2 := 200
	f1 := 1.2
	b1 := true
	fmt.Printf("string ptr is %T\n", &str) // *string
	fmt.Printf("int ptr is %T\n", &num2)   // *int
	fmt.Printf("float ptr is %T\n", &f1)   // *float64
	fmt.Printf("bool ptr is %T\n", &b1)    // *bool

	var nilPtr *int
	fmt.Println("nilPtr is ", nilPtr) // nil

	// 函数参数是指针类型，函数内对指针指向的变量做的修改最终都会体现在那个变量上
	test01(&num)
	fmt.Println(num)

	// 指针与切片、指针与数组，指针和切片一样，都可以引用一个数组
	// 如果想通过函数改变数组中的值，可以将该数组的切片作为函数的参数，或者将该数组的指针（地址）传入参数
	var arr [3]int
	fmt.Println(arr) // 0 0 0
	changeArrByPtr(&arr)
	fmt.Println(arr) // 1 2 3
	changeArrBySlice(arr[:])
	fmt.Println(arr) // 111 222 333

	// Go 不支持指针运算，如 prt++
}
