package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// 在 Go 中，字符串是只读的字节切片 []byte，是 UTF8 编码的文本容器
	// 在其他语言中，字符串被认为是由字符组成的
	// 在 Go 中，字符的概念被称为 “符文”，是代表 Unicode 代码点的整数

	const s = "สวัสดี" // 泰文
	// 由于字符串被认为是 []byte，所以 len 打印的是字节数量，而不是字符数量
	fmt.Println("Len:", len(s)) // 18 个字节

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i]) // 打印十六进制的字节码，即 for 循环不会解码
	}
	fmt.Println()

	// 使用 utf8 package 提供的方法获取字符数量
	fmt.Println("Rune count:", utf8.RuneCountInString(s)) // 6 个字符

	// for range 可以解码
	for i, runeVal := range s {
		fmt.Printf("%d, %#U\n", i, runeVal)
	}

	// 使用 utf8 package 提供的方法解码字符
	const s2 = "hello, 世界"
	for j, w := 0, 0; j < len(s2); j += w {
		// https://pkg.go.dev/unicode/utf8#DecodeRune
		// https://pkg.go.dev/unicode/utf8#DecodeRuneInString
		runeVal, width := utf8.DecodeRuneInString(s2[j:])
		fmt.Printf("%d, %#U, %d\n", j, runeVal, width) // 索引, 字符, 该字符的字节大小（宽度）
		w = width
	}
}
