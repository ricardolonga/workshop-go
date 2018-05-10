package main

import (
	"net/http"

	"github.com/NeowayLabs/logger"
	"github.com/gin-gonic/gin"
)

func postPessoa(c *gin.Context) {
	pessoa := &Pessoa{}

	if err := c.BindJSON(pessoa); err != nil {
		logger.Error("error: %q", err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, pessoa)
}
