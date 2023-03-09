package router

import (
	"airbnb-user-be/internal/pkg/http/server/middleware/cookie"

	"github.com/gin-gonic/gin"
)

func ProvideRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())

	// GIN apply CORS setting
	router.Use(CORSSetting())

	// GIN bind all cookie
	router.Use(cookie.BindAll())

	return router
}
