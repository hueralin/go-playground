package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	fs := http.FileServer(http.Dir("public"))

	http.Handle("/public", http.StripPrefix("/public", fs))

	http.HandleFunc("/upload", upload)

	log.Println("Running at http://localhost:8888")
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		return
	}
}

func upload(w http.ResponseWriter, r *http.Request) {
	var uploadPath = "uploads/"
	log.Printf("Upload Path: %v\n", uploadPath)
	log.Printf("Request Method: %v\n", r.Method)
	if r.Method == "POST" {
		// 将请求体按照 multipart/form-data 的形式解析，
		// 并设置文件最大存储内存，即这个范围内的文件会被存储在内存中，剩余部分会被存储在磁盘上的临时文件中
		err := r.ParseMultipartForm(1 << 20)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// 获取文件句柄和文件头信息
		file, fileHandler, err := r.FormFile("file")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()

		// 创建目标文件
		dst, err := os.Create(uploadPath + fileHandler.Filename)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		// 将上传的文件内容拷贝到目标文件
		_, err = io.Copy(dst, file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "文件上传成功：%s\n", fileHandler.Filename)
	} else {
		//
	}
}

func logHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 打印请求信息
		log.Printf("请求：%s %s", r.Method, r.URL.Path)

		// 调用下一个处理器
		handler.ServeHTTP(w, r)
	})
}
