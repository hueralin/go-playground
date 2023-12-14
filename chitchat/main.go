package main

import (
	"chitchat/data"
	"html/template"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// 设置静态文件服务
	fs := http.FileServer(http.Dir("public"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// 注册路由
	mux.HandleFunc("/", home)
	//mux.HandleFunc("/err", err)
	//mux.HandleFunc("/login", login)
	//mux.HandleFunc("/logout", logout)
	//mux.HandleFunc("/signup", signup)
	//mux.HandleFunc("/signup_account", signupAccount)
	//mux.HandleFunc("/authenticate", authenticate)
	//mux.HandleFunc("/thread/new", newThread)
	//mux.HandleFunc("/thread/create", createThread)
	//mux.HandleFunc("/thread/post", postThread)
	//mux.HandleFunc("/thread/read", readThread)

	server := &http.Server{
		Addr:    ":8888",
		Handler: mux,
	}

	log.Println("http://localhost:8888")
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"templates/layout.html",
		"templates/public.navbar.html",
		"templates/index.html",
	}
	// 解析模板
	templates := template.Must(template.ParseFiles(files...))
	// 获取 Threads
	threads, err := data.GetThreads()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// 执行模板（将数据和模板组合，然后写入到 writer）
	err = templates.ExecuteTemplate(w, "layout", threads)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
