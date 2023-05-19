package grpcserver

import (
	"airbnb-user-be/internal/pkg/log"
	"net"
)

func (s *Server) Start() error {
	log.Event(Instance, "starting server...")

	listener, err := net.Listen("tcp", s.address)
	if err != nil {
		log.Fatal(Instance, "failed to create listener", err)
	}

	s.Listener = listener

	return s.Server.Serve(listener)
}

func (s *Server) Stop() error {
	log.Event(Instance, "shutting down server...")

	s.Server.GracefulStop()
	return s.Listener.Close()
}
