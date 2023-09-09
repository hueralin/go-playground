package main

import (
	"fmt"
	"maps"
)

func main() {
	// map[key-type]value-type
	var a map[string]int
	fmt.Println(a, a == nil) // map[] true

	b := make(map[string]int)
	fmt.Println(b, b == nil) // map[] false

	b["height"] = 170
	b["age"] = 18
	fmt.Println(b) // map[age:18 height:170]

	// 如果 key 不存在，那么会返回零值
	fmt.Println(b["xxx"]) // 0

	// len 获取 map 中的键值对个数
	fmt.Println(len(b)) // 2

	// delete 删除键值对
	delete(b, "age")
	fmt.Println(b) // map[height:170]

	// clear 清空 map
	clear(b)
	fmt.Println(b) // map[]

	// 读取操作不光会返回值，还会返回该 key 是否存在，作为第二个返回值
	// 这用于消除一些零值如 0 或 "" 或 false 带来的歧义
	// 这里不需要值，可以使用空白符（_）省略掉
	_, isExisted := b["xxx"]
	fmt.Println(isExisted) // false

	// 声明并初始化
	c := map[string]string{"name": "tom", "age": "18"}
	fmt.Println(c) // map[age:18 name:tom]

	// maps package 中有一些实用的方法
	// maps.Equal 用来判断两个 map 的内容是否相等
	d := map[string]string{"name": "tom", "age": "18", "addr": "USA"}
	fmt.Println(maps.Equal(c, d)) // false
}
