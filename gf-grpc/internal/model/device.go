package model

type DeviceRegisterReq struct {
	Name string `json:"name"`
}

type DeviceRegisterResp struct {
	DeviceId int64 `json:"deviceId"`
}
