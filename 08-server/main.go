package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.New()

	r.POST("/pessoas", postPessoa())

	r.Run()
}
