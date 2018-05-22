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

	if !h.userService.IsValid(user) {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusCreated, nil)
}
