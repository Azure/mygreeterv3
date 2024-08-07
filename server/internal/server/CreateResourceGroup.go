package server

import (
	"context"

	"github.com/Azure/aks-middleware/ctxlogger"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	pb "go.goms.io/aks/rp/mygreeterv3/api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) CreateResourceGroup(ctx context.Context, in *pb.CreateResourceGroupRequest) (*emptypb.Empty, error) {
	logger := ctxlogger.GetLogger(ctx)
	if s.ResourceGroupClient == nil {
		logger.Error("ResourceGroupClient is nil in CreateResourceGroup(), azuresdk feature is likely disabled")
		return &emptypb.Empty{}, status.Errorf(codes.Unimplemented, "ResourceGroupClient is nil in CreateResourceGroup(), azuresdk feature is likely disabled")
	}
	resourceGroup, err := s.ResourceGroupClient.CreateOrUpdate(
		ctx,
		in.GetName(),
		armresources.ResourceGroup{
			Location: to.Ptr(in.GetRegion()),
		},
		nil)

	if err != nil {
		logger.Error("CreateOrUpdate() error: " + err.Error())
		return &emptypb.Empty{}, HandleError(err, "CreateOrUpdate()")
	}

	logger.Info("Created resource group: " + *resourceGroup.ResourceGroup.ID)
	return &emptypb.Empty{}, nil
}
