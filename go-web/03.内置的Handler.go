package main

import (
	"net/http"
)

func main03() {
	// 内置的 Handler

	// func NotFoundHandler() Handler
	// 给每个请求都响应 "404 page not found"

	// func RedirectHandler(url string, code int) Handler
	// 把请求使用指定的状态码跳转到其他 URL

	// func StripPrefix(prefix string, h Handler) Handler
	// 从 URL 中去掉指定的前缀，再调用另一个 Handler h
	// 如果请求的 URL 和提供的前缀不符，则返回 404，略像中间件，修饰了另外一个 Handler

	// func TimeoutHandler(h Handler, dt time.Duration, msg string) Handler
	// time.Duration 是 int64 的别名，表示一段时间（两个时间之间的纳秒数）
	// 用来在指定时间内执行传入的 Handler，相当于一个修饰器
	// 即 h 的允许处理的时间，如果超时，就返回错误信息给请求

	// func FileServer(root FileSystem) Handler
	// 使用基于 root 的文件系统来响应请求
	//type FileSystem interface {
	//	Open(name string) (File, error)
	//}
	// 我们一般使用操作系统的文件系统，os.Dir 实现了上述接口
	//type Dir string
	//func (d Dir) Open(name string) (File, error)

	// 实现静态文件服务的方法一
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	http.ServeFile(w, r, "www"+r.URL.Path)
	//})
	//err := http.ListenAndServe("localhost:8888", nil)
	//if err != nil {
	//	return
	//}

	// 实现静态文件服务的方法二
	err := http.ListenAndServe("localhost:8888", http.FileServer(http.Dir("www")))
	if err != nil {
		return
	}
}
