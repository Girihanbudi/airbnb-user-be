package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const Instance = "HTTP Server"

type Options struct {
	Config
	Router *gin.Engine
}

type Server struct {
	address string
	server  *http.Server
	Options
}

func ProvideServer(options Options) *Server {
	return &Server{Options: options}
}
