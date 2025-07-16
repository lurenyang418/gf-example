package captcha

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "gf-captcha/api/captcha/v1"
	"gf-captcha/internal/service"
)

// CaptchaGet 函数用于获取验证码
func (c *ControllerV1) CaptchaGet(ctx context.Context, req *v1.CaptchaGetReq) (res *v1.CaptchaGetRes, err error) {
	resp, err := service.Captcha().CaptchaGet(ctx)
	if err != nil {
		return nil, gerror.NewCode(gcode.CodeInternalError, "failed to get captcha")
	}

	return &v1.CaptchaGetRes{CaptchaGetResp: resp}, nil
}
