package main

import (
	"go-web/middlewares"
	"net/http"
)

func main11() {
	// 中间件，能对请求做处理，也能对响应做处理
	// 应用场景：日志、安全、请求超时、响应压缩...
	server := http.Server{Addr: "localhost:8888", Handler: new(middlewares.AuthMw)}
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
