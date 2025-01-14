package notfound

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NotFoundHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"path":    c.Request.URL.Path,
			"verb":    c.Request.Method,
			"rawPath": c.Request.URL.RawPath,
			"headers": c.Request.Header,
			"URI":     c.Request.RequestURI,
		})
	}
}
