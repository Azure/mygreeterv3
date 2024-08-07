package server

import (
	"context"

	"github.com/Azure/aks-middleware/ctxlogger"
	pb "go.goms.io/aks/rp/mygreeterv3/api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) DeleteStorageAccount(ctx context.Context, in *pb.DeleteStorageAccountRequest) (*emptypb.Empty, error) {
	logger := ctxlogger.GetLogger(ctx)
	if s.AccountsClient == nil {
		logger.Error("AccountsClient is nil in DeleteStorageAccount(), azuresdk feature is likely disabled")
		return &emptypb.Empty{}, status.Errorf(codes.Unimplemented, "AccountsClient is nil in DeleteStorageAccount(), azuresdk feature is likely disabled")
	}

	_, err := s.AccountsClient.Delete(context.Background(), in.GetRgName(), in.GetSaName(), nil)
	if err != nil {
		logger.Error("Delete() error: " + err.Error())
		return &emptypb.Empty{}, HandleError(err, "Delete()")
	}

	logger.Info("Deleted storage account: " + in.GetSaName())
	return &emptypb.Empty{}, nil
}
