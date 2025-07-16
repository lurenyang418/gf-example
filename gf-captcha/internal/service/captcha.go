// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"gf-captcha/internal/model"
)

type (
	ICaptcha interface {
		CaptchaGet(ctx context.Context) (*model.CaptchaGetResp, error)
		CaptchaVerify(ctx context.Context, req *model.CaptchaVerifyReq) error
	}
)

var (
	localCaptcha ICaptcha
)

func Captcha() ICaptcha {
	if localCaptcha == nil {
		panic("implement not found for interface ICaptcha, forgot register?")
	}
	return localCaptcha
}

func RegisterCaptcha(i ICaptcha) {
	localCaptcha = i
}
