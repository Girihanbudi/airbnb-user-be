package server

import (
	"airbnb-user-be/internal/pkg/http/server/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

const Instance = "HTTP Server"

type Options struct {
	config.Config
	Router *gin.Engine
}

type Server struct {
	Options
	address string
	server  *http.Server
}

func NewServer(options Options) *Server {
	return &Server{Options: options}
}
