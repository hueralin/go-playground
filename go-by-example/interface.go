package main

import "fmt"

// 接口是方法的集合

// 几何图形接口
type geometry interface {
	// 求面积的方法
	area2() float64
}

type rect2 struct {
	width, height float64
}

// rect2 实现了 geometry 接口
func (r rect2) area2() float64 {
	return r.width * r.height
}

type circle struct {
	radius float64
}

// circle 实现了 geometry 接口
func (c circle) area2() float64 {
	return c.radius * c.radius
}

// 如果一个变量是接口类型，那么我们就可以调用这个命名接口中的方法
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area2())
}

func main() {
	var r = rect2{width: 10, height: 10}
	var c = circle{radius: 4}

	measure(r) // {10 10} 100
	measure(c) // {4} 16
}
