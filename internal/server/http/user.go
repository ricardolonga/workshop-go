package http

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/ricardolonga/workshop-go/domain"
)

func (h *handler) postUser(c *gin.Context) {
	user := &domain.User{}

	if err := c.BindJSON(&user); err != nil {
		return
	}

	c.JSON(http.StatusOK, nil)
}
