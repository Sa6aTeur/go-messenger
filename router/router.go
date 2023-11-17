package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello my friend")
	})

	return r
}
