package router

import (
	"github.com/gin-gonic/gin"
)

func ProvideRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())

	// GIN apply CORS setting
	router.Use(CORSSetting())

	return router
}
