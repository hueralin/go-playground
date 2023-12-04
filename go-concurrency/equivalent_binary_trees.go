package main

//import (
//	"fmt"
//	"golang.org/x/tour/tree"
//)
//
//// 练习题：https://go.dev/tour/concurrency/8
//
//func walk(t *tree.Tree, ch chan int) {
//	walkHelper(t, ch)
//	close(ch)
//}
//
//func walkHelper(t *tree.Tree, ch chan int) {
//	if t == nil {
//		return
//	}
//	if t.Left != nil {
//		walkHelper(t.Left, ch)
//	}
//	ch <- t.Value
//	if t.Right != nil {
//		walkHelper(t.Right, ch)
//	}
//}
//
//func same(t1, t2 *tree.Tree) bool {
//	t1Ch := make(chan int)
//	t2Ch := make(chan int)
//	go walk(t1, t1Ch)
//	go walk(t2, t2Ch)
//
//	for {
//		// ok 表示取出这次值之后，通道的状态（是否被关闭）
//		a, aOk := <-t1Ch
//		b, bOk := <-t2Ch
//
//		// a、b 不相等
//		// 或者相等后，通道的状态不一致，有的被关闭了，说明树的节点数量不一样，即两棵树不相等
//		if a != b || aOk != bOk {
//			return false
//		}
//		// 走到这，说明 a == b, aOk == bOk
//		// 如果 t1Ch 被关闭，则 t2Ch 也被关闭，说明两棵树都遍历完了，于是退出循环
//		if !aOk {
//			break
//		}
//	}
//	return true
//}
//
//func ShowEquivalentBinaryTrees() {
//	fmt.Println(same(tree.New(1), tree.New(1))) // true
//	fmt.Println(same(tree.New(1), tree.New(2))) // false
//}
