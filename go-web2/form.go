package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {

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
		err := r.ParseForm()
		if err != nil {
			log.Println(err.Error())
			return
		}
		//fmt.Println("username: ", r.FormValue("username"))
		//fmt.Println("password: ", r.FormValue("password"))
		fmt.Println("username: ", r.Form["username"]) // [tom jack]
		fmt.Println("password: ", r.Form["password"]) // [123]
	}
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/login", login)
	fmt.Println("Server at http://localhost:8888")
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		return
	}
}