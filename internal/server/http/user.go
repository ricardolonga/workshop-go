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

	user, err := h.userService.Create(user)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (h *handler) getUser(c *gin.Context) {
	userID := c.Param("id")

	user, err := h.userService.Retrieve(userID)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *handler) deleteUser(c *gin.Context) {
	userID := c.Param("id")

	if err := h.userService.Delete(userID); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.AbortWithStatus(http.StatusNoContent)
}
