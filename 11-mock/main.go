package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.New()

	pessoaService := &PessoaServiceImpl{}
	r.POST("/pessoas", postPessoa(pessoaService))

	r.Run()
}
