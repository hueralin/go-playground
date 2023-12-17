package main

import (
	"fmt"
	"log"
	"os"
)

// WriteFile 使用 os.WriteFile 写入文件
func WriteFile() {
	data := []byte("Hello")
	// 如果文件不存在，则会创建
	err := os.WriteFile("data/file1", data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

// ReadFile 使用 os.ReadFile 读取文件
func ReadFile() {
	content, err := os.ReadFile("data/file1")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(content))
}

// CreateAndWriteFile 使用 os.Create 创建文件
func CreateAndWriteFile() {
	// 返回一个 File 结构
	file, err := os.Create("data/file2.txt")
	defer file.Close()
	if err != nil {
		log.Fatalln(err)
	}
	// 往文件里写入数据
	bytesLen, err := file.Write([]byte("haha"))
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(bytesLen)
}

// OpenFile 使用 os.Open 打开文件，然后读出数据
func OpenFile() {
	// 文件的内容是 "haha"
	file, err := os.Open("data/file2.txt")
	defer file.Close()
	if err != nil {
		log.Fatalln(err)
	}
	// 切片的容量和长度是 3
	v := make([]byte, 3)
	bytesLen, err := file.Read(v)
	if err != nil {
		log.Fatalln(err)
	}
	// 就读出来三个字节的数据： hah 3 3
	fmt.Println(string(v), bytesLen, len(v))
}
