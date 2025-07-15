// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"gf-grpc/internal/model"
)

type (
	IDevice interface {
		Register(ctx context.Context, in *model.DeviceRegisterReq) (*model.DeviceRegisterResp, error)
	}
)

var (
	localDevice IDevice
)

func Device() IDevice {
	if localDevice == nil {
		panic("implement not found for interface IDevice, forgot register?")
	}
	return localDevice
}

func RegisterDevice(i IDevice) {
	localDevice = i
}
