package main

import (
	"chitchat/data"
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
	// 准备模板
	publicTmplFiles := []string{"layout", "public.navbar", "index"}
	privateTmplFiles := []string{"layout", "private.navbar", "index"}
	// 获取 Threads
	threads, err := data.GetThreads()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// 检查登录状态
	_, err = Session(r)
	// 根据状态解析对应的模板文件（进行语法分析），得到模板
	if err != nil {
		_err := GenerateHTML(w, threads, publicTmplFiles...)
		if _err != nil {
			log.Println(_err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else {
		_err := GenerateHTML(w, threads, privateTmplFiles...)
		if _err != nil {
			log.Println(_err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
