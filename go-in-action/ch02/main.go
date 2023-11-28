package main

import (
	_ "go-playground/go-in-action/ch02/matchers"
	"go-playground/go-in-action/ch02/search"
	"log"
	"os"
)

// 在 main 之前调用
func init() {
	// 将日志输出到标准输出
	log.SetOutput(os.Stdout)
}

func main() {
	// 使用特定的项做搜索
	search.Run("president")
}
