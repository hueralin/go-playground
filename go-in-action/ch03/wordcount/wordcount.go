package main

import (
	"fmt"
	"go-playground/go-in-action/ch03/words"
	"os"
)

func main() {
	// 构建后的程序是一个可执行文件(命令)
	// 执行时，可执行文件名作为 os.Args 的第 0 个值
	// 其余的参数按顺序被放在 os.Args 的其他索引上
	filename := os.Args[1]

	contents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	text := string(contents)
	count := words.CountWords(text)
	fmt.Printf("There are %d words in your text.\n", count)
}
