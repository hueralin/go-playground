package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"html/template"
	"log"
	"net/http"
)

type User struct {
	id       uint `gorm:"primaryKey"`
	username string
	email    string
	password string
}

type Post struct {
	Title   string
	Content string
}

var posts []Post

var db *gorm.DB

func main() {
	// 连接数据库
	dsn := "malin:123456789@tcp(localhost:3306)/blog"
	var dbErr error
	db, dbErr = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if dbErr != nil {
		log.Fatal(dbErr)
		return
	}

	// 注册路由
	http.HandleFunc("/", listPosts)
	http.HandleFunc("/create", createPost)

	// 启动服务
	log.Println("Server is running on http://localhost:8080")
	runErr := http.ListenAndServe(":8080", nil)
	if runErr != nil {
		log.Fatal(runErr)
		return
	}
}

func listPosts(w http.ResponseWriter, r *http.Request) {
	var posts []Post
	result := db.Find(&posts)
	if result.Error != nil {
		log.Fatal(result.Error)
		return
	}

	tmpl := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Simple Blog</title>
	</head>
	<body>
		<h1>Simple Blog</h1>
		<ul>
			{{range .}}
			<li>
				<h2>{{.Title}}</h2>
				<p>{{.Content}}</p>
			</li>
			{{end}}
		</ul>
		<a href="/create">Create New Post</a>
	</body>
	</html>
	`

	t := template.Must(template.New("").Parse(tmpl))
	err := t.Execute(w, result)
	if err != nil {
		return
	}
}

func createPost(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		title := r.FormValue("title")
		content := r.FormValue("content")
		posts = append(posts, Post{Title: title, Content: content})
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tmpl := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Create New Post</title>
	</head>
	<body>
		<h1>Create New Post</h1>
		<form method="post" action="/create">
			<label for="title">Title:</label><br>
			<input type="text" id="title" name="title"><br><br>
			<label for="content">Content:</label><br>
			<textarea id="content" name="content" rows="4" cols="50"></textarea><br><br>
			<input type="submit" value="Create Post">
		</form>
		<a href="/">Back to Home</a>
	</body>
	</html>
	`

	t := template.Must(template.New("").Parse(tmpl))
	err := t.Execute(w, nil)
	if err != nil {
		return
	}
}
