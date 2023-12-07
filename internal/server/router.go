package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) setupApis(group *gin.RouterGroup) {
	group.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
}
