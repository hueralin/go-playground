package main

import "net/http"

func main05() {
	// 请求头的 Content-Type 决定了 Post 请求的数据格式
	// 表单 enctype 的默认值就是 application/x-www-form-urlencoded
	// 简单文本可以使用表单 URL 编码的方式发送
	// 大量数据，如上传文件，可以使用 multipart-data 的方式
	// 表单的 Get 请求，没有 Body，所有的数据都是通过 URL 编码的方式发送的
	// 表单数据在 Form, PostForm, MultipartForm 字段上，都是 map 类型
	// 通常的做法是先通过 ParseForm, ParseMultipartForm 方法解析 Request
	// 然后再访问相应的 Form, PostForm, MultipartForm 字段
	server := http.Server{Addr: "localhost:8888"}
	http.Handle("/", http.FileServer(http.Dir("www")))
	http.HandleFunc("/form", func(w http.ResponseWriter, r *http.Request) {
		//r.ParseForm()
		//r.ParseMultipartForm(1024)

		// r.Form 是 url.Values 类型，其本质是 map[string][]string
		// 其值是解析后的表单数据，包括 URL 字段的查询字符串，POST、PUT、PATCH 请求的表单数据
		//fmt.Fprintln(w, r.Form)         // map[city:[shijiazhuang] name:[wiwi]]
		//fmt.Fprintln(w, r.Form["name"]) // [wiwi]

		// Form 的使用会有一个问题，如果请求 URL 和表单里面有相同的 key，如 name，那么它俩的值都会被放进 r.Form
		// 其中表单里的值靠前，URL 中的值靠后，如果只想要表单里的 k-v，不想要 URL 里的，那就用 PostForm
		//fmt.Fprintln(w, r.Form)     // map[city:[shijiazhuang] name:[wiwi tom nike]] 多了请求 URL 的 name 值
		//fmt.Fprintln(w, r.PostForm) // map[city:[shijiazhuang] name:[wiwi]]

		// Form 和 PostForm 只支持 application/x-www-form-urlencoded
		// MultipartForm 支持 application/multipart-form
		// 先调用 ParseMultipartForm 方法，再访问 MultipartForm 字段
		// ParseMultipartForm 方法的参数是要读取的数据的长度
		// MultipartForm 只包含表单里的 k-v
		// MultipartForm 是个结构类型，包含两个 map，一个是类似于 Form/PostForm 的 map，一个是文件 map，没上传文件时是个空 map
		//fmt.Fprintln(w, r.MultipartForm) // &{map[city:[shijiazhuang] name:[wiwi]] map[]}

		// FormValue 和 PostFormValue 方法会返回指定 key 的第一个值，
		// 而且无需用户自己调用 ParseForm 和 ParseMultipartForm 方法，它会自己调用
		//fmt.Fprintln(w, r.FormValue("name")) // wiwi
	})
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
