package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ricardolonga/workshop-go/domain"
)

func (h *handler) postUser(c *gin.Context) {
	var u domain.User
	if err := c.BindJSON(&u); err != nil {
		return
	}

	if !h.userService.IsValid(&u) {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user, err := h.userService.Create(&u)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, user)
}
