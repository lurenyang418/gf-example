package device

import "gf-grpc/internal/service"

type (
	sDevice struct{}
)

func init() {
	service.RegisterDevice(New())
}

func New() *sDevice {
	return &sDevice{}
}
