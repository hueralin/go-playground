package main

import (
	"context"
	"fmt"
	"net/http"
)

func main12() {
	// 请求上下文
	// func (r *Request) Context() context.Context 返回当前请求的上下文
	// func (r *Request) WithContext(ctx context.Context, k, v) context.Context 扩展上下文（上下文本身不可修改）
	server := http.Server{Addr: ":8888"}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 返回当前请求的上下文
		ctx := r.Context()
		// 在上下文中存储一些该请求范围内的数据（请求上下文是不可修改的，所以每次都会创建新的）
		ctx = context.WithValue(ctx, "k1", "v1")
		ctx = context.WithValue(ctx, "k2", "v2")
		ctx = context.WithValue(ctx, "k3", "v3")
		// 读取指定 key 的 value
		_, err := fmt.Fprintln(w, ctx.Value("k1"))
		if err != nil {
			return
		}
	})
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
