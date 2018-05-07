package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func postPessoa() func(c *gin.Context) {
	return func(c *gin.Context) {
		pessoa := &Pessoa{}

		if err := c.BindJSON(&pessoa); err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusCreated, pessoa)
	}
}
