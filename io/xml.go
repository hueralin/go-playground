package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Person struct {
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age"`
}

func main() {
	// 获取当前目录
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("无法获取当前目录:", err)
		return
	}

	// 打开 XML 文件
	//file, err := os.Open("person.xml")
	//if err != nil {
	//	fmt.Printf("打开文件失败: %v\n", err)
	//	return
	//}
	//defer file.Close()

	// 读取 XML 文件内容
	content, err := os.ReadFile(dir + "/io/person.xml")
	if err != nil {
		fmt.Printf("读取文件失败: %v\n", err)
		return
	}

	var person Person

	err = xml.Unmarshal(content, &person)
	if err != nil {
		fmt.Printf("解析 XML 失败: %v\n", err)
		return
	}

	// 输出解析后的数据
	fmt.Printf("Name: %v\n", person.Name)
	fmt.Printf("Age: %v\n", person.Age)
}
