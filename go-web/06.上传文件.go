package main

import (
	"fmt"
	"io"
	"net/http"
)

func main06() {
	// 文件上传，Content-Type 为 multipart/form-data
	// 先调用 ParseMultipart 方法解析表单
	// 再从 r.MultipartForm.File["xxx"][0] 获取 fileHeader
	// 再通过 fileHeader 的 Open 方法获得文件
	// 再通过 io.ReadAll 方法将文件内容读到 []byte 里面
	server := http.Server{Addr: "localhost:8888"}
	http.Handle("/", http.FileServer(http.Dir("www")))
	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		//r.ParseMultipartForm(5 * 1024)
		//fileHeader := r.MultipartForm.File["file"][0]
		//file, err := fileHeader.Open()
		//if err != nil {
		//	fmt.Fprintln(w, err.Error())
		//	return
		//}
		//data, err := io.ReadAll(file)
		//fmt.Fprintln(w, string(data))

		// 当然还有更方便的形式，FormFile 直接读取给定 field 的第一个文件
		// 更适合只上传一个文件的情况
		file, _, err := r.FormFile("file")
		if err != nil {
			fmt.Fprintln(w, err.Error())
			return
		}
		data, err := io.ReadAll(file)
		fmt.Fprintln(w, string(data))
	})
	err := server.ListenAndServe()
	if err != nil {
		return
	}

	// 并非所有的 POST 请求都来自 Form，还有可能来自 ajax、fetch，
	// 它们采用的 Content-Type 不一定是 x-www-form-urlencoded 或 multipart/form-data
	// 所以 Form, PostForm, MultipartForm 解析不了 JSON 请求
}
