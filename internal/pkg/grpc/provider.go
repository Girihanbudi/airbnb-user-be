package grpc

import (
	"airbnb-user-be/internal/pkg/grpc/config"
	"airbnb-user-be/internal/pkg/log"
	"crypto/tls"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const Instance string = "GRPC"

type Options struct {
	config.Config
}

type Server struct {
	Options
	Server *grpc.Server
}

func NewRpcServer(options Options) *Server {
	tlsCredentials, err := loadTLSCredentials(options.PublicCert, options.PrivateKey)
	if err != nil {
		log.Fatal(Instance, "cannot load TLS credentials", err)
	}

	// interceptor := service.NewAuthInterceptor(jwtManager, accessibleRoles())

	server := grpc.NewServer(
		grpc.Creds(tlsCredentials),
		// grpc.UnaryInterceptor(interceptor.Unary()),
		// grpc.StreamInterceptor(interceptor.Stream()),
	)

	return &Server{
		Server: server,
	}
}

func loadTLSCredentials(publicCert, privateKey string) (credentials.TransportCredentials, error) {
	// Load server's certificate and private key
	serverCert, err := tls.LoadX509KeyPair(publicCert, privateKey)
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
	}

	return credentials.NewTLS(config), nil
}
