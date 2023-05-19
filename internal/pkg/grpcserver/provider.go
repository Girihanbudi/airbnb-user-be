package grpcserver

import (
	"airbnb-user-be/internal/pkg/credential"
	"airbnb-user-be/internal/pkg/grpcserver/config"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const Instance string = "GRPC"

type Options struct {
	config.Config
	Creds credential.TlsCredentials
}

type Server struct {
	Options
	address  string
	Server   *grpc.Server
	Listener net.Listener
}

func NewRpcServer(options Options) *Server {

	var server *grpc.Server
	if options.Creds.TlsCerts == nil {
		server = grpc.NewServer()
	} else {
		// interceptor := service.NewAuthInterceptor(jwtManager, accessibleRoles())
		tls := credentials.NewTLS(options.Creds.TlsConfig)
		server = grpc.NewServer(
			grpc.Creds(tls),
			// grpc.UnaryInterceptor(interceptor.Unary()),
			// grpc.StreamInterceptor(interceptor.Stream()),
		)
	}

	return &Server{
		Options: options,
		address: fmt.Sprintf("%s:%s", options.Host, options.Port),
		Server:  server,
	}
}
