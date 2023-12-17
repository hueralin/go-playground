package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"time"
)

type Post struct {
	User    string
	Threads []string
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		t, err := template.ParseFiles("templates/login.html")
		if err != nil {
			log.Println(err.Error())
			return
		}
		err = t.Execute(w, nil)
		if err != nil {
			log.Println(err.Error())
			return
		}
	} else {
		// ParseForm 会解析 URL 上的查询字符串和请求体数据，如果有冲突的 key，则请求体中的在前
		// ParseForm 后, r.Form 才可用
		//err := r.ParseForm()
		//err := r.ParseMultipartForm(1024)
		//if err != nil {
		//	log.Println(err.Error())
		//	return
		//}
		//fmt.Println(r.Form) // map[password:[123] username:[tom xxx]]
		// FormValue 会自动调用 ParseForm，且只会返回对应 key 的 value 中的第一个，如果不存在，则返回空字符串
		//fmt.Println("username: ", r.FormValue("username"))
		//fmt.Println("password: ", r.FormValue("password"))
		//fmt.Println("username: ", r.Form["username"]) // [tom xxx]
		//fmt.Println("password: ", r.Form["password"]) // [123]
		//fmt.Println(r.MultipartForm)

		// 测试 FormValue 和 PostFormValue
		body4(w, r)
	}
}

func body(w http.ResponseWriter, r *http.Request) {
	// get the length of request body
	length := r.ContentLength
	fmt.Println(length)
	// create a byte slice
	_body := make([]byte, length)
	// read from r.Body to _body
	r.Body.Read(_body)
	// write _body to writer
	_, err := fmt.Fprintln(w, string(_body))
	if err != nil {
		log.Println("write failed: ", err)
		return
	}
}

func body2(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// 包含 URL 键值对和表单键值对
	fmt.Println(r.Form)
	// 仅包含表单键值对
	fmt.Println(r.PostForm)
}

func body3(w http.ResponseWriter, r *http.Request) {
	// 参数表示想要从 multipart 编码的表单里面取出多少字节的数据
	r.ParseMultipartForm(1024)
	// 仅包含表单键值对
	// &{map[password:[123] username:[tom]] map[]}
	// 一个包含了两个映射的结构，其中第一个映射的键为字符串，值为字符串组成的切片，
	// 而第二个映射则是空的, 这个映射之所以会为空, 是因为它是用来记录用户上传的文件的
	fmt.Println(r.MultipartForm)
}

func body4(w http.ResponseWriter, r *http.Request) {
	// 假设 Content-Type: multipart/form-data
	// 请求 URL 为 /login?username=urltom&password=urlpwd
	// 表单 form 为 username=tom, password 123
	//fmt.Println(r.FormValue("username")) // urltom
	fmt.Println(r.PostFormValue("username")) // "tom"
	fmt.Println(r.Form)                      // map{username:[urltom tom] password:[urlpwd 123]}
	fmt.Println(r.PostForm)                  // map{username: [tom], password: [123]}
	fmt.Println(r.MultipartForm)             // &{map{username:[tom] password:[urlpwd]} map[]}

	// r.MultipartForm 能打印值，说明 FormValue or PostFormValue 会按需调用 ParseMultipartForm
	// r.Form 保存了 URL 上的键值对
	// r.PostForm 保存了 urlencoded 表单上的键值对，但是本次请求是 form-data 表单，所以为空 map
}

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		t, err := template.ParseFiles("templates/upload.html")
		if err != nil {
			log.Println(err.Error())
			return
		}
		err = t.Execute(w, nil)
		if err != nil {
			log.Println(err.Error())
			return
		}
	} else {
		//r.ParseMultipartForm(32 << 20)
		//fileHeader := r.MultipartForm.File["secret"][0]
		//file, err := fileHeader.Open()
		//if err != nil {
		//	log.Println(err)
		//	return
		//}
		//b, err := io.ReadAll(file)
		//if err != nil {
		//	log.Println(err)
		//	return
		//}
		//fmt.Fprintln(w, string(b))

		// 更方便的做法
		file, _, err := r.FormFile("secret")
		if err != nil {
			log.Println(err)
			return
		}
		b, err := io.ReadAll(file)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Fprintln(w, string(b))
	}
}

func write(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("hello"))
	// 多余的 WriterHeader 调用, 因为 Write 默认设置了 200
	// http: superfluous response.WriteHeader call from main.write (form.go:142)
	//w.WriteHeader(http.StatusNotImplemented)
	//fmt.Fprintln(w, "no service")

	// 临时重定向，先写入首部，再调用 WriteHeader 方法
	//w.Header().Set("Location", "https://google.com")
	//w.WriteHeader(http.StatusFound)

	// 写入 JSON
	post := Post{
		User: "tom",
		Threads: []string{
			"React",
			"Vue",
		},
	}
	json, err := json.Marshal(post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func cookie(w http.ResponseWriter, r *http.Request) {
	//c1 := http.Cookie{
	//	Name:     "c1",
	//	Value:    "v1",
	//	HttpOnly: true,
	//}
	//c2 := http.Cookie{
	//	Name:     "c2",
	//	Value:    "v2",
	//	HttpOnly: true,
	//}

	// 设置 cookie 的方法 1
	//header := w.Header()
	//header.Set("Set-Cookie", c1.String())
	//header.Add("Set-Cookie", c2.String())

	// 设置 cookie 的方法 2
	//http.SetCookie(w, &c1)
	//http.SetCookie(w, &c2)

	// 从客户端获取 cookie
	//c := r.Header["Cookie"]
	// 一个切片，里面一个字符串，包含多个 cookie，得自己解析
	// [c1=v1;c2=v2]
	//fmt.Fprintln(w, c)

	// 从客户端获取 cookie 的方法 2
	// 返回 Cookie 结构的指针
	c, err := r.Cookie("c1")
	if err != nil {
		log.Println(err.Error())
	}
	// 返回 Cookie 结构的指针切片
	cs := r.Cookies()
	// c1=v1
	fmt.Fprintln(w, c)
	// [c1=v1 c2=v2]
	fmt.Fprintln(w, cs)
}

func setMsg(w http.ResponseWriter, r *http.Request) {
	msg := []byte("Hello~")
	c := http.Cookie{
		Name:  "flash",
		Value: base64.URLEncoding.EncodeToString(msg),
	}
	http.SetCookie(w, &c)
}

func showMsg(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("flash")
	if err != nil {
		// http.ErrNoCookie
		log.Println(err.Error())
		fmt.Fprintln(w, "no message")
		return
	}
	// 用新 cookie 冲掉之前的 cookie，并且瞬间过期，即刷新后就没了
	rc := http.Cookie{
		Name:    "flash",
		MaxAge:  -1,
		Expires: time.Unix(1, 0),
	}
	http.SetCookie(w, &rc)
	val, err := base64.URLEncoding.DecodeString(c.Value)
	fmt.Fprintln(w, string(val))
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/login", login)
	http.HandleFunc("/body", body)
	http.HandleFunc("/body2", body2)
	http.HandleFunc("/body3", body3)
	http.HandleFunc("/upload", upload)
	http.HandleFunc("/write", write)
	http.HandleFunc("/cookie", cookie)
	http.HandleFunc("/set-msg", setMsg)
	http.HandleFunc("/show-msg", showMsg)

	WriteFile()
	ReadFile()
	CreateAndWriteFile()
	OpenFile()

	//fmt.Println("Server at http://localhost:8888")
	//err := http.ListenAndServe(":8888", nil)
	//if err != nil {
	//	return
	//}
}
