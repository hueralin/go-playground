package main

import (
	"fmt"
	"go-calc/simplemath"
	"os"
	"strconv"
)

// Usage 打印程序使用指南
func Usage() {
	fmt.Println("USAGE: calc command [arguments] ...")
	fmt.Println("\nThe commands are:\n\tadd\t计算两个数值相加\n\tsqrt\t计算一个非负数的平方根")
}

func main() {
	// 获取命令行参数，注意程序名本身是第一个参数
	// 例如 calc add 1 2，calc 是第一个参数
	args := os.Args
	// 除程序名本身外，至少需要传入两个其它参数，否则退出
	if args == nil || len(args) < 3 {
		Usage()
		return
	}

	// 第二个参数表示计算方法
	switch args[1] {
	case "add":
		if len(args) < 4 {
			fmt.Println("USAGE: calc add <integer1> <integer2>")
			return
		}
		v1, err1 := strconv.Atoi(args[2])
		v2, err2 := strconv.Atoi(args[3])
		if err1 != nil || err2 != nil {
			fmt.Println("USAGE: calc add <integer1> <integer2>")
			return
		}
		ret := simplemath.Add(v1, v2)
		fmt.Println("Result: ", ret)
	case "sqrt":
		if len(args) != 3 {
			fmt.Println("USAGE: calc sqrt <integer1>")
			return
		}
		v, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("USAGE: calc sqrt <integer1>")
			return
		}
		ret := simplemath.Sqrt(v)
		fmt.Println("Result: ", ret)
	default:
		Usage()
	}
}
