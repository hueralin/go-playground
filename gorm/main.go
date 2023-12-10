package main

import (
	"github.com/gin-gonic/gin"
	"go-playground/gorm/router"
	"log"
)

func main() {
	r := gin.Default()
	router.InitRouter(r)
	log.Println("http://localhost:8888")
	err := r.Run(":8888")
	if err != nil {
		log.Fatalln(err)
	}
}
