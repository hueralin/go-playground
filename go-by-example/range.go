package main

import "fmt"

func main() {
	// range 用来遍历一组元素
	a := []int{1, 2, 3, 4, 5}
	for _, val := range a {
		fmt.Println(val)
	}

	b := map[string]string{"name": "tom", "addr": "USA"}
	for k, v := range b {
		fmt.Println(k, v)
	}

	// 遍历字符串，第一个值是字符在字符串中的索引，第二个值是字符在 Unicode 中的码
	for i, c := range "Hello, world" {
		fmt.Println(i, c)
	}
}
