package main

import (
	"go-web/handlers"
	"net/http"
)

func main09() {
	// Go 内置路由并不强大，很多功能都需要自己实现
	// 有些第三方的实现，如 gorilla/mux, httprouter
	server := http.Server{Addr: "localhost:8888"}
	handlers.RegisterHandlers()
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
