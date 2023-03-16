package router

import (
	"airbnb-user-be/internal/pkg/env"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CORSSetting() gin.HandlerFunc {

	var origin string
	if env.Stage(env.CONFIG.Stage) == env.StageLocal {
		origin = env.CONFIG.LocalServer
	} else {
		origin = env.CONFIG.Domain
	}

	return func(c *gin.Context) {
		c.Writer.Header().Set("SameSite", "None")
		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, HEAD, POST, PUT, DELETE, OPTIONS, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
