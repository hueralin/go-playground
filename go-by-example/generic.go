package main

import "fmt"

// MapKeys 该函数接收一个 map 类型的参数，返回 key 的切片
// 该函数有两个类型参数（泛型），K 有可比较的限制（即这个类型的值可以使用 == 和 !=）
// V 没有任何限制，即 any（interface{} 的别名）
func MapKeys[K comparable, V any](m map[K]V) []K {
	s := make([]K, 0, len(m))
	for k := range m {
		s = append(s, k)
	}
	return s
}

// Node 结构的 val 值是任意类型
type Node[T any] struct {
	val  T
	next *Node[T]
}

type List[T any] struct {
	head *Node[T]
	tail *Node[T]
}

func (l *List[T]) Push(v T) {
	if l.tail == nil {
		l.head = &Node[T]{val: v}
		l.tail = l.head
	} else {
		l.tail.next = &Node[T]{val: v}
		l.tail = l.tail.next
	}
}

func (l *List[T]) GetAll() []T {
	var res []T
	for p := l.head; p != nil; p = p.next {
		res = append(res, p.val)
	}
	return res
}

func main() {
	// 注意：map 的遍历是无序的
	var p = map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println(MapKeys(p))
	var q = map[int]string{1: "a", 2: "b", 3: "c"}
	fmt.Println(MapKeys(q))

	l := List[int]{}
	fmt.Println(l.head, l.tail)
	l.Push(1)
	l.Push(2)
	l.Push(3)
	l.Push(4)
	l.Push(5)
	fmt.Println(l.GetAll()) // 1 2 3 4 5
}
