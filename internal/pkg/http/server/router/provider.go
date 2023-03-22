package router

import (
	"airbnb-user-be/internal/pkg/env"
	"airbnb-user-be/internal/pkg/http/server/middleware/cookie"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	if env.CONFIG.Stage == string(env.StageLocal) {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.Recovery())

	// GIN apply CORS setting
	router.Use(CORSSetting())

	// GIN bind all cookie
	router.Use(cookie.BindAll())

	return router
}
