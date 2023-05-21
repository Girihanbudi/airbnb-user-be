package grpcserver

import (
	"airbnb-user-be/internal/pkg/log"
	"fmt"
	"net"
)

func (s *Server) Start() {
	log.Event(Instance, "starting rpc listener...")

	listener, err := net.Listen("tcp", s.address)
	if err != nil {
		log.Fatal(Instance, "failed to create listener", err)
	}

	s.Listener = listener

	if err := s.Server.Serve(listener); err != nil {
		log.Fatal(Instance, "failed to start rpc listener", err)
	}

	log.Event(Instance, fmt.Sprintf("listening on %s", s.address))
}

func (s *Server) Stop() {
	log.Event(Instance, "shutting down rpc listener...")
	s.Server.GracefulStop()
	log.Event(Instance, "rpc listener has been shutted down")
}
