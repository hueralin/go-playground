package main

import (
	"go-web/middlewares"
	"net/http"
)

// 实现了 ServeHTTP 方法就是一个 handler
type myHandler struct{}

func (mh myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}

type homeHandler struct{}

func (h homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home"))
}

type aboutHandler struct{}

func (h aboutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About"))
}

type Company struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
}

func main() {
	// 创建 HTTP WebServer 方式一
	//err := http.ListenAndServe("localhost:8888", nil)
	//if err != nil {
	//	return
	//}

	// 创建 HTTP WebServer 方式二
	//server := http.Server{Addr: "localhost:8888", Handler: nil} // Handler 为 nil，则表明使用的是 DefaultServeMux
	//err := server.ListenAndServe()
	//if err != nil {
	//	return
	//}

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
	//server := http.Server{Addr: "localhost:8888", Handler: nil}
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte("Home"))
	//})
	//http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte("About"))
	//})
	//// 我们使用 http.HandlerFunc 将函数适配成 Handler，然后再注册
	//http.Handle("/welcome", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte("Welcome"))
	//}))
	//err := server.ListenAndServe()
	//if err != nil {
	//	return
	//}

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
	//err := http.ListenAndServe("localhost:8888", http.FileServer(http.Dir("www")))
	//if err != nil {
	//	return
	//}

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
	//server := http.Server{Addr: "localhost:8888"}
	//http.HandleFunc("/query", func(w http.ResponseWriter, r *http.Request) {
	//	query := r.URL.Query()
	//	// 返回 id 切片数组
	//	id := query["id"]
	//	// 返回 city 切片数组的第一个值
	//	city := query.Get("city")
	//	fmt.Fprintln(w, r.URL.RawQuery) // id=123&id=456&city=beijing&city=nanjing
	//	fmt.Fprintln(w, id)             // [123 456]
	//	fmt.Fprintln(w, city)           // beijing
	//})
	//err := server.ListenAndServe()
	//if err != nil {
	//	return
	//}

	// 请求头的 Content-Type 决定了 Post 请求的数据格式
	// 表单 enctype 的默认值就是 application/x-www-form-urlencoded
	// 简单文本可以使用表单 URL 编码的方式发送
	// 大量数据，如上传文件，可以使用 multipart-data 的方式
	// 表单的 Get 请求，没有 Body，所有的数据都是通过 URL 编码的方式发送的
	// 表单数据在 Form, PostForm, MultipartForm 字段上，都是 map 类型
	// 通常的做法是先通过 ParseForm, ParseMultipartForm 方法解析 Request
	// 然后再访问相应的 Form, PostForm, MultipartForm 字段
	//server := http.Server{Addr: "localhost:8888"}
	//http.Handle("/", http.FileServer(http.Dir("www")))
	//http.HandleFunc("/form", func(w http.ResponseWriter, r *http.Request) {
	//	//r.ParseForm()
	//	// 参数是指一次性加载到内存的最大字节数
	//	//r.ParseMultipartForm(1024)
	//
	//	// r.Form 是 url.Values 类型，其本质是 map[string][]string
	//	// 其值是解析后的表单数据，包括 URL 字段的查询字符串，POST、PUT、PATCH 请求的表单数据
	//	//fmt.Fprintln(w, r.Form)         // map[city:[shijiazhuang] name:[wiwi]]
	//	//fmt.Fprintln(w, r.Form["name"]) // [wiwi]
	//
	//	// Form 的使用会有一个问题，如果请求 URL 和表单里面有相同的 key，如 name，那么它俩的值都会被放进 r.Form
	//	// 其中表单里的值靠前，URL 中的值靠后，如果只想要表单里的 k-v，不想要 URL 里的，那就用 PostForm
	//	//fmt.Fprintln(w, r.Form)     // map[city:[shijiazhuang] name:[wiwi tom nike]] 多了请求 URL 的 name 值
	//	//fmt.Fprintln(w, r.PostForm) // map[city:[shijiazhuang] name:[wiwi]]
	//
	//	// Form 和 PostForm 只支持 application/x-www-form-urlencoded
	//	// MultipartForm 支持 application/multipart-form
	//	// 先调用 ParseMultipartForm 方法，再访问 MultipartForm 字段
	//	// ParseMultipartForm 方法的参数是要读取的数据的长度
	//	// MultipartForm 只包含表单里的 k-v
	//	// MultipartForm 是个结构类型，包含两个 map，一个是类似于 Form/PostForm 的 map，一个是文件 map，没上传文件时是个空 map
	//	//fmt.Fprintln(w, r.MultipartForm) // &{map[city:[shijiazhuang] name:[wiwi]] map[]}
	//
	//	// FormValue 和 PostFormValue 方法会返回指定 key 的第一个值，
	//	// 而且无需用户自己调用 ParseForm 和 ParseMultipartForm 方法，它会自己调用
	//	//fmt.Fprintln(w, r.FormValue("name")) // wiwi
	//})
	//err := server.ListenAndServe()
	//if err != nil {
	//	return
	//}

	// 文件上传，Content-Type 为 multipart/form-data
	// 先调用 ParseMultipart 方法解析表单
	// 再从 r.MultipartForm.File["xxx"][0] 获取 fileHeader
	// 再通过 fileHeader 的 Open 方法获得文件
	// 再通过 io.ReadAll 方法将文件内容读到 []byte 里面
	//server := http.Server{Addr: "localhost:8888"}
	//http.Handle("/", http.FileServer(http.Dir("www")))
	//http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
	//	//r.ParseMultipartForm(5 * 1024)
	//	//fileHeader := r.MultipartForm.File["file"][0]
	//	//file, err := fileHeader.Open()
	//	//if err != nil {
	//	//	fmt.Fprintln(w, err.Error())
	//	//	return
	//	//}
	//	//data, err := io.ReadAll(file)
	//	//fmt.Fprintln(w, string(data))
	//
	//	// 当然还有更方便的形式，FormFile 直接读取给定 field 的第一个文件
	//	// 更适合只上传一个文件的情况
	//	file, _, err := r.FormFile("file")
	//	if err != nil {
	//		fmt.Fprintln(w, err.Error())
	//		return
	//	}
	//	data, err := io.ReadAll(file)
	//	fmt.Fprintln(w, string(data))
	//})
	//err := server.ListenAndServe()
	//if err != nil {
	//	return
	//}

	// 并非所有的 POST 请求都来自 Form，还有可能来自 ajax、fetch，
	// 它们采用的 Content-Type 不一定是 x-www-form-urlencoded 或 multipart/form-data
	// 所以 Form, PostForm, MultipartForm 解析不了 JSON 请求

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
	//server := http.Server{Addr: "localhost:8888"}
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	header := w.Header()
	//	header.Set("Location", "https://baidu.com")
	//	w.WriteHeader(http.StatusTemporaryRedirect)
	//})
	//err := server.ListenAndServe()
	//if err != nil {
	//	return
	//}

	// 内置的 Response
	// NotFound 函数，包装一个 404 状态码和一个额外的信息ServeFile 函数，从文件系统提供文件，返回给请求者
	// ServeContent 函数，它可以把实现了 io.ReadSeeker 接口的任何东西里面的内容返回给请求者
	// 还可以处理 Range 请求 (范围请求) ，如果只请求了资源的一部分内容，那么ServeContent 就可以如此响应。而 ServeFile 或 io.Copy 则不行。
	// Redirect 函数，告诉客户端重定向到另一个 URL

	// TODO: 数据库操作
	// Open 方法并不会连接到数据库，也不会验证参数，它只是把后续连接到数据库所需的 struct 设置好了
	// sql.DB 是用来处理数据库的，并不是真正的连接，它维护数据库连接池
	// sql.DB 可以全局使用，也可以传递到函数中

	//// Go 内置路由并不强大，很多功能都需要自己实现
	//// 有些第三方的实现，如 gorilla/mux, httprouter
	//server := http.Server{Addr: "localhost:8888"}
	//handlers.RegisterHandlers()
	//err := server.ListenAndServe()
	//if err != nil {
	//	return
	//}

	// JSON
	// 类型映射
	// Go float64: JSON 数值
	// Go string: JSON 字符串
	// Go bool: JSON 布尔值
	// Go nil: JSON null
	// 在 Go 结构体上使用 Tag 的形式，将 JSON 的数据映射到结构体上
	// 对于未知结构的 JSON，可以使用以下两种方式
	// map[string]interface{}: 存储任意 JSON 对象
	// []interface{}: 存储任意 JSON 数组
	// 读取 JSON:
	// 先创建一个解码器 dec := json.NewDecoder(r.Body)，NewDecoder 的参数是一个 Reader 接口值（表明可从里面读数据）
	// 再使用解码器将 JSON 内容输出到变量 dec.decode(&xxx)
	// 写入 JSON:
	// 先创建一个编码器 enc := json.NewEncoder(w), NewEncoder 的参数是一个 Writer 接口值（表明可往里面写数据）
	// 再使用编码器将数据变成 JSON 格式写入 w
	//server := http.Server{Addr: "localhost:8888"}
	//http.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
	//	switch r.Method {
	//	case http.MethodPost:
	//		company := Company{}
	//		dec := json.NewDecoder(r.Body)
	//		err := dec.Decode(&company)
	//		if err != nil {
	//			log.Println(err.Error())
	//			w.WriteHeader(http.StatusInternalServerError)
	//			return
	//		}
	//
	//		enc := json.NewEncoder(w)
	//		err = enc.Encode(company)
	//		if err != nil {
	//			log.Println(err.Error())
	//			w.WriteHeader(http.StatusInternalServerError)
	//			return
	//		}
	//	default:
	//		w.WriteHeader(http.StatusMethodNotAllowed)
	//	}
	//})
	//err := server.ListenAndServe()
	//if err != nil {
	//	return
	//}

	//server := http.Server{Addr: "localhost:8888"}
	//http.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
	//	jsonStr := `{
	//		"id": 456,
	//		"name": "Oracle",
	//		"country": "USA"
	//	}`
	//	c := Company{}
	//	// 将 JSON 字符串解码到 c
	//	err := json.Unmarshal([]byte(jsonStr), &c)
	//	if err != nil {
	//		return
	//	}
	//
	//	// 将 c 编码为字节数组
	//	bytes, err := json.Marshal(c)
	//	if err != nil {
	//		return
	//	}
	//	w.Write(bytes)
	//})
	//err := server.ListenAndServe()
	//if err != nil {
	//	return
	//}

	// 第二种方式是使用 Marshal（编码） 和 Unmarshal（解码）
	// Marshal：把 go struct 转为 json 格式
	// Unmarshal：把 json 格式转化为 go struct

	// 两种方式的区别
	// 针对 string 和 byte，使用 Marshal 和 Unmarshal
	// Marshal => string
	// Unmarshal <= string
	// 针对 stream，使用 Encoder 和 Decoder
	// Encode => Stream，把数据写入 io.Writer
	// Decoder <= Stream，从 io.Reader 读取数据

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
