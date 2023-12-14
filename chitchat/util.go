package main

import (
	"chitchat/data"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
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
