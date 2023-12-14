package main

import (
	"chitchat/data"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"reflect"
	"runtime"
)

func Session(r *http.Request) (data.Session, error) {
	var sess data.Session
	cookie, err := r.Cookie("sid")
	if err != nil {
		log.Println("Session err:", err)
		return sess, err
	}
	sess = data.Session{Uuid: cookie.Value}
	ok, err := sess.Check()
	if err != nil {
		return sess, err
	}
	if !ok {
		return sess, errors.New("invalid session")
	}
	return sess, nil
}

func GenerateHTML(w http.ResponseWriter, data interface{}, fileNames ...string) error {
	files := make([]string, len(fileNames))
	for i, file := range fileNames {
		files[i] = fmt.Sprintf("templates/%s.html", file)
	}
	// 解析模板文件（进行语法分析），得到模板
	tmpl := template.Must(template.ParseFiles(files...))
	// 执行模板，将模板和数据拼接成 HTML 文件，并写入到 writer
	return tmpl.ExecuteTemplate(w, "layout", data)
}

// Log 串联处理器函数，类似于中间件
func Log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name :=
			runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("Handler function called - " + name)
		h(w, r)
	}
}

// Log2 串联处理器，类似于中间件
func Log2(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name :=
			runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("Handler function called - " + name)
		h.ServeHTTP(w, r)
	})
}
