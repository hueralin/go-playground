package main

import (
	"errors"
	"fmt"
)

// error 一般作为函数的最后一个返回值返回
func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

// 一个类型实现了 Error 方法，那么可以被当作错误类型来使用
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with is"}
	}
	return arg + 3, nil
}

func main() {
	for _, i := range []int{7, 42} {
		r, e := f1(i)
		if e == nil {
			fmt.Println("f1 worked: ", r)
		} else {
			fmt.Println("f1 failed: ", e)
		}
	}

	for _, i := range []int{7, 42} {
		r, e := f2(i)
		if e == nil {
			fmt.Println("f1 worked: ", r)
		} else {
			fmt.Println("f1 failed: ", e)
		}
	}

	_, e := f2(42)
	// 类型断言，又称
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
