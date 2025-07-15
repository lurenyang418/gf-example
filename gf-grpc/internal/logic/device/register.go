package device

import (
	"context"
	"gf-grpc/internal/dao"
	"gf-grpc/internal/model"
	"gf-grpc/internal/model/do"
)

func (s *sDevice) Register(ctx context.Context, in *model.DeviceRegisterReq) (*model.DeviceRegisterResp, error) {
	deviceId, err := dao.Device.Ctx(ctx).InsertAndGetId(
		&do.Device{Name: in.Name},
	)
	if err != nil {
		return nil, err
	}

	return &model.DeviceRegisterResp{DeviceId: deviceId}, nil
}
