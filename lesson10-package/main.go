package main

// 引入包，包名冲突时可以起别名
import (
	"fmt"
	"go-playground/lesson09-func/book"
)

// 包
// Go 语言有超过 100 个的标准包，可以用 go list std | wc -l 命令查看标准包的具体数目，标准库为大多数的程序提供了必要的基础组件

// 每个包都允许有一个或多个 init 函数
// init 函数不应该有任何参数和返回值
// 在代码中也不能显式调用它
// 当这个包被导入时，就会执行这个包的 init 函数，做初始化任务
// init 函数优先于 main 函数执行

func init() {
	fmt.Println("main init")
}

// 之前说过, 导入一个没有使用的包编译会报错, 但有时候我们只是想执行包里的 init 函数来执行一些初始化任务的话应该怎么办呢?
// 我们可以使用匿名导入的方法, 使用空白标识符(Blank Identifier)
// import _ "fmt"
// 由于导入时会执行该包里的 init 函数, 所以编译仍会将此包编译到可执行文件中

// main 函数是程序运行的入口
func main() {
	// 使用包内的函数
	bookInfo, err := book.ShowBookInfo("Go Go Go!", "Jack")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(bookInfo)
	}
}
