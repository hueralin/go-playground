package main

import (
	"fmt"
	"math"
	"unsafe"
)

// Integer 有符号整型
func Integer() {
	var num8 int8 = 127
	var num16 int16 = 32767
	var num32 int32 = math.MaxInt32
	var num64 int64 = math.MaxInt64
	var num int = math.MaxInt

	fmt.Printf("num8's type is %T, num8's size is %d, num8's value is %d\n", num8, unsafe.Sizeof(num8), num8)
	fmt.Printf("num16's type is %T, num16's size is %d, num16's value is %d\n", num16, unsafe.Sizeof(num16), num16)
	fmt.Printf("num32's type is %T, num32's size is %d, num32's value is %d\n", num32, unsafe.Sizeof(num32), num32)
	fmt.Printf("num64's type is %T, num64's size is %d, num64's value is %d\n", num64, unsafe.Sizeof(num64), num64)
	fmt.Printf("num's type is %T, num's size is %d, num's value is %d\n", num, unsafe.Sizeof(num), num)
}

// UnsignedInteger 无符号整型
func UnsignedInteger() {
	var num8 uint8 = 128
	var num16 uint16 = 32768
	var num32 uint32 = math.MaxUint32
	var num64 uint64 = math.MaxUint64
	var num uint = math.MaxUint

	fmt.Printf("num8's type is %T, num8's size is %d, num8's value is %d\n", num8, unsafe.Sizeof(num8), num8)
	fmt.Printf("num16's type is %T, num16's size is %d, num16's value is %d\n", num16, unsafe.Sizeof(num16), num16)
	fmt.Printf("num32's type is %T, num32's size is %d, num32's value is %d\n", num32, unsafe.Sizeof(num32), num32)
	fmt.Printf("num64's type is %T, num64's size is %d, num64's value is %d\n", num64, unsafe.Sizeof(num64), num64)
	fmt.Printf("num's type is %T, num's size is %d, num's value is %d\n", num, unsafe.Sizeof(num), num)
}

// Float 浮点型
func Float() {
	// 小数点后 6 位
	var num32 float32 = math.MaxFloat32
	// 小数点后 15 位
	var num64 float64 = math.MaxFloat64

	fmt.Printf("num32's type is %T, num32's size is %d, num32's value is %g\n", num32, unsafe.Sizeof(num32), num32)
	fmt.Printf("num64's type is %T, num64's size is %d, num64's value is %g\n", num64, unsafe.Sizeof(num64), num64)
}

func main() {
	// 一般使用 int 表示整型，在 32 位系统上就是 int32，在 64 位系统上就是 int64
	Integer()
	fmt.Println("------------------")
	UnsignedInteger()
	fmt.Println("------------------")
	Float()
}
