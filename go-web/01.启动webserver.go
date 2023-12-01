package main

import (
	"net/http"
)

func main01() {
	// 创建 HTTP WebServer 方式一
	//err := http.ListenAndServe("localhost:8888", nil)
	//if err != nil {
	//	return
	//}

	// 创建 HTTP WebServer 方式二
	// 自己创建 server，Handler 为 nil，则表明使用的是 DefaultServeMux
	server := http.Server{
		Addr:    "localhost:8888",
		Handler: nil,
	}
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
