package server

import (
	log "log/slog"
	"os"

	"github.com/Azure/aks-middleware/interceptor"
	pb "go.goms.io/aks/rp/mygreeterv3/api/v1"
	"go.goms.io/aks/rp/mygreeterv3/server/internal/logattrs"

	"go.goms.io/aks/rp/mygreeterv3/api/v1/client"
)

type Server struct {
	// When the UnimplementedMyGreeterServer struct is embedded,
	// the generated method/implementation in .pb file will be associated with this struct.
	// If this struct doesn't implment some methods,
	// the .pb ones will be used. If this struct implement the methods, it will override the .pb ones.
	// The reason is that anonymous field's methods are promoted to the struct.
	//
	// When this struct is NOT embedded,, all methods have to be implemented to meet the interface requirement.
	// See https://go.dev/ref/spec#Struct_types.
	pb.UnimplementedMyGreeterServer
	client pb.MyGreeterClient
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) init(options Options) {
	var err error

	logger := log.New(log.NewTextHandler(os.Stdout, nil).WithAttrs(logattrs.GetAttrs()))
	if options.JsonLog {
		logger = log.New(log.NewJSONHandler(os.Stdout, nil).WithAttrs(logattrs.GetAttrs()))
	}

	if options.RemoteAddr != "" {
		s.client, err = client.NewClient(options.RemoteAddr, interceptor.GetClientInterceptorLogOptions(logger, logattrs.GetAttrs()))
		// logging the error for transparency, retry interceptor will handle it
		if err != nil {
			log.Error("did not connect: " + err.Error())
		}
	}
}
