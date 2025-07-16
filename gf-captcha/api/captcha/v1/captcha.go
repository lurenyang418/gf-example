package v1

import (
	"github.com/gogf/gf/v2/frame/g"

	"gf-captcha/internal/model"
)

type CaptchaGetReq struct {
	g.Meta `path:"/captcha" tags:"captcha" method:"get" sm:"获取验证码"`
}
type CaptchaGetRes struct {
	*model.CaptchaGetResp
}

type CaptchaVerifyReq struct {
	g.Meta `path:"/captcha/verify" tags:"captcha" method:"post" sm:"验证验证码"`
	*model.CaptchaVerifyReq
}
type CaptchaVerifyRes struct{}
