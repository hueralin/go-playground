package main

import (
	"net/http"
)

//// 实现了 ServeHTTP 方法就是一个 handler
//type myHandler struct{}
//
//func (mh myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	w.Write([]byte("Hello world"))
//}
//
//type homeHandler struct{}
//
//func (h homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	w.Write([]byte("Home"))
//}
//
//type aboutHandler struct{}
//
//func (h aboutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	w.Write([]byte("About"))
//}

func main02() {
	// 使用自定义的 Handler
	//mh := myHandler{}
	//server := http.Server{Addr: "localhost:8888", Handler: mh}
	//err := server.ListenAndServe()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	// 添加多个 Handler 的方法一
	// DefaultServeMux 是个多路复用器，可往里面注册多个 handler
	//hh := homeHandler{}
	//ah := aboutHandler{}
	//// 使用 DefaultServeMux
	//server := http.Server{Addr: "localhost:8888", Handler: nil}
	//// 使用 http.Handle 向 DefaultServeMux 注册 Handler
	//http.Handle("/", hh)
	//http.Handle("/about", ah)
	//err := server.ListenAndServe()
	//if err != nil {
	//	return
	//}

	// 添加多个 Handler 的方法二
	// Handler 函数就是与 Handler 行为类似的函数，Handler 本身是一个实现了 ServeHTTP 方法的任意类型
	// Handler 函数的签名和 ServeHTTP 方法的签名一样
	// HandleFunc 将具有适当签名的函数 f 适配为 Handler
	// 源码中有个 HandlerFunc 实际上是个函数类型，签名和 ServeHTTP 一样，并且实现了 ServeHTTP 方法
	// 源码中通过使用 HandlerFunc(f)，将用户传递的函数 f 适配成一个 Handler
	// 这不是函数调用，而是一次类型转换，就像 []byte("hello") 一样
	// 然后调用 http.Handle 方法注册 Handler
	server := http.Server{Addr: "localhost:8888", Handler: nil}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Home"))
	})
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("About"))
	})
	// 我们使用 http.HandlerFunc 将函数适配成 Handler，然后再注册
	http.Handle("/welcome", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome"))
	}))
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
