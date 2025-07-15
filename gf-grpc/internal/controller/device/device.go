package device

import (
	"context"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "gf-grpc/api/device/v1"
	"gf-grpc/internal/model"
	"gf-grpc/internal/service"
)

type Controller struct {
	v1.UnimplementedDeviceServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterDeviceServer(s.Server, &Controller{})
}

func (*Controller) Register(ctx context.Context, req *v1.RegisterReq) (*v1.RegisterResp, error) {
	resp, err := service.Device().Register(ctx, &model.DeviceRegisterReq{
		Name: req.Name,
	})
	if err != nil {
		return nil, gerror.Wrap(err, "")
	}

	return &v1.RegisterResp{DeviceId: resp.DeviceId}, nil
}
