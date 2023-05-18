package grpc

import (
	"airbnb-user-be/internal/pkg/credential"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const Instance string = "GRPC"

type Options struct {
	Creds credential.TlsCredentials
}

type Server struct {
	Options
	Server *grpc.Server
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
		Server: server,
	}
}
