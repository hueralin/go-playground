package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("len = %d, cap = %d, value = %v\n", len(s), cap(s), s)
}

func main() {
	a := make([]int, 0, 5) // len =
	printSlice(a)
	b := a[:2]
	// len = 2, cap = 5, value = [0 0]
	printSlice(b)
}
