package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct{}

func NewHandler() http.Handler {
	handler := &handler{}

	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(handler.recovery())

	v1 := router.Group("/v1")
	v1.POST("/users", handler.postUser)

	return router
}

func (h *handler) recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
