package main

import (
	"fmt"
	"net/http"
)

func main04() {
	// HTTP 消息：请求和响应
	// Request 是个 struct，包含 URL, Header, Body, Form, PostForm, MultipartForm
	//type URL struct {
	//	Schema   string
	//	Opaque   string
	//	User     *Userinfo
	//	Host     string
	//	Path     string
	//	RawQuery string 即 name=tom&age=18
	//	Fragment string
	//}
	// 从浏览器发出的请求不会携带 Fragment，但其他的 HTTP 客户端可能会

	// Header 是一个 map, map[string][]string
	// 设置 key 时会创建一个空的 []string 作为 value
	// 获取 Header: r.Header，返回一个 map
	// 获取指定的 Header: r.Header["Content-Type"]，返回一个 []string
	// 获取指定的 Header: r.Header.Get("Content-Type")，返回一个字符串，多个值用 ", " 分隔
	//server := http.Server{Addr: "localhost:8888", Handler: nil}
	//http.HandleFunc("/header", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintln(w, r.Header)
	//	fmt.Fprintln(w, r.Header["Accept-Encoding"])
	//	fmt.Fprintln(w, r.Header.Get("Accept-Encoding"))
	//})
	//err := server.ListenAndServe()
	//if err != nil {
	//	return
	//}

	// 请求和响应的消息体都是 Body 字段
	// Body 实现了 io.ReadCloser 接口
	// 实际上是两个接口：Reader、Closer
	// Reader 接口定义了 Read 方法，参数是 []byte，返回值是 byte 的数量和可选的错误
	// Closer 接口定义了 Close 方法，没有参数，返回可选的错误
	//server := http.Server{Addr: "localhost:8888"}
	//http.HandleFunc("/body", func(w http.ResponseWriter, r *http.Request) {
	//	// 根据请求内容的长度创建切片
	//	body := make([]byte, r.ContentLength)
	//	// 将请求内容读入 body 切片
	//	r.Body.Read(body)
	//	fmt.Fprintf(w, "res: %v\n", string(body))
	//})
	//err := server.ListenAndServe()
	//if err != nil {
	//	return
	//}

	// r.URL.RawQuery 返回原始 query 参数
	// r.URL.Query() 返回 map[string][]string
	server := http.Server{Addr: "localhost:8888"}
	http.HandleFunc("/query", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		// 返回 id 切片数组
		id := query["id"]
		// 返回 city 切片数组的第一个值
		city := query.Get("city")
		fmt.Fprintln(w, r.URL.RawQuery) // id=123&id=456&city=beijing&city=nanjing
		fmt.Fprintln(w, id)             // [123 456]
		fmt.Fprintln(w, city)           // beijing
	})
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
