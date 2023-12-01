package main

import "net/http"

func main07() {
	// ResponseWriter 是一个接口，幕后的 struct 是非导出的 http.response
	//type ResponseWriter interface {
	//	Header()
	//	Write()
	//	WriteHeader()
	//}
	// http.response 指针实现了这个接口，所以 ServeHTTP 的 w 也是个指针
	// Write 方法，参数为 []byte，写入到 HTTP 响应的 Body 里面
	// 如果在 Write 方法被调用时还没设置 Content-Type，那么数据的前 512 字节会被用来检测（推断）Content-Type
	// WriteHeader(code) 用来设置状态码，如果没有被显示设置，那么在第一次调用 Write 时会隐式调用 WriteHeader(http.StatusOK)
	// 调用完 WriteHeader 后仍可以写入到响应，但是不能再修改状态码了
	// Header 方法返回一个 http.Header 类型的数据（map[string][]string），调用其 Set 方法可以设置响应头
	// 注意一定要在 WriteHeader 方法之前设置响应头，之后不允许修改了
	// w.Header().Set(key, value)
	server := http.Server{Addr: "localhost:8888"}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()
		header.Set("Location", "https://baidu.com")
		w.WriteHeader(http.StatusTemporaryRedirect)
	})
	err := server.ListenAndServe()
	if err != nil {
		return
	}

	// 内置的 Response
	// NotFound 函数，包装一个 404 状态码和一个额外的信息ServeFile 函数，从文件系统提供文件，返回给请求者
	// ServeContent 函数，它可以把实现了 io.ReadSeeker 接口的任何东西里面的内容返回给请求者
	// 还可以处理 Range 请求 (范围请求) ，如果只请求了资源的一部分内容，那么ServeContent 就可以如此响应。而 ServeFile 或 io.Copy 则不行。
	// Redirect 函数，告诉客户端重定向到另一个 URL
}
