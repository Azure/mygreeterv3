package client

import (
	"github.com/Azure/aks-middleware/interceptor"
	pb "go.goms.io/aks/rp/mygreeterv3/api/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	log "log/slog"
)

// NewClient returns a client that has all the interceptors registered.
func NewClient(remoteAddr string, options interceptor.ClientInterceptorLogOptions) (pb.MyGreeterClient, error) {
	conn, err := grpc.Dial(
		remoteAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(
			interceptor.DefaultClientInterceptors(options)...,
		),
	)
	if err != nil {
		// logging for transparency, error handled by retry interceptor
		log.Error("did not connect: " + err.Error())
	}

	return pb.NewMyGreeterClient(conn), err
	// TODO: Figure out how to close the connection when the program exits.
	// defer conn.Close()
}
