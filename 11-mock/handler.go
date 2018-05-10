package main

import (
	"net/http"

	"github.com/NeowayLabs/logger"
	"github.com/gin-gonic/gin"
)

func postPessoa(pessoaService PessoaService) func(c *gin.Context) {
	return func(c *gin.Context) {
		pessoa := &Pessoa{}

		if err := c.BindJSON(pessoa); err != nil {
			logger.Error("error: %q", err)
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		if err := pessoaService.Save(pessoa); err != nil {
			logger.Error("error: %q", err)
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusCreated, pessoa)
	}
}
