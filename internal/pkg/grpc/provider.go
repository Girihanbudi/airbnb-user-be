package grpc

import (
	"airbnb-user-be/internal/pkg/credential"

	"google.golang.org/grpc"
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
	if options.Creds.Tls == nil {
		server = grpc.NewServer()
	} else {
		// interceptor := service.NewAuthInterceptor(jwtManager, accessibleRoles())

		server = grpc.NewServer(
			grpc.Creds(*options.Creds.Tls),
			// grpc.UnaryInterceptor(interceptor.Unary()),
			// grpc.StreamInterceptor(interceptor.Stream()),
		)
	}

	return &Server{
		Server: server,
	}
}
