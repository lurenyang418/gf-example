package captcha

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "gf-captcha/api/captcha/v1"
	"gf-captcha/internal/service"
)

func (c *ControllerV1) CaptchaVerify(ctx context.Context, req *v1.CaptchaVerifyReq) (res *v1.CaptchaVerifyRes, err error) {
	if err := service.Captcha().CaptchaVerify(ctx, req.CaptchaVerifyReq); err != nil {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "captcha verification failed")
	}

	return
}
